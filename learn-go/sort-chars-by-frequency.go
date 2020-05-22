/**
Given a string, sort it in decreasing order based on the frequency of characters.

Example 1:

Input:
"tree"

Output:
"eert"

Explanation:
'e' appears twice while 'r' and 't' both appear once.
So 'e' must appear before both 'r' and 't'. Therefore "eetr" is also a valid answer.
*/
package exercise

import (
	"sort"
	"strings"
)

type keyVal struct {
	Key   string
	Value int
}

func frequencySort(s string) string {
	charCount := make(map[string]int)
	res := ""
	for _, c := range s {
		charCount[string(c)]++
	}

	var sortedArr []keyVal
	for k, v := range charCount {
		sortedArr = append(sortedArr, keyVal{k, v})
	}

	sort.Slice(sortedArr, func(a, b int) bool {
		return sortedArr[a].Value > sortedArr[b].Value
	})

	for _, c := range sortedArr {
		res += strings.Repeat(c.Key, c.Value)
	}

	return res
}
