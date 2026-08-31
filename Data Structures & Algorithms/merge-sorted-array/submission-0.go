func merge(nums1 []int, m int, nums2 []int, n int) {
	j := m
	for i:=0;i<n;i++{
		nums1[j]=nums2[i]
		j++
	}
	sort.Ints(nums1)
}
