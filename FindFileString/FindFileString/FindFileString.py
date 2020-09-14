import os

FileItem = {}


def isFile(path):
    if not os.path.exists(path):
        return False
    else:
        return True


def ReadFile(path):
    try:
        f = open(path, 'rt', encoding='UTF8')
        contents = f.read()
        return contents
    except:
        return ""


def SearchFiles(path):
    for root, subFolders, files in os.walk(path):
        for f in files:
            filepath = root + os.sep + f
            if isFile(filepath):
                if FileItem.get(f) is None:
                    contents = ReadFile(root + os.sep + f)
                    FileItem[root + os.sep + f] = contents #FileItem[f] = contents


if __name__ == '__main__':

    SearchFiles('E:\DaeseongLib')
    for key in FileItem:
        val = FileItem[key]
        #print("%s : %s" % (key, val))

        if "Split" not in val:
            continue
        else:
            print("%s" % (key))

    pass
