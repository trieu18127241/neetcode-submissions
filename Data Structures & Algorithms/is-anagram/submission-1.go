import (
	"cmp"
	"slices"
)
func isAnagram(s string, t string) bool {
	runeS, runeT := []rune(s), []rune(t)
	slices.SortFunc(runeS, func(i, j rune) int {
		return cmp.Compare(i, j)
	})
	slices.SortFunc(runeT, func(i, j rune) int {
		return cmp.Compare(i, j)
	})
	if string(runeS) != string(runeT) {
		return false
	}
	return true
}