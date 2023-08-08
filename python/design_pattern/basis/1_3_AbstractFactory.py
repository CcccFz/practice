# -- coding: utf-8 --

from random import choice

class PetShop(object):
    def __init__(self):
        self.pet_factory = None

    def show_pet(self):
        pet = self.pet_factory.get_pet()
        print 'Here is %s' % pet
        print 'She speak %s' % pet.speak()
        print 'She eat %s in shop' % self.pet_factory.get_food()

class Dog(object):
    def speak(self):
        return 'Wang~'

    def __str__(self):
        return 'Dog'

class Cat(object):
    def speak(self):
        return 'Miao~'

    def __str__(self):
        return 'Cat'

class DogFactory(object):
    def get_pet(self):
        return Dog()

    def get_food(self):
        return 'Dog Food'

class CatFactory(object):
    def get_pet(self):
        return Cat()

    def get_food(self):
        return 'Cat Food'

def get_factory():
    return choice([CatFactory, DogFactory])()

if __name__ == '__main__':
    shop = PetShop()
    shop.pet_factory = get_factory()
    shop.show_pet()
