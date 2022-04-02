package main

import "fmt"

var print = fmt.Println

func selectionSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		var min = i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		if i != min {
			var aux = nums[i]
			nums[i] = nums[min]
			nums[min] = aux
		}
	}
	return nums
}

func medianOfTwoSortedArrays(nums1 []int, nums2 []int) float64 {
	var nums []int

	nums = append(nums, nums1...)
	nums = append(nums, nums2...)

	selectionSort(nums)

	var isOddSize bool = len(nums)%2 == 0

	var size = len(nums)
	if isOddSize {
		var middle = size / 2
		var left = nums[middle-1]
		var right = nums[middle]
		var sum = left + right
		return float64(sum) / 2
	}
	var middle int = size / 2
	return float64(nums[middle])
}

func main() {
	var nums1 = []int{1, 2}
	var nums2 = []int{3, 4}
	var result = medianOfTwoSortedArrays(nums1, nums2)
	print(result)
}
