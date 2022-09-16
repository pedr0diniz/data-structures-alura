package alurastructs

import "fmt"

type Student struct {
	Name string
}

// O(1) complexity. Compares values that can be directly accessed.
func (s *Student) equals(anotherStudent Student) bool {
	return s.Name == anotherStudent.Name
}

type Vector struct {
	students         []Student
	amountOfstudents int
}

// Able to initialize vector without exposing private fields outside the package
func BuildVector(students []Student, amountOfstudents int) Vector {
	return Vector{
		students:         students,
		amountOfstudents: amountOfstudents,
	}
}

func (v *Vector) AddToNextEmptyIndex(s Student) {
	// Naive approach, O(n) complexity
	// Execution time will increase the bigger the students array is.

	// for i := 0; i < len(v.students); i++ {
	// 	if v.students[i].Name == "" {
	// 		v.students[i] = s
	// 		break
	// 	}
	// }

	// O(1), constant time complexity.
	// Will always find the next index right away.
	v.students[v.amountOfstudents] = s
	fmt.Printf("\nAdded %s to the list of students at position %d!", s.Name, v.amountOfstudents)
	v.amountOfstudents++
}

// O(n) complexity. Will have to go through the whole array to shift everyone forward.
func (v *Vector) AddAtGivenIndex(s Student, index int) {
	if index < 0 || index > v.amountOfstudents {
		fmt.Println("Cannot add to index", index, "because it is out of reach.")
		return
	}

	// Shifts the array elements one slot forward
	for i := v.amountOfstudents - 1; i >= index; i-- {
		v.students[i+1] = v.students[i]
	}
	// Another way to shift the array forward is to use the append function
	// v.students = append(v.students[:index+1], v.students[index:]...)
	// v.students[index] = s

	v.students[index] = s
	v.amountOfstudents++
}

// O(1) complexity. The student is returned straight away.
func (v *Vector) Grab(index int) Student {
	if index < 0 || index >= v.amountOfstudents {
		fmt.Println("Cannot grab at index", index, "because it is out of reach.")
		return Student{}
	}

	student := v.students[index]
	fmt.Println("\nFound", student, "at index", index)
	return student
}

// O(n) complexity. Will have to go through the whole array to shift everyone backwards.
func (v *Vector) RemoveAtGivenIndex(index int) {
	if index < 0 || index >= v.amountOfstudents {
		fmt.Println("Cannot remove at index", index, "because it is out of reach.")
		return
	}

	for i := index; i < v.amountOfstudents; i++ {
		v.students[i] = v.students[i+i]
	}

	v.amountOfstudents--
}

// O(n) complexity. Complexity grows as our array grows here.
func (v *Vector) Contains(s Student) bool {
	for i := 0; i < v.amountOfstudents; i++ {
		if s.equals(v.students[i]) {
			fmt.Println("Student", s.Name, "found in the students list at index", i)
			return true
		}
	}

	fmt.Println("Student", s.Name, " could not be found in the students list.")
	return false
}

// O(1) complexity. We always return the size straight away.
func (v *Vector) Size() int {
	return v.amountOfstudents
}
