package main

func mapGroupAnswersCount(groupAnswers []string) map[rune]int {
	groupAnswersCount := make(map[rune]int)
	for _, groupMemberAnswer := range groupAnswers {
		for _, answerCode := range groupMemberAnswer {
			groupAnswersCount[answerCode]++
		}
	}
	return groupAnswersCount
}

func countGroupUnanimousAnswers(groupAnswers []string) (count int) {
	groupAnswersCount := mapGroupAnswersCount(groupAnswers)
	for _, answerCount := range groupAnswersCount {
		if answerCount == len(groupAnswers) {
			count++
		}
	}
	return count
}

func sumUpGroupsUnanimousAnswers(groupsAnswers [][]string) (sum int) {
	for _, groupAnswers := range groupsAnswers {
		sum += countGroupUnanimousAnswers(groupAnswers)
	}
	return sum
}
