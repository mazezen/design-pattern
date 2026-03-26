package simplefactory

import (
	"fmt"
)

// API is interface
type API interface {
	Say(name string) string
}

// NewAPI return API instance by type
func NewAPI(t int) API {
	switch t {
	case 1:
		return &hiAPI{}
	case 2:
		return &helloAPI{}
	default:
		return nil
	}
}

// hiAPI is one of API implement
type hiAPI struct{}

func (*hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

// helloAPI is one of API implement
type helloAPI struct{}

func (*helloAPI) Say(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
