from bs4 import BeautifulSoup
import requests


def between(sInput, sFirst, sEnd):
    indexf = sInput.find(sFirst)
    indexl = sInput.rfind(sEnd)
    return sInput[indexf + 1:indexl]


def getVal(sInput, sFirst):
    indexf = sInput.find(sFirst)
    return sInput[indexf + 1:]


def loadurl():
    url = "http://movie.naver.com/movie/running/current.nhn#"
    responses = requests.get(url)
    soup = BeautifulSoup(responses.text, 'html.parser')

    count = 1
    for item in soup.findAll('ul', {'class': 'lst_detail_t1'}):
        lis = item.findAll('li')
        for item1 in lis:

            if count != 0:
                # print(item1)
                thumb = item1.find('div', {'class': 'thumb'})
                tit = item1.find('dt', {'class': 'tit'})
                star_t1 = item1.find('div', {'class': 'star_t1'})
                b_star = item1.find('div', {'class': 'star_t1 b_star'})

                title = thumb.find("img").get('alt')

                ico_rating = ''
                if str(tit).find('<span') > 0:
                    ico_rating = getVal(between(str(tit), '<span', '</span>'), '>')
                else:
                    ico_rating = getVal(between(str(tit), '<a', '</a>'), '>')

                emTemp = star_t1.find('em')
                em = between(str(emTemp), '>', '<')

                num = ''
                if b_star is not None:
                    numTemp = b_star.find('span', {'class': 'num'})
                    num = between(str(numTemp), '>', '<')
                else:
                    num = ''

                print("title:" + title + " ico_rating:" + ico_rating + " em:" + em + " num:" + num)

                count = count + 1


if __name__ == '__main__':
    loadurl()

    pass
