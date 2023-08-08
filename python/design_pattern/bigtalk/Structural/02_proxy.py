# -*- coding: utf-8 -*-

# 模式特点：为其他对象提供一种代理以控制对这个对象的访问。

class Interface(object):
    def request(self):
        pass

class RealSubject(Interface):
    def request(self):
        print 'request ok !'

class Proxy(Interface):
    def request(self):
        self.real = RealSubject()
        self.real.request()


if __name__ == '__main__':
    proxy = Proxy()
    proxy.request()

