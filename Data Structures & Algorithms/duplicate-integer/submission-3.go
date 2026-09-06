func hasDuplicate(nums []int) bool {
	sort.Ints(nums)
	if len(nums) <=1{
		return false
	}
	for i:=0;i<len(nums)-1;i++{
		if nums[i] == nums[i+1]{
			return true
		}
	}
	return false
}
