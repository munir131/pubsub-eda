# Project Context

Last updated: 2025-11-14 12:40 (IST)

## Business Context
- Learning-focused repo to explore event-driven architecture (EDA) with pub/sub.
- Emphasis on simplicity, reproducibility, and teaching core concepts.

## Domain Context
- Events represent domain facts (e.g., OrderCreated, PaymentProcessed).
- Consumers react independently; no direct coupling to publishers.

## Constraints & Assumptions
- Local-first development; internet access may be limited for CI.
- Minimal dependencies; prefer containers for broker.
- Eventually consistent flows are acceptable.

## Risks
- Over-scoping into production-grade features.
- Broker-specific lock-in if abstractions are weak.

## Glossary
- Event: Immutable record of something that happened.
- Topic/Stream/Queue: Logical channel carrying events.
- Producer/Consumer: Writers/readers of events.
- DLQ: Dead-letter queue for failed messages.
