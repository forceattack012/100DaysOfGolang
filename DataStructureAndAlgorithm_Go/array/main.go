package main

import (
	"errors"
	"fmt"
	"github/forceattack/array/findMaxMin"
)

func main() {
	arr := [5]int{2, -1, 5, -1, 5}
	fmt.Printf("max is %d and index is %d \n", findMaxMin.FindMaxUseMath(arr), findMaxMin.FindIndexMax(arr))
	fmt.Printf("max is %d and index is %d\n", findMaxMin.FindMax(arr), findMaxMin.FindIndexMin(arr))
	fmt.Printf("min is %d \n", findMaxMin.FindMin(arr))
	fmt.Printf("min is %d \n", findMaxMin.FindMinUseMath(arr))

	MultipleTwoEachElement(&arr)

	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}

	fmt.Println("")
	//arr = [5]int{5, 8, 9, 20, 1}
	sortArr(&arr)

	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}

	fmt.Println("")

	reverseTwoPointer(&arr)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}

	fmt.Println("")

	if v, err := find(arr, 4); err == nil {
		fmt.Printf("find is : %d", v)
	}

	fmt.Println("")

	sumP := prefixSumArray(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", sumP[i])
	}

	fmt.Println("")

	fmt.Printf("sum %d \n", sumArray(arr))

	fmt.Println("")
	a := 1
	b := 2
	swap(&a, &b)

	fmt.Printf("a : %d b : %d\n", a, b)

	arr2D := [][]int{
		{4, 2, 3},
		{1, 8, 3},
		{9, 2, 3},
	}

	for i := 0; i < len(arr2D); i++ {
		fmt.Printf("Row : %d ", i)
		for j := 0; j < len(arr2D[0]); j++ {
			fmt.Printf(" %d ", arr2D[i][j])
		}
		fmt.Println()
	}

	fmt.Println()

	var sum2D int
	var j int
	k := len(arr2D)
	for i := 0; i < len(arr2D); i++ {
		sum2D += arr2D[j][j]
		sum2D += arr2D[i][k-1-i]
		fmt.Printf(" %d ", arr2D[j][j])
		fmt.Printf(" %d ", arr2D[i][k-1-i])
		j = i + 1
	}

	fmt.Printf("\n %d", sum2D)
}

func MultipleTwoEachElement(arr *[5]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * 2
	}
}

func sortArr(arr *[5]int) {
	indexMin := 0

	for i := 0; i < len(arr); i++ {
		indexMin = i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[indexMin] {
				indexMin = j
			}
		}
		swap(&arr[indexMin], &arr[i])
	}
}

func reverse(arr *[5]int) {
	revers := [5]int{}

	for i := len(arr) - 1; i >= 0; i-- {
		revers[len(arr)-1-i] = arr[i]
	}
	*arr = revers
}

func reverseTwoPointer(arr *[5]int) {
	start := 0
	end := len(arr) - 1

	for start < end {
		swap(&arr[start], &arr[end])
		start++
		end--
	}
}

func swap(a, b *int) {
	temp := *b
	*b = *a
	*a = temp
}

func find(arr [5]int, num int) (int, error) {
	for i := 0; i < len(arr); i++ {
		if arr[i] == num {
			return arr[i], nil
		}
	}
	return -1, errors.New("not found")
}

func sumArray(arr [5]int) int {
	var sum int

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	return sum
}

func prefixSumArray(arr [5]int) []int {
	sum := make([]int, len(arr))
	sum[0] = arr[0]

	for i := 1; i < len(arr); i++ {
		sum[i] = sum[i-1] + arr[i]
	}

	return sum
}
