func getConcatenation(nums []int) []int {
    n := len(nums)
	ln := 2*n
	ans := make([]int,ln)
	l :=0
	r:=0
	for l<len(nums)&&r<ln{
		ans[r] = nums[l]
		l = (l+1)%len(nums)
		r++
	}
	return ans
}
