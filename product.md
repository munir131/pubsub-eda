# Product Brief

Last updated: 2025-11-14 12:40 (IST)

## 1. Vision
A hands-on, minimal product to demonstrate event-driven architecture (EDA) using pub/sub, enabling developers to learn, prototype, and benchmark asynchronous workflows.

## 2. Target Users
- Backend engineers exploring EDA patterns
- Platform teams prototyping messaging topologies
- Educators and students learning distributed systems

## 3. Problem & Value Proposition
- Problem: Hard to learn EDA end-to-end without heavy setup.
- Value: A small, reproducible playground showing producers → broker → consumers with reliability patterns.

## 4. Key Use Cases
- Publish an event and fan out to N consumers
- Introduce retries and dead-letter handling
- Measure latency and throughput under load

## 5. Success Metrics
- Time-to-first-event < 10 minutes
- Sample flows and docs that a new dev can follow without help
- Optional: scriptable load test results included in docs

## 6. Scope
- In: local development, minimal scripts, clear docs
- Out (for now): multi-tenant auth, complex UIs, cross-region HA

## 7. Non-Goals
- Competing with full-featured platforms
- Production-hardening beyond basics

## 8. Release Plan
- v0.1: Hello-world producer + 2 consumers + local broker
- v0.2: Retries + DLQ + minimal observability
- v0.3: Throughput test, basic dashboards
