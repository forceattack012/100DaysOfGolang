package main

func main() {

}

func BinarySearch(nums []int, target int) int {
	start := 0
	end := len(nums) - 1

	for start < end {
		pivot := (start + end) / 2

		if nums[pivot] == target {
			return pivot
		} else if nums[pivot] < target {
			start = pivot + 1
		} else {
			end = pivot - 1
		}
	}
	return -1
}
