package alurastructs

import (
	"fmt"
)

// Uses the LinkedList in the background.
type Stack struct {
	elements LinkedList
}

// O(1) complexity, easily accesses the first element because of its pointer
func (s *Stack) Insert(element string) {
	s.elements.AddToStart(element)
}

// O(1) complexity. Has direct access to the first element.
func (s *Stack) Remove() string {
	stackTop := s.elements.firstElement.Element
	s.elements.RemoveFromStart()
	return stackTop
}

// O(1) complexity. Has access to the size address.
func (s *Stack) IsEmpty() bool {
	return s.elements.Length() == 0
}

// O(n) complexity. Loops through the entire stack.
func (s *Stack) ShowContents() {
	current := s.elements.firstElement
	contents := ""
	for i := 0; i < s.elements.amountOfElements; i++ {
		contents += current.Element + ", "
		current = current.Next
	}

	fmt.Printf("contents: %v\n", contents[:len(contents)-2])
}
