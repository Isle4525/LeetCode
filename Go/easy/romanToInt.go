package easy

func RomanToInt(s string) int {
	roman := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	n := len(s)

	for i := 0; i < n; i++ {
		curentVal := roman[s[i]]

		if i+1 < n && curentVal < roman[s[i+1]] {
			result -= curentVal
		} else {
			result += curentVal
		}

	}

	return result

}
