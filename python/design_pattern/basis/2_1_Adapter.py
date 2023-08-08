# -- coding: utf-8 --


class Human(object):
    def speak(self):
        return 'Hello'

    def __str__(self):
        return 'human'

class Dog(object):
    def bark(self):
        return 'Wang'

    def __str__(self):
        return 'dog'

class Adapter(object):
    def __init__(self, obj, adapted_methods):
        self.obj = obj
        self.__dict__.update(adapted_methods)

    def __getattr__(self, item):
        return getattr(self.obj, item)


if __name__ == '__main__':
    human, dog = Human(), Dog()
    for x in [Adapter(human, dict(speak=human.speak)), Adapter(dog, dict(speak=dog.bark))]:
        print 'a %s speak %s' % (x.obj, x.speak())
