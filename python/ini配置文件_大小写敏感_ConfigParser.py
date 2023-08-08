import ConfigParser


class CaseConfigParser(ConfigParser.ConfigParser):
    def __init__(self):
        ConfigParser.ConfigParser.__init__(self)

    def optionxform(self, optionstr):
        return optionstr


conf = CaseConfigParser()
conf.read('test.ini')

# 获取指定的section， 指定的option的值
name = conf.get('section1', 'name')
print(name)
age = conf.get("section1", "age")
print(age)

#获取所有section
sections = conf.sections()
print(sections)


# 更新指定section, option的值
conf.set('section2', 'port', '8081')

# 写入指定section, 增加新option的值
conf.set("section2", 'IEPort', '80')

# 添加新的 section
conf.add_section('new_section')
conf.set('new_section', 'new_option', 'http://www.cnblogs.com/tankxiao')

# 写回配置文件
conf.write(open('test.ini', 'w'))