package words

func newNumberString() string {
	return string(randomIntOf(makeRange(0, 10000)))
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}
