# -- coding: utf-8 --

import weakref


class Card(object):
    _CardPool = weakref.WeakValueDictionary()

    def __new__(cls, val, suit):
        obj = Card._CardPool.get(val+suit, None)
        if not obj:
            obj = object.__new__(cls)
            Card._CardPool[val+suit] = obj
            obj.val, obj.suit = val, suit
        return obj

    def __repr__(self):
        return "<Card: %s%s>" % (self.val, self.suit)

if __name__ == '__main__':
    c1 = Card('9', 'h')
    c2 = Card('9', 'h')
    print(c1, c2)
    print(c1 == c2)
    print(id(c1), id(c2))
