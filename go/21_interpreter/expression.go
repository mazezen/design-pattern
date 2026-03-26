package interpreter

// Expression 表达式接口
type Expression interface {
	Interpret() int
}