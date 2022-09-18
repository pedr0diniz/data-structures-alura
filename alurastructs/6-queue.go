package alurastructs

import "fmt"

// Uses the LinkedList in the background.
type Queue struct {
	elements LinkedList
}

// O(1) complexity. We have direct access to the last element of our queue.
func (q *Queue) Enqueue(element string) {
	q.elements.AddToEnd(element)
	fmt.Printf("%s has joined the queue. The queue is now: %s\n", element, q.ShowContents())
}

// O(1) complexity. We have direct access to the first element of our queue.
func (q *Queue) Dequeue() string {
	current := q.elements.firstElement.Element
	q.elements.RemoveFromStart()

	fmt.Printf("It is %s's turn, they left the queue. The queue still has: %s\n", current, q.ShowContents())
	return current
}

func (q *Queue) Peek() string {
	return q.elements.firstElement.Element
}

// O(1) complexity. Has access to the size address.
func (q *Queue) IsEmpty() bool {
	return q.elements.Length() == 0
}

// O(n) complexity. Loops through the entire queue.
func (q *Queue) ShowContents() string {
	current := q.elements.firstElement
	contents := ""
	for i := 0; i < q.elements.amountOfElements; i++ {
		contents += current.Element + ", "
		current = current.Next
	}

	return contents[:len(contents)-2]
}
