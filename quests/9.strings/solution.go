package strings

type TextStats struct {
	ByteLength int
	RuneCount  int
}

func AnalyzeText(s string) TextStats {
	return TextStats{
		ByteLength: len([]byte(s)),
		RuneCount: len([]rune(s)),
	}
}

func RuneFrequencies(s string) map[rune]int {
	runes := make(map[rune]int)
	for _, ch := range s {
		if _, exists := runes[ch]; !exists {
			runes[ch] = 0
		}
		runes[ch]++;
	}
	return runes
}

func FirstRunePosition(s string, target rune) int {
	for i, ch := range s {
		if ch == target {
			return i
		}
	}
	return -1
}

func ExtractRunes(s string) []rune {
	return []rune(s)
}

func HasOnlyASCII(s string) bool {
	for _, ch := range s {
		if ch > 127 {
			return false
		}
	}
	return true
}
