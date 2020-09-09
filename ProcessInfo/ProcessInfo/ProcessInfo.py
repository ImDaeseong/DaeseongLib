import win32gui
import win32con
import win32api
import ctypes


def FindWindow(className, windowName):
    return win32gui.FindWindow(className, windowName)


def PostMessage(hwnd):
    win32gui.PostMessage(hwnd, win32con.WM_QUIT, 0, 0)
    win32gui.PostMessage(hwnd, win32con.WM_CLOSE, 0, 0)


if __name__ == '__main__':

    hwnd = FindWindow("#32770", "test")

    if hwnd != 0:
        PostMessage(hwnd)

    pass
