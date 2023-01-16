package stringmod

func FirstUniqChar(s string) int {

	if len(s) == 1 {
		return 0
	}

	check := -1

	for i := 0; i < len(s); i++ {
		if check != -1 {
			return check
		}
		for j := 0; j < len(s); j++ {
			if i != j {

				check = i
				if s[i] == s[j] {
					check = -1
					break
				}
			}

		}
	}

	return check
}
