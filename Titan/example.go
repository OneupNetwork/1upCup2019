/**
 * Accepted
 * 62 ms 	4400 KB
 * 2019-07-10
 */
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanln(&n)

	var line string
	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 12*200000)
	scanner.Buffer(buf, 12*200000)

	if scanner.Scan() {
		line = scanner.Text()
	}

	strList := strings.Split(line, " ")

	dArray := make([]int, n)
	for i := 0; i < n; i++ {
		dArray[i], _ = strconv.Atoi(strList[i])
	}

	si := 0
	ei := len(dArray) - 1

	var max, sum1, sum3 uint64
	max = 0
	sum1 = uint64(dArray[si])
	sum3 = uint64(dArray[ei])
	for si < ei {
		if sum1 > sum3 {
			ei--
			sum3 += uint64(dArray[ei])
		} else if sum3 > sum1 {
			si++
			sum1 += uint64(dArray[si])
		} else {
			max = sum1
			si++
			ei--
			sum1 += uint64(dArray[si])
			sum3 += uint64(dArray[ei])
		}
	}

	fmt.Println(max)
}
