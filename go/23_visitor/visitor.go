package visitor

type Visitor interface {
	VisitUser(*User)
	VisitOrder(*Order)
}