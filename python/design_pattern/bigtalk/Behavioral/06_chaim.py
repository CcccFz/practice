# -*- coding: utf-8 -*-

# 模式特点：使多个对象都有机会处理请求，从而避免发送者和接收者的耦合关系。将对象连成链并沿着这条链传递请求直到被处理。
# 程序实例：请假和加薪等请求发给上级，如果上级无权决定，那么递交给上级的上级。


class Request(object):
    def __init__(self, sth, num):
        self._something = sth
        self._num = num

class Major(object):
    def set_next(self, m):
        self._next = m

    def get_request(self, r):
        pass

class TeamMajor(Major):
    def get_request(self, r):
        if r._num < 10:
            print 'TeamMajor: You can %s' % r._something
        else:
            self._next.get_request(r)

class DepartmentMajor(Major):
    def get_request(self, r):
        print 'DepartmentMajor: You can %s' % r._something


if __name__ == '__main__':
    department = DepartmentMajor()
    team = TeamMajor()
    team.set_next(department)
    team.get_request(Request('go home', 5))
    team.get_request(Request('add salary', 12))
