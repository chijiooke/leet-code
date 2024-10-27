package main

import (
	"fmt"
	romanToInt "leet-code/solutions/roman-numeral-to-interger/pkg"
)

func main() {

	var r string //roman numeral string

	fmt.Print("Enter roman numeral: ")
	fmt.Scanln(&r)

	romanToInt.RomanToInteger(r) //prints true if number is a palindrome or false otherwise
}
