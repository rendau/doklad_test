package calc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	calc := &Calculator{}

	result := calc.Add(2, 3)
	assert.Equal(t, 5, result, "they should be equal")
}

func TestSubtract(t *testing.T) {
	calc := &Calculator{}

	result := calc.Subtract(5, 3)
	assert.Equal(t, 2, result, "they should be equal")
}

func TestMultiply(t *testing.T) {
	calc := &Calculator{}

	result := calc.Multiply(2, 3)
	assert.Equal(t, 6, result, "they should be equal")
}

func TestDivide(t *testing.T) {
	calc := &Calculator{}

	result, err := calc.Divide(6, 3)
	assert.NoError(t, err, "should not return error")
	assert.Equal(t, 2, result, "they should be equal")
}

func TestDivideByZero(t *testing.T) {
	calc := &Calculator{}

	_, err := calc.Divide(6, 0)
	assert.Error(t, err, "should return error")
}
