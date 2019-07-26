/**
 * Accepted
 * 31 ms 	0 KB
 * 2019-07-12
 */
package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	fmt.Scanln(&n)

	list := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&list[i])
	}

	var steps []int
	last := -1
	for _, v := range list {
		if v == 1 && last != -1 {
			steps = append(steps, last)
		}

		last = v
	}

	steps = append(steps, last)
	output := fmt.Sprint(steps)
	output = strings.Replace(output, "[", "", -1)
	output = strings.Replace(output, "]", "", -1)

	fmt.Println(len(steps))
	fmt.Println(output)
}
