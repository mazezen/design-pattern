package interpreter

import "testing"

// go test -v ./...

func TestSimpleAdd(t *testing.T) {
	expr := Parse("1 + 2 + 3")

	result := expr.Interpret()

	if result != 6 {
		t.Errorf("expected 6, got %d", result)
	}
}

func TestSingleNumber(t *testing.T) {
	expr := Parse("5")

	result := expr.Interpret()

	if result != 5 {
		t.Errorf("expected 5, got %d", result)
	}
}

func TestComplex(t *testing.T) {
	expr := Parse("10 + 20 + 30 + 40")

	result := expr.Interpret()

	if result != 100 {
		t.Errorf("expected 100, got %d", result)
	}
}