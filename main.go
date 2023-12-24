package main

import (
	"fmt"
)

/*

You are given N counters, initially set to 0, and you have two possible operations on them:

increase(X) − counter X is increased by 1,
max counter − all counters are set to the maximum value of any counter.
A non-empty array A of M integers is given. This array represents consecutive operations:

if A[K] = X, such that 1 ≤ X ≤ N, then operation K is increase(X),
if A[K] = N + 1 then operation K is max counter.
For example, given integer N = 5 and array A such that:

    A[0] = 3
    A[1] = 4
    A[2] = 4
    A[3] = 6
    A[4] = 1
    A[5] = 4
    A[6] = 4
the values of the counters after each consecutive operation will be:

    (0, 0, 1, 0, 0)
    (0, 0, 1, 1, 0)
    (0, 0, 1, 2, 0)
    (2, 2, 2, 2, 2)
    (3, 2, 2, 2, 2)
    (3, 2, 2, 3, 2)
    (3, 2, 2, 4, 2)
The goal is to calculate the value of every counter after all operations.

*/

func main() {

	a := []int{3, 4, 4, 6, 1, 4, 4}
	//counterLogic(5, a)

	fmt.Println(Solution(5, a))
}

func counterLogic(N int, A []int) []int {

	counters := make([]int, N)
	fmt.Println(counters)

	for i, v := range A {

		fmt.Printf("%d -- %d \n", i, v)
		if v < N {
			counters[v-1] = counters[v-1] + 1
			fmt.Printf("  i  %d ----- %v ", i, counters)
		} else {
			maxC := max(counters)
			setMaxCounterValue(counters, maxC)

		}
	}

	fmt.Println(counters)

	return counters
}

func setMaxCounterValue(counters []int, maxC int) {
	fmt.Println("maxc ", maxC)

	fmt.Println("before increasing counter ", counters)
	for i, _ := range counters {

		counters[i] = maxC
	}
	fmt.Println("after increasing counter ", counters)

}

func max(c []int) int {
	for i := 0; i <= len(c)-1; i++ {
		for j := i + 1; j <= len(c)-1; j++ {
			if c[i] < c[j] {
				c[i], c[j] = c[j], c[i]
			}
		}
	}
	return c[0]
}

func Solution(N int, A []int) []int {
	// Create a result array with N elements initialized to 0
	result := make([]int, N)

	// Iterate through each element in A
	for _, value := range A {
		// Decrement the value by 1 to adjust for 0-based indexing
		index := value - 1

		// Ensure the index is within bounds to avoid potential errors
		if 0 <= index && index < N {
			// Increment the counter at the corresponding index in the result array
			result[index]++
		}
	}

	return result
}
