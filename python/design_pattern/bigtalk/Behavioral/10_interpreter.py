# -*- coding: utf-8 -*-

# 模式特点：给定一个语言，定义它的文法的一种表示，并定义一个解释器，这个解释器使用该表示来解释语言中的句子。


class Expression(object):
    def interpret(self):
        print 'base',


class TerminalExpression(Expression):
    def interpret(self):
        print 'terminal',


class KdeExpression(Expression):
    def interpret(self):
        print 'kde',


if __name__ == '__main__':
    expressions = [TerminalExpression(), Expression(), TerminalExpression(), KdeExpression()]
    for ex in expressions:
        ex.interpret()
