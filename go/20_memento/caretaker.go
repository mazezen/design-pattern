package memento

type Caretaker struct {
	mementos []*Memento
}

func (c *Caretaker) addMemento(m *Memento) {
	c.mementos = append(c.mementos, m)
}

func (c *Caretaker) getMemento(index int) *Memento {
	return c.mementos[index]
}