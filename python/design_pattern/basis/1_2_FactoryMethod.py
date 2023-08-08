# -- coding: utf-8 --

class GeekTranslator(object):
    def __init__(self):
        self.trans = dict(dog="σκύλος", cat="γάτα")

    def get(self, msg):
        try:
            return self.trans[msg]
        except KeyError:
            return str(msg)

class EnglishTranslator(object):
    def get(self, msg):
        return str(msg)

# Factory method
def get_translator(language='english'):
    translator = dict(english=EnglishTranslator, geek=GeekTranslator)
    return translator[language.lower()]()

if __name__ == '__main__':
    sentence = 'dog and cat'
    translator = get_translator('Geek')
    print ' '.join([translator.get(word) for word in sentence.split()])
