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
