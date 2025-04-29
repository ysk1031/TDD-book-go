package money

import "testing"

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	if five.Times(2) != NewDollar(10) {
		t.Errorf("Expected amount to be 10, but got %d", five.Times(2).amount)
	}

	if five.Times(3) != NewDollar(15) {
		t.Errorf("Expected amount to be 15, but got %d", five.Times(3).amount)
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

func TestFrancMultiplication(t *testing.T) {
	five := NewFranc(5)
	if five.Times(2) != NewFranc(10) {
		t.Errorf("Expected amount to be 10, but got %d", five.Times(2).amount)
	}

	if five.Times(3) != NewFranc(15) {
		t.Errorf("Expected amount to be 15, but got %d", five.Times(3).amount)
	}
}
