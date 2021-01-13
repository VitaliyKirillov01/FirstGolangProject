package main

import "fmt"

func swap(x1 *int, x2 *int) {
	temp := *x1
	*x1 = *x2
	*x2 = temp
}

func main() {
	x1 := 1
	x2 := 2
	fmt.Println(x1, x2)
	swap(&x1, &x2)
	fmt.Println(x1, x2)
}
