// file: main.go
package main

import (
	"context"
	"fmt"
)


func main() {
	fmt.Println("=== Orchestrator Saga Example ===")

	// success run
	ctx := context.Background()
	if err := orchestrateSaga(ctx, "order-123"); err != nil {
		fmt.Println("Saga ended with error:", err)
	} else {
		fmt.Println("Saga successful for order-123")
	}

	fmt.Println()

	// failure run (inventory fails)
	if err := orchestrateSaga(ctx, "fail-inventory"); err != nil {
		fmt.Println("Saga ended with error (expected):", err)
	} else {
		fmt.Println("Saga successful for fail-inventory (unexpected)")
	}
}
