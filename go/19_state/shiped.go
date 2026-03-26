package state

type ShippedState struct{}

func (s *ShippedState) Pay(o *Order) string {
	return "already paid"
}

func (s *ShippedState) Ship(o *Order) string {
	return "already shipped"
}

func (s *ShippedState) Finish(o *Order) string {
	o.SetState(&FinishedState{})
	return "finished"
}