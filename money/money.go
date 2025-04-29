package money

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
