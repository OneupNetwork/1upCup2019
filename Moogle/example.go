/**
 * Accepted
 * 31 ms 	0 KB
 * 2019-07-11
 */
package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var k, n int
	fmt.Scanf("%d %d\n", &k, &n)

	var str string
	fmt.Scanln(&str)

	unsorted := strings.Split(str, "")
	sort.Strings(unsorted)
	sorted := strings.Join(unsorted, "")
	s := []rune(sorted)

	weight := s[0] - 96
	count := 1
	prev := 0
	for i := 1; i < k; i++ {
		if count == n {
			break
		}

		if s[i]-s[prev] < 2 {
			continue
		}

		weight += s[i] - 96
		count++
		prev = i
	}

	if count < n {
		fmt.Println("-1")
	} else {
		fmt.Println(weight)
	}
}
