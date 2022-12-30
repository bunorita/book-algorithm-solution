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

func (a *Array) Append(x int) {
	*a = append(*a, x)
}

// Insert x after y once
func (a *Array) Insert(x int, y int) {
	for i := range *a { // O(N)
		if (*a)[i] == y {
			tmp := make(Array, i+1)
			copy(tmp, (*a)[:i+1])

			*a = append(append(tmp, x), (*a)[i+1:]...)
			break
		}
	}
}

// Remove x once
func (a *Array) Remove(x int) {
	for i := range *a { // O(N)
		if (*a)[i] == x {
			*a = append((*a)[:i], (*a)[i+1:]...)
			break
		}
	}
}
