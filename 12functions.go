package main

import (
	"fmt"
)

func main() {
	superman()
	result := multiplyme(3, 6)
	fmt.Println(result)
	myresult, length, name := adder(3, 4, 5, 6)
	fmt.Println(myresult, length, name)
	mulresult, loc := multiplier(2, 5, 6, 7, 8, 4, 5, 6, 3, 7, 4, 6)
	fmt.Println(mulresult, loc)
}

func multiplier(values ...int) (int, string) {
	mul := 1
	for i := range values {
		mul = mul * values[i]
	}
	loc := "Multiplyier String"
	return mul, loc
}

func superman() {
	fmt.Println("I am a superman")
}

func multiplyme(v1 int, v2 int) int {
	return v1 * v2
}

func adder(values ...int) (int, int, string) {
	sum := 0
	for i := range values {
		sum = sum + values[i]
	}
	length := len(values)
	name := "just for fun"
	return sum, length, name
}
