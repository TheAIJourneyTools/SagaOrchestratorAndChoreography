package main

import (
	"context"
	"fmt"
)

func OrderService(ctx context.Context, bus *Bus) {
	ch := bus.Subscribe("OrderCreated")
	for {
		select {
		case <-ctx.Done():
			return
		case evt := <-ch:
			orderID := evt.Payload["orderID"]
			fmt.Println("[Order] Received OrderCreated for", orderID)
			bus.Publish(Event{Name: "PaymentRequested", Payload: map[string]string{"orderID": orderID}})
		}
	}
}
