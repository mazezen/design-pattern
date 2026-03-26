package state

type Order struct {
	state State
}

func NewOrder() *Order {
	return &Order{state: &CreatedState{}}
}

func (o *Order) SetState(s State) {
	o.state = s
}

func (o *Order) Pay() string {
	return o.state.Pay(o)
}

func (o *Order) Ship() string {
	return o.state.Ship(o)
}

func (o *Order) Finish() string {
	return o.state.Finish(o)
}