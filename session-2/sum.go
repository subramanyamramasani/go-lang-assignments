package main

import (
	"fmt"
	"sort"
)

func add(num1 []int, num2 []int) []int {
	res := make([]int, len(num1))
	sort.Sort(sort.Reverse(sort.IntSlice(num2)))
	for s, _ := range num1 {
		res[s] = num1[s] + num2[s]
	}

	return res
}

func main() {
	fmt.Println("addition of two slices")
	num1 := []int{1, 2, 3, 4, 5, 6}
	num2 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("result", add(num1, num2))

}
