func majorityElement(nums []int) []int {
	count := make(map[int]int)
	res := []int{}
	for _,v:=range nums{
		count[v]++
	}
	for v, c := range count {
		if c > len(nums)/3 {
			res = append(res, v)
		}
	}
	return res
}
