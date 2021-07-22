package leetcode_go

/*在一个长度为n+1的数组里的所有数字都在 1~n 的范围内，所以数组中至少有一个数字是重复的，
请找出数组中 任意一个重复的数字，但不能修改输入的数组。*/
func findDuplication(nums []int) int {
	if len(nums) < 1 {
		return -1
	}
	small := 1
	big := len(nums) - 1
	dup := -1
	for small < big {
		mid := (small + big) / 2
		count := countRange(nums, small, mid)
		if count > (mid - small + 1) {
			dup = mid
			big = mid
		} else {
			dup = mid + 1
			small = mid + 1
		}

	}
	return dup
}
func countRange(nums []int, small int, big int) (count int) {
	count = 0
	for _, v := range nums {
		if v <= big && v >= small {
			count++
		}
	}
	return
}
//func main() {
//	nums := []int{1, 2, 3, 4, 5, 6, 7, 6, 5}
//	print(findDuplication(nums))
//}
