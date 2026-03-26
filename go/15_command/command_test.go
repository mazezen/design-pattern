package command

import "testing"

// go test -v command_test.go hello_command.go add_command.go command.go

func TestHelloCommand(t *testing.T) {
	cmd := &HelloCommand{}

	if res := cmd.Execute([]string{}); res != "hello world" {
		t.Errorf("exprected hello world, got %s", res)
	}

	if res := cmd.Execute([]string{"mazezen"}); res != "hello mazezen" {
		t.Errorf("unexpected result: %s", res)
	}
}

func TestAddCommand(t *testing.T) {
	cmd := &AddCommand{}

	if res := cmd.Execute([]string{}); res != "need 2 numbers" {
		t.Fatalf("unexpected result: %s", res)
	}

	if res := cmd.Execute([]string{"1"}); res != "need 2 numbers" {
		t.Errorf("unexpected result: %s", res)
	}

	if res := cmd.Execute([]string{"2", "8"}); res != "10" {
		t.Errorf("expected 10, got %s", res)
	}

	if res := cmd.Execute([]string{"a", "5"}); res != "invalid number" {
		t.Errorf("unexpected result: %s", res)
	}
}

