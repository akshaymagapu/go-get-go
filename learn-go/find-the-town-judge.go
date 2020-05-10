/**
In a town, there are N people labelled from 1 to N.
There is a rumor that one of these people is secretly the town judge.

If the town judge exists, then:

The town judge trusts nobody.
Everybody (except for the town judge) trusts the town judge.
There is exactly one person that satisfies properties 1 and 2.
You are given trust, an array of pairs trust[i] = [a, b] representing that the person labelled a trusts the person labelled b.

If the town judge exists and can be identified, return the label of the town judge.  Otherwise, return -1.

Example 1:

Input: N = 2, trust = [[1,2]]
Output: 2
*/
package exercise

func findJudge(N int, trust [][]int) int {

	if len(trust) == 0 {
		return 1
	}

	degrees := make(map[int]int, N)

	// if vertex is connected to other vertex it is outward and it is getting connected it is inward
	for _, eachPerson := range trust {
		outward, inward := eachPerson[0], eachPerson[1]
		degrees[outward]--
		degrees[inward]++
	}

	for i := 1; i <= N; i++ {
		if degrees[i] == N-1 {
			return i
		}
	}

	return -1
}
