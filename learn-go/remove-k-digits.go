/**
Given a non-negative integer num represented as a string, remove k digits from the number so that the new number
is the smallest possible.

Note:
The length of num is less than 10002 and will be ≥ k.
The given num does not contain any leading zero.
Example 1:

Input: num = "1432219", k = 3
Output: "1219"
Explanation: Remove the three digits 4, 3, and 2 to form the new number 1219 which is the smallest.
*/
package exercise

import "regexp"

func removeKdigits(num string, k int) string {

	if len(num) == k {
		return "0"
	}

	result := make([]byte, 0)

	for i := range num {
		// find if current number is less than prev num in result
		for len(result) > 0 && k > 0 && result[len(result)-1] > num[i] {
			result = result[:len(result)-1]
			k--
		}

		result = append(result, num[i])
	}

	for k > 0 {
		result = result[:len(result)-1]
		k--
	}

	// convert to string
	byteToString := string(result)

	// remove leading zeroes
	regexMatch := regexp.MustCompile(`([^0].*)`)
	output := regexMatch.FindString(byteToString)

	if len(output) == 0 {
		return "0"
	}

	return output
}
