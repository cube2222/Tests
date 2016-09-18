package main

import (
	"time"
	"fmt"
	"math/rand"
)

func main() {
	nums := make([]int, 0, 10000000)
	for i := 0; i < 10000000; i++ {
		nums = append(nums, rand.Int()%10000000)
	}
	start := time.Now()
	nums = Quicksort(nums)
	end := time.Now()
	fmt.Printf("%f", end.Sub(start).Seconds())
}

func Quicksort(tab []int) []int {
	if len(tab) <= 1 {
		return tab
	}
	if len(tab) == 2 {
		if tab[0] > tab[1] {
			return []int{tab[1], tab[0]}
		}
		return tab
	}
	left, right := partition(tab, len(tab)-1)
	left = Quicksort(left)
	right = Quicksort(right)
	return append(append(left, tab[len(tab)-1]), right...)
}

func partition(tab []int, midIndex int) ([]int, []int) {
	left := make([]int, 0, len(tab)/2)
	right := make([]int, 0, len(tab)/2)
	for index, item := range tab {
		if index != midIndex {
			if item < tab[midIndex] {
				left = append(left, item)
			} else {
				right = append(right, item)
			}
		}
	}
	return left, right
}