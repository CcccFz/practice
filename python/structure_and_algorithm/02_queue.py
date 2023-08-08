class SQueue(object):
    def __init__(self):
        self._items = []

    def __len__(self):
        return len(self._items)

    def is_empty(self):
        return len(self) == 0

    def enqueue(self, item):
        return self._items.append(item)

    def dequeue(self):
        assert not self.is_empty()
        return self._items.pop(0)


class LQueue(object):
    def __init__(self):
        self._front = None
        self._rear = None
        self._size = 0

    def __len__(self):
        return self._size

    def is_empty(self):
        return len(self) == 0

    def enqueue(self, item):
        self._front.next = Node(item, None)
        self._size += 1

    def dequeue(self):
        assert not self.is_empty()
        self._size -= 1

        node = self._rear
        self._rear = self._rear.next
        return node.item


class Node(object):
    def __init__(self, item, next):
        self._item = item
        self._next = next

    @property
    def next(self):
        return self._next

    @next.setter
    def next(self, node):
        self._next = node
