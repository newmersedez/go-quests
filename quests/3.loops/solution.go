package loops

func SumEvenNumbers(n int) int {
	res := 0

	if n <= 0 {
		return res
	}

	for number := range (n + 1) {
		if number % 2 == 0 {
			res += number
		}
	}

	return res
}


func KeepOnlyConsonants(strs []string) []string {
	res := []string{}
	
	if len(strs) == 0 {
		return res
	}

	for _, str := range strs {
		newStr := make([]rune, 0, len(str))

		for _, ch := range str {
			switch ch {
			case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
				continue
			default:
				newStr = append(newStr, ch)
			}
		}

		if len(newStr) > 0 {
			res = append(res, string(newStr))
		}
 	}

	return res
}
