import os
import hashlib


def isFile(path):
    if not os.path.exists(path):
        return False
    else:
        return True


def FindFileList(path):
    fileList = []
    for root, subFolders, files in os.walk(path):
        for f in files:
            if isFile(root + os.sep + f):
                fileList.append(root + os.sep + f)
                # fileList.append(print("{0}\\{1}".format(root, f)))
    return fileList


def generate_file_md5(filename, blocksize=2 ** 20, debug=True):
    m = hashlib.md5()
    size = 0
    with open(filename, "rb") as f:
        while True:
            buf = f.read(blocksize)
            if not buf:
                break
            m.update(buf)
            size += len(buf)
    return m.hexdigest(), size


if __name__ == '__main__':

    FileList = FindFileList('E:\\test')
    for i in range(len(FileList)):
        # print(FileList[i])
        md5, size = generate_file_md5(FileList[i])
        print("md5:" + md5 + " size:" + str(size))

    pass
