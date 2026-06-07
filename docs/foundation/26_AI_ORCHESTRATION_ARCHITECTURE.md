# 26_AI_ORCHESTRATION_ARCHITECTURE.md

## Foundation Document for NUSA Education Platform

**Version**: 1.0
**Date**: June 2026
**Status**: FOUNDATION DOCUMENT
**Alignment**: Aligned with 05_AI_ARCHITECTURE.md, 15_AI_PROMPT_SPECIFICATION.md, 06_APPLICATION_ARCHITECTURE.md

**Purpose**: Define the AI execution architecture for NUSA MVP, establishing the orchestration layer that manages AI provider interactions, prompt building, response processing, and cost governance while maintaining MVP-friendly implementation.

---

# SECTION 1 — Executive Summary

## Why AI Orchestration Architecture Matters

NUSA MVP uses AI to generate educational artifacts (TP, and eventually ATP, Modul Ajar, etc.). A well-designed orchestration layer ensures:

- Provider-agnostic integration (switch between OpenAI, Gemini, etc.)
- Consistent prompt management and versioning
- Reliable failure handling and fallback mechanisms
- Complete audit trail of AI generations
- Cost tracking and governance
- Synchronous and asynchronous execution modes

## Core Principles

- **AI Assists, Humans Decide**: AI generates recommendations, teachers approve
- **Explainability Over Autonomy**: System shows what AI generated and why
- **Provider Agnostic**: Architecture supports multiple AI providers
- **MVP-Friendly**: Simple, implementable architecture without over-engineering
- **Cost-Aware**: Track and govern AI usage costs
- **Reliable**: Handle failures gracefully with fallbacks

---

# SECTION 2 — AI Principles

## Principle 1: AI Assists Teachers

**Rule**: AI generates artifact recommendations, but teachers make final decisions.

**Implications**:
- AI output is always a draft, never final
- Teachers can edit, reject, or regenerate AI output
- AI confidence scores are displayed but not used for auto-decisions
- AI suggestions are explainable (teachers can see reasoning)

## Principle 2: Human Approval Required

**Rule**: AI-generated artifacts require human approval before becoming official.

**Implications**:
- AI output goes to DRAFT status
- DRAFT artifacts cannot be used for downstream generation
- Approval workflow is mandatory (see 25_WORKFLOW_ARCHITECTURE.md)
- No auto-approval based on confidence or other metrics

## Principle 3: Explainability Preferred Over Autonomy

**Rule**: System prioritizes showing teachers what AI generated and why, over making autonomous decisions.

**Implications**:
- Prompt snapshots stored for transparency
- AI reasoning included in output when available
- Teachers can see which CP objectives informed TP objectives
- Confidence scores displayed but not used for auto-actions

---

# SECTION 3 — AI Components

## Component Overview

The AI orchestration layer consists of five core components:

```
┌─────────────────────────────────────────────────────────┐
│                    Application Layer                      │
│              (TP Service, ATP Service, etc.)              │
└────────────────────┬────────────────────────────────────┘
                     │
                     ↓
┌─────────────────────────────────────────────────────────┐
│                    AI Gateway                             │
│              (Entry point for AI requests)                │
└────────────────────┬────────────────────────────────────┘
                     │
        ┌────────────┼────────────┐
        ↓            ↓            ↓
┌─────────────┐ ┌──────────┐ ┌──────────────┐
│Prompt Builder│ │Provider  │ │Response      │
│              │ │Adapter   │ │Processor     │
└─────────────┘ └──────────┘ └──────────────┘
        │            │            │
        └────────────┼────────────┘
                     ↓
            ┌──────────────┐
            │Generation    │
            │Logger        │
            └──────────────┘
```

## Component 1: AI Gateway

**Purpose**: Entry point for all AI generation requests from application services.

**Responsibilities**:
- Receive generation requests from application services
- Validate request parameters
- Route to appropriate prompt builder
- Coordinate prompt building, AI call, and response processing
- Return formatted response to calling service
- Handle synchronous and asynchronous execution modes

**Interface**:

```typescript
interface IAIGateway {
  generate(request: AIGenerationRequest): Promise<AIGenerationResponse>;
  generateAsync(request: AIGenerationRequest): Promise<string>; // Returns job ID
  getJobStatus(jobId: string): Promise<JobStatus>;
}

interface AIGenerationRequest {
  artifact_type: 'TP' | 'ATP' | 'MODUL_AJAR' | 'ASSESSMENT' | 'RUBRIC';
  prompt_template: string;
  prompt_variables: Record<string, any>;
  provider_preference?: string; // Optional provider override
  execution_mode?: 'sync' | 'async';
}

interface AIGenerationResponse {
  success: boolean;
  artifact_id?: string;
  draft_content?: any;
  metadata: {
    generation_id: string;
    provider: string;
    model: string;
    tokens_used: number;
    cost: number;
    response_time_ms: number;
    confidence_score?: number;
  };
  error?: string;
}
```

## Component 2: Prompt Builder

**Purpose**: Construct AI prompts from templates and variables.

**Responsibilities**:
- Load prompt template based on artifact type
- Substitute variables into template
- Validate prompt completeness
- Store prompt snapshot for audit trail
- Support prompt versioning

**Interface**:

```typescript
interface IPromptBuilder {
  buildPrompt(template: string, variables: Record<string, any>): Promise<Prompt>;
  getTemplate(artifact_type: string, version: string): Promise<string>;
  validatePrompt(prompt: string): boolean;
}

interface Prompt {
  content: string;
  template_version: string;
  variables: Record<string, any>;
  snapshot: string; // For audit trail
}
```

## Component 3: Provider Adapter

**Purpose**: Abstract AI provider implementations (OpenAI, Gemini, etc.).

**Responsibilities**:
- Implement provider-specific API calls
- Handle provider-specific response formats
- Normalize responses to common format
- Manage provider authentication
- Handle provider-specific rate limits

**Interface**:

```typescript
interface IProviderAdapter {
  generate(prompt: string, options: GenerationOptions): Promise<ProviderResponse>;
  validateConfiguration(): Promise<boolean>;
  estimateCost(prompt: string, model: string): number;
}

interface GenerationOptions {
  model: string;
  temperature?: number;
  max_tokens?: number;
  top_p?: number;
}

interface ProviderResponse {
  content: string;
  tokens_used: number;
  model: string;
  provider: string;
  raw_response: any; // For debugging
}
```

**Provider Implementations**:

- OpenAIAdapter (GPT-4, GPT-3.5)
- GeminiAdapter (Gemini Pro, Gemini Flash)
- Future: AnthropicAdapter, etc.

## Component 4: Response Processor

**Purpose**: Validate and process AI provider responses.

**Responsibilities**:
- Validate response structure (JSON schema validation)
- Extract artifact content from response
- Calculate confidence scores if provided by provider
- Handle malformed responses
- Format response for application consumption

**Interface**:

```typescript
interface IResponseProcessor {
  process(response: ProviderResponse, schema: JSONSchema): ProcessedResponse;
  validateSchema(content: any, schema: JSONSchema): ValidationResult;
  extractConfidence(response: ProviderResponse): number | null;
}

interface ProcessedResponse {
  success: boolean;
  content: any;
  confidence_score?: number;
  validation_errors?: string[];
}
```

## Component 5: Generation Logger

**Purpose**: Log all AI generation requests and responses for audit and cost tracking.

**Responsibilities**:
- Log generation requests with full context
- Log generation responses with metadata
- Track tokens, cost, response time
- Store prompt snapshots
- Provide query capabilities for analytics

**Interface**:

```typescript
interface IGenerationLogger {
  logRequest(request: AIGenerationRequest, generation_id: string): Promise<void>;
  logResponse(response: AIGenerationResponse, generation_id: string): Promise<void>;
  logError(error: Error, generation_id: string): Promise<void>;
  getGenerationHistory(filters: GenerationFilters): Promise<GenerationLog[]>;
}

interface GenerationFilters {
  user_id?: string;
  school_id?: string;
  artifact_type?: string;
  provider?: string;
  date_from?: Date;
  date_to?: Date;
}
```

---

# SECTION 4 — Provider Strategy

## Provider Agnostic Design

**Principle**: Architecture must support multiple AI providers without code changes.

**Implementation**:
- Provider adapters implement common interface
- Provider selection via configuration
- Fallback chain supports multiple providers
- Provider-specific logic isolated in adapters

## Provider Configuration

### Primary Provider

**Definition**: Default provider used for most generations.

**Configuration**:

```yaml
ai_providers:
  primary:
    name: "openai"
    models:
      default: "gpt-4"
      fallback: "gpt-3.5-turbo"
    api_key_env: "OPENAI_API_KEY"
    rate_limit:
      requests_per_minute: 3500
      tokens_per_minute: 90000
```

### Fallback Provider

**Definition**: Secondary provider used if primary fails.

**Configuration**:

```yaml
ai_providers:
  fallback:
    name: "gemini"
    models:
      default: "gemini-pro"
    api_key_env: "GEMINI_API_KEY"
    rate_limit:
      requests_per_minute: 60
      tokens_per_minute: 32000
```

## Provider Selection Logic

```typescript
async selectProvider(request: AIGenerationRequest): Promise<string> {
  // 1. Check if user has provider preference
  if (request.provider_preference) {
    return request.provider_preference;
  }
  
  // 2. Check if primary provider is healthy
  if (await isProviderHealthy(config.primary.name)) {
    return config.primary.name;
  }
  
  // 3. Fallback to secondary provider
  if (config.fallback && await isProviderHealthy(config.fallback.name)) {
    return config.fallback.name;
  }
  
  // 4. All providers unavailable
  throw new AllProvidersUnavailableError();
}
```

## Provider Health Checks

**Health Check Indicators**:
- API key validity
- Rate limit status
- Recent error rate
- Response time (p95)

**Health Check Interval**: Every 60 seconds

**Health Status Cache**: TTL 30 seconds

---

# SECTION 5 — Generation Flow

## Synchronous Generation Flow (MVP Default)

```
┌─────────────────────────────────────────────────────────┐
│ 1. Teacher Request                                      │
│    - Teacher selects CP and context                     │
│    - Clicks "Generate TP"                               │
└────────────────────┬────────────────────────────────────┘
                     │
                     ↓
┌─────────────────────────────────────────────────────────┐
│ 2. Application Service (TP Service)                     │
│    - Validates request                                  │
│    - Calls AI Gateway.generate()                        │
└────────────────────┬────────────────────────────────────┘
                     │
                     ↓
┌─────────────────────────────────────────────────────────┐
│ 3. AI Gateway                                          │
│    - Receives request                                  │
│    - Calls Prompt Builder                              │
└────────────────────┬────────────────────────────────────┘
                     │
                     ↓
┌─────────────────────────────────────────────────────────┐
│ 4. Prompt Builder                                      │
│    - Loads TP template                                 │
│    - Substitutes curriculum hierarchy and context       │
│    - Validates prompt                                   │
│    - Returns prompt                                    │
└────────────────────┬────────────────────────────────────┘
                     │
                     ↓
┌─────────────────────────────────────────────────────────┐
│ 5. Provider Adapter                                    │
│    - Selects provider (primary or fallback)             │
│    - Calls provider API                                 │
│    - Returns raw response                               │
└────────────────────┬────────────────────────────────────┘
                     │
                     ↓
┌─────────────────────────────────────────────────────────┐
│ 6. Response Processor                                  │
│    - Validates JSON schema                             │
│    - Extracts TP Set content                            │
│    - Calculates confidence score                        │
│    - Returns processed response                         │
└────────────────────┬────────────────────────────────────┘
                     │
                     ↓
┌─────────────────────────────────────────────────────────┐
│ 7. Generation Logger                                   │
│    - Logs request with prompt snapshot                  │
│    - Logs response with metadata                        │
│    - Stores tokens, cost, response time                 │
└────────────────────┬────────────────────────────────────┘
                     │
                     ↓
┌─────────────────────────────────────────────────────────┐
│ 8. AI Gateway (Response)                               │
│    - Formats response                                   │
│    - Returns to TP Service                              │
└────────────────────┬────────────────────────────────────┘
                     │
                     ↓
┌─────────────────────────────────────────────────────────┐
│ 9. TP Service                                          │
│    - Creates TP Set in DRAFT status                    │
│    - Stores AI metadata                                 │
│    - Returns to UI                                      │
└────────────────────┬────────────────────────────────────┘
                     │
                     ↓
┌─────────────────────────────────────────────────────────┐
│ 10. UI Display                                          │
│     - Shows TP Set for review                           │
│     - Displays AI confidence score                     │
│     - Teacher reviews and approves                       │
└─────────────────────────────────────────────────────────┘
```

## Asynchronous Generation Flow (Future)

```
┌─────────────────────────────────────────────────────────┐
│ 1. Teacher Request                                      │
│    - Teacher submits generation request                 │
└────────────────────┬────────────────────────────────────┘
                     │
                     ↓
┌─────────────────────────────────────────────────────────┐
│ 2. Application Service                                  │
│    - Validates request                                  │
│    - Calls AI Gateway.generateAsync()                   │
└────────────────────┬────────────────────────────────────┘
                     │
                     ↓
┌─────────────────────────────────────────────────────────┐
│ 3. AI Gateway                                          │
│    - Creates job record                                 │
│    - Publishes job to RabbitMQ                          │
│    - Returns job ID to UI                              │
└────────────────────┬────────────────────────────────────┘
                     │
                     ↓
┌─────────────────────────────────────────────────────────┐
│ 4. RabbitMQ                                             │
│    - Queues job for background worker                   │
└────────────────────┬────────────────────────────────────┘
                     │
                     ↓
┌─────────────────────────────────────────────────────────┐
│ 5. Background Worker                                   │
│    - Consumes job from queue                            │
│    - Executes generation flow (steps 4-8 from sync)     │
│    - Updates job status                                 │
└────────────────────┬────────────────────────────────────┘
                     │
                     ↓
┌─────────────────────────────────────────────────────────┐
│ 6. UI Polling / WebSocket                              │
│    - UI polls job status or receives WebSocket update   │
│    - Displays completion notification                    │
└─────────────────────────────────────────────────────────┘
```

**When to Use Async**:
- Large TP Sets (10+ TPs)
- Batch generation for multiple CPs
- Resource-intensive generations
- User explicitly requests async mode

**MVP Approach**: Start with synchronous only, add async when needed.

---

# SECTION 6 — RabbitMQ Usage

## Purpose

RabbitMQ is used exclusively for long-running AI generation tasks.

## When to Use RabbitMQ

**Use RabbitMQ When**:
- Generation expected to take > 30 seconds
- Batch generation of multiple artifacts
- Resource-intensive operations
- User doesn't need immediate response

**Do NOT Use RabbitMQ When**:
- Single TP generation (typically < 10 seconds)
- User expects immediate response
- Simple, quick operations

## Queue Configuration

### Queue Definition

```yaml
rabbitmq:
  queues:
    ai_generation:
      name: "ai.generation"
      durable: true
      prefetch_count: 5
      max_priority: 10
```

### Job Message Format

```typescript
interface AIJobMessage {
  job_id: string;
  request: AIGenerationRequest;
  user_id: string;
  school_id: string;
  priority: number; // 1-10, higher = more urgent
  created_at: string;
}
```

## Worker Configuration

**Worker Pool Size**: 3-5 workers per instance

**Scaling**: Horizontal scaling based on queue length

**Job Timeout**: 5 minutes per job

**Retry Policy**: 3 retries with exponential backoff

---

# SECTION 7 — AI Generation Log

## ai_generation_logs Table

**Purpose**: Complete audit trail of all AI generations.

**Schema**:

```sql
CREATE TABLE ai_generation_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    generation_id UUID NOT NULL UNIQUE,
    
    -- Request Context
    user_id UUID NOT NULL,
    school_id UUID NOT NULL,
    artifact_type VARCHAR(50) NOT NULL,
    artifact_id UUID, -- Populated after artifact creation
    
    -- Provider Information
    provider VARCHAR(50) NOT NULL,
    model VARCHAR(100) NOT NULL,
    
    -- Prompt Information
    prompt_template VARCHAR(100),
    prompt_version VARCHAR(20),
    prompt_snapshot TEXT NOT NULL,
    
    -- Generation Metadata
    status VARCHAR(20) NOT NULL,
    tokens_used INTEGER,
    estimated_cost DECIMAL(10,6),
    actual_cost DECIMAL(10,6),
    response_time_ms INTEGER,
    confidence_score DECIMAL(3,2),
    
    -- Response Information
    response_snapshot TEXT,
    validation_errors TEXT[],
    
    -- Versioning
    prompt_template_version VARCHAR(20),
    model_version VARCHAR(50),
    
    -- Audit
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    completed_at TIMESTAMP WITH TIME ZONE,
    error_message TEXT,
    error_details JSONB,
    
    -- Foreign Keys
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (school_id) REFERENCES schools(id)
);

CREATE INDEX idx_ai_logs_generation_id ON ai_generation_logs(generation_id);
CREATE INDEX idx_ai_logs_user_id ON ai_generation_logs(user_id);
CREATE INDEX idx_ai_logs_school_id ON ai_generation_logs(school_id);
CREATE INDEX idx_ai_logs_artifact_type ON ai_generation_logs(artifact_type);
CREATE INDEX idx_ai_logs_provider ON ai_generation_logs(provider);
CREATE INDEX idx_ai_logs_status ON ai_generation_logs(status);
CREATE INDEX idx_ai_logs_created_at ON ai_generation_logs(created_at);
```

## Status Values

| Status | Description |
|--------|-------------|
| PENDING | Generation queued or in progress |
| COMPLETED | Generation completed successfully |
| FAILED | Generation failed |
| VALIDATION_FAILED | Response failed schema validation |
| TIMEOUT | Provider timeout |
| CANCELLED | Generation cancelled by user |

## Field Descriptions

| Field | Description |
|-------|-------------|
| `generation_id` | Unique identifier for this generation request |
| `user_id` | User who requested generation |
| `school_id` | School of requesting user |
| `artifact_type` | Type of artifact being generated |
| `artifact_id` | ID of generated artifact (populated after artifact creation) |

## Root Traceability Record

**ai_generation_logs as Root Traceability Record**

The ai_generation_logs table serves as the root traceability record for all AI-generated artifacts in the system. This ensures complete audit trail from AI generation to artifact to workflow history.

**Traceability Chain**:

```
ai_generation_logs (root)
  ↓ (referenced by)
artifact_sets (tp_sets.ai_generation_id, atp_sets.ai_generation_id, modul_ajar_sets.ai_generation_id)
  ↓ (referenced by)
workflow_history (workflow_history.artifact_id)
```

**Benefits**:

- **Complete Audit Trail**: Every AI generation can be traced from initial request through artifact creation to workflow actions
- **Cost Governance**: Enables cost tracking by user, school, artifact type, provider, and model
- **Debugging**: Provides complete context for troubleshooting failed generations
- **Quality Analysis**: Enables analysis of generation quality over time
- **Compliance**: Supports audit requirements for AI usage in educational content

**Implementation Notes**:

- All artifact sets (tp_sets, atp_sets, modul_ajar_sets) include ai_generation_id field
- ai_generation_id is populated when artifact set is created from AI generation
- Manual artifact creation has ai_generation_id = NULL
- Workflow history references artifact_id, which can be traced back to ai_generation_id via artifact set

---

# SECTION 8 — Failure Handling

## Failure Scenarios

### 1. Provider Timeout

**Detection**: Provider API call exceeds timeout threshold (default 30 seconds)

**Handling**:
- Log timeout event
- Attempt fallback provider if configured
- If fallback also times out, return error to user
- Do not retry automatically for timeout (may indicate provider issue)

**User Experience**:
- Display timeout error message
- Offer option to retry with different provider
- Preserve user input for retry

### 2. Provider Unavailable

**Detection**: Provider API returns 5xx error or connection failure

**Handling**:
- Log unavailability event
- Mark provider as unhealthy in health check cache
- Attempt fallback provider immediately
- If fallback also unavailable, return error

**User Experience**:
- Display provider unavailable error
- Show which provider failed
- Offer retry option

### 3. Validation Failed

**Detection**: Response Processor detects schema validation failure

**Handling**:
- Log validation errors
- Store response snapshot for debugging
- Return error to user with validation details
- Do not create artifact from invalid response

**User Experience**:
- Display validation error details
- Show which fields failed validation
- Offer option to regenerate

### 4. Partial Generation

**Detection**: Provider returns incomplete response (truncated, missing fields)

**Handling**:
- Log partial generation event
- Validate what was received
- If critical fields missing, treat as failure
- If non-critical fields missing, may proceed with warning

**User Experience**:
- Display partial generation warning
- Show which fields are missing
- Offer option to regenerate

### 5. Rate Limit Exceeded

**Detection**: Provider returns 429 rate limit error

**Handling**:
- Log rate limit event
- Implement exponential backoff retry
- Max 3 retries with delays: 1s, 2s, 4s
- If all retries fail, return error

**User Experience**:
- Display rate limit error
- Show estimated wait time
- Offer option to retry later

## Retry Strategy

### Retry Configuration

```yaml
ai_retry:
  max_attempts: 3
  backoff_strategy: exponential
  initial_delay_ms: 1000
  max_delay_ms: 4000
  retryable_errors:
    - 429 # Rate limit
    - 500 # Internal server error
    - 502 # Bad gateway
    - 503 # Service unavailable
    - 504 # Gateway timeout
  non_retryable_errors:
    - 400 # Bad request
    - 401 # Unauthorized
    - 403 # Forbidden
    - 404 # Not found
    - 422 # Unprocessable entity
```

### Retry Logic

```typescript
async executeWithRetry(
  operation: () => Promise<any>,
  context: string
): Promise<any> {
  let lastError: Error;
  
  for (let attempt = 1; attempt <= config.max_attempts; attempt++) {
    try {
      return await operation();
    } catch (error) {
      lastError = error;
      
      if (!isRetryable(error)) {
        throw error; // Don't retry non-retryable errors
      }
      
      if (attempt === config.max_attempts) {
        throw new MaxRetriesExceededError(context, lastError);
      }
      
      const delay = calculateBackoff(attempt);
      await sleep(delay);
    }
  }
  
  throw lastError;
}
```

---

# SECTION 9 — Prompt Versioning

## Version Storage

All prompt-related version information is stored in the generation log:

### Prompt Version Fields

| Field | Description |
|-------|-------------|
| `prompt_template` | Name/identifier of the prompt template |
| `prompt_version` | Version of the prompt template (e.g., "1.0", "1.1") |
| `prompt_template_version` | Version of the template file itself |
| `model_version` | Version of the AI model used |

### Version Tracking

**Template Versioning**:
- Prompt templates stored in version control (Git)
- Template version incremented on changes
- Template version referenced in generation log

**Model Versioning**:
- Model version tracked from provider response
- Some providers (e.g., OpenAI) return model version
- Stored for reproducibility analysis

**Prompt Snapshot**:
- Full prompt sent to provider stored as text
- Allows reconstruction of exact prompt used
- Critical for debugging and audit

## Version Query Patterns

### Get Generation History by Prompt Version

```sql
SELECT 
    prompt_version,
    COUNT(*) as generation_count,
    AVG(response_time_ms) as avg_response_time,
    AVG(confidence_score) as avg_confidence
FROM ai_generation_logs
WHERE prompt_version = :prompt_version
GROUP BY prompt_version;
```

### Compare Prompt Versions

```sql
SELECT 
    prompt_version,
    artifact_type,
    AVG(confidence_score) as avg_confidence,
    AVG(response_time_ms) as avg_response_time,
    COUNT(*) as count
FROM ai_generation_logs
WHERE prompt_version IN (:version1, :version2)
GROUP BY prompt_version, artifact_type;
```

---

# SECTION 10 — Cost Governance

## MVP Cost Governance

For MVP, cost governance is simplified to focus on essential tracking and reporting without complex enforcement mechanisms.

### MVP Scope

**What MVP Stores**:

The ai_generation_logs table stores the following cost-related fields:
- `provider` - AI provider used (openai, gemini, etc.)
- `model` - Specific model used (gpt-4, gemini-pro, etc.)
- `tokens_used` - Number of tokens consumed
- `estimated_cost` - Estimated cost based on token count
- `response_time_ms` - Time taken for generation in milliseconds
- `status` - Generation status (COMPLETED, FAILED, etc.)

**What MVP Does NOT Include**:

The following enterprise-level cost governance features are deferred to Wave 2:
- Budget Management
- Quota Management
- Cost Dashboard
- Budget Enforcement
- Chargeback
- Department Billing
- Provider Optimization Engine

### Cost Calculation

**Token-Based Cost**:

```typescript
function calculateCost(tokens: number, provider: string, model: string): number {
  const pricing = getPricing(provider, model);
  return (tokens / 1000) * pricing.per_1k_tokens;
}
```

**Pricing Configuration**:

```yaml
ai_pricing:
  openai:
    gpt-4:
      input_per_1k: 0.03
      output_per_1k: 0.06
    gpt-3.5-turbo:
      input_per_1k: 0.0015
      output_per_1k: 0.002
  gemini:
    gemini-pro:
      input_per_1k: 0.00025
      output_per_1k: 0.0005
```

### Cost Aggregation for Reporting

Costs can be aggregated by multiple dimensions for basic reporting:

**By School**:

```sql
SELECT 
    school_id,
    SUM(estimated_cost) as total_cost,
    COUNT(*) as generation_count
FROM ai_generation_logs
WHERE created_at >= :start_date
  AND created_at <= :end_date
GROUP BY school_id;
```

**By Teacher**:

```sql
SELECT 
    user_id,
    SUM(estimated_cost) as total_cost,
    COUNT(*) as generation_count
FROM ai_generation_logs
WHERE created_at >= :start_date
  AND created_at <= :end_date
GROUP BY user_id;
```

**By Module**:

```sql
SELECT 
    artifact_type,
    SUM(estimated_cost) as total_cost,
    COUNT(*) as generation_count,
    AVG(response_time_ms) as avg_response_time
FROM ai_generation_logs
WHERE created_at >= :start_date
  AND created_at <= :end_date
GROUP BY artifact_type;
```

**By Provider**:

```sql
SELECT 
    provider,
    model,
    SUM(estimated_cost) as total_cost,
    SUM(tokens_used) as total_tokens,
    COUNT(*) as generation_count
FROM ai_generation_logs
WHERE created_at >= :start_date
  AND created_at <= :end_date
GROUP BY provider, model;
```

### Cost Reporting

MVP provides basic cost reporting queries for administrators to monitor AI usage. No automated enforcement or alerts are implemented in MVP.

**Monthly Cost Summary**:

```sql
SELECT 
    DATE_TRUNC('month', created_at) as month,
    SUM(estimated_cost) as total_cost,
    SUM(tokens_used) as total_tokens,
    COUNT(*) as generation_count
FROM ai_generation_logs
GROUP BY DATE_TRUNC('month', created_at)
ORDER BY month DESC;
```

---

## Future Cost Governance (Wave 2)

Future enhancement will add enterprise-level cost governance features:

### Budget Management

- School-level monthly budget limits
- Teacher-level daily/weekly limits
- Budget allocation by department
- Budget rollover rules

### Quota Management

- Per-school generation quotas
- Per-teacher generation quotas
- Quota reset schedules (daily, weekly, monthly)
- Quota overage handling

### Cost Dashboard

- Real-time cost monitoring UI
- Cost trend visualization
- Per-school cost breakdown
- Per-teacher cost breakdown
- Provider cost comparison

### Budget Enforcement

- Automatic blocking when limits exceeded
- Configurable enforcement actions (alert, block, throttle)
- Override mechanisms for emergencies
- Approval workflows for over-budget requests

### Chargeback

- Cost allocation to departments
- Cost allocation to projects
- Chargeback reporting
- Integration with billing systems

### Provider Optimization Engine

- Automatic provider selection based on cost
- Cost-performance analysis
- A/B testing of providers
- Dynamic routing based on cost targets

**Implementation Notes**:

- Foundation preserved: ai_generation_logs table structure supports all future enhancements
- No schema changes required for Wave 2
- Cost aggregation queries already support multi-dimensional analysis
- Pricing configuration already supports multiple providers and models

---

# SECTION 11 — Component Diagram

```
┌──────────────────────────────────────────────────────────────┐
│                      Application Layer                         │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐        │
│  │ TP Service   │  │ ATP Service  │  │ Assessment    │        │
│  │              │  │              │  │ Service       │        │
│  └──────┬───────┘  └──────┬───────┘  └──────┬───────┘        │
└─────────┼──────────────────┼──────────────────┼────────────────┘
          │                  │                  │
          └──────────────────┼──────────────────┘
                             ↓
┌──────────────────────────────────────────────────────────────┐
│                        AI Gateway                             │
│  ┌──────────────────────────────────────────────────────┐   │
│  │  - Request Validation                                │   │
│  │  - Provider Selection                                 │   │
│  │  - Execution Mode (sync/async)                        │   │
│  │  - Response Formatting                                │   │
│  └──────────────────────────────────────────────────────┘   │
└────────────────────┬─────────────────────────────────────────┘
                     │
        ┌────────────┼────────────┐
        ↓            ↓            ↓
┌──────────────┐ ┌──────────┐ ┌──────────────┐
│Prompt Builder│ │Provider  │ │Response      │
│              │ │Adapter   │ │Processor     │
│- Load template│ │- OpenAI  │ │- Schema      │
│- Substitute   │ │- Gemini  │ │  validation  │
│- Validate     │ │- Fallback│ │- Extract     │
│- Snapshot     │ │          │ │  content     │
└──────────────┘ └──────────┘ └──────────────┘
        │            │            │
        └────────────┼────────────┘
                     ↓
            ┌──────────────┐
            │Generation    │
            │Logger        │
            │- Log request │
            │- Log response│
            │- Track cost  │
            │- Audit trail │
            └──────────────┘
                     │
                     ↓
            ┌──────────────┐
            │Database      │
            │- ai_generation_logs│
            │- artifacts    │
            └──────────────┘
```

---

# SECTION 12 — Sequence Diagram

## Synchronous Generation Sequence

```
Teacher          UI          TP Service    AI Gateway    Prompt Builder    Provider Adapter    Response Processor    Generation Logger    Database
  │                │               │              │                │                  │                    │                    │              │
  │Generate TP     │               │              │                │                  │                    │                    │              │
  ├───────────────>│               │              │                │                  │                    │                    │              │
  │                │Generate TP    │              │                │                  │                    │                    │              │
  │                ├───────────────>│              │                │                  │                    │                    │              │
  │                │               │generate()     │                │                  │                    │                    │              │
  │                │               ├──────────────>│                │                  │                    │                    │              │
  │                │               │              │buildPrompt()   │                  │                    │                    │              │
  │                │               │              ├───────────────>│                  │                    │                    │              │
  │                │               │              │                │load template      │                    │                    │              │
  │                │               │              │<───────────────┤                  │                    │                    │              │
  │                │               │              │substitute vars  │                  │                    │                    │              │
  │                │               │              ├───────────────>│                  │                    │                    │              │
  │                │               │              │<───────────────┤                  │                    │                    │              │
  │                │               │              │<───────────────┤                  │                    │                    │              │
  │                │               │              │generate()       │                  │                    │                    │              │
  │                │               │              ├──────────────────────────────────>│                    │                    │              │
  │                │               │              │                │call provider API │                    │                    │              │
  │                │               │              │<──────────────────────────────────┤                    │                    │              │
  │                │               │              │process()       │                  │                    │                    │              │
  │                │               │              ├───────────────────────────────────────────────────────>│                    │              │
  │                │               │              │                │                  │validate schema     │                    │              │
  │                │               │              │<───────────────────────────────────────────────────────┤                    │              │
  │                │               │              │logRequest()    │                  │                    │logRequest()         │              │
  │                │               │              ├──────────────────────────────────────────────────────────────────────────────>│
  │                │               │              │<──────────────────────────────────────────────────────────────────────────────┤
  │                │               │              │logResponse()   │                  │                    │logResponse()        │              │
  │                │               │              ├──────────────────────────────────────────────────────────────────────────────>│
  │                │               │              │<──────────────────────────────────────────────────────────────────────────────┤
  │                │               │<──────────────│                │                  │                    │                    │              │
  │                │create TP Set│               │                │                  │                    │                    │              │
  │                ├───────────────>│               │                │                  │                    │                    │              │
  │                │               │INSERT TP     │                │                  │                    │                    │              │
  │                │               ├──────────────────────────────────────────────────────────────────────────────────────────────>│
  │                │               │<──────────────│                │                  │                    │                    │              │
  │                │<──────────────│               │                │                  │                    │                    │              │
  │Show TP Set     │               │               │                │                  │                    │                    │              │
  │<───────────────│               │               │                │                  │                    │                    │              │
```

---

# SECTION 13 — Security Considerations

## API Key Management

**Storage**: API keys stored in environment variables, never in code

**Rotation**: Support for API key rotation without deployment

**Access**: Only AI Gateway has access to provider API keys

## Prompt Injection Prevention

**Validation**: Validate prompt variables before substitution

**Sanitization**: Sanitize user-provided prompt content

**Limits**: Enforce maximum prompt length limits

## Data Privacy

**Sensitive Data**: No student PII in prompts

**Audit**: Prompt snapshots logged but may be redacted for privacy

**Retention**: Prompt snapshots retained per audit retention policy

---

# SECTION 14 — Performance Considerations

## Response Time Targets

| Operation | Target P95 |
|-----------|------------|
| Single TP generation (sync) | < 10 seconds |
| TP Set generation (3-5 TPs) | < 20 seconds |
| TP Set generation (6-10 TPs) | < 30 seconds |
| Async job completion | < 5 minutes |

## Caching Strategy

**Prompt Templates**: Cached in memory with TTL 1 hour

**Provider Health Status**: Cached with TTL 30 seconds

**Pricing Configuration**: Cached with TTL 24 hours

## Connection Pooling

**Provider Connections**: HTTP connection pooling for provider APIs

**Database Connections**: Connection pooling for generation log inserts

---

# SECTION 15 — Summary

## Key Deliverables

✓ **AI Orchestration Architecture**: Complete architecture for AI generation orchestration
✓ **Component Diagram**: Visual representation of AI components and their interactions
✓ **Sequence Diagram**: Detailed sequence diagram for synchronous generation flow
✓ **Failure Handling Rules**: Comprehensive failure handling for all scenarios
✓ **Cost Governance Rules**: Cost tracking, limits, and reporting mechanisms
✓ **Provider Strategy**: Provider-agnostic design with primary/fallback configuration

## Design Principles Adhered To

- ✓ AI assists teachers, human approval required
- ✓ Explainability preferred over autonomy
- ✓ Provider-agnostic architecture
- ✓ MVP-friendly implementation
- ✓ Cost-aware with governance
- ✓ Reliable failure handling
- ✓ Complete audit trail
- ✓ Synchronous default with async option for long-running tasks

## Next Steps

1. Implement AI Gateway service
2. Implement Prompt Builder with template loading
3. Implement Provider Adapters (OpenAI, Gemini)
4. Implement Response Processor with schema validation
5. Implement Generation Logger with database table
6. Implement cost tracking and limit checking
7. Configure RabbitMQ for async generation (future)
8. Implement health checks for providers
9. Create cost reporting dashboard
10. Add prompt versioning to template management
