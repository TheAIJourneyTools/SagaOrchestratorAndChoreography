# Saga Orchestrator and Choreography Example (Go)

This project demonstrates the Saga pattern using the **Orchestrator** approach in Go.  
It simulates a distributed transaction across payment, inventory, and shipping services, with compensation logic for failures.

## Structure

- `orchestrator/main.go` - Entry point, runs the saga.
- `orchestrator/saga.go` - Saga orchestration logic.
- `orchestrator/services.go` - Simulated service actions and compensations.

## How It Works

The orchestrator runs three steps in order:

1. Reserve Payment
2. Reserve Inventory
3. Ship Order

If any step fails, previously completed steps are compensated in reverse order.

## Running the Example

Open a terminal and navigate to the `orchestrator` directory:

```sh
cd orchestrator
go run .
```

You should see output for both a successful and a failed saga run.

## Notes

- All files must be in the same directory and use `package main`.
- No external dependencies are required.
