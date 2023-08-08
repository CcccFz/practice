class SStack(object):
    def __init__(self):
        self._items = []

    def __len__(self):
        return len(self._items)

    def is_empty(self):
        return len(self) == 0

    def peak(self):
        assert not self.is_empty()
        return self._items[-1]

    def pop(self):
        assert not self.is_empty()
        return self._items.pop()

    def push(self, item):
        self._items.append(item)


class LStack(object):
    def __init__(self):
        self._top = None
        self._size = 0

    def __len__(self):
        return self._size

    def is_empty(self):
        return len(self) == 0

    def peak(self):
        assert not self.is_empty()
        return self._top.item

    def pop(self):
        assert not self.is_empty()
        self._size -= 1

        node = self._top
        self._top = self._top.next
        return node.item

    def push(self, item):
        self._size += 1
        self._top = Node(item, self._top)


class Node(object):
    def __init__(self, item, next):
        self._item = item
        self._next = next

    @property
    def next(self):
        return self._next
