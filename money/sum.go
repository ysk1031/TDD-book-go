package money

type Sum struct {
	augend Expression
	addend Expression
}

func NewSum(augend, addend Expression) Sum {
	return Sum{augend: augend, addend: addend}
}

func (s Sum) Plus(addend Expression) Expression {
	return NewSum(s, addend)
}

func (s Sum) Reduce(bank Bank, to string) Money {
	amount := s.augend.Reduce(bank, to).amount + s.addend.Reduce(bank, to).amount
	return NewMoney(amount, to)
}
