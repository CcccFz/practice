# -- coding: utf-8 --
import sys
import os
import re

class SqlserverUser():
    def insert(self):
        print 'SqlserverUser insert'

    def delete(self):
        print 'SqlserverUser delete'

class AccessUser():
    def insert(self):
        print 'AccessUser insert'

    def delete(self):
        print 'AccessUser delete'

class Factory:

    dictFactory = {'sqlserver': SqlserverUser, 'access': AccessUser}

    @classmethod
    def create_user(cls, db):
        db = db.lower()
        if db not in cls.dictFactory:
            print 'No product !!!'
            sys.exit()
        return cls.dictFactory[db]()

if __name__ == '__main__':
    user = Factory.create_user('Sqlserver')
    user.insert()
    user.delete()