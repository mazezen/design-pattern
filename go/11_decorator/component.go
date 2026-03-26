package decorator

type Service interface {
	Do(req string) string
}