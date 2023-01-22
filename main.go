package web

func Between(a, b int) []int {
	var c []int
	for i := a; i <= b; i++ {
		c = append(c, i)
	}
	return c
}
