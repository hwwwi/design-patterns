package transport

// transport interface
type Transport interface {
	SetSpeed(speed int)
	GetSpeed() int
}

// walkTransport
type WalkTransport struct {
	Speed int
}

func (t *WalkTransport) SetSpeed(speed int) {
	t.Speed = speed
}

func (t *WalkTransport) GetSpeed() int {
	return t.Speed
}

// MotorcycleTransport
type MotorcycleTransport struct {
	Speed int
}

func (t *MotorcycleTransport) SetSpeed(speed int) {
	t.Speed = speed
}

func (t *MotorcycleTransport) GetSpeed() int {
	return t.Speed
}

// CarTransport
type CarTransport struct {
	Speed int
}

func (t *CarTransport) SetSpeed(speed int) {
	t.Speed = speed
}

func (t *CarTransport) GetSpeed() int {
	return t.Speed
}
