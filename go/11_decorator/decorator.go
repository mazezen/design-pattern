package decorator

type Decorator struct {
	service Service
}

func (d *Decorator) Do(req string) string {
	return d.service.Do(req)
}