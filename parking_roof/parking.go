package main

import "sort"

func carParkingRoof(cars []int, k int) int {
	sort.Ints(cars)

	length := len(cars)
	min := cars[k-1] - cars[0] + 1
	for i := 0; i < length-k+1; i++ {
		temp := cars[i+k-1] - cars[i] + 1
		if temp < min {
			min = temp + 1
		}
	}

	return min
}
