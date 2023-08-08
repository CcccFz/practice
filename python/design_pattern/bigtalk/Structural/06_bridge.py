# -*- coding: utf-8 -*-

# 模式特点：将抽象部分与它的实现部分分离，使它们都可以独立地变化。
# 程序实例：两种品牌的手机，要求它们都可以运行游戏和通讯录两个软件，而不是为每个品牌的手机都独立编写不同的软件。
# 代码特点：虽然使用了object的新型类，不过在这里不是必须的，是对在Python2.2之后“尽量使用新型类”的建议的遵从示范。


class HandsetSoft(object):
    def run(self):
        pass


class HandsetGame(HandsetSoft):
    def run(self):
        print 'Start Game !'


class HandsetAddressList(HandsetSoft):
    def run(self):
        print 'Start Address List !'


class HandsetBrand(object):
    def __init__(self):
        self._app = None

    def install_app(self, app):
        self._app = app

    def run(self):
        pass


class HandsetBrandM(HandsetBrand):
    def run(self):
        print 'Brand M',
        self._app.run()


class HandsetBrandN(HandsetBrand):
    def run(self):
        print 'Brand N',
        self._app.run()


if __name__ == '__main__':
    brand = HandsetBrandM()
    brand.install_app(HandsetGame())
    brand.run()
