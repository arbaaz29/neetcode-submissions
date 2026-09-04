func characterReplacement(s string, k int) int {
	l:=0
	res := 0
	cnt := make(map[byte]int)
	maxFreq := 0
	for r:=0;r<len(s);r++{
		cnt[s[r]]++
		if cnt[s[r]] > maxFreq{
			maxFreq = cnt[s[r]]
		}
		wlen := r-l+1
		replace := wlen - maxFreq
		if replace > k{
			cnt[s[l]]--
			l++
		}
		wlen = r-l+1
		if wlen > res{
			res = wlen
		}
	}
	return res
}
