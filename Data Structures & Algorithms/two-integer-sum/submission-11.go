func twoSum(nums []int, target int) []int {
    freq := make(map[int]int)
	for idx:=0;idx<len(nums);idx++{
		freq[nums[idx]] = idx
	}
	rem := 0
	i :=0
	for i<len(nums){
		rem = target - nums[i]
		if  freq[rem] !=i && freq[rem] !=0{
			return []int{i,freq[rem]}
		}
		i++
	}
	return []int{}
}
