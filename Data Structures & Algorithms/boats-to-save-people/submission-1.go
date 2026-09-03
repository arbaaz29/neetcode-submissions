func numRescueBoats(people []int, limit int) int {
sort.Ints(people)
l,r := 0, len(people)-1
cnt := 0
	for l<=r{
		diff := limit-people[r]
		r--
		cnt++
		if l <=r && diff >= people[l]{
			l++
		}
	}
	return cnt
}
