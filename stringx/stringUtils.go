package stringx

func StartsWith(s,sub string) bool{
	l := len(sub)
	if l > len(s) {
		return false
	}

	for i := 0; i < l; i++ {
		if s[i] != sub[i] {
			return false
		}
	}

	return true
}
