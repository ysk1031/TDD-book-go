package money

type Pair struct {
	From string
	To   string
}

func NewPair(from, to string) Pair {
	return Pair{From: from, To: to}
}

func (p Pair) Equals(other any) bool {
	if otherPair, ok := other.(Pair); ok {
		return p.From == otherPair.From && p.To == otherPair.To
	}
	return false
}

func (p Pair) HashCode() int {
	return 0
}
