import configparser


def SetIniSection(section, path):
    config = configparser.ConfigParser()
    config.read(path)

    if not config.has_section(section):
        config.add_section(section)
        config.write(open(path, 'w'))


def SetIniValue(section, key, val, path):
    config = configparser.ConfigParser()
    config.read(path)
    try:
        config.add_section(section)
    except configparser.DuplicateSectionError:
        pass
    config.set(section, key, val)
    config.write(open(path, 'w'))


def GetIniValue(section, key, path):
    config = configparser.ConfigParser()
    config.read(path)
    if config.has_option(section, key):
        return config.get(section, key)
    else:
        return None


if __name__ == '__main__':

    path = "C:\\cfgtest.cfg"

    nGameCount = 3
    SetIniSection("GameList", path)
    SetIniValue("GameList", "GameCount", str(nGameCount), path)
    for i in range(nGameCount):
        key = "Gamekey%d" % (i)
        val = "Gamename%d" % (i)
        SetIniValue("GameList", key, val, path)

    nReadCount = GetIniValue("GameList", "GameCount", path)
    nCount = int(nReadCount)
    for i in range(nCount):
        key = "Gamekey%d" % (i)
        val = GetIniValue("GameList", key, path)
        print(val)


    pass
