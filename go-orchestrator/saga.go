package main

import (
	"context"
	"fmt"
)

func orchestrateSaga(ctx context.Context, orderID string) error {
	type step struct {
		name       string
		action     func(context.Context, string) error
		compensate func(context.Context, string) error
	}

	steps := []step{
		{"ReservePayment", reservePayment, compensatePayment},
		{"ReserveInventory", reserveInventory, compensateInventory},
		{"ShipOrder", shipOrder, compensateShipping},
	}

	completed := []step{}

	for _, s := range steps {
		fmt.Println("[Orchestrator] executing step:", s.name)
		err := s.action(ctx, orderID)
		if err != nil {
			fmt.Printf("[Orchestrator] step %s failed: %v\n", s.name, err)
			// run compensations in reverse order for completed steps
			fmt.Println("[Orchestrator] running compensating transactions...")
			for i := len(completed) - 1; i >= 0; i-- {
				cs := completed[i]
				if cerr := cs.compensate(ctx, orderID); cerr != nil {
					fmt.Printf("[Orchestrator] compensation %s failed: %v\n", cs.name, cerr)
				}
			}
			return fmt.Errorf("saga failed at step %s: %w", s.name, err)
		}
		completed = append(completed, s)
	}

	fmt.Println("[Orchestrator] saga completed successfully for order", orderID)
	return nil
}
