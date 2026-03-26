package visitor

type User struct {
	Name string
	Age int
}

func (u *User) Accept(v Visitor) {
	v.VisitUser(u)
}