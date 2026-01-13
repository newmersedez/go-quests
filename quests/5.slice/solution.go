package slice

func ProcessScores(scores []int) []int {
	res := make([]int, 0, len(scores))

	for _, score := range scores {
		if score < 0 || score > 100 {
			continue
		} else if score < 40 {
			score = 40
		}
		res = append(res, score)
	}

	if len(res) > 5 {
		for i, score := range res {
			res[i] = min(score + 5, 100)
		}
	}

	return res
}
