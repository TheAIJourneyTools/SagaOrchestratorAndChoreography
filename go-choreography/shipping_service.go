package main

import (
	"context"
	"fmt"
	"time"
)

func ShippingService(ctx context.Context, bus *Bus) {
	invResCh := bus.Subscribe("InventoryReserved")
	invFailCh := bus.Subscribe("InventoryFailed")
	compCh := bus.Subscribe("ShippingCompensate")
	for {
		select {
		case <-ctx.Done():
			return
		case evt := <-invResCh:
			orderID := evt.Payload["orderID"]
			fmt.Println("[Shipping] creating shipment for", orderID)
			time.Sleep(100 * time.Millisecond)
			fmt.Println("[Shipping] shipment created for", orderID)
			bus.Publish(Event{Name: "OrderCompleted", Payload: map[string]string{"orderID": orderID}})
		case evt := <-invFailCh:
			orderID := evt.Payload["orderID"]
			fmt.Println("[Shipping] inventory failed for", orderID, "triggering compensations")
			bus.Publish(Event{Name: "InventoryCompensate", Payload: map[string]string{"orderID": orderID}})
			bus.Publish(Event{Name: "PaymentCompensate", Payload: map[string]string{"orderID": orderID}})
		case evt := <-compCh:
			fmt.Println("[Shipping] compensating shipping for", evt.Payload["orderID"])
		}
	}
}
