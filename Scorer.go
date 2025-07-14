package main


func Scorer(choices []int) int {
	var correctAnswers = []int{1, 2, 1, 2, 1, 1, 2, 2, 1, 2, 1, 1, 2, 1, 1, 1, 1, 1, 0, 1}
	score := 0

	for i := range choices {
		if choices[i] == correctAnswers[i] {
			score++
		}
	}

	return score
}
