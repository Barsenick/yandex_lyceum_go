package main

type Student struct {
	name            string
	solvedProblems  int
	scoreForOneTask float64
	passingScore    float64
}

func (s Student) IsExcellentStudent() (is bool) {
	return float64(s.solvedProblems)*s.scoreForOneTask >= s.passingScore
}
