package main

import (
	"fmt"

	interpreter "github.com/mazezen/design-pattern/go/21_interpreter"
)


func main() {
	expr := interpreter.Parse("1 + 2 + 3")
	result := expr.Interpret()

	fmt.Println(result)
}
