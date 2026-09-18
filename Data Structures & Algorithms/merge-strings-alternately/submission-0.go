func mergeAlternately(word1 string, word2 string) string {
	maxLen := len(word1)
	if len(word1) < len(word2) {
		maxLen = len(word2)
	}
	var result string
	for i := 0; i < maxLen; i++ {
		fmt.Println("index: ", i)
		if i < len(word1) {
			result += string(word1[i])
			fmt.Println("result:", result)
		}
		if i < len(word2) {
			result += string(word2[i])
			fmt.Println("result:", result)
		}
	}
	return result
}