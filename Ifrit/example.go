/**
 * Accepted
 * 124 ms 	18200 KB
 * 2019-07-11
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
	buf := make([]byte, 0, 13*200000)
	scanner.Buffer(buf, 13*200000)

	if scanner.Scan() {
		line = scanner.Text()
	}

	strList := strings.Split(line, " ")

	list := make([]int, n)
	for i := 0; i < n; i++ {
		list[i], _ = strconv.Atoi(strList[i])
	}

	var pos []int
	for i := 0; i < n; i++ {
		if list[i] == 0 {
			pos = append(pos, i)
		}
	}

	distance := make([]int, n)
	posIndex := 0
	posLen := len(pos)
	for i := 0; i < n; i++ {
		leftPos := pos[posIndex]
		var rightPos int
		if posIndex+1 < posLen {
			rightPos = pos[posIndex+1]
		} else {
			rightPos = leftPos
		}

		d1 := (int)(math.Abs((float64)(leftPos - i)))
		d2 := (int)(math.Abs((float64)(rightPos - i)))
		if d2 <= d1 {
			distance[i] = d2
			if posIndex+1 < posLen {
				posIndex++
			}
		} else {
			distance[i] = d1
		}
	}

	output := fmt.Sprint(distance)
	output = strings.Replace(output, "[", "", -1)
	output = strings.Replace(output, "]", "", -1)

	fmt.Println(output)
}
