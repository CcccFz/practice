# -*- coding: utf-8 -*-


# 模式特点：将对象组合成成树形结构以表示“部分-整体”的层次结构
# 程序实例：公司人员的组织结构


class Component(object):
    def __init__(self, n):
        self._name = n

    def add(self, c):
        pass

    def display(self, ndepth):
        pass


class Leaf(Component):
    def add(self):
        print 'Leaf can not add component !'

    def display(self, ndepth):
        print '%s%s' % ('-' * ndepth, self._name)


class Composite(Component):
    def __init__(self, n):
        super(Composite, self).__init__(n)
        self._c = []

    def add(self, c):
        self._c.append(c)

    def display(self, ndepth):
        print '%s%s' % ('-' * ndepth, self._name)
        for x in self._c:
            x.display(ndepth+1)


if __name__ == '__main__':
    c = Composite('root')
    c.add(Leaf('a1'))
    b = Composite('b')
    b.add(Leaf('b1'))
    b.add(Leaf('b2'))
    c.add(b)
    c.display(1)

