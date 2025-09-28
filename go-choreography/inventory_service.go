package main

import (
	"context"
	"fmt"
	"time"
)

func InventoryService(ctx context.Context, bus *Bus) {
	payResCh := bus.Subscribe("PaymentReserved")
	payFailCh := bus.Subscribe("PaymentFailed")
	compCh := bus.Subscribe("InventoryCompensate")
	for {
		select {
		case <-ctx.Done():
			return
		case evt := <-payResCh:
			orderID := evt.Payload["orderID"]
			fmt.Println("[Inventory] Received PaymentReserved for", orderID, " — reserving inventory")
			time.Sleep(150 * time.Millisecond)
			if orderID == "fail-inventory" {
				fmt.Println("[Inventory] inventory reservation failed for", orderID)
				bus.Publish(Event{Name: "InventoryFailed", Payload: map[string]string{"orderID": orderID}})
			} else {
				fmt.Println("[Inventory] inventory reserved for", orderID)
				bus.Publish(Event{Name: "InventoryReserved", Payload: map[string]string{"orderID": orderID}})
			}
		case evt := <-payFailCh:
			fmt.Println("[Inventory] received PaymentFailed for", evt.Payload["orderID"], "- nothing to do")
		case evt := <-compCh:
			fmt.Println("[Inventory] compensating inventory for", evt.Payload["orderID"])
		}
	}
}
