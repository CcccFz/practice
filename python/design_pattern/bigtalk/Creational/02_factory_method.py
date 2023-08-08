# -*- coding: utf-8 -*-

# 模式特点：定义一个用于创建对象的接口，让子类决定实例化哪一个类。这使得一个类的实例化延迟到其子类。
# 程序实例：基类雷锋类，派生出学生类和志愿者类，由这两种子类完成“学雷锋”工作。子类的创建由雷锋工厂的对应的子类完成。


class LeiFeng(object):
    def sweep(self):
        pass

class Student(LeiFeng):
    def sweep(self):
        print 'Student Sweep !'

class Volenter(LeiFeng):
    def sweep(self):
        print 'Volenter Sweep !'

class LeiFengFactory(object):
    def create_leifeng(self):
        pass

class StudentFactory(LeiFengFactory):
    def create_leifeng(self):
        return Student()

class VolenterFactory(LeiFengFactory):
    def create_leifeng(self):
        return Volenter()


if __name__ == '__main__':
    factory = StudentFactory()
    leifeng = factory.create_leifeng()
    leifeng.sweep()

