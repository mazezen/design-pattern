package observer

type Observer interface {
	Update(productName string)
}