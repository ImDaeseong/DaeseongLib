import socket
import getmac


def IpAddress():
    return socket.gethostbyname(socket.gethostname())


def macAddress():
    return getmac.get_mac_address()


if __name__ == '__main__':

    # 아이피 조회
    print("IP: " + IpAddress())

    # mac address
    print("mac address: " + macAddress())

    pass
