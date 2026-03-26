package interpreter

// Number 数字 (终结符表达式)
type Number struct {
	value int
}

func NewNumber(val int) *Number {
	return &Number{value: val}
}

func (n *Number) Interpret() int {
	return n.value
}