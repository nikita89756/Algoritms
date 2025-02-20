package main

import "fmt"

func main() {
	var n, x int
	fmt.Scan(&n, &x)
	mas := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&mas[i])
	}
	index := searchNearest(mas, x, n)
	fmt.Println(index)
}

func searchNearest(mas []int, x int, n int) int {
	l:=0
	r:= n-1
	ind:=n-1
	diff:=0
	if x<mas[n-1]{
		diff= mas[n-1]-x
	}else{
		diff= x-mas[n-1]
	}
	for l<=r{
		m:=(l+r)/2
		if mas[m]>=x{
			r =m-1
			
			if diff==mas[m]-x{
				ind = min(m,ind)
				diff = mas[m]-x
			
			}else if diff>mas[m]-x{
				ind = m
				diff = mas[m]-x
			}
		}else{
			l=m+1
			if diff==x-mas[m]{
				ind = min(m,ind)
				diff = x - mas[m]

			}else if diff>x-mas[m]{
				ind = m
				diff = x - mas[m]
			}
		}
	}
	return ind
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}