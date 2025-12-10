# Go Pub/Sub CLI

This is a command-line application for interacting with Google Cloud Pub/Sub.

## Prerequisites

- Go 1.16 or later
- Google Cloud SDK (`gcloud`)

## Setup

1.  **Authenticate with Google Cloud:**

    ```bash
    gcloud auth application-default login
    ```

2.  **Set your Google Cloud Project ID:**

    Open `main.go` and replace `"your-gcp-project-id"` with your actual project ID.

3.  **Create a Pub/Sub topic and subscription:**

    You can do this via the `gcloud` CLI or the Google Cloud Console. The topic and subscription names are hardcoded in `main.go` as `test` and `test-sub`.

    ```bash
    gcloud pubsub topics create test
    gcloud pubsub subscriptions create test-sub --topic=test
    ```

## Usage

### Publish a Message

To publish a message to the `test` topic:

```bash
go run main.go publish "your message here"
```

### Listen for Messages

To listen for messages from the `test-sub` subscription:

```bash
go run main.go listen
```
