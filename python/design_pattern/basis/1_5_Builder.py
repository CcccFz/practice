# -- coding: utf-8 --


# Director
class Director(object):
    def __init__(self):
        self.builder = None

    def construct_building(self):
        self.builder.new_building()
        self.builder.build_floor()
        self.builder.build_size()

    def get_building(self):
        return self.builder.building

# Abstract Builder
class Builder(object):
    def __init__(self):
        self.building = None

    def new_building(self):
        self.building = Building()

# Concrete Builder
class BuilderHouse(Builder):
    def build_floor(self):
        self.building.floor = 1

    def build_size(self):
        self.building.size = 'Big'

# Concrete Builder
class BuilderFlat(Builder):
    def build_floor(self):
        self.building.floor = 16

    def build_size(self):
        self.building.size = 'Small'

# Product
class Building(object):
    def __init__(self):
        self.floor = 0
        self.size = ''

    def __repr__(self):
        return '%s: floor:%d size:%s' % (self.__class__.__name__, self.floor, self.size)


if __name__ == '__main__':
    director = Director()
    director.builder = BuilderHouse()
    director.construct_building()
    print director.get_building()
