package money

// Money
type Money struct {
	amount      int
	concurrency string
}

func NewMoney(amount int, name string) Money {
	return Money{amount: amount, concurrency: name}
}

func (m Money) Times(multiplier int) Money {
	return NewMoney(m.amount*multiplier, m.concurrency)
}

func (m Money) Equals(object any) bool {
	if otherMoney, ok := object.(Money); ok {
		return m.amount == otherMoney.amount && m.concurrency == otherMoney.concurrency
	}
	return false
}

func (m Money) Currency() string {
	return m.concurrency
}

func (m Money) plus(addend Money) Expression {
	return NewMoney(m.amount+addend.amount, m.concurrency)
}

// Dollar
func NewDollar(amount int) Money {
	return NewMoney(amount, "USD")
}

// Franc
func NewFranc(amount int) Money {
	return NewMoney(amount, "CHF")
}
