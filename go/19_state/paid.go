package state

type PaidState struct{}

func (p *PaidState) Pay(o *Order) string {
	return "already paid"
}

func (p *PaidState) Ship(o *Order) string {
	o.SetState(&ShippedState{})
	return "shipped"
}

func (p *PaidState) Finish(o *Order) string {
	return "cannot finish"
}