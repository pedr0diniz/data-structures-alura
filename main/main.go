package main

import (
	"fmt"

	"github.com/pedr0diniz/data-structures-alura/alurastructs"
)

func main() {
	// fmt.Println("Sequential storage and vector structs:")
	// vector()
	// fmt.Println()

	// fmt.Println("Linked List:")
	// linkedList()
	// fmt.Println()

	// fmt.Println("Double Linked List:")
	// doubleLinkedList()
	// fmt.Println()

	fmt.Println("Stack:")
	stack()
	fmt.Println()

	fmt.Println("Queue:")
	stack()
	fmt.Println()
}

func vector() {
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

func linkedList() {
	linkedList := alurastructs.LinkedList{}
	linkedList.AddToStart("Amanda")
	linkedList.AddToStart("Pedro")
	linkedList.AddToStart("Maru")

	linkedList.ShowContents()

	linkedList.AddToEnd("Jane")

	linkedList.ShowContents()

	linkedList.AddAtGivenIndex("Tia Ana", 1)
	linkedList.ShowContents()

	fmt.Printf("linkedList.GrabElement(2): %v\n", linkedList.GrabElementAtIndex(2))

	fmt.Printf("linkedList.Length(): %v\n", linkedList.Length())

	linkedList.RemoveFromStart()
	linkedList.ShowContents()
	fmt.Printf("linkedList.Length(): %v\n", linkedList.Length())

	linkedList.RemoveAtGivenIndex(2)
	linkedList.ShowContents()
	fmt.Printf("linkedList.Length(): %v\n", linkedList.Length())

	linkedList.RemoveFromTheEnd()
	linkedList.ShowContents()
	fmt.Printf("linkedList.Length(): %v\n", linkedList.Length())
}

func doubleLinkedList() {
	doubleLinkedList := alurastructs.DoubleLinkedList{}
	doubleLinkedList.AddToStart("Amanda")
	doubleLinkedList.AddToStart("Pedro")
	doubleLinkedList.AddToStart("Maru")

	doubleLinkedList.ShowContents()

	doubleLinkedList.AddToEnd("Jane")

	doubleLinkedList.ShowContents()

	doubleLinkedList.AddAtGivenIndex("Tia Ana", 1)
	doubleLinkedList.ShowContents()

	fmt.Printf("doubleLinkedList.GrabElement(2): %v\n", doubleLinkedList.GrabElementAtIndex(2))

	fmt.Printf("doubleLinkedList.Length(): %v\n", doubleLinkedList.Length())

	doubleLinkedList.RemoveFromStart()
	doubleLinkedList.ShowContents()
	fmt.Printf("doubleLinkedList.Length(): %v\n", doubleLinkedList.Length())

	doubleLinkedList.RemoveAtGivenIndex(2)
	doubleLinkedList.ShowContents()
	fmt.Printf("doubleLinkedList.Length(): %v\n", doubleLinkedList.Length())

	doubleLinkedList.RemoveFromTheEnd()
	doubleLinkedList.ShowContents()
	fmt.Printf("doubleLinkedList.Length(): %v\n", doubleLinkedList.Length())

	fmt.Printf("doubleLinkedList.Contains(\"Pedro\"): %v\n", doubleLinkedList.Contains("Pedro"))
	fmt.Printf("doubleLinkedList.Contains(\"Maru\"): %v\n", doubleLinkedList.Contains("Maru"))
}

func stack() {
	stack := alurastructs.Stack{}

	stack.Insert("Pedro")
	stack.ShowContents()

	stack.Insert("Amanda")
	stack.ShowContents()

	stack.Insert("Maru")
	stack.ShowContents()

	s1 := stack.Remove()
	fmt.Printf("s1: %v\n", s1)
	stack.ShowContents()

	s2 := stack.Remove()
	fmt.Printf("s2: %v\n", s2)
	stack.ShowContents()
}

func queue() {

}
