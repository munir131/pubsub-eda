# pubsub-eda — Project Brief

Last updated: 2025-11-14 11:39 (IST)

## Overview
A lightweight, event-driven architecture (EDA) playground using the publish/subscribe (pub/sub) pattern. The goal is to demonstrate how producers publish events to a broker and how multiple consumers react to those events independently, enabling loose coupling and horizontal scalability.

## Objectives
- Model core EDA concepts (producers, broker, consumers) in a minimal, easy-to-run project.
- Showcase asynchronous communication and eventual consistency patterns.
- Provide a foundation for experimenting with reliability (retries, DLQs), ordering, and scaling consumers.

## Architecture at a glance
Publisher(s) → Event Broker (topic/queue) → Subscriber(s)

Key properties:
- Decoupled producers and consumers
- Fan-out to multiple subscribers
- Asynchronous processing and backpressure handling

## Tech Stack (TBD)
- Language/Framework: TODO
- Broker (e.g., Kafka, NATS, RabbitMQ, Redis Streams): TODO
- Packaging/Infra: TODO (Docker Compose, K8s, or local-only)

## Local Development (TBD)
- Prerequisites: TODO
- Run locally: TODO
- Sample commands/events: TODO

## Repository Structure (high-level)
- producers/ — event publishers (examples, tests)
- consumers/ — event subscribers/handlers
- broker/ — local broker config or mocks for development
- scripts/ — helper scripts (run, seed, clean)
- docs/ — architecture notes and diagrams

## Next Steps
1. Lock the broker choice and minimal local setup (prefer Docker Compose for reproducibility).
2. Implement a hello-world producer and two independent consumers.
3. Add basic reliability: retry policy and a simple dead-letter strategy.
4. Provide Makefile/scripts for one-command bootstrap.
5. Document end-to-end flow with a small trace example.
