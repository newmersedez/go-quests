package generics

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// Push adds a value to the end of the list.
func (lst *List[T]) Push(v T) {
	newElemPtr := &element[T]{
		next:	nil,
		val:	v,
	}
	if lst.head == nil {
		lst.head = newElemPtr
		lst.tail = newElemPtr;
	} else {
		lst.tail.next = newElemPtr
		lst.tail = newElemPtr
	}
}

// Pop removes the last element and returns it.
func (lst *List[T]) Pop() (T, bool) {
	if lst.head == nil {
		return *new(T), false
	}

	value := lst.tail.val
	
	if lst.head == lst.tail {
		lst.head = nil
		lst.tail = nil
	} else {
		curr := lst.head
		for curr.next != lst.tail {
			curr = curr.next
		}
		curr.next = nil
		lst.tail = curr
	}

	return value, true

}

// AllElements returns a slice of all elements in the list.
func (lst *List[T]) AllElements() []T {
	elements := make([]T, 0, 10)

	head := lst.head
	for head != nil {
		elements = append(elements, head.val)
		head = head.next
	}

	return elements
}
