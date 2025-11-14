# Tech

Last updated: 2025-11-14 12:40 (IST)

## Stack (proposed/TBD)
- Language: TBD (Python/Node/Go)
- Broker: TBD (Kafka/NATS/RabbitMQ/Redis Streams)
- Containerization: Docker + Docker Compose
- Observability: Structured logs; optional OpenTelemetry for traces

## Versioning & Compatibility
- Pin dependencies (requirements.txt/package.json/go.mod)
- Prefer LTS versions for language/runtime

## Local Development
- Prereqs: Docker, Make, language runtime
- Make targets: `make up`, `make down`, `make logs`, `make test`
- Environment: `.env` for local overrides (never commit secrets)

## Configuration
- Use env vars for broker connection and topics
- Provide sane defaults for local runs

## CI (Optional)
- Lint, format, and test on pull requests
- Build images and run basic integration tests with ephemeral broker

## Security & Secrets
- No secrets in repo; use `.env.local` (ignored) or secret manager
- Rotate credentials for any shared demo infra

## Testing
- Unit tests for producers/consumers logic
- Integration tests with local broker via Docker Compose

## Observability
- Log fields: `event_id`, `event_type`, `correlation_id`, `latency_ms`, `outcome`
- Metrics: counts, lag, retries, DLQ
- Tracing: propagate `trace_id`/`correlation_id`
