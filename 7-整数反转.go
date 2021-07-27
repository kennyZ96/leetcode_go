package leetcode_go

import "math"

func reverse(x int) int {
	reversed := 0
	for x != 0 {
		if reversed < math.MinInt32/10 || reversed > math.MaxInt32/10 {
			//if temp:=int32(reversed);(temp*10)/10 != temp{
			//if int32(reversed*10)/10 != int32(reversed){
			return 0
		}
		tailNum := x % 10
		reversed = reversed*10 + tailNum
		x = x / 10
	}
	return reversed
}
