package strain

// Ints contain array of int
type Ints []int

// Lists contain array of array of int
type Lists [][]int

// Strings contain array of string
type Strings []string

// Keep will keep int that fulfill the bool returned by supplied function.
func (i Ints) Keep(f func(int) bool) Ints {
	var a Ints
	for _, v := range i {
		if f(v) {
			a = append(a, v)
		}
	}
	return a
}

// Discard will discard int that not fulfill the bool returned by supplied function.
func (i Ints) Discard(f func(int) bool) Ints {
	var a Ints
	for _, v := range i {
		if !f(v) {
			a = append(a, v)
		}
	}
	return a
}

// Keep will keep array of int that fulfill the bool returned by supplied function.
func (l Lists) Keep(f func([]int) bool) Lists {
	var a Lists
	for _, v := range l {
		if f(v) {
			a = append(a, v)
		}
	}
	return a
}

// Keep will keep string that fulfill the bool returned by supplied function.
func (s Strings) Keep(f func(string) bool) Strings {
	var a Strings
	for _, v := range s {
		if f(v) {
			a = append(a, v)
		}
	}
	return a
}
