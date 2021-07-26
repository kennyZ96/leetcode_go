//this package is used for kennyZ's review of sort algorithm
package main

import "fmt"

//冒泡排序，时间复杂度为 O(n^2),空间复杂度为O(1),排序稳定，原地排序。
//使用flag标识一次排序中是否发生过逆序，以此来略微提高效率
func bubbleSort(nums []int) []int {
	length := len(nums)
	for i := 0; i < length-1; i++ {
		flag := false
		for j := 0; j < length-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				flag = true
			}
		}
		if flag == false {
			return nums
		}
	}
	return nums
}

//选择排序，时间复杂度为 O(n^2),空间复杂度为O(1),排序不稳定，原地排序。
func selectionSort(nums []int) []int {
	length := len(nums)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i; j < length; j++ {
			if nums[j] < nums[min] {
				min = j
			}
		}
		nums[i], nums[min] = nums[min], nums[i]
	}
	return nums
}

//插入排序，时间复杂度为 O(n^2),空间复杂度为O(1),排序稳定，原地排序。
func insertionSort(nums []int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i; j > 0 && nums[j] < nums[j-1]; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}
	return nums
}

//
func main() {
	nums1 := []int{19, 2, 22, 1, 25, 5, 34, 32, 65, 32, 7, 7, 2, 5756, 899, 6, 46}
	insertionSort(nums1)
	fmt.Println(nums1)
	nums2 := []int{1, 2, 3, 4, 6, 6, 7, 9, 8}
	insertionSort(nums2)
	fmt.Println(nums2)
}
