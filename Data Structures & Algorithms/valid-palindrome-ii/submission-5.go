func validPalindrome(s string) bool {
	isPalindrome := func(s string) bool {
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
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			rightCheck := isPalindrome(s[l:r])
			leftCheck := isPalindrome(s[l+1:r+1])
			return rightCheck || leftCheck
		}
		l++
		r--
	}
	return true
}