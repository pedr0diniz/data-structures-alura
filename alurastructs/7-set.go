package alurastructs

import (
	"strings"
)

// Slice of DoubleLinkedLists
// Is firstly a slice because it gets O(1) insertion in first subgroup
// Subgroups are alphabetical here, ranging from A (0) to Z (25)
// Insertion in letter subgroups is also O(1), since it adds to the end as a DoubleLinkedList
// Given that a property of the Set is the fact that we cannot have duplicates
// The Contains method must be run before every single addition
// To simplify complexity, we need to read data as fast as possible
// That's why we have a slice of DoubleLinkedLists.
// In this case, we separate the DoubleLinkedLists alphabetically
// We access the DoubleLinkedList for each letter with an O(1) complexity
// We read the smaller amount of DoubleLinkedList elements with O(n) complexity
type Set struct {
	elements []DoubleLinkedList
}

// O(n) complexity, will run linearly for as long as we need to create DoubleLinkedLists to spread access.
// In this case, n caps at 26.
func CreateSet() Set {
	set := Set{}
	for i := 0; i < 26; i++ {
		set.elements = append(set.elements, DoubleLinkedList{})
	}

	return set
}

// O(1) complexity. We can directly access the slice index.
// We can also easily insert in the DoubleLinkedList
func (s *Set) Add(element string) {
	sliceIndex := s.getTableIndex(element)

	if !s.Contains(element) {
		s.elements[sliceIndex].AddToEnd(element)
	}

}

// O(1) due to the DoubleLinkedList
func (s *Set) Remove(element string) {
	sliceIndex := s.getTableIndex(element)
	if s.Contains(element) {
		s.elements[sliceIndex].RemoveFromTheEnd()
	}
}

// O(n) at the worst case scenario. Given that we only run through one DoubleLinkedList,
// The worst case scenario is achieved if this List contains all elements.
func (s *Set) Contains(element string) bool {
	sliceIndex := s.getTableIndex(element)
	return s.elements[sliceIndex].Contains(element)
}

// O(n) depending on how many DoubleLinkedLists we have.
// In this case, n caps at 26.
func (s *Set) getTableIndex(element string) int {
	return int(rune(strings.ToUpper(element)[0])) - 65
}

// O(n) complexity. Given that we're running two loops, it may seem like the complexity is O(n²).
// It remains O(n) due to the fact that there are no duplicate elements, so we only loop through each once.
func (s *Set) ShowContents() {
	for i := 0; i < len(s.elements); i++ {
		if s.elements[i].Length() > 0 {
			s.elements[i].ShowContents()
		}
	}
}
