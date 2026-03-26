package observer

type Subject interface {
	Subscribe(o Observer)
	Unsubscribe(o Observer)
	Notify()
}