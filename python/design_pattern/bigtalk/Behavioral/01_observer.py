# -*- coding: utf-8 -*-

# 模式特点：定义了一种一对多的关系，让多个观察对象同时监听一个主题对象，当主题对象状态发生变化时会通知所有观察者。
# 程序实例：公司里有两种上班时趁老板不在时偷懒的员工：看NBA的和看股票行情的，并且事先让老板秘书当老板出现时通知他们继续做手头上的工作。


class Observer(object):
    def __init__(self, n):
        self._name = n

    def update(self):
        pass


class StockObserver(Observer):
    def update(self):
        print('{}: stop watching Stock and go on work !'.format(self._name))


class NbaObserver(Observer):
    def update(self):
        print('{}: stop watching NBA and go on work !'.format(self._name))


class SecretaryBase(object):
    def __init__(self):
        self._observers = []

    def attach(self, p):
        pass

    def notify(self):
        pass


class Secretary(SecretaryBase):
    def attach(self, p):
        self._observers.append(p)

    def notify(self):
        for ob in self._observers:
            ob.update()


if __name__ == '__main__':
    stock = StockObserver('xiaofan')
    nba = NbaObserver('chengyang')
    secretar = Secretary()
    secretar.attach(stock)
    secretar.attach(nba)
    secretar.notify()
