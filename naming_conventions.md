# Naming Conventions

Last updated: 2025-11-14 12:41 (IST)

Consistent names improve readability, searchability, and tooling.

## General
- Use lowercase with hyphens for file/folder names: `event-producer`, `order-service`.
- Prefer descriptive over clever names.
- Avoid abbreviations unless common (e.g., `id`, `url`, `dlq`).

## Events
- Event type: `DomainAction` in PascalCase, e.g., `OrderCreated`, `PaymentProcessed`.
- Topic/stream names: kebab-case, plural where appropriate, e.g., `orders-created`, `payments-processed`.
- Event payload fields: snake_case or lowerCamel depending on language; keep consistent across a stack.

## Code (per language)

### Python
- Files/modules: `snake_case.py`
- Packages: `snake_case`
- Classes/Exceptions: `PascalCase`
- Functions/vars: `snake_case`
- Constants: `UPPER_SNAKE_CASE`

### Node/TypeScript
- Files: `kebab-case.ts` or `kebab-case.js`
- Classes/Types/Interfaces/Enums: `PascalCase`
- Functions/vars: `camelCase`
- Constants: `UPPER_SNAKE_CASE`

### Go
- Packages: short, all lowercase, no underscores
- Files: `snake_case.go` or meaningful lowercase
- Exported names: `PascalCase`; unexported: `camelCase`

## Directories
- Producers: `producers/` with subfolders per domain/component
- Consumers: `consumers/` with subfolders per domain/component
- Broker configs: `broker/`
- Scripts and tooling: `scripts/`
- Documentation: `docs/` and top-level `*.md`

## Tests
- Python: `tests/test_<module>.py`
- Node/TS: `__tests__/<name>.test.ts`
- Go: `<file>_test.go`

## Config & Env
- Example env file: `.env.example`
- Compose files: `docker-compose.yml`, variants `docker-compose.dev.yml`
- Makefile targets: `kebab-case` (e.g., `start-broker`, `run-tests`)

## Git Branches
- `feat/<short-description>`
- `fix/<short-description>`
- `chore/<short-description>`
- `docs/<short-description>`

## Misc
- Correlation/trace ids: `correlation_id`, `trace_id`
- Idempotency keys: `idempotency_key`
