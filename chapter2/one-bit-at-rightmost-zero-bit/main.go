// page 11
package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter a binary number:\n")

	var inNum uint16
	fmt.Scanf("%b", &inNum)

	fmt.Printf("%016b\n", inNum)

	outNum := (^inNum) & (inNum + 1)

	fmt.Printf("%016b\n", outNum)
}
