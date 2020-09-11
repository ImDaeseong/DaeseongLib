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

    nCount = GetIniValue("GameInfo", "GameFileCount", "c:\\readini.cfg")
    for i in range(int(nCount)):
        key = "File%d" % (i)
        val = GetIniValue("GameInfo", key, "c:\\readini.cfg")
        print(val)

    pass
