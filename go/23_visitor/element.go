package visitor

type Element interface {
	Accept(v Visitor)
}