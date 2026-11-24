package main

import "fmt"

// Partition returns every way to cut s into substrings that are all
// palindromes. At each position, try every palindromic prefix and recurse
// on the rest.
func Partition(s string) [][]string {
	out := [][]string{}
	cur := []string{}
	var walk func(start int)
	walk = func(start int) {
		if start == len(s) {
			out = append(out, append([]string{}, cur...))
			return
		}
		for end := start + 1; end <= len(s); end++ {
			if isPalindrome(s[start:end]) {
				cur = append(cur, s[start:end])
				walk(end)
				cur = cur[:len(cur)-1]
			}
		}
	}
	walk(0)
	return out
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Partition("aab"))
}
