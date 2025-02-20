package main

func binarySearch(mas []int, x int, n int) int {
	l := 0
	r := n - 1
	res := -1
	for l <= r {
		m := (r + l) / 2

		if mas[m] >= x {
			res = m
			r = m - 1
		} else {
			r = m + 1
		}

	}
	return res
}
