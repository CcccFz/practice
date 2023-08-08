# -- coding: utf-8 --

class Handler(object):
    def successor(self, successor):
        self.successor = successor

class ConcreteHandler1(Handler):
    def handle(self, req):
        if 0 <= req < 10:
            print "%d: in handler1" % req
        else:
            self.successor.handle(req)

class ConcreteHandler2(Handler):
    def handle(self, req):
        if 10 <= req < 50:
            print "%d: in handler2" % req
        else:
            self.successor.handle(req)

class ConcreteHandler3(Handler):
    def handle(self, req):
        if 50 <= req < 100:
            print "%d: in handler3" % req
        else:
            print 'not find'


if __name__ == '__main__':
    c1 = ConcreteHandler1()
    c2 = ConcreteHandler2()
    c3 = ConcreteHandler3()

    c1.successor = c2
    c2.successor = c3

    reqs = [2, 22, 56, 11, 7, 88, 103]
    for req in reqs:
        c1.handle(req)
