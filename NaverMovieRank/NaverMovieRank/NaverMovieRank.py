from bs4 import BeautifulSoup
import requests


def loadurl():
    responses = requests.get("http://movie.naver.com/movie/sdb/rank/rmovie.nhn")
    soup = BeautifulSoup(responses.text, 'html.parser')

    for item in soup.findAll('div', {'class': 'tit3'}):
        link = item.find("a").get('href')
        title = item.find("a").get('title')
        print("링크: " + link + " 영화제목: " + title)


if __name__ == "__main__":
    loadurl()
