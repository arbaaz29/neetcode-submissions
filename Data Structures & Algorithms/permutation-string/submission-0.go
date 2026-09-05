func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2){
		return false
	}
	s1cnt := make([]int,26)
	s2cnt := make([]int,26)
	match := 0
	// Initialize the count arrays with the respective counts
	for i:=0; i<len(s1);i++{
		s1cnt[s1[i]-'a']++
		s2cnt[s2[i]-'a']++
	}
	// The number of matches represent the matching elements between an array of a to z
	// if s1cnt[1] != s2cnt[1] then match won't be incremented
	for i:=0; i<26;i++{
		if s1cnt[i] == s2cnt[i]{
			match++
		}
	}
	//lower limit of the window
	l:=0
	for r:=len(s1);r<len(s2);r++{
			// if during the previous enum, the matches come out to be equal to 26 (because there are only 26 alphabets)
			if match == 26{
				return true
			}

			//increase the right limit of the window
			idx := s2[r]-'a'
			//since we have seen the current element increase its count by 1
			s2cnt[idx]++
			//if the elements count is equal increase the match
			if s1cnt[idx] == s2cnt[idx]{
				match++
			} else if s1cnt[idx]+1 == s2cnt[idx]{  // if the elements are unequal then lower the match as there are no matches to s1 in the current s2 window
				match--
			}
			// compare the elements at lower limit
			idx = s2[l]-'a'
			//drop the previous element, since we need to match the frequency of the alphabets within the current window limit. i.e. the subarray should be consecutive 
			s2cnt[idx]--
			// if the count is equal increment the match
			if s1cnt[idx] == s2cnt[idx]{
				match++
			}else if s1cnt[idx]-1 == s2cnt[idx]{
				// if the nums are unequal decrement match as there are no matches to s1 in the current s2 window
				match--
			}
			l++
			// increment the lower limit by one
	}
	// the match should always come out to 26 if the substring exists
	return match == 26
}
