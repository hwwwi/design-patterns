package main

import (
	"fmt"
	"github.com/hwwwi/design-patterns/factory/pkg/delivery"
	"github.com/hwwwi/design-patterns/factory/pkg/transport"
)

func main() {
	if t, err := createTransport(&delivery.CarDelivery{}); err == nil && t != nil{
		fmt.Printf("car: %d\n", t.GetSpeed())
	}

	if t, err := createTransport(&delivery.WalkDelivery{}); err == nil && t != nil{
		fmt.Printf("walk: %d\n", t.GetSpeed())
	}
}

// createTransport is the client code
// this function will not break even you create new delivery type
func createTransport(del delivery.Delivery) (transport.Transport, error) {
	if del == nil {
		return nil, fmt.Errorf("failed to create transport")
	}
	return del.CreateTransport(), nil
}
