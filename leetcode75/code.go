package leetcode75

/*
1768. Merge Strings Alternately
*/
// mergeAlternately merges two strings by alternating their characters.
// It preallocates a byte slice with the final capacity to avoid reallocations,
// appends two bytes per loop iteration for efficiency, and then appends
// any remaining characters from the longer string in bulk.
// Time complexity: O(n+m), Space complexity: O(n+m).
func MergeAlternately(word1 string, word2 string) string {
	i := 0
	r := make([]byte, 0, len(word1)+len(word2))

	for i < len(word1) && i < len(word2) {
		r = append(r, word1[i], word2[i])
		i++
	}

	return string(append(append(r, word1[i:]...), word2[i:]...))
}

/*
1071. Greatest Common Divisor of Strings
// gcdOfStrings returns the largest string that divides both input strings.
// It first checks if such a string is possible by verifying that str1+str2 == str2+str1.
// If not, it returns an empty string.
// If possible, it computes the GCD of the two string lengths using Euclid's algorithm,
// and returns the prefix of str1 with that GCD length, which is the greatest common divisor string.
*/
func GcdOfStrings(str1, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	g := gcd(len(str1), len(str2))
	return str1[:g]
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// 1431. Kids With the Greatest Number of Candies
// kidsWithCandies returns a boolean slice where each element indicates
// whether the corresponding kid can have the greatest number of candies
// after receiving all the extraCandies.
// It first finds the maximum number of candies among all kids,
// then uses a threshold (max - extraCandies) to determine in O(n) time
// and O(1) extra space if each kid meets or exceeds that maximum.
func KidsWithCandies(candies []int, extraCandies int) []bool {

	n := len(candies)
	res := make([]bool, n)

	max := candies[0]
	for i := 1; i < n; i++ {
		if candies[i] > max {
			max = candies[i]
		}
	}

	threshold := max - extraCandies

	for i := 0; i < n; i++ {
		res[i] = candies[i] >= threshold
	}
	return res
}

/*
605. Can Place Flowers
*/

func CanPlaceFlowers(flowerbed []int, n int) bool {
	if n <= 0 {
		return true
	}

	for i := 0; i < len(flowerbed); {
		if flowerbed[i] == 1 {
			i += 2
			continue
		}

		leftOK := i == 0 || flowerbed[i-1] == 0
		rightOK := i == len(flowerbed)-1 || flowerbed[i+1] == 0

		if leftOK && rightOK {
			flowerbed[i] = 1
			n--
			if n == 0 {
				return true
			}
			i += 2
		} else {
			i++
		}
	}

	return false
}

/*
345. Reverse Vowels of a String
*/

var vowel = func() (v [256]bool) {
	for _, c := range []byte("aeiouAEIOU") {
		v[c] = true
	}
	return
}()

func ReverseVowels(s string) string {
	b := []byte(s)
	i, j := 0, len(b)-1
	for i < j {
		for i < j && !vowel[b[i]] {
			i++
		}
		for i < j && !vowel[b[j]] {
			j--
		}
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
	return string(b)
}
