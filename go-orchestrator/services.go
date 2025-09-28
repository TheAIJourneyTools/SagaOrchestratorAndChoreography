package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func reservePayment(ctx context.Context, orderID string) error {
	fmt.Println("[Payment] reserving payment for order", orderID)
	time.Sleep(200 * time.Millisecond)
	// simulate success
	return nil
}
func compensatePayment(ctx context.Context, orderID string) error {
	fmt.Println("[Payment] compensating payment for order", orderID)
	time.Sleep(100 * time.Millisecond)
	return nil
}

func reserveInventory(ctx context.Context, orderID string) error {
	fmt.Println("[Inventory] reserving inventory for order", orderID)
	time.Sleep(200 * time.Millisecond)
	// simulate failure for demo when orderID == "fail-inventory"
	if orderID == "fail-inventory" {
		return errors.New("inventory unavailable")
	}
	return nil
}
func compensateInventory(ctx context.Context, orderID string) error {
	fmt.Println("[Inventory] compensating inventory for order", orderID)
	time.Sleep(100 * time.Millisecond)
	return nil
}

func shipOrder(ctx context.Context, orderID string) error {
	fmt.Println("[Shipping] creating shipment for order", orderID)
	time.Sleep(200 * time.Millisecond)
	return nil
}
func compensateShipping(ctx context.Context, orderID string) error {
	fmt.Println("[Shipping] compensating shipment for order", orderID)
	time.Sleep(100 * time.Millisecond)
	return nil
}
