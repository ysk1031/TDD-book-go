package money

import "testing"

func TestMultiplication(t *testing.T) {
	five := NewDollar(5)
	if five.Times(2) != NewDollar(10) {
		t.Errorf("Expected amount to be %v, but got %v", NewDollar(10), five.Times(2))
	}
	if five.Times(3) != NewDollar(15) {
		t.Errorf("Expected amount to be %v, but got %v", NewDollar(15), five.Times(3))
	}
}

func TestEquality(t *testing.T) {
	if !(NewDollar(5)).Equals(NewDollar(5)) {
		t.Errorf("Expected dollars to be equal, but they are not")
	}
	if NewDollar(5).Equals(NewDollar(6)) {
		t.Errorf("Expected dollars to be not equal, but they are")
	}

	if NewDollar(5).Equals(NewFranc(5)) {
		t.Errorf("Expected dollar and franc to be not equal, but they are")
	}
}

func TestCurrency(t *testing.T) {
	if NewDollar(1).Currency() != "USD" {
		t.Errorf("Expected currency to be USD, but got %s", NewDollar(1).Currency())
	}
	if NewFranc(1).Currency() != "CHF" {
		t.Errorf("Expected currency to be CHF, but got %s", NewFranc(1).Currency())
	}
}

func TestSimpleAddition(t *testing.T) {
	five := NewDollar(5)
	sum := five.Plus(five)
	bank := NewBank()
	reduced := bank.Reduce(sum, "USD")
	if reduced != NewDollar(10) {
		t.Errorf("Expected reduced amount to be 10, but got %d", reduced.amount)
	}
}

func TestPlusReturnsSum(t *testing.T) {
	five := NewDollar(5)
	result := five.Plus(five)
	sum, ok := result.(Sum)
	if !ok {
		t.Errorf("Expected result to be of type Sum, but got %T", result)
		return
	}

	if sum.augend != five {
		t.Errorf("Expected augend to be %v, but got %v", five, sum.augend)
	}
	if sum.addend != five {
		t.Errorf("Expected addend to be %v, but got %v", five, sum.addend)
	}
}

func TestReduceSum(t *testing.T) {
	sum := NewSum(NewDollar(3), NewDollar(4))
	bank := NewBank()
	result := bank.Reduce(sum, "USD")
	if result != NewDollar(7) {
		t.Errorf("Expected reduced amount to be 7, but got %d", result.amount)
	}
}

func TestReduceMoney(t *testing.T) {
	bank := NewBank()
	result := bank.Reduce(NewDollar(1), "USD")
	if result != NewDollar(1) {
		t.Errorf("Expected reduced amount to be %v, but got %v", NewDollar(1), result)
	}
}

func TestReduceMoneyDifferentCurrency(t *testing.T) {
	bank := NewBank()
	bank.AddRate("CHF", "USD", 2)
	result := bank.Reduce(NewFranc(2), "USD")
	if result != NewDollar(1) {
		t.Errorf("Expected reduced amount to be %v, but got %v", NewDollar(1), result)
	}
}

func TestIdentityRate(t *testing.T) {
	rate := NewBank().Rate("USD", "USD")
	if rate != 1 {
		t.Errorf("Expected identity rate to be 1, but got %d", rate)
	}
}

func TestMixedAddition(t *testing.T) {
	fiveDollars := NewDollar(5)
	tenFrancs := NewFranc(10)
	bank := NewBank()
	bank.AddRate("CHF", "USD", 2)
	result := bank.Reduce(fiveDollars.Plus(tenFrancs), "USD")
	if result != NewDollar(10) {
		t.Errorf("Expected reduced amount to be %v, but got %v", NewDollar(10), result)
	}
}

func TestSumPlusMoney(t *testing.T) {
	fiveDollars := NewDollar(5)
	tenFrancs := NewFranc(10)
	bank := NewBank()
	bank.AddRate("CHF", "USD", 2)
	sum := NewSum(fiveDollars, tenFrancs).Plus(fiveDollars)
	result := bank.Reduce(sum, "USD")
	if result != NewDollar(15) {
		t.Errorf("Expected reduced amount to be %v, but got %v", NewDollar(15), result)
	}
}

func TestSumTimes(t *testing.T) {
	fiveDollars := NewDollar(5)
	tenFrancs := NewFranc(10)
	bank := NewBank()
	bank.AddRate("CHF", "USD", 2)
	sum := NewSum(fiveDollars, tenFrancs).Times(2)
	result := bank.Reduce(sum, "USD")
	if result != NewDollar(20) {
		t.Errorf("Expected reduced amount to be %v, but got %v", NewDollar(20), result)
	}
}
