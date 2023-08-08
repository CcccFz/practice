# -*- coding: utf-8 -*-

# 模式特点：运用共享技术有效地支持大量细粒度的对象。
# 程序实例：一个网站工厂，根据用户请求的类别返回相应类别的网站。如果这种类别的网站已经在服务器上，那么返回这种网站并加上不同用户的独特的数据；如果没有，那么生成一个。
# 代码特点：为了展示每种网站的由用户请求的次数，这里为它们建立了一个引用次数的字典。之所以不用Python的sys模块中的sys.getrefcount()方法统计引用计数是因为有的对象可能在别处被隐式的引用，从而增加了引用计数。


class WebSite(object):
    def __init__(self, n):
        self._name = n


class ConcreteWebSite(WebSite):
    def use(self, user):
        print 'Website type: %s, user: %s' % (self._name, user)


class UnshareWebsite(WebSite):
    def use(self, user):
        print 'Unshare Website type: %s, user: %s' % (self._name, user)


class WebFactory(object):
    def __init__(self):
        self._types = {'test': ConcreteWebSite('test')}
        self._count = {'test': 1}

    def get_web(self, t):
        if t in self._types:
            self._count[t] += 1
        else:
            self._types[t] = ConcreteWebSite(t)
            self._count[t] = 1
        return self._types[t]

    def get_count(self):
        for k, v in self._count.items():
            print '%s: %s' % (k, v)


if __name__ == '__main__':
    factory = WebFactory()
    shop = factory.get_web('shop')
    test = factory.get_web('test')
    unshare = UnshareWebsite('unshare')
    factory.get_count()