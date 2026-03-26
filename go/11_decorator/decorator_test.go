package decorator

import "testing"

func TestDecoratorChain(t *testing.T) {
	var svc Service = &BaseService{}

	svc = NewAuthDecorator(svc)
	svc = NewLogDecorator(svc)

	res := svc.Do("test")

	if res != "handle:test" {
		t.Errorf("unexpected result: %s", res)
	}
}

func TestAuthFail(t *testing.T) {
	var svc Service = &BaseService{}

	svc = NewAuthDecorator(svc)

	res := svc.Do("")

	if res != "unauthorized" {
		t.Errorf("expected unauthorized, got %s", res)
	}
}

func TestMultiDecorator(t *testing.T) {
	var svc Service = &BaseService{}

	svc = NewAuthDecorator(svc)
	svc = NewLogDecorator(svc)
	svc = NewTimerDecorator(svc)

	res := svc.Do("ok")

	if res != "handle:ok" {
		t.Errorf("unexpected result: %s", res)
	}
}