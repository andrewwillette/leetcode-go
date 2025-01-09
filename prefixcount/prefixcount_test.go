package prefixcount

// # 2185
// https://leetcode.com/problems/counting-words-with-a-given-prefix
func prefixCount(words []string, pref string) int {
	p := []rune(pref)
	count := 0
	for i := 0; i < len(words); i++ {
		w := []rune(words[i])
		if len(w) < len(p) {
			continue
		}
		match := true
		for j := 0; j < len(p); j++ {
			if w[j] != p[j] {
				match = false
				break
			}
		}
		if match {
			count++
		}
	}
	return count
}
