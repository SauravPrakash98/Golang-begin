package main

import "fmt"

func diag(n int) int {
	if n == 1 {
		return 1
	}
	if n%2 == 0 {
		return diag(n - 1)
	}
	return diag(n-1) + 4*n*n - 6*(n-1)
}

func main() {
	fmt.Println(diag(3))
}
