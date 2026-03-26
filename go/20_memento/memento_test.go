package memento

import (
	"testing"
)

// go test -v ./...

func TestMemento(t *testing.T) {
	
	caretaker := &Caretaker{
		mementos: make([]*Memento, 0),
	}

	o := &Originator{
		state: "A",
	}

	if o.getState() != "A" {
		t.Fatalf("expected A, got %s", o.getState())
	}
	caretaker.addMemento(o.createMemento())
	

	o.setState("B")
	if o.getState() != "B" {
		t.Fatalf("expected B, got %s", o.getState())
	}
	caretaker.addMemento(o.createMemento())

	o.setState("C")
	if o.getState() != "C" {
		t.Fatalf("expected C, got %s", o.getState())
	}
	caretaker.addMemento(o.createMemento())


	o.restoreMemento(caretaker.getMemento(0))
	if o.getState() != "A" {
		t.Fatalf("expected A, got %s", o.getState())
	}

	o.restoreMemento(caretaker.getMemento(1))
	if o.getState() != "B" {
		t.Fatalf("expected A, got %s", o.getState())
	}

	o.restoreMemento(caretaker.getMemento(2))
	if o.getState() != "C" {
		t.Fatalf("expected A, got %s", o.getState())
	}
}