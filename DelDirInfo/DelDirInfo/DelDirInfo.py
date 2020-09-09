import os
import shutil


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


def FindDirList(path):
    dirList = []
    for root, subFolders, files in os.walk(path, topdown=False):
        dirList.append(root)
    return dirList


if __name__ == '__main__':

    # 파일, 디렉토리 모두 삭제
    # shutil.rmtree('E:\\test')

    # 파일만 먼저 지우고
    FileList = FindFileList('E:\\test')
    for i in range(len(FileList)):
        # print(FileList[i])
        os.remove(FileList[i])

    # 파일 삭제 완료후 제일 하위 폴더 부터 차례로 삭제
    dirList = FindDirList('E:\\test')
    for dir in dirList:
        # print(dir)
        os.rmdir(dir)

    pass
