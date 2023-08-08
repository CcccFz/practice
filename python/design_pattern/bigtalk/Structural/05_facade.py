# -*- coding: utf-8 -*-

# 模式特点：为一组调用提供一致的接口。
# 程序实例：接口将几种调用分别组合成为两组，用户通过接口调用其中的一组。


class SubSystemOne(object):
    def method_one(self):
        print 'SubSys One'


class SubSystemTwo(object):
    def method_two(self):
        print 'SubSys Two'


class SubSystemThree(object):
    def method_three(self):
        print 'SubSys Three'


class SubSystemFour(object):
    def method_four(self):
        print 'SubSys Four'


class Facade(object):
    def __init__(self):
        self.one = SubSystemOne()
        self.two = SubSystemTwo()
        self.three = SubSystemThree()
        self.four = SubSystemFour()

    def method_a(self):
        print 'method a'
        self.one.method_one()
        self.four.method_four()

    def method_b(self):
        print 'method b'
        self.two.method_two()
        self.three.method_three()


if __name__ == '__main__':
    facade = Facade()
    facade.method_a()
