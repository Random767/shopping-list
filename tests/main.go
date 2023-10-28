package main

import "fmt"

func calculate(x int, y int, f func(int, int) int) int {
	return f(x, y)
}

func add(x int, y int) int {
	return x + y
}

func subtract(x int, y int) int {
	return x - y
}

func main() {
	result := calculate(5, 3, add)
	fmt.Println(result) // Output: 8

	result = calculate(5, 3, subtract)
	fmt.Println(result) // Output: 2
}
