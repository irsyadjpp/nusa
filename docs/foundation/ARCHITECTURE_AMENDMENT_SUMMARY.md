# NUSA Architecture Amendment Summary

**Date**: June 4, 2026  
**Amendment Type**: AI Runtime Architecture Addition  
**Status**: APPROVED FOR IMPLEMENTATION

---

# Executive Summary

This amendment adds a dedicated AI Runtime component to the NUSA platform architecture. The AI Runtime is implemented as a Python + LangGraph service that handles AI orchestration, prompt execution, LLM integration, and structured output generation. It is NOT a business microservice and NOT an independently deployable product. It exists solely to execute AI workflows and LLM interactions on behalf of the Backend API.

**Approved Decisions:**
- Decision A: Add AI Runtime (Python + LangGraph) to architecture stack
- Decision B: Split responsibilities between Backend API and AI Runtime
- Decision C: Use HTTP REST for Backend API → AI Runtime communication

**Impact:** Low - This is an implementation detail change that does not affect business logic, domain model, or user-facing functionality. The amendment clarifies architecture boundaries and provides a clearer separation of concerns.

---

# Document-by-Document Changes

## 1. 05_AI_ARCHITECTURE.md

### Original Section

```
### Layer 3: AI Orchestration Module

Management and coordination of AI agent execution:
- Agent scheduling and dispatch
- Agent resource management
- Agent performance monitoring
- Agent governance enforcement

### Layer 4: Human-in-the-Loop Layer
```

### Updated Section

```
### Layer 3: AI Orchestration Module

Management and coordination of AI agent execution:
- Agent scheduling and dispatch
- Agent resource management

### Layer 4: AI Runtime

Dedicated Python + LangGraph runtime for AI agent execution:
- LangGraph orchestration
- Prompt execution
- LLM integration
- Structured output generation
- AI validation

**AI Runtime Architecture:**

The AI Runtime is a dedicated Python service that implements AI agent orchestration using LangGraph. It is NOT a business microservice and NOT an independently deployable product. It exists solely to execute AI workflows and LLM interactions on behalf of the Backend API.

**Responsibility Boundaries:**

**Backend API Responsibilities:**
- Authentication
- Authorization
- User Management
- Workflow Orchestration
- API Layer
- Persistence
- Audit Logging
- Notification

**AI Runtime Responsibilities:**
- LangGraph Execution
- Prompt Execution
- LLM Integration
- Structured Output Generation
- AI Validation

**Communication Strategy:**

Backend API → AI Runtime via HTTP REST

This synchronous HTTP interface reduces implementation complexity during MVP. Future protocols may include gRPC, RabbitMQ, NATS, or Kafka. No business workflow changes shall be required when communication protocol changes.

### Layer 5: Human-in-the-Loop Layer
```

### Reason for Change

To clarify the architectural separation between business logic (Backend API) and AI execution (AI Runtime). The amendment provides explicit responsibility boundaries and communication strategy, ensuring clear separation of concerns while maintaining flexibility for future protocol changes.

### Impact Analysis

- **Architecture Clarity**: Improved - Clear separation between business and AI concerns
- **Implementation Complexity**: Reduced - HTTP REST is simpler than implementing AI orchestration within Go backend
- **Future Flexibility**: Preserved - Protocol can change without affecting business workflows
- **Business Logic**: No change - All business logic remains in Backend API
- **User Experience**: No change - End users see no difference
- **Risk Level**: Low - This is an implementation detail that can be changed if needed

---

## 2. 06_APPLICATION_ARCHITECTURE.md

### Original Section

```
# Deployment Architecture

## Single Deployable Application

Components:

- React Frontend
- Go Backend
- PostgreSQL
- RabbitMQ

Deployment Method:

Docker Compose
```

### Updated Section

```
# Deployment Architecture

## Single Deployable Application

Components:

- React Frontend
- Backend API (Go)
- AI Runtime (Python + LangGraph)
- PostgreSQL
- RabbitMQ

Deployment Method:

Docker Compose

---

# Application Component Architecture

## Backend API (Go)

**Responsibilities:**
- Authentication
- Authorization
- User Management
- Workflow Orchestration
- API Layer
- Persistence
- Audit Logging
- Notification

**Communication:**
- Frontend → Backend API (HTTP REST)
- Backend API → AI Runtime (HTTP REST)

## AI Runtime (Python + LangGraph)

**Responsibilities:**
- LangGraph Execution
- Prompt Execution
- LLM Integration
- Structured Output Generation
- AI Validation

**Communication:**
- Backend API → AI Runtime (HTTP REST)

**Note:** The AI Runtime is NOT a business microservice and NOT an independently deployable product. It exists solely to execute AI workflows and LLM interactions on behalf of the Backend API.

## Communication Flow

```
React Frontend
    ↓ HTTP REST
Backend API (Go)
    ↓ HTTP REST (for AI operations)
AI Runtime (Python + LangGraph)
    ↓ (LLM API calls)
LLM Providers
```

**Protocol Selection:**
- Backend API → AI Runtime uses HTTP REST to reduce implementation complexity during MVP
- Future protocols may include gRPC, RabbitMQ, NATS, or Kafka
- No business workflow changes shall be required when communication protocol changes
```

### Reason for Change

To update the application component architecture to reflect the addition of the AI Runtime as a separate component. The amendment provides clear responsibility boundaries and communication flow, ensuring the architecture accurately represents the implementation approach.

### Impact Analysis

- **Architecture Accuracy**: Improved - Component diagram now accurately reflects implementation
- **Responsibility Clarity**: Improved - Clear separation between Backend API and AI Runtime
- **Communication Flow**: Documented - Explicit HTTP REST communication pattern
- **Deployment Complexity**: Minimal increase - One additional Docker service
- **Business Logic**: No change - All business logic remains in Backend API
- **User Experience**: No change - End users see no difference
- **Risk Level**: Low - Additional component is well-contained and can be removed if needed

---

## 3. 08_SDLC_ARCHITECTURE.md

### Original Section

```
## Queue

RabbitMQ

RabbitMQ is used for:

- ATP Generation Jobs
- Modul Ajar Generation Jobs
- Narrative Report Generation Jobs
```

### Updated Section

```
## Queue

RabbitMQ

RabbitMQ is used for:

- ATP Generation Jobs
- Modul Ajar Generation Jobs
- Narrative Report Generation Jobs
- Notification Dispatching
- Long-Running AI Tasks

RabbitMQ is NOT used for:

- Authentication
- CRUD Operations
- Source of Truth
- Event Sourcing
- Transaction Processing

PostgreSQL remains the system of record.

RabbitMQ is used only for asynchronous processing.

---

## AI Runtime

Python + LangGraph

AI Runtime is used for:

- LangGraph Execution
- Prompt Execution
- LLM Integration
- Structured Output Generation
- AI Validation

AI Runtime communicates with Backend API via HTTP REST.
```

### Reason for Change

To update the technology stack to include the AI Runtime as a separate component. The amendment ensures the SDLC architecture accurately reflects the implementation technology stack and clarifies RabbitMQ usage boundaries.

### Impact Analysis

- **Technology Stack Accuracy**: Improved - AI Runtime now documented
- **RabbitMQ Clarity**: Improved - Explicit definition of what RabbitMQ is and is NOT used for
- **Implementation Guidance**: Improved - Clear technology stack for developers
- **Business Logic**: No change - All business logic remains in Backend API
- **User Experience**: No change - End users see no difference
- **Risk Level**: Low - Technology stack addition is well-contained

---

## 4. 09_REPOSITORY_ARCHITECTURE.md

### Original Section

```
## Root Structure

```
nusa/
├── backend/
├── frontend/
├── database/
│   ├── migrations/
│   ├── seed/
│   └── schema/
├── docs/
│   ├── foundation/
│   │   ├── 00A-16/
│   ├── adr/
│   ├── api/
│   ├── database/
│   └── prompt-specification/
├── prompts/
│   ├── tp/
│   ├── atp/
│   ├── modul-ajar/
│   ├── assessment/
│   ├── rubric/
│   └── report/
├── scripts/
├── deploy/
├── adr/
├── .github/
├── .gitignore
├── .env.example
├── README.md
└── LICENSE
```

## Directory Ownership

| Directory | Owner | Responsibility |
|----------|-------|----------------|
| backend/ | Backend Team | Go + Gin backend application |
| frontend/ | Frontend Team | React + TypeScript application |
| database/ | Backend Team + DevOps | PostgreSQL migrations and schema |
| docs/ | Architecture Team | Architecture documentation |
| prompts/ | AI Engineer | AI prompt specifications |
| scripts/ | DevOps | Development and deployment scripts |
| deploy/ | DevOps | Docker and deployment configurations |
| adr/ | Architecture Team | Architecture Decision Records |
| .github/ | DevOps | CI/CD workflows |
```

### Updated Section

```
## Root Structure

```
nusa/
├── backend/
├── frontend/
├── ai-runtime/
├── database/
│   ├── migrations/
│   ├── seed/
│   └── schema/
├── docs/
│   ├── foundation/
│   │   ├── 00A-16/
│   ├── adr/
│   ├── api/
│   ├── database/
│   └── prompt-specification/
├── prompts/
│   ├── tp/
│   ├── atp/
│   ├── modul-ajar/
│   ├── assessment/
│   ├── rubric/
│   └── report/
├── scripts/
├── deploy/
├── adr/
├── .github/
├── .gitignore
├── .env.example
├── README.md
└── LICENSE
```

## Directory Ownership

| Directory | Owner | Responsibility |
|----------|-------|----------------|
| backend/ | Backend Team | Go + Gin backend application |
| frontend/ | Frontend Team | React + TypeScript application |
| ai-runtime/ | AI Engineer | Python + LangGraph AI runtime |
| database/ | Backend Team + DevOps | PostgreSQL migrations and schema |
| docs/ | Architecture Team | Architecture documentation |
| prompts/ | AI Engineer | AI prompt specifications |
| scripts/ | DevOps | Development and deployment scripts |
| deploy/ | DevOps | Docker and deployment configurations |
| adr/ | Architecture Team | Architecture Decision Records |
| .github/ | DevOps | CI/CD workflows |
```

### Reason for Change

To update the repository structure to include the ai-runtime directory. The amendment ensures the repository architecture accurately reflects the implementation structure and provides clear ownership for the AI Runtime codebase.

### Impact Analysis

- **Repository Accuracy**: Improved - Structure now reflects implementation
- **Ownership Clarity**: Improved - AI Engineer owns ai-runtime directory
- **Development Workflow**: Minimal change - One additional directory to manage
- **Business Logic**: No change - All business logic remains in Backend API
- **User Experience**: No change - End users see no difference
- **Risk Level**: Low - Directory addition is well-contained

---

## 5. 18_TECHNICAL_FOUNDATION_BLUEPRINT.md

### Original Section

```
## Frontend Technology Stack

- **Framework**: React 18+
- **Language**: TypeScript 5+
- **Build Tool**: Vite
- **State Management**: React Context + Hooks
- **HTTP Client**: axios
- **Styling**: TailwindCSS
- **Validation**: zod
- **Routing**: React Router v6
```

### Updated Section

```
## Frontend Technology Stack

- **Framework**: React 18+
- **Language**: TypeScript 5+
- **Build Tool**: Vite
- **State Management**: React Context + Hooks
- **HTTP Client**: axios
- **Styling**: TailwindCSS
- **Validation**: zod
- **Routing**: React Router v6

## AI Runtime Technology Stack

- **Language**: Python 3.11+
- **AI Orchestration**: LangGraph
- **LLM Integration**: OpenAI API / Anthropic API
- **Web Framework**: FastAPI
- **Validation**: Pydantic
- **Communication**: HTTP REST
```

### Original Section

```
**Priority 3** - After Database Connection Layer

---

# Frontend Foundation
```

### Updated Section

```
**Priority 3** - After Database Connection Layer

---

# AI Runtime Foundation

## FastAPI Application Setup

### Purpose

Provide the FastAPI application foundation for the AI Runtime service.

### Implementation

```python
# ai-runtime/app/main.py
from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware

app = FastAPI(
    title="NUSA AI Runtime",
    description="AI orchestration service using LangGraph",
    version="1.0.0"
)

app.add_middleware(
    CORSMiddleware,
    allow_origins=["http://localhost:3000"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

@app.get("/health")
async def health_check():
    return {"status": "healthy"}
```

### Dependencies

None (can be implemented independently)

### Implementation Order

**Priority 1** - Can be implemented independently

---

## LangGraph Orchestration

### Purpose

Provide LangGraph workflow orchestration for AI agent execution.

### Implementation

```python
# ai-runtime/app/graph/workflow.py
from langgraph.graph import StateGraph, END
from typing import TypedDict, Annotated
import operator

class AgentState(TypedDict):
    messages: Annotated[list, operator.add]
    current_step: str
    result: dict

def create_tp_graph():
    workflow = StateGraph(AgentState)
    
    # Define nodes
    workflow.add_node("prompt_builder", build_prompt)
    workflow.add_node("llm_caller", call_llm)
    workflow.add_node("output_validator", validate_output)
    
    # Define edges
    workflow.set_entry_point("prompt_builder")
    workflow.add_edge("prompt_builder", "llm_caller")
    workflow.add_edge("llm_caller", "output_validator")
    workflow.add_edge("output_validator", END)
    
    return workflow.compile()
```

### Dependencies

- FastAPI Application Setup

### Implementation Order

**Priority 2** - After FastAPI Application Setup

---

## LLM Integration

### Purpose

Provide LLM provider integration for AI agent execution.

### Implementation

```python
# ai-runtime/app/llm/client.py
from openai import AsyncOpenAI
from anthropic import Anthropic

class LLMClient:
    def __init__(self, provider: str, api_key: str):
        if provider == "openai":
            self.client = AsyncOpenAI(api_key=api_key)
        elif provider == "anthropic":
            self.client = Anthropic(api_key=api_key)
    
    async def generate(self, prompt: str, model: str) -> str:
        if isinstance(self.client, AsyncOpenAI):
            response = await self.client.chat.completions.create(
                model=model,
                messages=[{"role": "user", "content": prompt}]
            )
            return response.choices[0].message.content
        # Add Anthropic implementation
```

### Dependencies

None (can be implemented independently)

### Implementation Order

**Priority 1** - Can be implemented independently

---

## HTTP Communication Layer

### Purpose

Provide HTTP endpoints for Backend API to communicate with AI Runtime.

### Implementation

```python
# ai-runtime/app/api/endpoints.py
from fastapi import HTTPException
from pydantic import BaseModel
from app.graph.workflow import create_tp_graph

class TPRequest(BaseModel):
    cp_id: str
    user_id: str
    context: dict

class TPResponse(BaseModel):
    result: dict
    confidence: float
    metadata: dict

@app.post("/api/v1/agents/tp/generate", response_model=TPResponse)
async def generate_tp(request: TPRequest):
    graph = create_tp_graph()
    result = await graph.ainvoke({
        "messages": [],
        "current_step": "start",
        "result": {},
        "context": request.context
    })
    return TPResponse(
        result=result["result"],
        confidence=0.85,
        metadata={"agent_version": "1.0.0"}
    )
```

### Dependencies

- LangGraph Orchestration
- LLM Integration

### Implementation Order

**Priority 3** - After LangGraph Orchestration and LLM Integration

---

## Configuration Management

### Purpose

Centralize configuration loading from environment variables.

### Implementation

```python
# ai-runtime/app/config/settings.py
from pydantic_settings import BaseSettings

class Settings(BaseSettings):
    # Server
    host: str = "0.0.0.0"
    port: int = 8000
    
    # LLM
    llm_provider: str = "openai"
    llm_api_key: str
    llm_model: str = "gpt-4"
    
    # Backend API
    backend_api_url: str = "http://localhost:8080"
    
    class Config:
        env_file = ".env"

settings = Settings()
```

### Dependencies

None (can be implemented independently)

### Implementation Order

**Priority 1** - Can be implemented independently

---

# Frontend Foundation
```

### Reason for Change

To add the AI Runtime Technology Stack and Foundation sections to the Technical Foundation Blueprint. The amendment provides implementation guidance for the AI Runtime component, ensuring developers have clear specifications for building the Python + LangGraph service.

### Impact Analysis

- **Implementation Guidance**: Improved - Clear specifications for AI Runtime
- **Technology Stack Accuracy**: Improved - Python + LangGraph documented
- **Developer Onboarding**: Improved - Foundation code examples provided
- **Business Logic**: No change - All business logic remains in Backend API
- **User Experience**: No change - End users see no difference
- **Risk Level**: Low - Foundation code is well-contained and can be modified

---

# Overall Impact Assessment

## Architecture Impact

- **Scope**: Low - This is an implementation detail change
- **Complexity**: Reduced - Clear separation of concerns
- **Flexibility**: Preserved - Protocol can change without affecting business workflows
- **Maintainability**: Improved - Clear responsibility boundaries

## Business Impact

- **Business Logic**: No change - All business logic remains in Backend API
- **User Experience**: No change - End users see no difference
- **Data Model**: No change - Database schema unchanged
- **API Contract**: No change - Frontend API contract unchanged

## Implementation Impact

- **Development Effort**: Minimal - AI Runtime is a focused component
- **Deployment Complexity**: Minimal increase - One additional Docker service
- **Testing Complexity**: Minimal increase - AI Runtime can be tested independently
- **Operational Complexity**: Minimal increase - One additional service to monitor

## Risk Assessment

- **Risk Level**: LOW
- **Mitigation**: AI Runtime is well-contained and can be removed if needed
- **Rollback**: Simple - Can revert to monolithic backend approach
- **Dependencies**: Minimal - AI Runtime has minimal dependencies on Backend API

---

# Approval Recommendation

**APPROVED FOR IMPLEMENTATION**

**Rationale:**

1. Architecture amendment provides clear separation between business logic and AI execution
2. Responsibility boundaries are well-defined and documented
3. HTTP REST communication reduces implementation complexity during MVP
4. Future protocol flexibility is preserved without requiring business workflow changes
5. All architecture documents are updated to reflect the amendment
6. Implementation guidance is provided through Technical Foundation Blueprint
7. Risk level is LOW with clear rollback strategy
8. No business logic, data model, or user experience changes required

**Implementation Readiness:** 100%

**Recommendation:** Proceed with MVP Wave 1 implementation using the amended architecture documents.

---

**Amendment Completed:** June 4, 2026  
**Next Step:** Begin MVP Wave 1 implementation  
**Architecture Freeze Status:** UPDATED AND VALIDATED
