package leetcode_go

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
func findRepeatNumber2(nums []int) int {
	i := 0
	res := -1
	for i < len(nums) {
		temp := nums[i]
		if temp == i {
			i++
		} else if nums[temp] == temp{
			res = temp
			break
		}else{
			nums[i] = nums[temp]
			nums[temp] = temp
		}
	}
	return res
}
