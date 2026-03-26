package decorator

import (
	"log"
	"time"
)

type TimerDecorator struct {
	Decorator
}

func NewTimerDecorator(s Service) Service {
	return &TimerDecorator{
		Decorator{service: s},
	}
}

func (t *TimerDecorator) Do(req string) string {
	start := time.Now()
	res := t.service.Do(req)
	log.Println("[TIMER] cost: ", time.Since(start))
	return res
}