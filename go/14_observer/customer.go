package observer

type Customer struct {
	Name     string
	Messages []string 
}

func NewCustomer(name string) *Customer {
	return &Customer{Name: name}
}

func (c *Customer) Update(productName string) {
	msg := c.Name + " notified: " + productName + " is in stock"
	c.Messages = append(c.Messages, msg)
}