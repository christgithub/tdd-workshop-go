package number

import "testing"

func TestChrisInteger(t *testing.T) {
	ci := NewChrisInt(2)
	val := ci.GetNative()

	if val != 2 {
		t.Fatal("Failed to instantiate")
	}
}

func TestChrisInteger_IsEqualTo(t *testing.T) {
	ci := NewChrisInt(2)
	ci2 := NewChrisInt(2)
	ci.IsEqualTo(ci2)

	if !ci.IsEqualTo(ci2) {
		t.Fatal("Not equal")
	}
}

func TestIsGreaterThan(t *testing.T) {
	ci := NewChrisInt(4)
	ci2 := NewChrisInt(3)
	ci.IsGreaterThan(ci2)

	if !ci.IsGreaterThan(ci2) {
		t.Fatal("Not greater than")
	}
}

func TestIsLessThan(t *testing.T) {
	ci := NewChrisInt(2)
	ci2 := NewChrisInt(3)
	ci.IsGreaterThan(ci2)

	if !ci.IsLessThan(ci2) {
		t.Fatal("Not less than")
	}
}

func TestAdd(t *testing.T) {
	ci := NewChrisInt(2)
	ci2 := NewChrisInt(3)
	ci3 := NewChrisInt(5)

	result := ci.Add(ci2)
	if ci3.GetNative() != result.GetNative() {
		t.Fatalf("Expected %d but got %d", ci3.GetNative(), result.GetNative())
	}
}

func TestSubtract(t *testing.T) {
	ci := NewChrisInt(8)
	ci2 := NewChrisInt(5)
	ci3 := NewChrisInt(3)

	result := ci.Subtract(ci2)
	if ci3.GetNative() != result.GetNative() {
		t.Fatalf("Expected %d but got %d", ci3.GetNative(), result.GetNative())
	}
}

func TestMultiply(t *testing.T) {
	ci := NewChrisInt(8)
	ci2 := NewChrisInt(5)
	ci3 := NewChrisInt(40)

	result := ci.Multiply(ci2)
	if ci3.GetNative() != result.GetNative() {
		t.Fatalf("Expected %d but got %d", ci3.GetNative(), result.GetNative())
	}
}

func TestDivideBy(t *testing.T) {
	ci := NewChrisInt(40)
	ci2 := NewChrisInt(5)
	ci3 := NewChrisInt(8)

	result, err := ci.DivideBy(ci2)

	if err != nil {
		t.Fatalf("Unexpected result. %s", err)
	}

	if ci3.GetNative() != result.GetNative() {
		t.Fatalf("Expected %d but got %d", ci3.GetNative(), result.GetNative())
	}
}
