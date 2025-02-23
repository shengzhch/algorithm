package main

func max(vs ...int) int {
	if len(vs) == 0 {
		panic("func min args empty")
	}
	m := vs[0]
	for _, n := range vs[1:] {
		if n > m {
			m = n
		}
	}
	return m
}

func min(vs ...int) int {
	if len(vs) == 0 {
		panic("func min args empty")
	}
	m := vs[0]
	for _, n := range vs[1:] {
		if n < m {
			m = n
		}
	}
	return m
}

func maxDG(vs []int) int {
	if len(vs) == 1 {
		return vs[0]
	}
	if vs[0] > vs[1] {
		vs[1] = vs[0]
	}

	return maxDG(vs[1:])
}

func maxDG2(vs []int) int {
	if len(vs) == 1 {
		return vs[0]
	}
	if len(vs) == 2 {
		if vs[0] > vs[1] {
			return vs[0]
		}
		return vs[1]
	}
	if vs[0] > vs[1] {
		return maxDG2(append(vs[2:], vs[0]))
	}

	return maxDG2(vs[1:])
}
