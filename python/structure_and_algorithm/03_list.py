class SList(object):
    def __init__(self):
        self._items = []

    def __len__(self):
        return len(self._items)

    def is_empty(self):
        return len(self) == 0

    def __getitem__(self, idx):
        assert not self.is_empty() and 0 <= idx < len(self)
        return self._items[idx]

    def __setitem__(self, idx, item):
        assert not self.is_empty() and 0 <= idx < len(self)
        self._items[idx] = item

    def insert(self, idx, item):
        assert 0 <= idx <= len(self)
        if self.is_empty():
            assert idx == 0

        for i in range(idx, len(self))[::-1]:
            self._items[i + 1] = self._items[i]

        self._items[idx] = item

    def delete(self, idx):
        assert not self.is_empty() and 0 <= idx < len(self)

        try:
            for i in range(idx, len(self)):
                self._items[i] = self._items[i + 1]
        finally:
            del self._items[-1]


class LList(object):
    def __init__(self):
        self._head = None
        self._size = 0

    def __len__(self):
        return self._size

    def is_empty(self):
        return len(self) == 0

    def __getitem__(self, idx):
        assert not self.is_empty() and 0 <= idx < len(self)
        return self.get_node_at(idx).item

    def __setitem__(self, idx, item):
        assert not self.is_empty() and 0 <= idx < len(self)
        self.get_node_at(idx).item = item

    def insert(self, idx, item):
        assert 0 <= idx <= len(self)
        if self.is_empty():
            assert idx == 0

        if idx == 0:
            self._head = Node(item, self._head)
        else:
            cur = self.get_node_at(idx, True)
            cur.next = Node(item, cur.next)

    def delete(self, idx):
        assert not self.is_empty() and 0 <= idx < len(self)

        if idx == 0:
            self._head = self._head.next
        else:
            cur = self.get_node_at(idx, True)
            cur.next = cur.next.next

    def get_node_at(self, idx, prev=False):
        cur = self._head
        start = 1 if prev else 0

        for i in range(start, idx):
            cur = cur.next

        return cur


class Node(object):
    def __init__(self, item, next):
        self._item = item
        self._next = next

    @property
    def item(self):
        return self._item

    @property
    def next(self):
        return self._next

    @next.setter
    def next(self, node):
        self._next = node


# 另一种表示链式链表的简单形式
head = Node(1, Node(2, Node(3, Node(4, Node(5, None)))))


def print_llist(link):
    cur = link
    while cur:
        print(cur.item)
        cur = cur.next


# 单向链表的转置
def reverse_llist(link):
    cur = link.next
    link.next = None

    while cur:
        temp = cur.next

        cur.next = link
        link = cur
        cur = temp

    return link


# 链表相交
def cross_llist(head1, head2):
    if not head1 or not head2:
        return False

    l1, l2 = head1, head2
    length1, length2 = 1, 1

    while l1.next:
        l1 = l1.next
        length1 += 1

    while l2.next:
        l2 = l2.next
        length2 += 1

    if l1 != l2:
        return False

    l1, l2 = head1, head2

    if length1 > length2:
        for _ in range(length1 - length2):
            l1 = l1.next
    else:
        for _ in range(length2 - length1):
            l2 = l2.next

    # 是同一条链表
    if l1 == l2:
        raise Exception('是同一条链表')

    while l1.next:
        l1 = l1.next
        l2 = l2.next
        if l1 == l2:
            return True

    return False


# 链表有环
def check_circle_list(head):
    if head is None:
        return False

    slow = head.next
    if slow is None:
        return False

    fast = slow.next
    if fast is None:
        return False

    while slow and fast:
        if slow == fast:
            return True

        slow = slow.next
        fast = fast.next

        if fast:
            fast = fast.next

    return False


# 链表成对调换
def swap_pair(head):
    if head and head.next:
        next = head.next
        head.next = swap_pair(next.next)
        next.next = head
        return next
    return head


if __name__ == '__main__':
    print_llist(head)
    head = reverse_llist(head)
    print_llist(head)
