package stringmod

func CanConstruct(ransomNote string, magazine string) bool {
	// words := make(map[byte]int)

	// for _, el := range magazine {
	// 	if words[byte(el)] != 0 {
	// 		words[byte(el)] += 1
	// 	} else {
	// 		words[byte(el)] = 1
	// 	}

	// }

	// for _, el := range ransomNote {
	// 	if words[byte(el)] <= 0 {
	// 		return false
	// 	} else {
	// 		words[byte(el)] -= 1
	// 	}
	// }

	// return true
	ransome_note_nums := make([]int, 26)
	magazine_note_nums := make([]int, 26)

	for _, el := range ransomNote {
		ransome_note_nums[el-'a'] += 1
	}

	for _, el := range magazine {
		magazine_note_nums[el-'a'] += 1
	}

	for i := 0; i < 26; i++ {
		if ransome_note_nums[i] > magazine_note_nums[i] {
			return false
		}
	}
	return true
}
