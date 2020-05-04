package exercise

import (
	"strconv"
	"strings"
)

/**
Given a positive integer, output its complement number. The complement strategy is to flip the bits of its binary representation.
Input: 5
Output: 2
Explanation: The binary representation of 5 is 101 (no leading zero bits), and its complement is 010. So you need to output 2.
*/

func findComplement(num int) int {

	// convert given number to binary, then replace all 0's with 1's
	// convert replace string to int and XOR with input
	// it complements

	binaryNum := strconv.FormatInt(int64(num), 2)
	allOnes := strings.ReplaceAll(binaryNum, "0", "1")
	allOneInt, _ := strconv.ParseInt(allOnes, 2, 32)

	return int(allOneInt) ^ num
}
