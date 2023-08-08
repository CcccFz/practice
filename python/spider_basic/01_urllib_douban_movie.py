# -*- coding: utf-8 -*-

import urllib2
from HTMLParser import HTMLParser


def get_attr_value(key, attrs):
    for item in attrs:
        if key == item[0]:
            return item[1]
    return None


class WebClient(object):
    def __init__(self):
        self._headers = None
        self._url = None

    def _get_request(self):
        self._req = urllib2.Request(self._url, headers=self._headers)

    def _get_response(self):
        self._response = urllib2.urlopen(self._req)

    def parse_movie(self, parse):
        parse.feed(self._response.read())
        self._response.close()

class DoubanWebClient(WebClient):
    def __init__(self):
        self._headers = {'User-Agent': 'Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2924.87 Safari/537.36'}
        self._url = 'https://movie.douban.com/cinema/nowplaying/chengdu/'
        self._get_request()
        self._get_response()


class WebParser(HTMLParser):
    def __init__(self):
        HTMLParser.__init__(self)

class DoubanMovieParser(WebParser):
    def __init__(self):
        WebParser.__init__(self)

    def handle_starttag(self, tag, attrs):
        if tag == 'li' and get_attr_value('data-title', attrs) \
            and get_attr_value('data-category', attrs) == 'nowplaying':
            movie = []
            movie.append(get_attr_value('data-title', attrs))
            movie.append(get_attr_value('data-score', attrs))
            movie.append(get_attr_value('data-duration', attrs))
            movie.append(get_attr_value('data-release', attrs))
            movie.append(get_attr_value('data-director', attrs))
            movie.append(get_attr_value('data-actors', attrs))

            print ' | '.join(movie)


if __name__ == '__main__':
    parse = DoubanMovieParser()
    client = DoubanWebClient()
    client.parse_movie(parse)
