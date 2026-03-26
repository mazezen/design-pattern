package composite

import "fmt"

type Floder struct {
	components []Component
	name string
}

func (f *Floder) search(keyword string) {
	fmt.Printf("Searching recursively for keyword %s in floder %s\n", keyword, f.name)
	for _, composite := range f.components {
		composite.search(keyword)
	}
}

func (f *Floder) add(c Component) {
	f.components = append(f.components, c)
}