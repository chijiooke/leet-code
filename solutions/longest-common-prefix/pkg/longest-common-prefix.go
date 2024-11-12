package longestcommonprefix

import (
	"strings"
)

func LongestCommonPrefix(strs []string) string {

	var l string //longest common prefix

	var curPrefix string
	for _, str := range strs[0] {
		curPrefix = string(curPrefix) + string(str)

		for i := 0; i < len(strs); i++ {

			if !strings.HasPrefix(strs[i], curPrefix) {
				return l
			}
		}

		if len(l) < len(curPrefix) {
			l = curPrefix
		}
	}

	return l
}
