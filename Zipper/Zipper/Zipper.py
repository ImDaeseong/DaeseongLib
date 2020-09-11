import os
import zipfile


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
    return fileList


def Zipfile(folderpath, zipname):
    with zipfile.ZipFile(zipname, 'w', compression=zipfile.ZIP_BZIP2) as zipobj:
        FileList = FindFileList(folderpath)
        for i in range(len(FileList)):
            zipobj.write(FileList[i])


def UnZipfile(folderpath, unzippath):
    with zipfile.ZipFile(folderpath, 'r') as zipobj:
        filelist = zipobj.namelist()
        for item in filelist:
            # info = zipobj.getinfo(item)
            # print(info)
            zipobj.extract(item, unzippath)


if __name__ == '__main__':
    folderpath = 'E:\\test'
    zipname = 'E:\\test.zip'
    # zipfile.ZipFile('c:\\test.zip').extractall()

    # 압축
    # Zipfile(folderpath, zipname)

    # 압축 해제
    UnZipfile(zipname, folderpath)

    pass
