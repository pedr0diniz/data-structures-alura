package main

import (
	"fmt"

	"github.com/pedr0diniz/data-structures-alura/alurastructs"
)

func main() {
	startVector()
}

func startVector() {

	// Creating students
	s1 := alurastructs.Student{
		Name: "João",
	}
	s2 := alurastructs.Student{
		Name: "José",
	}

	studentsVector := alurastructs.BuildVector(make([]alurastructs.Student, 5), 0)

	studentsVector.AddToNextEmptyIndex(s1)
	studentsVector.AddToNextEmptyIndex(s2)

	s3 := alurastructs.Student{
		Name: "Kobra",
	}

	studentsVector.Grab(1)
	studentsVector.Grab(111)

	studentsVector.AddAtGivenIndex(s3, 1)
	studentsVector.AddAtGivenIndex(s3, 111)
	fmt.Println("studentsVector:", studentsVector)

	studentsVector.RemoveAtGivenIndex(1)
	studentsVector.RemoveAtGivenIndex(111)

	fmt.Println("studentsVector:", studentsVector)
}
