# -*- coding: utf-8 -*-

# 模式特点：表示一个作用于某对象结构中的各元素的操作。它使你可以在不改变各元素的类的前提下定义作用于这些元素的新操作。
# 程序实例：对于男人和女人（接受访问者的元素，ObjectStructure用于穷举这些元素），不同的遭遇（具体的访问者）引发两种对象的不同行为。


class Action(object):
    def get_man_conclusion(self):
        pass

    def get_woman_conclusion(self):
        pass


class Success(Action):
    def get_man_conclusion(self):
        print '男人成功时，背后有个伟大的女人'

    def get_woman_conclusion(self):
        print '女人成功时，背后有个不成功的男人'


class Failure(Action):
    def get_man_conclusion(self):
        print '男人失败时，闷头喝酒，谁也不用劝'

    def get_woman_conclusion(self):
        print '女人失败时，眼泪汪汪，谁也劝不了'


class Person(object):
    def accept(self, visitor):
        pass


class Man(Person):
    def accept(self, visitor):
        visitor.get_man_conclusion()


class Woman(Person):
    def accept(self, visitor):
        visitor.get_woman_conclusion()


if __name__ == '__main__':
    man, woman = Man(), Woman()
    woman.accept(Success())
    man.accept(Failure())
