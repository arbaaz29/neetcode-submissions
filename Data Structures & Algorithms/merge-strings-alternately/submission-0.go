func mergeAlternately(word1 string, word2 string) string {

	var result strings.Builder

	i := 0

	for i < len(word1) && i < len(word2) {
		result.WriteByte(word1[i])
		result.WriteByte(word2[i])
		i++
	}

	if i < len(word1) {
		result.WriteString(word1[i:])
	}

	if i < len(word2) {
		result.WriteString(word2[i:])
	}

	return result.String()
}
