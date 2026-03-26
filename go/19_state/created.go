package state

type CreatedState struct{}

func (c *CreatedState) Pay(o *Order) string {
	o.SetState(&PaidState{})
	return "paid"
}

func (c *CreatedState) Ship(o *Order) string {
	return "cannot ship, not paid"
}

func (c *CreatedState) Finish(o *Order) string {
	return "cannot finish"
}