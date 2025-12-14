# Go Pub/Sub CLI (Event-Driven Architecture)

This is a production-ready, event-driven Go application for interacting with Google Cloud Pub/Sub.
It follows 12-factor app principles and uses a modular architecture.

## Directory Structure

```
go/
├── cmd/
│   └── app/          # Application entry point
├── internal/
│   ├── config/       # Configuration loading
│   ├── event/        # Event interfaces and Dispatcher
│   ├── jobs/         # Message handlers (Business Logic)
│   └── platform/     # Infrastructure adapters (GCP Pub/Sub)
├── Dockerfile        # Docker build definition
└── .env.example      # Example environment variables
```

## Prerequisites

- Go 1.25 or later
- Docker (optional)
- Google Cloud Project with Pub/Sub API enabled

## Setup

1.  **Clone the repository:**
    ```bash
    git clone <repo-url>
    cd pubsub-eda/go
    ```

2.  **Configure Environment:**
    Copy `.env.example` to `.env` and fill in your values.
    ```bash
    cp .env.example .env
    ```
    Update `.env`:
    ```env
    PROJECT_ID=your-gcp-project-id
    TOPIC_NAME=test-topic
    SUBSCRIPTION_NAME=test-sub
    GOOGLE_APPLICATION_CREDENTIALS=/path/to/key.json # If running locally without gcloud auth
    ```

3.  **Install Dependencies:**
    ```bash
    go mod download
    ```

## Usage

### Local Run

**Publish a message:**
```bash
go run ./cmd/app publish "Hello World"
```

**Listen for messages:**
```bash
go run ./cmd/app listen
```

### Docker Run

1.  **Build the image:**
    ```bash
    docker build -t pubsub-app .
    ```

2.  **Run (Publish):**
    ```bash
    docker run --env-file .env pubsub-app publish "Hello from Docker"
    ```

3.  **Run (Listen):**
    ```bash
    docker run --env-file .env pubsub-app listen
    ```

## Testing

Run all unit tests:
```bash
go test ./...
```