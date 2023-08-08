# -*- coding: utf-8 -*-

# 模式特点：将请求封装成对象，从而使可用不同的请求对客户进行参数化；对请求排队或记录请求日志，以及支持可撤消的操作。
# 程序实例：烧烤店有两种食物，羊肉串和鸡翅。客户向服务员点单，服务员将点好的单告诉大厨，由大厨进行烹饪。
# 代码特点：注意在遍历列表时不要用注释的方式删除，否则会出现bug。bug示例程序附在后面，我认为这是因为remove打乱了for迭代查询列表的顺序导致的。


class Barbucer(object):
    def make_mutton(self):
        print 'make mutton !'

    def make_chicken(self):
        print 'make chicken !'


class Command(object):
    def __init__(self, barbucer):
        self._reciver = barbucer

    def excute(self):
        pass


class BakeMuttonCmd(Command):
    def execute(self):
        self._reciver.make_mutton()


class ChickenWingCmd(Command):
    def execute(self):
        self._reciver.make_chicken()


class Waiter(object):
    def __init__(self):
        self._order = []

    def set_cmd(self, order):
        self._order.append(order)

    def notify(self):
        for cmd in self._order:
            cmd.execute()


if __name__ == '__main__':
    waiter = Waiter()
    barbucer = Barbucer()
    waiter.set_cmd(BakeMuttonCmd(barbucer))
    waiter.set_cmd(ChickenWingCmd(barbucer))
    waiter.notify()
