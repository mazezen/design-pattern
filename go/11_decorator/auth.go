// decorator/auth.go
package decorator

type AuthDecorator struct {
	Decorator
}

func NewAuthDecorator(s Service) Service {
	return &AuthDecorator{
		Decorator{service: s},
	}
}

func (a *AuthDecorator) Do(req string) string {
	if req == "" {
		return "unauthorized"
	}
	return a.service.Do(req)
}