# -*- coding: utf-8 -*-

# 模式特点：将一个复杂对象的构建(Director)与它的表示(Builder)分离，使得同样的构建过程可以创建不同的表示(ConcreteBuilder)。
# 程序实例："画"出一个四肢健全（头身手腿）的小人


class Person(object):
    def create_head(self):
        pass

    def create_hand(self):
        pass

    def create_body(self):
        pass

    def create_foot(self):
        pass


class ThinPerson(Person):
    def create_head(self):
        print 'Thin create head'

    def create_hand(self):
        print 'Thin create hand'

    def create_body(self):
        print 'Thin create body'

    def create_foot(self):
        print 'Thin create foot'


class FatPerson(Person):
    def create_head(self):
        print 'Fat create head'

    def create_hand(self):
        print 'Fat create hand'

    def create_body(self):
        print 'Fat create body'

    def create_foot(self):
        print 'Fat create foot'


class Director(object):
    def __init__(self, p):
        self.p = p

    def create(self):
        self.p.create_head()
        self.p.create_hand()
        self.p.create_body()
        self.p.create_foot()


if __name__ == '__main__':
    director = Director(FatPerson())
    director.create()

