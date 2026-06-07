# NUSA AI Runtime

Python + LangGraph AI orchestration service.

## Technology Stack

- Python 3.11+
- FastAPI
- LangGraph
- OpenAI API / Anthropic API

## Responsibilities

- LangGraph Execution
- Prompt Execution
- LLM Integration
- Structured Output Generation
- AI Validation

## Communication

Backend API → AI Runtime via HTTP REST

## Installation

```bash
pip install -r requirements.txt
```

## Running

```bash
uvicorn app.main:app --reload --host 0.0.0.0 --port 8000
```

## Development

This is a skeleton only. No LangGraph workflows, nodes, agents, or prompts are implemented yet.
