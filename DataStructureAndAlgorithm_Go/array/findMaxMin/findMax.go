package findMaxMin

import (
	"math"
)

// 1 find max number in array
func FindMax(arr [5]int) int {
	max := arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
}

//2 find max number in array using math {0, 4, 5, 4, 1}

func FindMaxUseMath(arr [5]int) int {
	max := math.MinInt

	for _, v := range arr {
		max = Max(v, max)
	}

	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

//3 find min number in array

func FindMin(arr [5]int) int {
	min := arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
	}

	return min
}

//4 find min number in array use math

func FindMinUseMath(arr [5]int) int {
	min := arr[0]

	for _, v := range arr {
		min = Min(v, min)
	}

	return min
}

func Min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

//5 find int of max number in array

func FindIndexMax(arr [5]int) int {
	index := 0
	max := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] >= max {
			max = arr[i]
			index = i
		}
	}

	return index
}

//6 find int of max number in array

func FindIndexMin(arr [5]int) int {
	index := 0
	min := arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] <= min {
			min = arr[i]
			index = i
		}
	}

	return index
}
