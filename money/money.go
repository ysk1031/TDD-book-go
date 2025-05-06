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
type Dollar struct {
	Money
}

func NewDollar(amount int) Dollar {
	return Dollar{Money: NewMoney(amount)}
}

func (d Dollar) Times(multiplier int) Dollar {
	return NewDollar(d.Money.amount * multiplier)
}

func (d Dollar) Equals(object any) bool {
	if otherDollar, ok := object.(Dollar); ok {
		return d.Money.Equals(otherDollar.Money)
	}
	return false
}

// Franc
type Franc struct {
	Money
}

func NewFranc(amount int) Franc {
	return Franc{Money: NewMoney(amount)}
}

func (f Franc) Times(multiplier int) Franc {
	return NewFranc(f.Money.amount * multiplier)
}

func (f Franc) Equals(object any) bool {
	if otherFranc, ok := object.(Franc); ok {
		return f.Money.Equals(otherFranc.Money)
	}
	return false
}
