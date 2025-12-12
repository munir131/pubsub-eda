# Plan: Production-Ready Event-Driven Architecture Refactor (Go)

## 1. Context & Goals
The goal is to refactor the existing Go application in the `go/` directory into a production-ready, event-driven architecture (EDA).
The current implementation is a simple CLI with hardcoded values in `main.go`.

**Goals:**
- **Standard Directory Structure:** Adopt a standard Go project layout (`cmd`, `internal`, `pkg`).
- **12-Factor App Compliance:** Configuration via environment variables (`.env` support), dependency management, logs as event streams (standard out).
- **Event-Driven Architecture:** Decouple the Pub/Sub implementation from the business logic. Define clear interfaces for `Publisher`, `Subscriber`, and `Handler` (Jobs).
- **Extensibility:** Make it easy to add new message handlers/jobs with minimal code changes.
- **Testing:** High test coverage with unit tests for every atomic function.

## 2. High-level Architecture Decisions

### Directory Structure
```
go/
├── cmd/
│   └── app/
│       └── main.go       # Entry point
├── internal/
│   ├── config/           # Configuration management
|   |    └── pubsub.go    # Configuration loading
│   ├── platform/         # Infrastructure adapters
│   │   └── pubsub/       # GCP Pub/Sub implementation
│   ├── event/            # Event definitions and Dispatcher
│   └── jobs/             # Business logic handlers
└── .env                  # Environment variables (gitignored)
```

### Components
1.  **Config:** Loads configuration from environment variables (using `joho/godotenv` and standard `os`).
2.  **Event Core:**
    *   `Message`: Standard struct for events.
    *   `Publisher` Interface: Contract for sending messages.
    *   `Subscriber` Interface: Contract for listening to messages.
    *   `Handler` Interface: Contract for processing messages (`Handle(ctx, msg)`).
3.  **Platform (GCP Pub/Sub):**
    *   Implements `Publisher` and `Subscriber` using `cloud.google.com/go/pubsub`.
4.  **Dispatcher:**
    *   Registry for mapping topics/events to Handlers.
    *   Manages the subscription loop and invokes the correct handler.
5.  **Jobs:**
    *   Refactor the current "print message" logic into a reusable Job/Handler.

## 3. Local Development & Docker Setup
- **Environment:** Use `.env` file for local secrets (Project ID, Topic names).
- **Docker:** Add a `Dockerfile` to the `go/` directory for containerization.
- **Dependencies:** Ensure `go.mod` is tidy.

## 4. Implementation Steps (Backend)

1.  [x] **Initialize Directory Structure & Config**:
    *   Create `go/cmd/app`, `go/internal/config`.
    *   Implement configuration loading.
    *   Create `.env.example`.

2.  [x] **Define Event Interfaces**:
    *   Create `go/internal/event/interfaces.go`.
    *   Define `Publisher`, `Subscriber`, `Handler`, and `Message` types.

3.  [x] **Implement GCP Pub/Sub Adapter**:
    *   Create `go/internal/platform/pubsub/client.go`.
    *   Implement `NewPublisher` and `NewSubscriber`.
    *   Implement `Publish` and `Subscribe` methods satisfying the interfaces.

4.  [x] **Implement Dispatcher/Router**:
    *   Create `go/internal/event/dispatcher.go`.
    *   Allow registering Handlers for specific topics.
    *   Implement logic to start listening and route incoming messages to handlers.

5.  **Implement Sample Job**:
    *   Create `go/internal/jobs/logger_job.go`.
    *   Implement the `Handler` interface to log the received message.

6.  **Refactor Main Entrypoint**:
    *   Update `go/cmd/app/main.go`.
    *   Wire up Config -> Pub/Sub Adapter -> Dispatcher -> Jobs.
    *   Use flags/args to determine mode (Publish vs Listen) via `flag` package or subcommands.

## 5. Implementation Steps (Tests)

1.  **Config Tests**:
    *   Test loading values from environment.

2.  **Dispatcher Tests**:
    *   Unit test the Dispatcher's routing logic using mock Publishers/Subscribers.

3.  **Job Tests**:
    *   Unit test the `LoggerJob` to ensure it processes messages correctly.

4.  **Pub/Sub Adapter Tests**:
    *   Add tests for the adapter (mocking the underlying client if possible, or skip if purely integration).

## 6. Verification Checklist
- [ ] Application builds successfully (`go build -o app ./cmd/app`).
- [ ] All unit tests pass (`go test ./...`).
- [ ] Configuration is loaded correctly from `.env`.
- [ ] Can publish a message using the refactored code.
- [ ] Can receive and process a message using the new Dispatcher and Job.
- [ ] Directory structure is clean and strictly follows the plan.
