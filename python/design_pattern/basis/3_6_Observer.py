# -- coding: utf-8 --

class Subject(object):
    def __init__(self):
        self._observers = []

    def attach(self, observer):
        if not observer in self._observers:
            self._observers.append(observer)

    def detach(self, observer):
        self._observers.remove(observer)

    def notify(self, modifier=None):
        for observer in self._observers:
            if modifier != observer:
                observer.update(self)

class Data(Subject):
    def __init__(self, name=''):
        Subject.__init__(self)
        self.name = name
        self._data = 0

    @property
    def data(self):
        return self._data

    @data.setter
    def data(self, value):
        self._data = value
        self.notify()

class HexViewer:
    def update(self, subject):
        print 'HexViewer: Subject %s has data 0x%x' % (subject.name, subject.data)

class DecimalViewer:
    def update(self, subject):
        print 'DecimalViewer: Subject %s has data %d' % (subject.name, subject.data)


if __name__ == '__main__':
    data1 = Data('Data 1')
    data2 = Data('Data 2')
    view1 = DecimalViewer()
    view2 = HexViewer()
    data1.attach(view1)
    data1.attach(view2)
    data2.attach(view2)
    data2.attach(view1)

    data1.data = 3
    data2.data = 5
    data1.detach(view2)
    data2.detach(view2)
    data1.data = 10
    data2.data = 15
    