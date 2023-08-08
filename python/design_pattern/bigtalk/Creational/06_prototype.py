# -*- coding: utf-8 -*-

# 模式特点：用原型实例指定创建对象的种类，并且通过拷贝这些原型创建新的对象。
# 程序实例：从简历原型，生成新的简历
# 代码特点：简历类Resume提供的Clone()方法其实并不是真正的Clone，只是为已存在对象增加了一次引用。Python为对象提供的copy模块中的copy方法和deepcopy方法已经实现了原型模式，但由于例子的层次较浅，二者看不出区别。

import copy


class Total(object):
    def __init__(self, t):
        self._total = t


class Resume(object):
    def __init__(self, n):
        self._name = n

    def set_age(self, a):
        self._age = a

    def set_total(self, t):
        self._total = t

    def clone(self):
        return self


if __name__ == '__main__':
    a = Resume('english')
    a.set_age(18)
    a.set_total(Total('A'))
    b = a.clone()
    c = copy.copy(a)
    d = copy.deepcopy(a)

    print id(a._total)
    print id(b._total)
    print id(c._total)
    print id(d._total)