package money

import "testing"

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	five.times(2)
	if five.amount != 10 {
		t.Errorf("Expected amount to be 10, but got %d", five.amount)
	}
}
