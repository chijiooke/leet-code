package romanToInt

import (
	"fmt"
)

func RomanToInteger(str string) {

	romanIntMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var intConv int
	r := []rune(str)

	for i := range r {
		if i < len(r)-1 && romanIntMap[r[i]] < romanIntMap[r[i+1]] {
			intConv -= romanIntMap[r[i]]
		} else {
			intConv += romanIntMap[r[i]]
		}
	}

	fmt.Println(intConv) //returns integer conversion of roma numeral entered
}
