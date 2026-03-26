package main

import (
	"fmt"

	iterator "github.com/mazezen/design-pattern/go/16_iterator"
)

func main() {

	users := []iterator.User{
		{"A", 18},
		{"B", 25},
		{"C", 30},
	}


	it := iterator.NewUserIterator(users)

	// 过滤年龄 > 20
	filerIt := iterator.NewFilterIterator(it, func (v interface{}) bool {
		return v.(iterator.User).Age > 20
	})

	for filerIt.HasNext() {
		u := filerIt.Next().(iterator.User)
		fmt.Println(u.Name)
	}
}
