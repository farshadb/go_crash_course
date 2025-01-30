package main

import "fmt"

type Student struct {
	Name   string
	ID     string
	Scores []float64 // and slice of int
}

func PrintStudent(s Student) {
	fmt.Printf("The student name is %s and the his/her ID is %s.\n", s.Name, s.ID)
}

func studentAvgScore(s Student) float64 {
	sum := 0.0
	for _, s := range s.Scores {
		sum += s
	}
	return float64(sum) / float64(len(s.Scores))
}

func isStudentEligible(s Student) bool {
	if studentAvgScore(s) > 12 {
		fmt.Println("This sdudent is eligible")
		return true
	}
	return false
}

func main() {
	// create an instance of the student type
	st := Student{
		Name:   "Farshad Bolouri",
		ID:     "34135",
		Scores: []float64{18, 20, 15, 14, 17.5},
	}
	PrintStudent(st)
	fmt.Printf("Student average score is %.2f\n", studentAvgScore(st))
	isStudentEligible(st)
}
