package main

import (
	"fmt"
	"math"
)

func main() {

	f1, f2, f3 := 23.5, 65.1, 76.3
	sum := f1 + f2 + f3
	fmt.Println("Float sum:", sum)
	fmt.Printf("Float sum with 2 decimals: %.2f\n ", sum)
	sumRounded := math.Round(sum * 100) / 100
	fmt.Println("Float sum rounded: ", sumRounded)

}
