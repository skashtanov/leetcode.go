func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func findLengthOfLCIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	res, current := 0, 0
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] <= nums[i-1] {
			current = i
		}
		res = max(res, i-current+1)
	}
	return res
}

