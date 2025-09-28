# Saga Pattern: Choreography Example in Go

This project demonstrates the Saga pattern using the **Choreography** approach in Go. It simulates a distributed transaction across multiple services (Order, Payment, Inventory, Shipping) using an in-memory event bus for communication.

## Project Structure

- `main.go` - Entry point; starts all services and triggers example orders.
- `bus.go` - Simple in-memory pub/sub event bus.
- `order_service.go` - Handles order creation and triggers payment.
- `payment_service.go` - Handles payment reservation and compensation.
- `inventory_service.go` - Handles inventory reservation and compensation.
- `shipping_service.go` - Handles shipping and compensation.

## How It Works

1. **Order Service** receives an order and requests payment.
2. **Payment Service** reserves payment and notifies success/failure.
3. **Inventory Service** reserves inventory and notifies success/failure.
4. **Shipping Service** creates shipment or triggers compensations if previous steps fail.
5. Compensation events are published if any step fails, triggering rollback actions in other services.

## Running the Example

1. Make sure you have Go installed.
2. Run the example:

   ```sh
   go run .
   ```

3. You will see logs for three scenarios:
   - Successful order
   - Order where inventory reservation fails
   - Order where payment reservation fails

## Notes

- This is a simplified, in-memory demonstration. In real-world systems, services would be separate processes and communicate via message brokers.
- The choreography approach means services react to events and coordinate via the event bus, without a central orchestrator.

---
