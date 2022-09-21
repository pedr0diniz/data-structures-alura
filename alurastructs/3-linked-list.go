package alurastructs

import "fmt"

type Node struct {
	Element string
	Next    *Node
}

type LinkedList struct {
	firstElement     *Node
	lastElement      *Node
	amountOfElements int
}

// O(1) complexity. The beginning of the list is always easily reachable.
// Better than the Vector.
func (ll *LinkedList) AddToStart(element string) {
	node := Node{
		Element: element,
		Next:    ll.firstElement,
	}

	ll.firstElement = &node

	if ll.amountOfElements == 0 {
		ll.lastElement = ll.firstElement
	}

	ll.amountOfElements++
}

// O(1) complexity. Given that we know who our last element is, we can point directly to it.
// Better than the Vector.
func (ll *LinkedList) AddToEnd(element string) {
	if ll.amountOfElements == 0 {
		ll.AddToStart(element)
		return
	}

	node := Node{
		Element: element,
	}

	ll.lastElement.Next = &node
	ll.lastElement = &node

	ll.amountOfElements++
}

// O(n) complexity. Has to go through all elements to show them.
func (ll *LinkedList) ShowContents() {
	current := ll.firstElement
	contents := "List contents: "
	for i := 0; i < ll.amountOfElements; i++ {
		contents += current.Element + ", "
		current = current.Next
	}

	fmt.Printf("contents: %v\n", contents)
}

// O(n) complexity. Might have to loop through all elements in order to grab at the given position.
func (ll *LinkedList) grab(position int) *Node {
	if position < 0 || position >= ll.amountOfElements {
		fmt.Printf("Cannot grab at position %d. Try grabbing between positions 0 and %d", position, ll.amountOfElements-1)
		return nil
	}
	current := ll.firstElement

	for i := 0; i < position; i++ {
		current = current.Next
	}

	return current
}

// O(n) complexity because of the O(n) complexity of the grab method.
func (ll *LinkedList) AddAtGivenIndex(element string, position int) {
	if position == 0 {
		ll.AddToStart(element)
		return
	} else if position == ll.amountOfElements {
		ll.AddToEnd(element)
		return
	}
	previousNode := ll.grab(position - 1)

	newNode := Node{
		Element: element,
		Next:    previousNode.Next,
	}

	previousNode.Next = &newNode
	ll.amountOfElements++
}

// O(n) complexity because of the grab method.
// Worse than the Vector.
func (ll *LinkedList) GrabElementAtIndex(position int) string {
	return ll.grab(position).Element
}

// O(1) complexity. Finds the value directly.
func (ll *LinkedList) Length() int {
	return ll.amountOfElements
}

// O(1) complexity. Finds the first element directly.
// Better than the Vector.
func (ll *LinkedList) RemoveFromStart() {
	if ll.amountOfElements == 0 {
		fmt.Println("Cannot remove from empty linked list!")
		return
	}

	ll.firstElement = ll.firstElement.Next

	ll.amountOfElements--

	if ll.amountOfElements == 0 {
		ll.lastElement = nil
	}
}

// O(n) complexity because of the grab method.
// The worst scenario is always triggered here, because we have to get the second last element.
// Same complexity as the Vector.
func (ll *LinkedList) RemoveFromTheEnd() {
	if ll.amountOfElements == 0 {
		fmt.Println("Cannot remove from empty linked list!")
		return
	}

	//
	newLastElement := ll.grab(ll.amountOfElements - 2)
	newLastElement.Next = nil
	ll.lastElement = nil
	ll.amountOfElements--
}

// O(n) complexity because of the grab method. Same complexity as the Vector.
func (ll *LinkedList) RemoveAtGivenIndex(position int) {
	if ll.amountOfElements == 0 {
		fmt.Println("Cannot remove from empty linked list!")
		return
	}

	if position < 0 || position >= ll.amountOfElements {
		fmt.Printf("Cannot remove from position %d. Try grabbing between positions 0 and %d", position, ll.amountOfElements-1)
		return
	}

	previousNode := ll.grab(position - 1)
	previousNode.Next = ll.grab(position).Next

	ll.amountOfElements--
}

// O(n) complexity. Potentially has to loop through all nodes.
func (ll *LinkedList) Contains(element string) bool {
	current := ll.firstElement
	for {
		if current == nil {
			return false
		}

		if current.Element == element {
			return true
		}

		current = current.Next
	}
}
