package money

// Dollar
type Dollar struct {
	amount int
}

func NewDollar(amount int) Dollar {
	return Dollar{amount: amount}
}

func (d Dollar) Times(multiplier int) Dollar {
	return NewDollar(d.amount * multiplier)
}

func (d Dollar) Equals(object any) bool {
	if otherDollar, ok := object.(Dollar); ok {
		return d.amount == otherDollar.amount
	}
	return false
}

// Franc
type Franc struct {
	amount int
}

func NewFranc(amount int) Franc {
	return Franc{amount: amount}
}

func (f Franc) Times(multiplier int) Franc {
	return NewFranc(f.amount * multiplier)
}

func (f Franc) Equals(object any) bool {
	if otherFranc, ok := object.(Franc); ok {
		return f.amount == otherFranc.amount
	}
	return false
}
