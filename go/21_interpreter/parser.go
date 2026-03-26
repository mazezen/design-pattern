package interpreter

import (
	"strconv"
	"strings"
)

// Parse 加法表达式解释器
func Parse(input string) Expression {
	tokens := strings.Split(input, " ")

	var expr Expression

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]

		if token == "+" {
			continue
		}

		num, _ := strconv.Atoi(token)
		number := NewNumber(num)

		if expr == nil {
			expr = number
		} else {
			expr = NewAdd(expr, number)
		}
	}

	return expr
}