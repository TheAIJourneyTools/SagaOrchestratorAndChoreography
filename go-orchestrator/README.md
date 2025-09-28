# Saga Orchestrator Pattern Example (Go)

This example demonstrates the **Saga Orchestrator** pattern for managing distributed transactions in microservices using Go.

## What is the Saga Pattern?

The Saga pattern is a way to manage data consistency across multiple services in distributed systems without using distributed transactions. It breaks a transaction into a series of smaller, isolated steps (local transactions), each with a corresponding compensating action to undo its effect if needed.

There are two main approaches:

- **Orchestrator-based Saga**: A central coordinator (orchestrator) tells each service what to do and when to execute compensating actions.
- **Choreography-based Saga**: Each service produces and listens to events, coordinating the saga without a central controller.

This example implements the **Orchestrator-based Saga**.

## How it Works

- The orchestrator executes a series of steps (e.g., reserve payment, reserve inventory, ship order).
- If a step fails, the orchestrator triggers compensating actions for all previously completed steps in reverse order.
- Each step and its compensation are implemented as functions.

### Example Flow

1. **Reserve Payment**
2. **Reserve Inventory**
3. **Ship Order**

If any step fails (e.g., inventory is unavailable), the orchestrator compensates all completed steps to maintain consistency.

## Files

- `main.go`: Entry point, runs saga scenarios.
- `saga.go`: Contains the orchestrator logic.
- `services.go`: Simulates service actions and compensations.

## Running the Example

1. Run `go run .` in this directory.
2. Observe the output for both a successful saga and a failed saga (with compensation).

## Output Example

```
=== Orchestrator Saga Example ===
[Orchestrator] executing step: ReservePayment
[Payment] reserving payment for order order-123
[Orchestrator] executing step: ReserveInventory
[Inventory] reserving inventory for order order-123
[Orchestrator] executing step: ShipOrder
[Shipping] creating shipment for order order-123
[Orchestrator] saga completed successfully for order order-123
Saga successful for order-123

[Orchestrator] executing step: ReservePayment
[Payment] reserving payment for order fail-inventory
[Orchestrator] executing step: ReserveInventory
[Inventory] reserving inventory for order fail-inventory
[Orchestrator] step ReserveInventory failed: inventory unavailable
[Orchestrator] running compensating transactions...
[Payment] compensating payment for order fail-inventory
Saga ended with error (expected): saga failed at step ReserveInventory: inventory unavailable
```

## When to Use

- When you need to maintain data consistency across multiple microservices.
- When distributed transactions (2PC) are not feasible.

## References

- [Microservices Patterns: With examples in Java (Chris Richardson)](https://microservices.io/patterns/data/saga.html)
- [Saga Pattern - Microsoft Docs](https://learn.microsoft.com/en-us/azure/architecture/reference-architectures/saga/saga)

---

## Further Reading

For a detailed explanation of the Saga pattern, including orchestrator and choreography variants, see the official documentation:

- [Saga Pattern on microservices.io](https://microservices.io/patterns/data/saga.html)

> The Saga pattern manages failures in distributed transactions by breaking them into a sequence of local transactions. Each local transaction updates data within a single service and publishes an event or triggers the next transaction. If a step fails, the saga executes compensating transactions to undo the impact of previous steps. This pattern is essential for maintaining data consistency in microservices architectures.
