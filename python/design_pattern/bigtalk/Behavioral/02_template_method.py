# -*- coding: utf-8 -*-

# 模式特点：定义一个操作中的算法骨架，将一些步骤延迟至子类中。
# 程序实例：考试时使用同一种考卷（父类），不同学生上交自己填写的试卷（子类方法的实现）


class TestPaper(object):
    def question_1(self):
        print 'question_1: A, B, C, D ?'
        self.answer_1()

    def answer_1(self):
        pass

    def question_2(self):
        print 'question_2: A, B, C, D ?'
        self.answer_2()

    def answer_2(self):
        pass


class TestPaper_xiaofan(TestPaper):
    def answer_1(self):
        print 'C'

    def answer_2(self):
        print 'A'


class TestPaper_chengyang(TestPaper):
    def answer_1(self):
        print 'D'

    def answer_2(self):
        print 'C'


if __name__ == '__main__':
    xiaofan = TestPaper_xiaofan()
    chengyang = TestPaper_chengyang()
    print '------------------'
    xiaofan.question_1()
    xiaofan.question_2()
    print '------------------'
    chengyang.question_1()
    chengyang.question_2()
