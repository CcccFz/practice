# -*- coding: utf-8 -*-

# 模式特点：将一个类的接口转换成为客户希望的另外一个接口。
# 程序实例：用户通过适配器使用一个类的方法。

class Target(object):
    def request(self):
        print 'target request !'

class Adaptee(object):
    def special_request(self):
        print 'special request !'

class Adapter(Target):
    def __init__(self, adaptee):
        self._adaptee = adaptee

    def request(self):
        self._adaptee.special_request()


if __name__ == '__main__':
    adapter = Adapter(Adaptee())
    adapter.request()

