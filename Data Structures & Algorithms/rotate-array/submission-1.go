func rotate(nums []int, k int) {
  ln := len(nums)
  temp := k%ln
  l,r := 0,len(nums)-1
  for l<r{
	nums[l],nums[r] = nums[r],nums[l]
	l++
	r--
  }
  m,n := 0,temp-1
  for m<n{
	nums[m],nums[n] = nums[n],nums[m]
	m++
	n--
  }
  i,j := temp,len(nums)-1
  for i<j{
	nums[i],nums[j] = nums[j],nums[i]
	i++
	j--
  } 
}
