# -*- coding: utf-8 -*-

# 模式特点：工厂根据条件产生不同功能的类。
# 程序实例：四则运算计算器，根据用户的输入产生相应的运算类，用这个运算类处理具体的运算。
# 代码特点：C/C++中的switch...case...分支使用字典的方式代替。使用异常机制对除数为0的情况进行处理。


class Operation(object):
    def __init__(self, x, y):
        self._x = x
        self._y = y

    def get_result():
        print 'Error: Undefine Operation !'

class OperationAdd(Operation):
    def get_result(self):
        print self._x + self._y

class OperationSub(Operation):
    def get_result(self):
        print self._x - self._y

class OperationMul(Operation):
    def get_result(self):
        print self._x * self._y

class OperationDiv(Operation):
    def get_result(self):
        try:
            print self._x / self._y
        except ZeroDivisionError:
            print 'Error: divided by zero !'

class OperationFactory():
    operations = {'+': OperationAdd, '-': OperationSub, '*': OperationMul, '/':OperationDiv}
    def get_operation(self, op):
        return self.operations.get(op, Operation)


if __name__ == '__main__':
    factory = OperationFactory()
    op = factory.get_operation('/')(10, 2)
    op.get_result()

