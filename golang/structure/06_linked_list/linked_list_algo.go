package _6_linked_list

// 升序
func Merge(l1, l2 *LinkedList) *LinkedList {
	if l1 == nil || l1.head == nil || l1.head.next == nil {
		return l2
	}
	if l2 == nil || l2.head == nil || l2.head.next == nil {
		return l1
	}

	l := NewLinkedList()
	cur := l.head
	cur1 := l1.head.next
	cur2 := l2.head.next

	for cur1 != nil && cur2 != nil {
		if cur1.val.(int) < cur2.val.(int) {
			cur.next = cur1
			cur1 = cur1.next
		} else {
			cur.next = cur2
			cur2 = cur2.next
		}

		cur = cur.next
	}

	if cur1 != nil {
		cur.next = cur1
	} else {
		cur.next = cur2
	}

	return l
}

// O(N)
func (l *LinkedList) Reverse() {
	if l == nil || l.head == nil || l.head.next == nil {
		return
	}

	var pre *Node = nil

	cur := l.head.next
	for cur != nil {
		tmp := cur.next
		cur.next = pre
		pre = cur
		cur = tmp
	}

	l.head.next = pre
}

func (l *LinkedList) DeleteBottomN(n uint) {
	if n == 0 {
		return
	}

	fast := l.head
	for i := 0; uint(i) < n && fast != nil; i++ {
		fast = fast.next
	}

	if fast == nil {
		return
	}

	slow := l.head
	for fast != nil {
		fast = fast.next
		slow = slow.next
	}

	slow.next = slow.next.next
}

func (l *LinkedList) HasCircle() bool {
	if l == nil || l.head == nil || l.head.next == nil {
		return false
	}

	fast := l.head
	slow := l.head

	for fast != nil && fast.next != nil {
		fast = fast.next.next
		slow = slow.next

		if fast == slow {
			return true
		}
	}

	return false
}

func (l *LinkedList) FindMiddleNode() *Node {
	if l == nil || l.head == nil || l.head.next == nil {
		return nil
	}

	fast := l.head
	slow := l.head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow
}
