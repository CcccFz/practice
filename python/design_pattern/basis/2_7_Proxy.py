# -- coding: utf-8 --

class Manager(object):
    def talk(self):
        print 'main is a good worker'

class Proxy(object):
    def __init__(self):
        self.person = Manager()
        self.busy = False

    def talk(self):
        if self.busy:
            print 'I am busy'
        else:
            self.person.talk()

if __name__ == '__main__':
    p = Proxy()
    p.talk()
    p.busy = True
    p.talk()
