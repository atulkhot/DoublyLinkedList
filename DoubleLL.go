package main

import (
	"fmt"
	"strings"
)

type DoublyLinkedList struct {
	top_sentinel    *Cell
	bottom_sentinel *Cell
}

type Cell struct {
	data string
	next *Cell
	prev *Cell
}

func make_doubly_linked_list() DoublyLinkedList {
	top_sentinel := Cell{}
	bottom_sentinel := Cell{next: &top_sentinel, prev: &top_sentinel}

	top_sentinel.next = &bottom_sentinel
	top_sentinel.prev = &bottom_sentinel

	list := DoublyLinkedList{top_sentinel: &top_sentinel, bottom_sentinel: &bottom_sentinel}
	return list
}

func (me *Cell) add_before(cell *Cell) {
	//me.prev.next = cell
	//cell.prev = me.prev
	//me.prev = cell
	//cell.next = me
	me.prev.add_after(cell)
	me.prev = cell
}

func (me *Cell) add_after(after *Cell) {
	after.next = me.next
	after.prev = me
	me.next.prev = after
	me.next = after
}

func (me *Cell) delete() *Cell {
	me.prev.next = me.next
	me.next.prev = me.prev
	return me
}

func (list *DoublyLinkedList) add_range(values []string) {
	for _, v := range values {
		cell := Cell{data: v}
		list.bottom_sentinel.add_before(&cell)
	}
}

func (list *DoublyLinkedList) to_string(separator string) string {
	sbuf := make([]string, 0)
	for p := list.top_sentinel.next; p != list.bottom_sentinel; p = p.next {
		sbuf = append(sbuf, p.data)
	}
	return strings.Join(sbuf[:], " ")
}

func (list *DoublyLinkedList) length() int {
	len := 0
	for p := list.top_sentinel.next; p != list.bottom_sentinel; p = p.next {
		len++
	}
	return len
}

func (list *DoublyLinkedList) is_empty() bool {
	return list.top_sentinel.next == list.bottom_sentinel
}

func (list *DoublyLinkedList) push_top(v string) {
	cell := Cell{data: v}
	list.top_sentinel.add_after(&cell)
}

func (list *DoublyLinkedList) enqueue(v string) {
	list.push_top(v)
}

func (list *DoublyLinkedList) dequeue() string {
	if list.is_empty() {
		panic("No item to dequeue - Empty queue")
	}
	elem := list.bottom_sentinel.prev.delete()
	return elem.data
}

func (list *DoublyLinkedList) pop_bottom() string {
	cell := list.bottom_sentinel.prev.delete()
	return cell.data
}

func (list *DoublyLinkedList) push_bottom(v string) {
	cell := Cell{data: v}
	list.bottom_sentinel.add_before(&cell)
}

func main() {
	// Test queue functions.
	fmt.Printf("*** Queue Functions ***\n")
	queue := make_doubly_linked_list()
	queue.enqueue("Agate")
	queue.enqueue("Beryl")
	fmt.Printf("%s ", queue.dequeue())
	queue.enqueue("Citrine")
	fmt.Printf("%s ", queue.dequeue())
	fmt.Printf("%s ", queue.dequeue())
	queue.enqueue("Diamond")
	queue.enqueue("Emerald")
	for !queue.is_empty() {
		fmt.Printf("%s ", queue.dequeue())
	}
	fmt.Printf("\n\n")

	// Test deque functions. Names starting
	// with F have a fast pass.
	fmt.Printf("*** Deque Functions ***\n")
	deque := make_doubly_linked_list()
	deque.push_top("Ann")
	deque.push_top("Ben")
	fmt.Printf("%s ", deque.pop_bottom())
	deque.push_bottom("F-Cat")
	fmt.Printf("%s ", deque.pop_bottom())
	fmt.Printf("%s ", deque.pop_bottom())
	deque.push_bottom("F-Dan")
	deque.push_top("Eva")
	for !deque.is_empty() {
		fmt.Printf("%s ", deque.pop_bottom())
	}
	fmt.Printf("\n")
}
