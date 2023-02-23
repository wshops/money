package money

func isValidStrAmount(s string) bool {
	dotCount := 0
	// check empty or blank string
	if len(s) == 0 || s == " " || s == "" {
		return false
	}
	for i := 0; i < len(s); i++ {
		// check empty or blank string
		if s[i] == ' ' {
			return false
		}
		// check if there is more than one dot
		if s[i] == '.' {
			dotCount++
			if dotCount > 1 {
				return false
			}
		}
		// check if there is a non-digit character
		if s[i] != '.' && (s[i] < '0' || s[i] > '9') {
			return false
		}
	}
	return true
}
