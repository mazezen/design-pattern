package state

type State interface {
	Pay(o *Order) string
	Ship(o *Order) string
	Finish(o *Order) string
}