func isAnagram(s string, t string) bool {
	if len(s)>len(t) || len(t)>len(s){
		return false
	}
	lw := strings.ToLower(s)
	rw := strings.ToLower(t)
	freq1 := make(map[byte]int)
	freq2 := make(map[byte]int)
	for i:=0;i<len(s);i++{
		freq1[lw[i]-'a']++
		freq2[rw[i]-'a']++
	}

	for i:=0;i<len(s);i++{
		if freq1[lw[i]-'a'] != freq2[lw[i]-'a']{
			return false
		}
	}
	return true
}
