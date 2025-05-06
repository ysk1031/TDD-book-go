package money

// Money
type Money struct {
	amount int
}

func NewMoney(amount int) Money {
	return Money{amount: amount}
}

func (m Money) Times(multiplier int) Money {
	return NewMoney(m.amount * multiplier)
}

func (m Money) Equals(object any) bool {
	if otherMoney, ok := object.(Money); ok {
		return m.amount == otherMoney.amount
	}
	return false
}

// Dollar
func NewDollar(amount int) Money {
	return NewMoney(amount)
}

// Franc
func NewFranc(amount int) Money {
	return NewMoney(amount)
}
