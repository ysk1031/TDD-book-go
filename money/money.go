package money

// Money
type Money struct {
	amount   int
	currency string
}

func NewMoney(amount int, name string) Money {
	return Money{amount: amount, currency: name}
}

func (m Money) Times(multiplier int) Expression {
	return NewMoney(m.amount*multiplier, m.currency)
}

func (m Money) Currency() string {
	return m.currency
}

func (m Money) Plus(addend Expression) Expression {
	return NewSum(m, addend)
}

func (m Money) Reduce(bank Bank, to string) Money {
	rate := bank.Rate(m.currency, to)
	return NewMoney(m.amount/rate, to)
}

// Dollar
func NewDollar(amount int) Money {
	return NewMoney(amount, "USD")
}

// Franc
func NewFranc(amount int) Money {
	return NewMoney(amount, "CHF")
}
