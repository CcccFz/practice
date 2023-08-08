# -*- coding: utf-8 -*-

# 模式特点：在不破坏封装性的前提下捕获一个对象的内部状态，并在该对象之外保存这个状态，以后可以将对象恢复到这个状态。
# 程序实例：将Originator对象的状态封装成Memo对象保存在Caretaker内


class Originator(object):
    def __init__(self):
        self._state = None

    def show(self):
        print self._state

    def set_state(self, mem):
        self._state = mem.state

    def create_mem(self):
        return Memo(self._state)


class Memo(object):
    state = ''

    def __init__(self, state):
        self.state = state


class Caretaker(object):
    memo = ''


if __name__ == '__main__':
    o = Originator()
    o._state = 'on'
    o.show()

    c = Caretaker()
    c.memo = o.create_mem()

    o._state = 'off'
    o.show()

    o.set_state(c.memo)
    o.show()

