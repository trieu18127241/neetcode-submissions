func validPalindrome(s string) bool {
	l, r := 0, len(s)-1
	flagRemove := false
	for l < r {
		fmt.Println(string(s[l]), string(s[r]))
		if s[l] != s[r] {
			if flagRemove {
				return false
			}
			flagRemove = true
			if l+1 < r && s[l+1] == s[r] && s[l+2] == s[r-1] {
				l++
				fmt.Println("left increase: ", l, r, s[l], s[r])
				continue
			} else if l < r-1 && s[l] == s[r-1] && s[l+1] == s[r-2] {
				r--
				fmt.Println("right decrease: ", l, r, s[l], s[r])
				continue
			} else if l+1 == r || l == r-1 {
				fmt.Println("meet middle string")
				return true
			}
			fmt.Println("cannot remove 1 element")
			return false
		}
		l++
		r--
	}
	return true
}