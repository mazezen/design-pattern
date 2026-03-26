package visitor

import "fmt"

type PrintVisitor struct {}

func (p *PrintVisitor) VisitUser(u *User) {
	fmt.Printf("User: %s, Age: %d\n", u.Name, u.Age)
}

func (p *PrintVisitor) VisitOrder(o *Order) {
	fmt.Printf("Order: %d, Amount: %d\n", o.ID, o.Amount)
}