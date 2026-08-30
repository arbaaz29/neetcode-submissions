func productExceptSelf(nums []int) []int {
	 n := len(nums)
	 res := make([]int, n)
	 pre := make([]int, n)
	 suff := make([]int, n)

	 pre[0],suff[n-1] = 1,1
	 
	 //Compute only the left side of the index, for every iteration pre[i] will store the product of elements that come before `i`th index
	 // i=1, since (i-1) -ve index is not valid we initialized i to 1
	 for i:=1;i<n;i++{
		pre[i] = pre[i-1]*nums[i-1]
	}
	 //Compute only the right side of the index, for every iteration pre[i] will store the product of elements that come after `i`th index
	// i= n-2, since (i+1) is out of bound we initialized i to n-2 (2nd last index)
	for i:=n-2;i>=0;i--{
		suff[i] = suff[i+1]*nums[i+1]
	}
	// For the final result we want to multiply the left of a index (pre) with the right (suff) of the index.
	for i:=0;i<n;i++{
		res[i] = pre[i]*suff[i]
	}
	return res
}
