package main

import "fmt"

func longestPalindrome(s string) string {
	var maxLen int = 1
	var startIdx int = 0
	if len(s) <= 1 {
		return s
	}

	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
	}
	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			if s[i] == s[j] {
				if j-i+1 <= 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
				if dp[i][j] == true {
					if j-i+1 > maxLen {
						maxLen = j - i + 1
						startIdx = i
					}
				}
			} else {
				dp[i][j] = false
			}
		}
	}
	return s[startIdx : startIdx+maxLen]
}

func main() {
	res := longestPalindrome("babad")
	fmt.Println(res)
}
