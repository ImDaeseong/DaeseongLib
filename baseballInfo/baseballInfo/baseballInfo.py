from bs4 import BeautifulSoup
import requests


def between(sInput, sFirst, sEnd):
    indexf = sInput.find(sFirst)
    indexl = sInput.rfind(sEnd)
    return sInput[indexf:indexl]


def rfind(sInput, sEnd):
    indexl = sInput.rfind(sEnd)
    return sInput[indexl + 1:]


def loadurlTeamRank():
    url = "https://www.koreabaseball.com/TeamRank/TeamRank.aspx"
    responses = requests.get(url)
    soup = BeautifulSoup(responses.text, 'html.parser')

    table = soup.find('table', {'summary': '순위, 팀명,승,패,무,승률,승차,최근10경기,연속,홈,방문'})
    # print(table)

    tbody = table.find('tbody')
    # print(tbody)

    trs = table.findAll('tr')
    for item in trs:
        # print("tr: " + str(item) + "\n")
        tds = item.find_all('td')
        if len(tds) == 12:
            val = [rankdata.text.strip() for rankdata in tds]
            print(val[0], val[1], val[2], val[3], val[4], val[5], val[6], val[7], val[8], val[9], val[10], val[11])


def loadurlRecord():
    url = "https://www.koreabaseball.com/Record/Main.aspx"
    responses = requests.get(url)
    soup = BeautifulSoup(responses.text, 'html.parser')

    for item in soup.findAll('ol', {'class': 'rankList'}):
        lis = item.findAll('li')

        # rankcount = 1
        for item1 in lis:
            # print(item1)
            """ 
            rankname = "rank%d name" % (rankcount)
            print(rankname)
            rankcount = rankcount + 1
            """

            link = between(str(item1), '<a href', '</a>')
            rank = rfind(link, '>')
            team = item1.find('span', attrs={'class': 'team'}).text
            rr = item1.find('span', attrs={'class': 'rr'}).text
            print("rank:" + rank + " team:" + team + " rr:" + rr)


if __name__ == '__main__':
    # loadurlTeamRank()
    loadurlRecord()

    pass
