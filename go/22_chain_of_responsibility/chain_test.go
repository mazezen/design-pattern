package chainofresponsibility

import "testing"

// go test -v ./...

func buildChain() Handler {
	auth := &AuthHandler{}
	role := &RoleHandle{}
	param := &ParamHandler{}
	final := &FinalHandler{}

	auth.SetNext(role)
	role.SetNext(param)
	param.SetNext(final)

	return auth
}

func TestSuccess(t *testing.T) {
	chain := buildChain()

	req := &Request{
		IsLogin: true,
		Role: "admin",
		Param: "ok",
	}

	res := chain.Handle(req)

	if res != "success" {
		t.Errorf("expected success, got %s", res)
	}
}

func TestNotLogin(t *testing.T) {
	chain := buildChain()

	req := &Request{
		IsLogin: false,
		Role: "admin",
		Param: "ok",
	}

	res := chain.Handle(req)
	
	if res != "not login" {
		t.Errorf("unexpected result: %s", res)
	}
}

func TestNotPermission(t *testing.T) {
	chain := buildChain()

	req := &Request{
		IsLogin: true,
		Param: "ok",
	}

	res := chain.Handle(req)

	if res != "no permision" {
		t.Errorf("unexpected result: %s", res)
	}
}

func TestNotParam(t *testing.T) {
	chain := buildChain()

	req := &Request{
		IsLogin: true,
		Role: "admin",
	}

	res := chain.Handle(req)

	if res != "invalid param" {
		t.Errorf("unexpected result: %s", res)
	}
}