# Saga Pattern Examples in Go

This repository demonstrates the Saga pattern for managing distributed transactions in microservices, using both the **Orchestrator** and **Choreography** approaches.

## What is the Saga Pattern?

The Saga pattern manages data consistency across multiple services by breaking a transaction into a series of smaller, isolated steps (local transactions), each with a corresponding compensating action to undo its effect if needed.

There are two main approaches:

- **Orchestrator-based Saga**: A central coordinator (orchestrator) tells each service what to do and when to execute compensating actions.
- **Choreography-based Saga**: Each service produces and listens to events, coordinating the saga without a central controller.

## Project Structure

- [`orchestrator/`](./orchestrator/) — Implements the Orchestrator-based Saga pattern.
- [`go-choreography/`](./go-choreography/) — Implements the Choreography-based Saga pattern.

Each subdirectory contains its own README with details and instructions.

## How to Run

1. Choose the pattern you want to explore:

   - For Orchestrator:
     ```
     cd orchestrator
     go run .
     ```
   - For Choreography:
     ```
     cd go-choreography
     go run .
     ```

2. See the respective subdirectory README for more information and sample output.

## References

- [Saga Pattern on microservices.io](https://microservices.io/patterns/data/saga.html)
- [Saga Pattern - Microsoft Docs](https://learn.microsoft.com/en-us/azure/architecture/reference-architectures/saga/saga)

---
