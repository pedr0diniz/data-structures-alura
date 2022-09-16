package alurastructs

import "fmt"

type DoubleLinkedNode struct {
	Element  string
	Previous *DoubleLinkedNode
	Next     *DoubleLinkedNode
}

type DoubleLinkedList struct {
	firstElement     *DoubleLinkedNode
	lastElement      *DoubleLinkedNode
	amountOfElements int
}

// O(1) complexity. The beginning of the list is always easily reachable.
// Better than the Vector.
func (dll *DoubleLinkedList) AddToStart(element string) {
	newNode := DoubleLinkedNode{
		Element: element,
	}

	if dll.amountOfElements == 0 {
		dll.firstElement = &newNode
		dll.lastElement = &newNode
	} else {
		newNode.Next = dll.firstElement
		dll.firstElement.Previous = &newNode
		dll.firstElement = &newNode
	}

	dll.amountOfElements++
}

// O(1) complexity. Given that we know who our last element is, we can point directly to it.
// Better than the Vector.
func (dll *DoubleLinkedList) AddToEnd(element string) {
	if dll.amountOfElements == 0 {
		dll.AddToStart(element)
		return
	}

	newNode := DoubleLinkedNode{
		Element:  element,
		Previous: dll.lastElement,
	}

	newNode.Previous = dll.lastElement
	dll.lastElement.Next = &newNode
	dll.lastElement = &newNode

	dll.amountOfElements++
}

// O(n) complexity. Has to go through adll elements to show them.
func (dll *DoubleLinkedList) ShowContents() {
	current := dll.firstElement
	contents := "List contents: "
	for i := 0; i < dll.amountOfElements; i++ {
		contents += current.Element + ", "
		current = current.Next
	}

	fmt.Printf("contents: %v\n", contents)
}

// O(n) complexity. Might have to loop through adll elements in order to grab at the given position.
func (dll *DoubleLinkedList) grab(position int) *DoubleLinkedNode {
	if position < 0 || position >= dll.amountOfElements {
		fmt.Printf("Cannot grab at position %d. Try grabbing between positions 0 and %d", position, dll.amountOfElements-1)
		return nil
	}
	current := dll.firstElement

	for i := 0; i < position; i++ {
		current = current.Next
	}

	return current
}

// O(n) complexity because of the O(n) complexity of the grab method.
func (dll *DoubleLinkedList) AddAtGivenIndex(element string, position int) {
	if position == 0 {
		dll.AddToStart(element)
		return
	} else if position == dll.amountOfElements {
		dll.AddToEnd(element)
		return
	}
	previousNode := dll.grab(position - 1)
	nextNode := previousNode.Next

	newNode := DoubleLinkedNode{
		Element:  element,
		Next:     nextNode,
		Previous: previousNode,
	}

	nextNode.Previous = &newNode
	previousNode.Next = &newNode
	dll.amountOfElements++
}

// O(n) complexity because of the grab method.
// Worse than the Vector.
func (dll *DoubleLinkedList) GrabElementAtIndex(position int) string {
	return dll.grab(position).Element
}

// O(1) complexity. Finds the value directly.
func (dll *DoubleLinkedList) Length() int {
	return dll.amountOfElements
}

// O(1) complexity. Finds the first element directly.
// Better than the Vector.
func (dll *DoubleLinkedList) RemoveFromStart() {
	if dll.amountOfElements == 0 {
		fmt.Println("Cannot remove from empty linked list!")
		return
	}

	dll.firstElement = dll.firstElement.Next

	dll.amountOfElements--

	if dll.amountOfElements == 0 {
		dll.lastElement = nil
	}
}

// O(1) complexity. The last element is easily accessible here.
// Better than the Linked List and the Vector.
func (dll *DoubleLinkedList) RemoveFromTheEnd() {
	if dll.amountOfElements == 0 {
		fmt.Println("Cannot remove from empty linked list!")
		return
	} else if dll.amountOfElements == 1 {
		dll.RemoveFromStart()
		return
	}

	dll.lastElement = dll.lastElement.Previous
	dll.lastElement.Next = nil

	dll.amountOfElements--
}

// O(n) complexity because of the grab method. Same complexity as the Vector.
func (dll *DoubleLinkedList) RemoveAtGivenIndex(position int) {
	if dll.amountOfElements == 0 {
		fmt.Println("Cannot remove from empty linked list!")
		return
	} else if position < 0 || position >= dll.amountOfElements {
		fmt.Printf("Cannot remove from position %d. Try grabbing between positions 0 and %d", position, dll.amountOfElements-1)
		return
	} else if position == 1 {
		dll.RemoveFromStart()
		return
	} else if position == dll.amountOfElements-1 {
		dll.RemoveFromTheEnd()
		return
	}

	previousNode := dll.grab(position - 1)
	currentNode := previousNode.Next
	nextNode := currentNode.Next

	previousNode.Next = nextNode
	nextNode.Previous = previousNode

	dll.amountOfElements--
}

// O(n) complexity. Potentially has to loop through all nodes.
func (dll *DoubleLinkedList) Contains(element string) bool {
	current := dll.firstElement
	for {
		if current.Element == element {
			return true
		}

		current = current.Next
		if current == nil {
			return false
		}
	}
}
