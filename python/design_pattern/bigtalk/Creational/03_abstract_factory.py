# -*- coding: utf-8 -*-

# 模式特点：提供一个创建一系列相关或相互依赖对象的接口，而无需指定它们的类。
# 程序实例：提供对不同的数据库访问的支持。IUser和IDepartment是两种不同的抽象产品，它们都有Access和SQL。Server这两种不同的实现；IFactory是产生IUser和IDepartment的抽象工厂，根据具体实现（AccessFactory和SqlFactory）产生对应的具体的对象（CAccessUser与CAccessDepartment，或者CSqlUser与CSqlDepartment）

class User(object):
    def get_user(self):
        pass

    def insert_user(self):
        pass

class Department(object):
    def get_department(self):
        pass

    def insert_department(self):
        pass

class AccessUser(User):
    def get_user(self):
        print 'Get Access User !'

    def insert_user(self):
        print 'Insert Access User !'

class SqlUser(User):
    def get_user(self):
        print 'Get Sql User !'

    def insert_user(self):
        print 'Insert Sql User !'

class AccessDepartment(Department):
    def get_department(self):
        print 'Get Access Department !'

    def insert_department(self):
        print 'Insert Access Department !'

class SqlDepartment(Department):
    def get_department(self):
        print 'Get Sql Department !'

    def insert_department(self):
        print 'Insert Sql Department !'

class Factory(object):
    def create_user(self):
        pass

    def create_department(self):
        pass

class AccessFactory(Factory):
    def create_user(self):
        return AccessUser()

    def creaet_department(self):
        return AccessDepartment()

class SqlFactory(Factory):
    def create_user(self):
        return SqlUser()

    def create_department(self):
        return SqlDepartment()


if __name__ == '__main__':
    factory = SqlFactory()
    user = factory.create_user()
    user.insert_user()

