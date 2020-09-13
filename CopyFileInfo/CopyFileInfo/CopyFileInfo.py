import os
import shutil


def isFile(path):
    if not os.path.exists(path):
        return False
    else:
        return True


def isDir(path):
    if not os.path.isdir(path):
        return False
    else:
        return True


def createFolder(path):
    if not isDir(path):
        os.makedirs(path)


def FindImageFiles(path):
    fileList = []
    for root, subFolders, files in os.walk(path):
        for f in files:
            if f.upper().endswith(".GIF") or f.upper().endswith(".PNG"):
                if isFile(root + os.sep + f):
                    fileList.append(root + os.sep + f)
    return fileList


def CopyFile(fromPath, toPath):
    shutil.copy(fromPath, toPath)


def CopyFolder(fromPath, toPath):
    # 폴더 생성
    for root, subFolders, files in os.walk(fromPath):
        drive, path = os.path.splitdrive(root)
        createFolder(toPath + path)

    # 파일 복사
    for root, dirs, files in os.walk(fromPath):
        rootpath = os.path.join(os.path.abspath(fromPath), root)

        for file in files:

            # 원본 파일 위치
            filepath = os.path.join(rootpath, file)

            # 복사할 위치
            drive, path = os.path.splitdrive(filepath)

            # 파일 복사
            shutil.copy(filepath, (toPath + path))


if __name__ == '__main__':
    # shutil.copytree('C:\\test', 'C:\\test1')

    # CopyFile('C:\\test\\a.png', 'C:\\test1')

    # createFolder('C:\\test1')

    """
    fileList = FindImageFiles('C:\\test')
    for file in fileList:
        print(file)
    """

    CopyFolder('C:\\test', 'C:\\test1\\v1')

    pass
