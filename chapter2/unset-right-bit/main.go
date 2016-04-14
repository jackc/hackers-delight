package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter a binary number:\n")

	var inNum int32
	fmt.Scanf("%b", &inNum)

	outNum := inNum & (inNum - 1)

	fmt.Printf("%b\n", outNum)
}
