package money

type Bank struct {
}

func NewBank() Bank {
	return Bank{}
}

func (b Bank) Reduce(source Expression, to string) Money {
	return source.Reduce(to)
}
