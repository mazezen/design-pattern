package observer

import "testing"

func TestObserverNotify(t *testing.T) {
	product := NewProduct("iPhone")

	u1 := NewCustomer("A")
	u2 := NewCustomer("B")

	product.Subscribe(u1)
	product.Subscribe(u2)

	product.SetStock(true)

	if len(u1.Messages) != 1 {
		t.Errorf("u1 should receive notification")
	}

	if len(u2.Messages) != 1 {
		t.Errorf("u2 should receive notification")
	}
}

func TestUnsubscribe(t *testing.T) {
	product := NewProduct("MacBook")

	u1 := NewCustomer("A")
	u2 := NewCustomer("B")

	product.Subscribe(u1)
	product.Subscribe(u2)

	product.Unsubscribe(u1)

	product.SetStock(true)

	if len(u1.Messages) != 0 {
		t.Errorf("u1 should not receive notification")
	}

	if len(u2.Messages) != 1 {
		t.Errorf("u2 should receive notification")
	}
}