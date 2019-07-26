/**
 * Accepted
 * 46 ms 	0 KB
 * 2019-07-12
 */
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanln(&n)

	var line string
	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 12*100000)
	scanner.Buffer(buf, 12*100000)

	if scanner.Scan() {
		line = scanner.Text()
	}

	strList := strings.Split(line, " ")

	list := make([]int, n)
	for i := 0; i < n; i++ {
		list[i], _ = strconv.Atoi(strList[i])
	}

	if n < 3 {
		fmt.Println("0")
		return
	}

	changes := [][]int{
		{0, 0},
		{1, 0},
		{0, 1},
		{1, 1},
		{-1, 0},
		{0, -1},
		{-1, -1},
		{1, -1},
		{-1, 1},
	}

	min := n + 1
	for _, change := range changes {
		//fmt.Print(change)
		first := list[0] + change[0]
		last := list[n-1] + change[1]

		diff := (float64)(last-first) / (float64)(n-1)
		if diff != math.Trunc(diff) {
			//fmt.Println()
			continue
		}

		count := 0
		for i, v := range list {
			listN := first + i*(int)(diff)
			//fmt.Printf(" -> %d %d", listN, v)
			listN -= v
			if listN > 1 || listN < -1 {
				count = -1
				break
			} else if listN != 0 {
				count++
			}
		}

		//fmt.Println()

		if min > count && count >= 0 {
			min = count
		}
	}

	if min == n+1 {
		fmt.Println("-1")
	} else {
		fmt.Println(min)
	}

}
