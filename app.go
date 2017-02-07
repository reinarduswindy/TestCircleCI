package main

import (
	"fmt"
)

func main() {
	abc := GetABC()
	fmt.Println(abc)

	a := 5
	b := 6
	c := AddSquare(a, b)
	fmt.Println(c)
}

func GetABC() string {
	return "ABC"
}

func AddSquare(a, b int) int {
	tempA, tempB := Square(a, b)

	return tempA + tempB
}

func Square(a, b int) (int, int) {
	return a * a, b * b
}
