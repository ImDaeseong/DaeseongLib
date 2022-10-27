import datetime
import json
import time
import requests
from bs4 import BeautifulSoup


def SetDay(nDay):
    now = datetime.datetime.now()
    data = now + datetime.timedelta(days=nDay)
    return data.strftime("%Y%m%d")


def SetMonth(nMonth):
    week = nMonth * 4
    now = datetime.datetime.now()
    data = now + datetime.timedelta(weeks=week)
    return data.strftime("%Y%m%d")


def GetUrlString(nType):
    spage = 'https://apis.data.go.kr/1360000/AsosHourlyInfoService/getWthrDataList?'
    serviceKey = '%2FSWbuoncrZtSM3DaBUA4PJVxqJMFKs0Eu%2F%2FzgFQf8dvVjzIi8ESOjmRaQtAkLKoQUS3S%2BZy%2FwLwR08%2BCT9BWuA%3D%3D'
    pageNo = 1
    numOfRows = 10

    if nType == 1:
        dataType = 'XML'
    else:
        dataType = 'JSON'

    dataCd = 'ASOS'
    dateCd = 'HR'
    startDt = SetDay(-1)  # SetMonth(-1)  # 한달전
    startHh = '01'
    endDt = SetDay(-1)  # 하루전
    endHh = '01'
    stnIds = '108'  # 서울

    sUrl = "{0}serviceKey={1}&pageNo={2}&numOfRows={3}&dataType={4}&dataCd={5}&dateCd={6}&startDt={7}&startHh={8}&endDt={9}&endHh={10}&stnIds={11}".format(
        spage, serviceKey, pageNo, numOfRows, dataType, dataCd, dateCd, startDt, startHh, endDt, endHh, stnIds)
    # print(sUrl)
    return sUrl


def loadUrl_Xml():
    try:
        responses = requests.get(GetUrlString(1), verify=False)
    except:
        time.sleep(2)

    # print(responses)
    # print(responses.text)
    soup = BeautifulSoup(responses.content, 'lxml-xml')

    for item in soup.findAll('item'):
        # print(item)
        print('시간:' + item.find('tm').get_text())
        print('지점번호:' + item.find('stnId').get_text())
        print('지점명:' + item.find('stnNm').get_text())
        print('기온:' + item.find('ta').get_text())
        print('풍속:' + item.find('ws').get_text())
        print('강수량:' + item.find('rn').get_text())
        print('풍향:' + item.find('wd').get_text())
        print('습도:' + item.find('hm').get_text())


def loadUrl_JSON():
    try:
        responses = requests.get(GetUrlString(2), verify=False)
    except:
        time.sleep(2)

    # print(responses)
    # print(responses.text)
    jData = json.loads(responses.content)

    for item in jData['response']['body']['items']['item']:
        # print(item)
        print('시간:' + item['tm'])
        print('지점번호:' + item['stnId'])
        print('지점명:' + item['stnNm'])
        print('기온:' + item['ta'])
        print('풍속:' + item['ws'])
        print('강수량:' + item['rn'])
        print('풍향:' + item['wd'])
        print('습도:' + item['hm'])


if __name__ == '__main__':
    loadUrl_Xml()
    loadUrl_JSON()
