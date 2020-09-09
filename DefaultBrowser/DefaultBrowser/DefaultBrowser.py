import win32api


def ShellExecute_url(url):
    win32api.ShellExecute(0, 'open', url, '', '', 1)


if __name__ == '__main__':
    ShellExecute_url('https://time.navyism.com/?host=http://naver.com')
    pass
