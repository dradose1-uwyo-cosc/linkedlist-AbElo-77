// Abdalla Elokely
// COSC 3750
// 2/5/2026

package main

import (
	"fmt"
	"linkedlist-AbElo-77/ds"
)

func main() {

	fmt.Println("---------------------")
	fmt.Println("Testing General Insertions and Deletions")
	fmt.Println("---------------------")
	fmt.Println()

	insert_removelist := &ds.LinkedList{}
	insert_removelist.InsertAt(0, "A")
	insert_removelist.Insert("A")
	insert_removelist.Insert("A")
	insert_removelist.Insert("B")
	insert_removelist.Insert("C")
	insert_removelist.Insert("D")
	insert_removelist.Insert("E")
	insert_removelist.PrintList()
	fmt.Println("-------------")

	insert_removelist.RemoveAt(4)
	insert_removelist.PrintList()
	fmt.Println("-------------")
	fmt.Println("Size after RemoveAt(4):", insert_removelist.GetSize())
	fmt.Println("-------------")
	insert_removelist.RemoveAll("A")
	insert_removelist.PrintList()
	fmt.Println("-------------")	
	fmt.Println("Size after RemoveAll(A):", insert_removelist.GetSize())

	fmt.Println("-------------")
	fmt.Println("Testing isEmpty():", insert_removelist.IsEmpty())
	fmt.Println("-------------")

	curSize := insert_removelist.GetSize()
	for i := 0; i < curSize; i++ {
		insert_removelist.RemoveAt(0)
		insert_removelist.PrintList()
		fmt.Println("-------------")
	}
	fmt.Println("Testing isEmpty():", insert_removelist.IsEmpty())
	fmt.Println("-------------")
	fmt.Println("Thus, size must be:", insert_removelist.GetSize())

	fmt.Println()
	fmt.Println("---------------------")
	fmt.Println("Testing Reversal")
	fmt.Println("---------------------")
	fmt.Println()

	reversedlist := &ds.LinkedList{}
	reversedlist.Insert("A")
	reversedlist.Insert("C")
	reversedlist.InsertAt(1, "B")
	reversedlist.Insert("Tail")
	reversedlist.PrintList()

	fmt.Println("------ First Reversal ------")
	reversedlist.Reverse()
	reversedlist.PrintList()

	fmt.Println("------ Second Reversal ------")
	reversedlist.Reverse()
	reversedlist.PrintList()

	fmt.Println()
	fmt.Println("---------------------")
	fmt.Println("Testing Error Messages")
	fmt.Println("---------------------")
	fmt.Println()

	errorlist := &ds.LinkedList{}
	errorlist.Insert("A")
	errorlist.Insert("B")
	errorlist.Insert("A")
	errorlist.Insert("B")
	errorlist.Insert("A")
	errorlist.Insert("B")

	err := errorlist.InsertAt(8, "Out of Bounds")
	fmt.Println("The Insert returns error:",  err)
	fmt.Println("-------------")

	err = errorlist.Remove("Not in list")
	fmt.Println("The Remove returns error:",  err)
	fmt.Println("-------------")

	err = errorlist.RemoveAll("Not in list")
	fmt.Println("The RemoveAll returns error:",  err)
	fmt.Println("-------------")

	err = errorlist.RemoveAt(8)
	fmt.Println("The RemoveAt returns error:",  err)

	fmt.Println()
	fmt.Println("---------------------")
	fmt.Println("Testing Stack")
	fmt.Println("---------------------")
	fmt.Println()

	stack := &ds.Stack{}
	stack.Push("first")
	stack.Push("second")
	stack.Push("third")

	data, _ := stack.Pop()
	fmt.Println("Popped:", data)
	fmt.Println("-------------")

	data, bool := stack.Pop()
	fmt.Println("Worked?", bool)
	fmt.Println("Yes, popped:", data)
	fmt.Println("-------------")

	data, bool = stack.Pop()
	fmt.Println("Worked?", bool)
	fmt.Println("Yes, popped:", data)
	fmt.Println("-------------")

	data, bool = stack.Pop()
	fmt.Println("Worked?", bool)
	fmt.Println("-------------")

	fmt.Println()
	fmt.Println("---------------------")
	fmt.Println("Testing Queue")
	fmt.Println("---------------------")
	fmt.Println()

	queue := &ds.Queue{}
	queue.Push("first")
	queue.Push("second")
	queue.Push("third")

	data, _ = queue.Pop()
	fmt.Println("Popped:", data)
	fmt.Println("-------------")

	data, err = queue.Pop()
	fmt.Println("Error?", err)
	fmt.Println("No, then popped:", data)
	fmt.Println("-------------")

	data, err = queue.Pop()
	fmt.Println("Error?", err)
	fmt.Println("No, then popped:", data)
	fmt.Println("-------------")
	
	data, err = queue.Pop()
	fmt.Println("Error?", err)
	fmt.Println("-------------")

}
