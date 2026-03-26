package decorator

import "log"

type LogDecorator struct {
	Decorator
}

func NewLogDecorator(s Service) Service {
	return &LogDecorator{
		Decorator{service: s},
	}
}

func (l *LogDecorator) Do(req string) string {
	log.Println("[LOG] request: ", req)
	res := l.service.Do(req)
	log.Println("[LOG] response: ", res)
	return res
}