# -*- coding: utf-8 -*-

# 模式特点：定义算法家族并且分别封装，它们之间可以相互替换而不影响客户端。
# 程序实例：商场收银软件，需要根据不同的销售策略方式进行收费。
# 代码特点：不同于同例1，这里使用字典是为了避免关键字不在字典导致bug的陷阱。


class CashBase(object):
    def accept_cash(self, money):
        pass


class CashNormal(CashBase):
    def accept_cash(self, money):
        return money


class CashReturn(CashBase):
    def accept_cash(self, money):
        if money > 200:
            return money - 50
        return money


class CashRebate(CashBase):
    def accept_cash(self, money):
        return money * 0.8


class CasshContext(object):
    def __init__(self, strategy):
        self.strategy = strategy

    def get_result(self, money):
        return self.strategy.accept_cash(money)


if __name__ == '__main__':
    strategies = [CashNormal, CashReturn, CashRebate]
    strategy = strategies[int(raw_input('Plese select: 1) CashNormal, 2) CashReturn, 3) CashRebate'))]()
    context = CasshContext(strategy)
    print context.get_result(300)
