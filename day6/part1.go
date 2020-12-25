package main

func countGroupPositiveAnswers(groupAnswers []string) int {
	answersSet := make(map[rune]bool)
	for _, groupMemberAnswer := range groupAnswers {
		for _, answerCode := range groupMemberAnswer {
			answersSet[answerCode] = true
		}
	}
	return len(answersSet)
}

func sumUpGroupsPositiveAnswers(groupsAnswers [][]string) (sum int) {
	for _, groupAnswers := range groupsAnswers {
		sum += countGroupPositiveAnswers(groupAnswers)
	}
	return sum
}
