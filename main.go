package main

import "fmt"

func main() {
	num := sum(5, 3)
	fmt.Println(num)
	PrintMeta()
}

func sum(num1 int, num2 int) int {
	return num1 + num2
}
