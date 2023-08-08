# -*-coding: utf-8 -*-

import os
import shutil
import requests
from HTMLParser import HTMLParser

def get_attr(key, attrs):
    for item in attrs:
        if key == item[0]:
            return item[1]
    return None

def create_dir(name):
    if os.path.isdir(name):
        shutil.rmtree(name)
    os.mkdir(name)

class TaobaoWebClient(object):
    goods = {
        u'华为荣耀6X': 'https://item.taobao.com/item.htm?spm=a230r.1.14.283.qKSQD0&id=541904092467&ns=1&abbucket=20',
    }

    def __init__(self):
        self._headers = {'User-Agent': 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36'}
        self._parser = TaobaoImageParser()

    def _get_response(self, url):
        return requests.get(url, headers=self._headers)

    def _parse_content(self, content):
        self._parser.feed(content)

    def _save_images(self):
        for path, src in self._parser.images.iteritems():
            content = requests.get(src).content
            with open(path, 'wb') as wf:
                wf.write(content)

    def parse_webs(self):
        for name, url in self.goods.iteritems():
            content = self._get_response(url).content
            with open('1.html', 'w') as wf:
                wf.write(content)
            create_dir('img/%s' % name)

            self._parser.in_description = False
            self._parser.images = {}
            self._parser.filename = name
            self._parser.id = 0

            self._parse_content(content)
            self._save_images()

        self._parser.close()

class TaobaoImageParser(HTMLParser):
    in_description = False
    images = {}
    filename = ''
    id = 0

    def __init__(self):
        HTMLParser.__init__(self)

    def handle_starttag(self, tag, attrs):
        if not self.in_description:
            if tag == 'div' and get_attr('id', attrs) == 'description':
                self.in_description = True
        else:
            if tag == 'img':
                src = get_attr('src', attrs)
                if src:
                    self.id += 1
                    sufix = os.path.splitext(src)[-1]
                    path = r'img/%s%d%s' % (self.filename, self.id, sufix)

                    if path in images:
                        print 'error: %s exists!' % path
                    else:
                        images[path] = src

if __name__ == '__main__':
    client = TaobaoWebClient()
    client.parse_webs()
    print '*' * 10, ' finish ', '*' * 10
