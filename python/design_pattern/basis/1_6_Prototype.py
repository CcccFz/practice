# -- coding: utf-8 --

import copy


class WorkExp:
    def __init__(self):
        self.place = None
        self.year = 0

class Resume:
    def __init__(self):
        self.name = 0
        self.age = 0
        self.work = WorkExp()

    def setInfo(self, *args):
        self.name, self.age, self.work.place, self.work.year = args

    def Display(self):
        print self.name, self.age, self.work.place, self.work.year

    def Clone(self):
        return  self

if __name__ == '__main__':
    a = Resume()
    a.setInfo('xiaoming', 18, 'shanghai', 2)
    a.Display()
    print '*' * 30

    b = a.Clone()
    b.setInfo('xiaoyu', 19, 'beijing', 3)
    a.Display()
    b.Display()
    print '*' * 30

    c = copy.copy(a)
    c.setInfo('xiaoma', 20, 'shenzhen', 4)
    a.Display()
    b.Display()
    c.Display()
    print '*' * 30

    d = copy.deepcopy(a)
    d.setInfo('xiaofan', 24, 'chendu', 8)
    a.Display()
    b.Display()
    c.Display()
    d.Display()