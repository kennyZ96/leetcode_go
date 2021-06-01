package leetcode_go

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	max := 0
	for left < right {
		w := right - left
		h := 0
		if height[left] < height[right] {
			h = height[left]
			left++
		} else {
			h = height[right]
			right--
		}
		temp := w * h
		if temp > max {
			max = temp
		}
	}

	return max
}
