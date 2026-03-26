package visitor

type Order struct {
	ID int
	Amount int
}

func (o *Order) Accept(v Visitor) {
	v.VisitOrder(o)
}