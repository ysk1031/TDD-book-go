package money

type Pair struct {
	From string
	To   string
}

func NewPair(from, to string) Pair {
	return Pair{From: from, To: to}
}
