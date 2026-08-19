func isPalindrome(s string) bool {
	reg := regexp.MustCompile("[^0-9a-zA-Z]+")
	s = reg.ReplaceAllString(s, "")
	s = strings.ToLower(strings.TrimSpace(s))
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}