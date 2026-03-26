package visitor

type SumVisitor struct {
	Total int
}

func (s *SumVisitor) VisitUser(u *User) {
	// 用户不参与统计
}

func (s *SumVisitor) VisitOrder(o *Order) {
	s.Total += o.Amount
}