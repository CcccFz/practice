# -*- coding: utf-8 -*-

# 模式特点：当一个对象的内在状态改变时允许改变其行为，这个对象看起来像是改变了其类。
# 程序实例：描述一个程序员的工作状态，当需要改变状态时发生改变，不同状态下的方法实现不同。

import time

class State(object):
    def write_program(self):
        pass

class AmState(State):
    def write_program(self, work):        
        if work.hour > 12:
            work.set_state(PmState())
        else:
            print('code in Am !')


class PmState(State):
    def write_program(self, work):
        if work.hour > 18:
            work.hour = 9
            work.set_state(AmState())
        else:
            print('code in Pm !')

class Work(object):
    def __init__(self):
        self.hour = 9
        self.current = AmState()

    def set_state(self, st):
        self.current = st

    def write_program(self):
        self.current.write_program(self)

    def start(self):
        while True:
            time.sleep(1)
            self.hour += 1
            self.write_program()


if __name__ == '__main__':
    work = Work()
    work.start()

