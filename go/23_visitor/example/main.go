package main

import (
	"fmt"

	visitor "github.com/mazezen/design-pattern/go/23_visitor"
)


func main() {

	elements := []visitor.Element{
		&visitor.User{Name: "A", Age: 20},
		&visitor.Order{ID: 1, Amount: 100},
		&visitor.Order{ID: 2, Amount: 200},
	}


	// 打印
	printVisitor := &visitor.PrintVisitor{}
	for _, e := range elements {
		e.Accept(printVisitor)
	}

	// 统计
	sumVisitor := &visitor.SumVisitor{}
	for _, e := range elements {
		e.Accept(sumVisitor)
	}
	
	fmt.Println("total:", sumVisitor.Total)
}
