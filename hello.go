package main

import (
	"example/hello/others"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Starting the application...")

	name := "Stanley"

	var age int = time.Now().Year() - 2001

	fmt.Println("Hello, my name is", name, "and I am", age, "years old.")

	fmt.Println("The product of 3 and 4 is", multiply(3, 4))

	numArray := []int{21, 1, 23, 78, 25, 12, 74, 5, 9}

	numArraySorted := quicksort(numArray)

	fmt.Println(numArraySorted)

	searchValue := binarySearch(numArraySorted, 1)
	fmt.Println(searchValue)

	myLinkedList := LinkedList{}
	myLinkedList.InsertAtBegining(21)
	myLinkedList.InsertAtEnd(33)
	myLinkedList.insertAtPosition(1, 17)
	myLinkedList.deleteAtBegining()
	myLinkedList.deleteAtEnding()
	result := myLinkedList.getAllValues()
	fmt.Println("elements inside the linked list are ", result)

	others.MainForOthers()
}

func multiply(a, b int) int {
	return a * b
}

func quicksort(num []int) []int {
	if len(num) <= 0 {
		return num
	}

	middlePoint := len(num) / 2

	pivot := num[middlePoint]

	remaining := append(num[:middlePoint], num[middlePoint+1:]...)
	left := []int{}
	right := []int{}

	for i := 0; i < len(remaining); i++ {
		if remaining[i] < pivot {
			left = append(left, remaining[i])
		} else {
			right = append(right, remaining[i])
		}
	}

	return append(append(quicksort(left), pivot), quicksort(right)...)
}

func binarySearch(num []int, target int) bool {
	first := 0
	last := len(num) - 1

	for first <= last {
		middlePoint := (first + last) / 2
		if num[middlePoint] == target {
			return true
		} else if num[middlePoint] < target {
			first = middlePoint + 1
		} else {
			last = middlePoint - 1
		}
	}

	return false
}

type Node struct {
	value    int
	next     *Node
	previous *Node
}

type LinkedList struct {
	length int
	head   *Node
	tail   *Node
}

func (ll *LinkedList) InsertAtEnd(value int) {
	newNode := &Node{value: value}
	ll.length += 1
	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
	} else {
		ll.tail.next = newNode
		newNode.previous = ll.tail
		ll.tail = newNode
	}
}

func (ll *LinkedList) InsertAtBegining(value int) {
	newNode := &Node{value: value}
	ll.length += 1
	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
	} else {
		ll.head.previous = newNode
		newNode.next = ll.head
		ll.head = newNode
	}
}

func (ll *LinkedList) deleteAtBegining() {
	if ll.head == nil || ll.length == 0 {
		return
	}

	if ll.head.next == nil {
		ll.head = nil
		ll.tail = nil
		ll.length = 0
		return
	}
	ll.head = ll.head.next
	ll.head.previous = nil
	ll.length -= 1
}

func (ll *LinkedList) deleteAtEnding() {
	if ll.tail == nil || ll.length == 0 {
		return
	}

	if ll.tail.previous == nil {
		ll.tail = nil
		ll.head = nil
	} else {
		ll.tail = ll.tail.previous
		ll.tail.next = nil
	}
	ll.length -= 1
}

func (ll *LinkedList) insertAtPosition(position int, value int) {
	if position <= 0 {
		return
	}
	lengthOfList := ll.length
	if position > lengthOfList {
		return
	}
	newNode := &Node{value: value}

	if position == 1 {
		ll.InsertAtBegining(value)
		return
	}
	if position == lengthOfList {
		ll.InsertAtEnd(value)
		return
	}

	curr := ll.head
	counter := 1
	for curr != nil {
		if counter == position-1 {
			newNode.next = curr.next
			curr.next = newNode
			newNode.previous = curr
			if newNode.next != nil {
				newNode.next.previous = newNode
			}
			ll.length++
			return
		}
		counter++
		curr = curr.next
	}
}

func (ll *LinkedList) getAllValues() []int {
	listToReturn := []int{}
	if ll.length == 0 {
		return listToReturn
	}

	curr := ll.head
	for curr != nil {
		listToReturn = append(listToReturn, curr.value)
		curr = curr.next
	}

	return listToReturn
}
