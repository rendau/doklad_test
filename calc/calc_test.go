package calc

import "testing"

func TestAdd(t *testing.T) {
	calc := &Calculator{}

	result := calc.Add(2, 3)
	if result != 5 {
		t.Errorf("Expected 5, but got %d", result)
	}
}

func TestSubtract(t *testing.T) {
	calc := &Calculator{}

	result := calc.Subtract(5, 3)
	if result != 2 {
		t.Errorf("Expected 2, but got %d", result)
	}
}

func TestMultiply(t *testing.T) {
	calc := &Calculator{}

	result := calc.Multiply(2, 3)
	if result != 6 {
		t.Errorf("Expected 6, but got %d", result)
	}
}

func TestDivide(t *testing.T) {
	calc := &Calculator{}

	result, err := calc.Divide(6, 3)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != 2 {
		t.Errorf("Expected 2, but got %d", result)
	}
}

func TestDivideByZero(t *testing.T) {
	calc := &Calculator{}

	_, err := calc.Divide(6, 0)
	if err == nil {
		t.Error("Expected error but got nil")
	}
}
