package tools

func StringInclude(s [2]string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}