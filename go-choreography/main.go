// file: choreography_saga.go
package main

import (
	"context"
	"fmt"
	"time"
)

// Very small in-memory pub/sub

func main() {
	fmt.Println("=== Choreography Saga Example ===")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	bus := NewBus()

	// start services
	go OrderService(ctx, bus)
	go PaymentService(ctx, bus)
	go InventoryService(ctx, bus)
	go ShippingService(ctx, bus)

	// Start a successful order
	bus.Publish(Event{Name: "OrderCreated", Payload: map[string]string{"orderID": "order-1"}})
	time.Sleep(1 * time.Second)
	fmt.Println()

	// Start an order where inventory will fail
	bus.Publish(Event{Name: "OrderCreated", Payload: map[string]string{"orderID": "fail-inventory"}})
	time.Sleep(1 * time.Second)
	fmt.Println()

	// Start an order where payment will fail
	bus.Publish(Event{Name: "OrderCreated", Payload: map[string]string{"orderID": "fail-payment"}})
	time.Sleep(1 * time.Second)
}
