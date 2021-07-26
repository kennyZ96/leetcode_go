package leetcode_go

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		k := len(nums) - 1
		temp := -1 * nums[i]
		for j := i + 1; j < len(nums); j++ {
			//j > i + 1,用于判定a==b的符合解，将i后第一个与i相同的j值进行+1枚举。
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k > j && (nums[k]+nums[j] > temp) {
				k--
			}
			if k == j {
				break
			}
			if nums[j]+nums[k] == temp {
				res = append(res, []int{nums[i], nums[j], nums[k]})
			}
		}
	}
	return res
}
