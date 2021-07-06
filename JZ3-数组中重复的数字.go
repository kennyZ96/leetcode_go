package main

//数组中重复的数字
//method 1 hash
func findRepeatNumber(nums []int) int {
	temp := make(map[int]int)
	for _, item := range nums {
		_, ok := temp[item]
		if ok {
			return item
		} else {
			temp[item]++
		}
	}
	return -1
}
