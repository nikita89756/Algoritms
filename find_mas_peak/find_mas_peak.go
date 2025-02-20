package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	mas := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&mas[i])
	}
	index := findPeak(mas, n)
	fmt.Println(index)
}

func findPeak(mas []int, n int) int {
	l := 0
	r := n - 2
	for l <= r {
		m := (l + r) / 2
		if mas[m] < mas[m+1] {
			l = m + 1
		} else {
			r = m-1
		}
	}
	return l
}