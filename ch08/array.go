package ch08

type Array []int

func NewArray(ints ...int) *Array {
	arr := Array(ints)
	return &arr
}

func (a *Array) Get(i int) int {
	return (*a)[i] // O(1)
}

func (a *Array) Set(i, x int) {
	(*a)[i] = x // O(1)
}

func (a *Array) Append(ints ...int) {
	tmp := make(Array, len(*a)+len(ints))
	copy(tmp, *a)

	for i := range ints {
		tmp[len(*a)+i] = ints[i]
	}
	*a = tmp
}

// Insert x after y once
func (a *Array) Insert(x int, y int) {
	for i := range *a { // O(N)
		if (*a)[i] == y {
			tmp := make(Array, 0, len(*a)+1)
			tmp.Append((*a)[:i+1]...)
			tmp.Append(x)
			tmp.Append((*a)[i+1:]...)
			*a = tmp
			break
		}
	}
}

// Remove x once
func (a *Array) Remove(x int) {
	for i := range *a { // O(N)
		if (*a)[i] == x {
			tmp := Array((*a)[:i])
			tmp.Append((*a)[i+1:]...)
			*a = tmp
			break
		}
	}
}
