package ds

import (
	"fmt"
	"errors"
)

type LinkedList struct {
    Head *Node
    Tail *Node
    Size int
}

func (list *LinkedList) Insert(value string) {
    newNode := &Node{data: value, Next: nil}

    if list.Head == nil {
        list.Head = newNode
		list.Size = list.Size + 1
        return
    }

    current := list.Head
    for current.Next != nil {
        current = current.Next 
    }
    current.Next = newNode
	list.Size = list.Size + 1
}

func (list *LinkedList) InsertAt(position int, value string) error {
    newNode := &Node{data: value, Next: nil}

    if list.Head == nil {
		if position == 0 {
			list.Head = newNode
			list.Size = list.Size + 1
        	return nil
		} else {
			return errors.New("This position does not exist.")
		}
    }

    current := list.Head
	curr := 1
    for curr <= position {
		if curr == position {
			curNext := current.Next

			current.Next = newNode
			newNode.Next = curNext
			list.Size = list.Size + 1

			return nil
		}
		
		if current.Next == nil {
			return errors.New("This position does not exist.")
		}

        current = current.Next 
		curr++
    }

	return nil
}

func (list *LinkedList) Remove(value string) error {

	if list.Head == nil {
		return errors.New("Value does not occur.")
	} else if list.Head.data == value {
		list.Head = list.Head.Next
		list.Size = list.Size - 1
		return nil
	}

	prev := list.Head
	current := list.Head.Next
    for current.Next != nil {
		if current.data == value {
			prev.Next = current.Next
			list.Size = list.Size - 1
			return nil
		}

		prev = current
        current = current.Next 
    }

	return errors.New("Value does not occur.")
}

func (list *LinkedList) RemoveAll(value string) error {

	if list.Head == nil {
		return errors.New("Value does not occur.")
	}

	var prev *Node = nil
	current := list.Head

	removed := 0
    for current.Next != nil {
		rmv := false
		if current.data == value {

			if prev != nil {
				prev.Next = current.Next
			} else {
				list.Head = list.Head.Next
			}

			rmv = true
			removed++
		}

		if rmv {
			current = current.Next
		} else {
			prev = current
        	current = current.Next
		} 
    }

	if removed != 0 {
		list.Size -= removed
		return nil
	} else {
		return errors.New("Value does not occur.")
	}
	
}

func (list *LinkedList) RemoveAt(pos int) error {

    if list.Head == nil {
		return errors.New("This position does not exist.")
    } 

	if pos == 0 {
		list.Head = list.Head.Next
		list.Size = list.Size - 1
		return nil
	}

	prev := list.Head
	current := list.Head.Next

	curr := 1
    for current.Next != nil {
		if curr == pos {
			prev.Next = current.Next
			list.Size = list.Size - 1
			return nil
		}

		prev = current
		current = current.Next
		curr++
    }

	return errors.New("This position does not exist.")
}

func (list *LinkedList) IsEmpty() bool {
	return list.Size == 0
}

func (list *LinkedList) GetSize() int {
	return list.Size
}

func (list *LinkedList) Reverse() {

	if list.Head == nil || list.Head.Next == nil {
		return
	}

	var prev *Node = nil
	current := list.Head

	for current != nil {
		curNext := current.Next

		current.Next = prev
		prev = current 
		current = curNext

		if current == nil {
			list.Head = prev
		}
	}
}

func (list *LinkedList) PrintList() {

	if list.Head == nil {
		return 
	}
	
	current := list.Head
	for true {
		if current.Next == nil {
			fmt.Print(current.data)
			break
		}

		fmt.Print(current.data + " -> ")
		current = current.Next

		if current == nil {
			break 
		}

	}
	fmt.Println()
}
