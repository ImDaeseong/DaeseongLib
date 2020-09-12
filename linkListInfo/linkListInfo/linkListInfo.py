from bs4 import BeautifulSoup
import requests
import codecs


urls = ["https://www.naver.com", "https://www.daum.net"]
imglist = []


def WriteString(path, text):
    file = codecs.open(path, 'a', encoding='utf-8')
    file.write(text)
    file.close()


def betweenStartEnd(sInput, sFirst, sEnd):
    index = sInput.find(sFirst)
    sStart = sInput[index + len(sFirst) + 1:]
    sEnd = sStart.find(sEnd)
    return sStart[:sEnd]


def FileUrl(Url):
    if Url.rfind('/') != -1:
        return Url[Url.rfind('/') + 1:]
    else:
        return Url


def GetDownloadFile(sUrl):
    url = FileUrl(sUrl)
    filename = url.replace('?', '')
    filepath = 'c:\\imglist\\' + filename

    with open(filepath, "wb") as file:
        response = requests.get(sUrl)
        if response.status_code == 200:
            file.write(response.content)


if __name__ == '__main__':

    for item in urls:
        result = requests.get(item)
        # print(result.status_code)
        # print(result.headers)
        # print(result.content)

        if result.status_code == 200:
            soup = BeautifulSoup(result.content, 'html.parser')
            imgs = soup.findAll('img')

            for item in imgs:
                # print(item)
                src = betweenStartEnd(str(item), 'src=', '"')
                if src is not None:
                    if src.startswith('https'):
                        # WriteString('C:\\linklist.txt', src + '\n')
                        imglist.append(src)

                datasrc = betweenStartEnd(str(item), 'data-src=', '"')
                if datasrc is not None:
                    if datasrc.startswith('https'):
                        # WriteString('C:\\linklist.txt', datasrc + '\n')
                        imglist.append(datasrc)

            for imgitem in imglist:
                # print(imgitem)
                GetDownloadFile(imgitem)

    pass
