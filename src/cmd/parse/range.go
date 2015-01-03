package parse

type RuneRange struct {
	From rune
	To   rune
}

func NewRuneRange(from rune, to rune) *RuneRange {
	if from > to {
		from = from ^ to
		to = from ^ to
		from = from ^ to
	}
	return &RuneRange{From: from, To: to}
}

func (r *RuneRange) Contains(target rune) bool {
	return r.From <= target && target <= r.To
}
