package main

import (
	"fmt"
	longestcommonprefix "leet-code/solutions/longest-common-prefix/pkg"
)

func main() {

	s := []string{"ab", "a"} //list of strings - change at will

	prefix := longestcommonprefix.LongestCommonPrefix(s) //computes for the longest common prefix in a list of strings
	fmt.Println("Longest common prefix:", prefix)
}
