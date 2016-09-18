package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main() {
	nums := make([]int, 0, 10000000)
	for i := 0; i < 10000000; i++ {
		nums = append(nums, rand.Int()%10000000)
	}
	start := time.Now()
	nums = mergeSort(nums)
	end := time.Now()
	fmt.Printf("%f", end.Sub(start).Seconds())
}

func mergeSort(tab []int) []int {
	if len(tab) > 2 {
		left := tab[:len(tab) / 2]
		right := tab[len(tab) / 2:]
		left = mergeSort(left)
		right = mergeSort(right)
		return merge(left, right)
	} else  if len(tab) == 2 {
		if tab[0] > tab[1] {
			newTab := make([]int, 2, 2)
			newTab[0] = tab[1]
			newTab[1] = tab[0]
			return newTab
		}
		return tab
	} else {
		return tab
	}
}

func merge(left []int, right []int) []int {
	tab := make([]int, 0, len(left) + len(right))
	for i, j := 0, 0; i < len(left) || j < len(right); {
		if i < len(left) && j < len(right) {
			if left[i] < right[j] {
				tab = append(tab, left[i])
				i++
			} else {
				tab = append(tab, right[j])
				j++
			}
			continue
		}
		if i < len(left) {
			tab = append(tab, left[i:]...)
			break
		}
		if j < len(right) {
			tab = append(tab, right[j:]...)
			break
		}
	}
	return tab
}
