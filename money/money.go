package money

// Money
type Money struct {
	amount int
	name   string
}

func NewMoney(amount int, name string) Money {
	return Money{amount: amount, name: name}
}

func (m Money) Times(multiplier int) Money {
	return NewMoney(m.amount*multiplier, m.name)
}

func (m Money) Equals(object any) bool {
	if otherMoney, ok := object.(Money); ok {
		return m.amount == otherMoney.amount && m.name == otherMoney.name
	}
	return false
}

// Dollar
func NewDollar(amount int) Money {
	return NewMoney(amount, "USD")
}

// Franc
func NewFranc(amount int) Money {
	return NewMoney(amount, "CHF")
}
