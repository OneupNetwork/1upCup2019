/**
 * Accepted
 * 31 ms 	0 KB
 * 2019-07-07
 */
package main

import (
	"fmt"
	"sort"
	"strings"
)

func SortString(s string) string {
	sa := strings.Split(s, "")
	sort.Strings(sa)
	return strings.Join(sa, "")
}

func SwapItem(stringArray []string, i int) []string {
	tmp := stringArray[i]
	stringArray[i] = stringArray[i-1]
	stringArray[i-1] = tmp
	return stringArray
}

func SearchItemIndex(stringArray []string, s string, startIndex int) int {
	l := len(stringArray)
	for i := startIndex; i < l; i++ {
		if stringArray[i] == s {
			return i
		}
	}
	return -1
}

func main() {
	var n int
	var s, t string
	fmt.Scanln(&n)
	fmt.Scanln(&s)
	fmt.Scanln(&t)

	if s == t {
		fmt.Println("0")
		return
	}

	sortedS := SortString(s)
	sortedT := SortString(t)

	if sortedS != sortedT {
		fmt.Println("-1")
		return
	}

	sArray := strings.Split(s, "")
	tArray := strings.Split(t, "")

	var result []int

	for k, v := range tArray {
		if sArray[k] == v {
			continue
		}

		index := SearchItemIndex(sArray, v, k+1)
		for i := index; i > k; i-- {
			sArray = SwapItem(sArray, i)
			result = append(result, i)
		}
	}

	output := fmt.Sprint(result)
	output = strings.Replace(output, "[", "", -1)
	output = strings.Replace(output, "]", "", -1)

	fmt.Println(len(result))
	fmt.Println(output)
}
