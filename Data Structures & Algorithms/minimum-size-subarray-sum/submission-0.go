func minSubArrayLen(target int, nums []int) int {
	l, total := 0,0
	res:=len(nums)+1

	for r:=0;r<len(nums);r++{
		total +=nums[r]
		for total >= target{
			if r-l+1 < res {
				res = r-l+1
			}
			total -=nums[l]
			l++
		}
	}
	if res == len(nums)+1{
		return 0
	}
	return res
}
