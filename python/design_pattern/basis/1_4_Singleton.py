# -- coding: utf-8 --
class Singleton(object):
    def __new__(cls, *args, **kwargs):
        if not hasattr(cls, '_instance'):
            cls._instance = super(Singleton, cls).__new__(cls, *args, **kwargs)
        return cls._instance

class MyClass(Singleton):
    def __init__(self):
        self.attr = 0

a = MyClass()
b = MyClass()

b.attr = 1

print a.attr
print id(a)
print id(b)
