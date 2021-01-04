package delivery

import "github.com/hwwwi/design-patterns/factory/pkg/transport"

// Delivery interface
type Delivery interface {
	CreateTransport() transport.Transport
}

// WalkDelivery
type WalkDelivery struct {}

func (d *WalkDelivery) CreateTransport() transport.Transport {
	return &transport.WalkTransport{Speed: 5}
}

// CarDelivery
type CarDelivery struct {}

func (d *CarDelivery) CreateTransport() transport.Transport {
	return &transport.CarTransport{Speed: 60}
}

// MotorcycleDelivery
type MotorcycleDelivery struct {}

func (d *MotorcycleDelivery) CreateTransport() transport.Transport {
	return &transport.WalkTransport{Speed: 30}
}
