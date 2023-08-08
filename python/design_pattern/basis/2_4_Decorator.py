# -- coding: utf-8 --


class Foo(object):
    def f1(self):
        print 'Original f1'

    def f2(self):
        print 'Original f2'

class Decorator(object):
    def __init__(self, decoratee):
        self.decoratee = decoratee

    def f1(self):
        print 'decorator f1'
        self.decoratee.f1()

    def f2(self):
        print 'decorator f2'
        self.decoratee.f2()


if __name__ == '__main__':
    origin = Foo()
    new_Foo =  Decorator(origin)
    new_Foo.f1()
