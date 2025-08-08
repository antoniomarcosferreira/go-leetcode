package main

import (
	"fmt"

	"go-leetcode/leetcode75"
)

func main() {
	fmt.Println("Start leetcode")

	fmt.Println(leetcode75.MergeAlternately("abc", "pqr"))
	fmt.Println(leetcode75.GcdOfStrings("ABCABC", "ABC"))
	fmt.Println(leetcode75.KidsWithCandies([]int{2, 3, 5, 1, 3}, 3))
}
