package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello there :D") // This code for print "Hello there :D" in console
	var x int = 10                // This code for declare variable x with value 10
	const y int = 54              // This code for declare variable y with value 54
	fmt.Println(x + y)
	var array = [3]int{1, 2, 3}          // This code for declare array with 3 element
	array2 := [...]int{4, 5, 6, 7, 8, 9} // This code for declare array with 6 element
	fmt.Println(array)
	fmt.Print(array2)

	var numbers = [4]int{2, 4, 7, 2}
	var max int
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > max {
			max = numbers[i]
		}
	}
	fmt.Println("\nMax value is ", max)

	myslice := [5]int{1, 2, 3, 4, 5} // This code for declare myslice with 5 element and capacity 7
	fmt.Println(cap(myslice))        // This code for print capacity of myslice result is 7 because myslice have 7 element
	fmt.Println(len(myslice))        // This code for print length of myslice result is 5 because myslice have 5 element
	fmt.Println(myslice)             // This code for print all element of myslice
	array_slice := array2[3:5]       // This code for slice array2 from index 3 to 5
	fmt.Println(array_slice)         // This code for print array_slice

	neededNumbes := myslice[:len(myslice)-2]      // This code for slice myslice from index 0 to 3
	numbersCopy := make([]int, len(neededNumbes)) // This code for copy neededNumbes to numbersCopy
	copy(numbersCopy, neededNumbes)               // This code for copy neededNumbes to numbersCopy
	fmt.Println(numbersCopy)                      // This code for print neededNumbes
}
