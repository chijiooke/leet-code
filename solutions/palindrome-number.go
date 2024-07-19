package solution

import "fmt"

func IsPalindrome(x int) {

	if x < 0 {
		fmt.Println(false)
		return
	}

	if x > 0 && x < 10 {
		fmt.Println(true)
		return
	}

	stash := x
	reversed := 0

	for {

		if stash == 0 {
			break
		}
		reversed = (reversed * 10) + (stash % 10)
		stash = stash / 10
	}

	fmt.Println(reversed == x) //is the number same with the reversed format
	

}