package observer

type Product struct {
	Name      string
	InStock   bool
	observers []Observer
}

func NewProduct(name string) *Product {
	return &Product{Name: name}
}

func (p *Product) Subscribe(o Observer) {
	p.observers = append(p.observers, o)
}

func (p *Product) Unsubscribe(o Observer) {
	for i, ob := range p.observers {
		if ob == o {
			p.observers = append(p.observers[:i], p.observers[i+1:]...)
			break
		}
	}
}

func (p *Product) Notify() {
	for _, o := range p.observers {
		o.Update(p.Name)
	}
}

// ⭐ 补货时触发通知
func (p *Product) SetStock(inStock bool) {
	p.InStock = inStock
	if inStock {
		p.Notify()
	}
}