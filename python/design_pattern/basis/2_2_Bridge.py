# -- coding: utf-8 --


class DrawAPI1(object):
    def draw(self, x, y, radius):
        print 'API1.circle at %d:%d radius %s' % (x, y, radius)

class DrawAPI2(object):
    def draw(self, x, y, radius):
        print 'API2.circle at {}:{} radius {}'.format(x, y, radius)


class CircleShape(object):
    def __init__(self, x, y, radius, drawing_api):
        self._x = x
        self._y = y
        self._radius = radius
        self._drawing_api = drawing_api

    def scale(self, pct):
        self._radius *= pct

    def draw(self):
        self._drawing_api.draw(self._x, self._y, self._radius)


if __name__ == '__main__':
    shapes = [CircleShape(1, 2, 1.1, DrawAPI1()), CircleShape(4, 5, 3.1, DrawAPI2())]
    for shape in shapes:
        shape.scale(2)
        shape.draw()
