package main

import "fmt"

func main() {

}

// {1,2,3,4,5,6,7}
func LeftRotationArray(nums []int, n int) []int {
	var temp []int

	for i := 0; i < n; i++ {
		temp = nums[1:]
		temp = append(temp, nums[0])
		nums = temp

		fmt.Printf("num %d %d\n", i, nums)
	}

	return nums
}

// {1,2,3,4,5,6,7}
func RightRotationArray(nums []int, n int) []int {
	var temp []int

	for i := 0; i < n; i++ {
		temp = nums[:len(nums)-1]
		temp = append([]int{nums[len(nums)-1]}, temp...)
		nums = temp
	}

	fmt.Printf("num %d \n", nums)
	return nums
}

//input [1,2,3,4,5,6,7] , n = 2
//Reverse start 0 to size-1 [7,6,5,4,3,2,1]
//Reverse start size-2 to size [7,6,5,4,3,1,2]
//Reverse start 0 to size-2 [3,4,5,6,7,1,2]
//output [3,4,5,6,7,1,2]

func LeftRotationArrayByReverse(nums []int, n int) []int {
	temp := reverse(nums, 0, len(nums)-1)
	temp = reverse(temp, len(nums)-n, len(nums)-1)
	temp = reverse(temp, 0, len(nums)-1-n)
	return temp
}

func reverse(nums []int, start int, end int) []int {
	for start < end {
		temp := nums[start]
		nums[start] = nums[end]
		nums[end] = temp
		start++
		end--
	}

	return nums
}
