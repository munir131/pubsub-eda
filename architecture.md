# Architecture

Last updated: 2025-11-14 12:40 (IST)

## Overview
Event-driven architecture (EDA) using publish/subscribe. Producers emit domain events to a broker; multiple consumers subscribe and process independently.

```
+-----------+        +------------------+        +----------------+
| Producer  |  --->  |  Event Broker    |  --->  | Consumer A     |
| (one/more)|  --->  | (topics/streams) |  --->  | Consumer B     |
+-----------+        +------------------+        +----------------+
                             |   |                         
                             v   v                         
                           DLQ  Metrics/Logs/Traces
```

## Components
- Producers: Services or scripts that publish events.
- Broker: Provides topics/streams, delivery semantics, retention.
- Consumers: Independent workers that react to events and perform side effects.
- DLQ: Stores failed messages after retry budgets are exhausted.
- Observability: Logs, metrics, traces for end-to-end visibility.

## Event Flow
1. Producer creates an immutable event (with id, type, timestamp, payload).
2. Event is published to a topic/stream.
3. The broker persists and delivers to all subscribed consumers.
4. Consumers process; on success, they acknowledge/commit; on failure, they retry then route to DLQ.

## Reliability & Semantics
- Delivery: At-least-once by default; idempotent consumers recommended.
- Ordering: Best-effort or partitioned ordering based on key.
- Retries: Exponential backoff; max attempts configurable per consumer.
- DLQ: Dedicated topic/stream; include error reason and original payload.

## Scalability
- Horizontal scale by increasing consumer instances/partitions.
- Backpressure handled via broker offsets/ack policies.

## Security (TBD)
- Local dev often unauthenticated; secure configs for CI/CD if needed.

## Observability
- Structured logs (event_id, type, latency_ms, outcome).
- Metrics: processed_count, error_count, retry_count, dlq_count, lag.
- Tracing: propagate correlation_id across producer/consumer boundaries.

## Open Decisions
- Broker choice and client libs
- Exactly-once vs at-least-once tradeoffs
- Schema validation strategy (e.g., JSON Schema, Protobuf)
