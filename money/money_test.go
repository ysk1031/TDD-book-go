package money

import "testing"

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	product := five.Times(2)
	if product.amount != 10 {
		t.Errorf("Expected amount to be 10, but got %d", product.amount)
	}

	product = five.Times(3)
	if product.amount != 15 {
		t.Errorf("Expected amount to be 15, but got %d", product.amount)
	}
}

func TestEquality(t *testing.T) {
	if !(NewDollar(5)).Equals(NewDollar(5)) {
		t.Errorf("Expected dollars to be equal, but they are not")
	}

	if NewDollar(5).Equals(NewDollar(6)) {
		t.Errorf("Expected dollars to be not equal, but they are")
	}
}
