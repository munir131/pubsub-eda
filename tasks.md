# Tasks Backlog

Last updated: 2025-11-14 12:41 (IST)

Use this file to track work items. Keep tasks small, testable, and outcome-focused.

## Conventions
- Status: TODO | DOING | BLOCKED | REVIEW | DONE
- Each task should include acceptance criteria and, if useful, test notes.

## Now
- [TODO] Choose broker (Kafka/NATS/RabbitMQ/Redis Streams)
  - Acceptance: Broker runs locally via `make up`; topic(s) created.
- [TODO] Hello-world producer and two consumers
  - Acceptance: Publishing an event triggers both consumers; logs show outcomes.
- [TODO] Add retries and DLQ
  - Acceptance: Failing handler retries N times, then routes to DLQ with reason.

## Next
- [TODO] Add basic metrics (processed_count, error_count, retry_count)
- [TODO] Provide Makefile and scripts for local dev
- [TODO] Document end-to-end example with trace

## Later
- [TODO] Load testing script and results
- [TODO] Optional dashboards (Grafana/Tempo/Prom)

## Ideas/Parking Lot
- Schema validation (JSON Schema/Protobuf)
- Exactly-once demo vs at-least-once idempotency

## Changelog
- 2025-11-14: Initialized backlog
