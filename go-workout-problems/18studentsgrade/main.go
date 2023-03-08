package main

import "fmt"

func main() {
	var written, lab, assignment float64
	fmt.Print("Enter the marks scored by the students:\nWritten Test:")
	fmt.Scan(&written)
	fmt.Print("Lab Exams:")
	fmt.Scan(&lab)
	fmt.Print("Assignments:")
	fmt.Scan(&assignment)
	grade(written, lab, assignment)
}
func grade(writ float64, lab float64, ass float64) {
	studentGrade := (writ*70)/100 + (lab*20)/100 + (ass*10)/100
	fmt.Println("The overall grade of the student is ", studentGrade)
}
