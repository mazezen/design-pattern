package interpreter

type Add struct {
	left Expression
	right Expression
}

func NewAdd(left, right Expression) *Add {
	return &Add{left: left, right: right}
}

func (a *Add) Interpret() int {
	return a.left.Interpret() + a.right.Interpret()
}