package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("Enter a decimal number: ")
	var binary int
	fmt.Scan(&binary)
	decimalNumber := decimalNumber(binary)
	fmt.Printf("Binary representation of %d is %s\n", binary, decimalNumber)
}

func decimalNumber(binary int) string {
	return strconv.FormatInt(int64(binary), 2)
}
