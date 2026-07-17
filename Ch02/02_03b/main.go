package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	bufReader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter text: ")
	inputText, _ := bufReader.ReadString('\n')
	fmt.Println(inputText)
	fmt.Println("Enter float number:")
	inputText, _ = bufReader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(inputText), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Number: %f\n", f)
	}
}
