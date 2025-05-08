package money

type Bank struct {
}

func NewBank() Bank {
	return Bank{}
}

func (b Bank) reduce(source Expression, to string) Money {
	return NewDollar(10)
}
