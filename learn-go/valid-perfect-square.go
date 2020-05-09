/**
Given a positive integer num, write a function which returns True if num is a perfect square else False.

Note: Do not use any built-in library function such as sqrt.

Example 1:

Input: 16
Output: true
*/

package exercise

func isPerfectSquare(num int) bool {

	if num < 1 {
		return false
	}
	if num == 1 {
		return true
	}

	// Initial guess with half as square of greater than half is not valid
	x := num / 2

	// using netwon iterative method
	// x is guess and b is num
	// i.e. x = 1/2(x+b/x) https://pages.mtu.edu/~shene/COURSES/cs201/NOTES/chap06/sqrt-1.html
	for x*x > num {
		x = (x + (num / x)) / 2
	}

	return x*x == num

}
