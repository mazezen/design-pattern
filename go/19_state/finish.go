package state

type FinishedState struct{}

func (f *FinishedState) Pay(o *Order) string {
	return "order done"
}

func (f *FinishedState) Ship(o *Order) string {
	return "order done"
}

func (f *FinishedState) Finish(o *Order) string {
	return "already finished"
}