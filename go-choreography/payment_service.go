package main

import (
	"context"
	"fmt"
	"time"
)

func PaymentService(ctx context.Context, bus *Bus) {
	ch := bus.Subscribe("PaymentRequested")
	compCh := bus.Subscribe("PaymentCompensate")
	for {
		select {
		case <-ctx.Done():
			return
		case evt := <-ch:
			orderID := evt.Payload["orderID"]
			fmt.Println("[Payment] trying to reserve payment for", orderID)
			time.Sleep(150 * time.Millisecond)
			if orderID == "fail-payment" {
				fmt.Println("[Payment] failed to reserve payment", orderID)
				bus.Publish(Event{Name: "PaymentFailed", Payload: map[string]string{"orderID": orderID}})
			} else {
				fmt.Println("[Payment] payment reserved for", orderID)
				bus.Publish(Event{Name: "PaymentReserved", Payload: map[string]string{"orderID": orderID}})
			}
		case evt := <-compCh:
			fmt.Println("[Payment] compensating payment for", evt.Payload["orderID"])
		}
	}
}
