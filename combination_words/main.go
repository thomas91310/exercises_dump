package main

import "fmt"

func main() {
	hello := "abcdef"
	result := []string{}
	possibilities := []string{}
	for _, c := range hello {
		rest := getRest(hello, string(c))
		result = append(result, wordCombinations(rest, possibilities, string(c))...)
	}

	//remove duplicates
	mem := make(map[string]bool)
	actual := []string{}
	for _, sol := range result {
		exists := mem[sol]
		if exists {
			continue
		}
		actual = append(actual, sol)
		mem[sol] = true
	}
	fmt.Println(actual)
	fmt.Println(len(actual))
}

func getRest(a string, c string) []string {
	rest := []string{}
	for _, matched := range a {
		if string(matched) == string(c) {
			continue
		}
		rest = append(rest, string(matched))
	}
	return rest
}

func wordCombinations(rest []string, possibilities []string, currentWord string) []string {
	if len(rest) == 2 {
		possibilities = append(possibilities, currentWord+rest[0]+rest[1])
		possibilities = append(possibilities, currentWord+rest[1]+rest[0])
		return possibilities
	}
	for idx, c := range rest {
		tmp := make([]string, len(rest))
		copy(tmp, rest)
		newRest := append(tmp[:idx], tmp[idx+1:]...)
		// fmt.Printf("Calling with %v and %v and %v\n", newRest, possibilities, currentWord+string(c))
		possibilities = wordCombinations(newRest, possibilities, currentWord+string(c))
	}
	return possibilities
}

// a b c d 						a				b				c				d
//							 b  c 	d		a	c 	d	 	a 	b	d		a	b	c
//						c	d
