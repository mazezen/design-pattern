package state

import "testing"

func TestOrderFlow(t *testing.T) {
	o := NewOrder()

	if o.Pay() != "paid" {
		t.Error("pay failed")
	}

	if o.Ship() != "shipped" {
		t.Error("ship failed")
	}

	if o.Finish() != "finished" {
		t.Error("finish failed")
	}
}

func TestInvalidFlow(t *testing.T) {
	o := NewOrder()

	if o.Ship() != "cannot ship, not paid" {
		t.Error("should not ship")
	}
}

func TestRepeatAction(t *testing.T) {
	o := NewOrder()

	o.Pay()

	if o.Pay() != "already paid" {
		t.Error("repeat pay failed")
	}
}