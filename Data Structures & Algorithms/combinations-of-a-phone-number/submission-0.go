func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	phone := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var res []string
	var backtrack func(index int, path []byte)

	backtrack = func(index int, path []byte) {
		if index == len(digits) {
			res = append(res, string(path))
			return
		}

		letters := phone[digits[index]]

		for i := 0; i < len(letters); i++ {
			path = append(path, letters[i])
			backtrack(index+1, path)
			path = path[:len(path)-1] 
		}
	}

	backtrack(0, []byte{})
	return res
}