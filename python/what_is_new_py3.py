import sys
from functools import total_ordering

# print
print('The answer is', 2 * 2)
print('The answer is', 2 * 2, end='')
print('ok')

print((1,))
print('The', 'answer', 'is', 2 * 2, sep=',')
print('The' 'answer' 'is')
print("fatal error", file=sys.stderr)
#####################################


# Views And Iterators Instead Of Lists
dic = {'c': 3, 'b': 2, 'a': 1}
keys = dic.keys()
print(keys)
print(sorted(keys))  # can't use keys.sort(), because of keys is view instead of list in py3
print(range(10))
print(zip(['a', 'b', 'c'], [1, 2, 3]))
#####################################


# Ordering Comparisons
print(2 < 3)  # 2 < '3' is true in py2, but will rails TypeError im py3


# rich comparisons instead of __cmp__
@total_ordering
class Person(object):
    def __init__(self, firstname, lastname):
        self.first = firstname
        self.last = lastname

    def __eq__(self, other):
        return (self.last, self.first) == (other.last, other.first)

    def __lt__(self, other):
        return (self.last, self.first) < (other.last, other.first)
#####################################


# Integers
# long becomes int
print(1/2)
print(1//2)
print(123456789987654321)  # print(123456789987654321L) is need in py2
print(0o10)  # print(010) is need in py2
#####################################


# Text Vs. Data Instead Of Unicode Vs. 8-bit
s = 'abc'
b = b'a'
print(type(s))
print(type(b))
print(s.encode())  # str 2 bytes
print(b.decode())  # bytes 2 str
print(len('\u20ac'))
print(len(r'\u20ac'))
#####################################


# New Syntax

def func():
    aa = 1

    def inner():
        nonlocal aa
        aa = 2
        print('aa', aa)

    inner()
    print('aa', aa)


func()

print({1, 2, 3})  # set
print(0b10)  # 二进制
print(0o10)  # 八进制
print(0x10)  # 十进制

print({1, 2})  # set instead of sets function

# ABC抽象基类
#####################################


# Changed Syntax

# as instead of ,
# try:
#     print([][2])
# except (IndexError, KeyError) as e:
#     raise Exception('Bad custom error') from e

# metaclass
# class M(object):
#     pass
#
# class C(metaclass=M):
#     pass


print([x for x in (1, 2, 3)])  # use [x for x in 1, 2, 3] in py2

aa = [[1, 2, 3], [4, 5, 6]]
# print(aa[..., 0])

i = 0
for i in range(3):
    pass
print(i)  # i is 2 inn py2, but i is 0 in py3
# for iterator, only next(iter), can't iter.next()
#####################################


# Removed Syntax
# def foo(a, (b, c)) must be def foo(a, b_c): b, c = b_c
# no <>
# no 100L or 100l or u'abc' or U'abc'
# exec() no support stream argument, instead of exec(f.read())
# 'from module import *' cat't use in func
# still can 'from .[module] import name', but can't 'from module import .[name]', import must be absolute path
# classic class is gone
# str still support format, but bytes can't
#####################################


# Changes To Exceptions
# must be raise Exception('bad exception'), not raise Exception, 'bad exception'
# sys.exc_info() removed, instead of e.exc_info
# next() method has been renamed to __next__()
#####################################


# Builtins
# super() without arguments, 自动去找父类
# raw_input() was renamed to input()
print(round(2.5))  # 以前是3
# Removed execfile(). Instead of execfile(fn) use exec(open(fn).read())
# Removed reduce(). Use functools.reduce()
# Removed. dict.has_key() – use the in operator instead
#####################################


# performance
# is slow 10% better than py2
# but json, 以二进制的方式读写文件， 编码的转换(str 2 bytes)非常快
#####################################
