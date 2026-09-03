func lengthOfLongestSubstring(s string) int {
	seen := make(map[byte]int)
	l:=0
	res := 0
	for r:=0;r<len(s);r++{
		if idx, found:= seen[s[r]]; found{
			l = max(idx+1,l)
		}
		seen[s[r]] = r
		if r-l+1 > res{
			res = r-l+1
		}
	}
	return res
}
