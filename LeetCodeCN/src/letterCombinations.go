package main

import "fmt"

var mapping = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	result := make([]string, 0)
	return dfs(digits, "", &result)
}

func dfs(digits string, value string, result *[]string) []string {
	if len(digits) > 0 {
		firstDigit := string(digits[0])
		for _, v := range mapping[firstDigit] {
			dfs(string(digits[1:]), value+v, result)
		}
	} else {
		if len(value) > 0 {
			*result = append(*result, value)
		}
	}
	return *result
}

func main() {
	fmt.Println(letterCombinations("23"))
}

