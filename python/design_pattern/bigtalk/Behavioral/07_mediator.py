# -*- coding: utf-8 -*-

# 模式特点：用一个对象来封装一系列的对象交互，中介者使各对象不需要显示地相互引用，从而使耦合松散，而且可以独立地改变它们之间的交互。
# 程序实例：两个对象通过中介者相互通信。


class Colleague:
    def __init__(self, med):
        self.mediator = med

    def notify(self, message):
        pass


class Colleague1(Colleague):
    def send(self, message):
        self.mediator.send(message, self)

    def notify(self, message):
        print 'Colleague1 get a message: %s' % message


class Colleague2(Colleague):
    def send(self, message):
        self.mediator.send(message, self)

    def notify(self, message):
        print 'Colleague2 get a message: %s' % message


class Mediator:
    def send(self, message, col):
        pass


class ConcreteMediator(Mediator):
    def send(self, message, col):
        if col.__class__.__name__ == 'Colleague1':
            self.col2.notify(message)
        else:
            self.col1.notify(message)


if __name__ == '__main__':
    m = ConcreteMediator()
    col1 = Colleague1(m)
    col2 = Colleague2(m)
    m.col1 = col1
    m.col2 = col2

    col1.send("How are you?")
    col2.send("Fine.")
