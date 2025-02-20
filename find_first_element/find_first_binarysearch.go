package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n, &x)
	mas := make([]int, n)
	for i := 0; i < n; i++ {

		fmt.Scan(&mas[i])
	}
	index := binarySearch(mas, x, n)
	fmt.Println(index)
}

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
			l = m + 1
		}

	}
	return res
}
