# -*- coding: utf-8 -*-

# 模式特点：保证类仅有一个实例，并提供一个访问它的全局访问点。
# 类本身根本不是单例, 不支持带参数的实例化


def singleton(cls):
    _ins = {}

    def _singleton():
        if cls not in _ins:
            _ins[cls] = cls()
        return _ins[cls]

    return _singleton


@singleton
class A(object):
    pass


if __name__ == '__main__':
    a, b = A(), A()
    print id(a)
    print id(b)

