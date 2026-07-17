package main

import (
	"fmt"
)

func main() {

	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	aNumber := 42

	fmt.Println(str1, str2, str3)
	stringLen, err := fmt.Println(aNumber)
	if err != nil {
		panic(err)
	}
	fmt.Println("number of bytes written:", stringLen)
}