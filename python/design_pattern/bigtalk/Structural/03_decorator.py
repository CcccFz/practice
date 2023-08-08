# -*- coding: utf-8 -*-

# 模式特点：动态地为对象增加额外的职责。
# 程序实例：展示一个人一件一件穿衣服的过程。

class Person(object):
    def __init__(self, name):
        self.name = name

    def show(self):
        print 'dressed %s' % self.name

class Decorator(object):
    def __init__(self):
        self.dt = None

    def decorate(self, dt):
        self.dt = dt

class Tshirt(Decorator):
    def show(self):
        print 'Tshirt, ',
        self.dt.show()

class Trouser(Decorator):
    def show(self):
        print 'Trouser, ',
        self.dt.show()


if __name__ == '__main__':
    person = Person('xiaofan')
    tshirt = Tshirt()
    trouser = Trouser()
    trouser.decorate(person)
    tshirt.decorate(trouser)
    tshirt.show()
