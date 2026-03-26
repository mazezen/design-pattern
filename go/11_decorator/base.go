package decorator

type BaseService struct{}

func (b *BaseService) Do(req string) string {
	return "handle: " + req
}