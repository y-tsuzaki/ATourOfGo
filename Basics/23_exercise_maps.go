package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	str_arr := strings.Fields(s)
	str_counts := make(map[string]int)
	for  _, v := range str_arr {
		str_counts[v]++
	}
	
	return str_counts
}

func main() {
	wc.Test(WordCount)
}
