import requests
from requests import get


def GetWebPage(Url):
    response = requests.get('https://github.com/ImDaeseong')
    # print(response.status_code)
    # print(response.text)
    # print(type(response))
    # print(type(response.text))
    # print(response.cookies)
    # print(response.content)
    if response.status_code == 200:
        return response.text
    else:
        return ""


def FileUrl(Url):
    if Url.rfind('/') != -1:
        return Url[Url.rfind('/') + 1:]
    else:
        return Url


def GetDownloadFile(Url):
    path = 'c:\\' + FileUrl(Url)
    with open(path, "wb") as file:
        response = get(Url)
        file.write(response.content)


if __name__ == '__main__':
    req = GetWebPage('https://github.com/ImDaeseong')
    # print(req)

    GetDownloadFile("https://github.com/ImDaeseong/DaeseongLib/blob/master/DaeseongLib/BluestackInfo.go")

    pass
