type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	var res strings.Builder
	for _, str := range strs {
		//append string length i.e. if 1st element in the array is hello, the lenght will be 4
		res.WriteString(strconv.Itoa(len(str)))
		//this will be used as a delimiter
		//"4"#
		res.WriteByte('#')
		//iterate through the array and append each string following the above pattern using string builder 
		//"4"#"Hello""5"#"World"
		res.WriteString(str)
	}
	return res.String()
}

func (s *Solution) Decode(encoded string) []string {
	res := []string{}
	// this pointer is used to denote the start of the string
	i := 0
	for i < len(encoded) {
		//this pointer is used to search for the delimiter
		j := i
		for encoded[j] != '#' {
			j++
		}
		// once we hit the delimiter we use the slice that represents the length of a respective string and convert it to integer
		length, _ := strconv.Atoi(encoded[i:j])
		// now we increment the first pointer to point to the start of 1st string
		i = j + 1
		// we create a string array (res) and select the slice where i represents the start of the string and i+length represents the end of the string
		res = append(res, encoded[i:i+length])
		// increment the 1st pointer to point towards the next string patterned 4#Hello
		i += length
	}
	// return the string array
	return res
}
