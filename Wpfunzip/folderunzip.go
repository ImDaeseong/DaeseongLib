// folderunzip
package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
	"unsafe"
)

type (
	HANDLE uintptr
	HWND   HANDLE
	Hwnd   HWND
	BOOL   int32
	CSIDL  uint32
)

const (
	MAX_PATH = 260

	WM_USER = 1024

	ZIP_COMPLETEUNZIP = WM_USER + 1

	CSIDL_DESKTOPDIRECTORY        = 0x10
	CSIDL_COMMON_DESKTOPDIRECTORY = 0x19
)

var (
	user32 = syscall.NewLazyDLL("user32.dll")

	shell32 = syscall.NewLazyDLL("shell32.dll")

	procFindWindow = user32.NewProc("FindWindowW")

	procSendMessageW = user32.NewProc("SendMessageW")

	procSHGetSpecialFolderPathW = shell32.NewProc("SHGetSpecialFolderPathW")
)

func FindWindowTitle(sTitle string) HWND {
	ret, _, _ := procFindWindow.Call(0, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(sTitle))))
	return HWND(ret)
}

func FindProcessIdByName(name string) (int, string, error) {
	cmd := exec.Command("tasklist.exe", "/fo", "csv", "/nh")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := cmd.Output()
	if err != nil {
		return -1, "", err
	}

	for _, line := range strings.Split(string(out), "\n") {
		infs := strings.Split(line, ",")
		if len(infs) >= 2 {
			pName := strings.Trim(infs[0], "\"")
			pId, _ := strconv.Atoi(strings.Trim(infs[1], "\""))

			if strings.HasPrefix(pName, name) {
				return pId, pName, nil
			}
		}
	}
	return -1, "", err
}

func SendMessage(hwnd HWND, msg uint32, wParam uintptr, lParam uintptr) uintptr {

	ret, _, _ := procSendMessageW.Call(uintptr(hwnd), uintptr(msg), wParam, lParam)
	return ret
}

func BoolToBOOL(value bool) BOOL {
	if value {
		return 1
	}
	return 0
}

func SHGetSpecialFolderPath(hwndOwner HWND, lpszPath *uint16, csidl CSIDL, fCreate bool) bool {

	ret, _, _ := procSHGetSpecialFolderPathW.Call(uintptr(hwndOwner), uintptr(unsafe.Pointer(lpszPath)), uintptr(csidl), uintptr(BoolToBOOL(fCreate)))

	return ret != 0
}

func getDesktopDir() string {

	var buf [MAX_PATH]uint16

	if !SHGetSpecialFolderPath(0, &buf[0], CSIDL_DESKTOPDIRECTORY, false) {
		return ""
	}

	return (syscall.UTF16ToString(buf[0:]))
}

func folderZipflag() string {

	sPath := flag.String("Path", "C:\\Users\\Daeseong\\Downloads", "Path")
	flag.Parse()

	return *sPath
}

func FindZipList(sPath string) []string {

	fileList := []string{}
	zipLst, err := ioutil.ReadDir(sPath)
	if err != nil {
		return nil
	}

	for _, file := range zipLst {
		sExt := strings.ToUpper(file.Name())
		if strings.HasSuffix(sExt, ".ZIP") {
			fileList = append(fileList, file.Name())
		}
	}
	return fileList
}

func UnZipfile(sSource, sTarget string) error {
	reader, err := zip.OpenReader(sSource)
	if err != nil {
		return err
	}
	defer reader.Close()

	if err := os.MkdirAll(sTarget, 0755); err != nil {
		return err
	}

	for _, file := range reader.File {
		path := filepath.Join(sTarget, file.Name)

		if file.FileInfo().IsDir() {
			os.MkdirAll(path, file.Mode())
			continue
		}

		fileReader, err := file.Open()
		if err != nil {
			return err
		}
		defer fileReader.Close()

		targetFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer targetFile.Close()

		if _, err := io.Copy(targetFile, fileReader); err != nil {
			return err
		}
	}

	return nil
}

func writeUnzipLogString(sPath, sText string) {

	file, err := os.OpenFile(sPath, os.O_RDWR|os.O_APPEND, 0660)
	if os.IsNotExist(err) {
		file, err = os.Create(sPath)
	}
	defer file.Close()

	if err != nil {
		return
	}

	n, err := io.WriteString(file, sText)
	if err != nil {
		fmt.Println(n, err)
		return
	}
}

func main() {

	//folderZipflag()
	//writeUnzipLogString("date.txt", sPath)

	//Arg1 := fmt.Sprintf("%s", os.Args[0:1])
	//Arg2 := fmt.Sprintf("%s", os.Args[1:2])

	var sPathFlag string
	if len(os.Args) < 2 {
		sPathFlag = "C:\\Users\\Daeseong\\Downloads"
	} else {
		sPathFlag = fmt.Sprintf("%s", os.Args[1:2])
	}
	replacer := strings.NewReplacer("<br>", " ", "]", "", "[", "")
	sPath := replacer.Replace(sPathFlag)

	var bFile bool = false
	if sPath != "" {
		sExt := strings.ToUpper(sPath)
		if strings.HasSuffix(sExt, ".ZIP") {
			bFile = true
		}
	}

	if bFile {

		sZipFile := fmt.Sprintf("%s", sPath)
		file := filepath.Base(sPath)
		sUnZipFile := fmt.Sprintf("%s\\%s", getDesktopDir(), file[:strings.LastIndex(file, ".")])

		err := UnZipfile(sZipFile, sUnZipFile)
		if err != nil {
			fmt.Println("UnZipfile error")
		} else {
			err = os.Remove(sZipFile)
			if err != nil {
				fmt.Println(err)
			}
		}

	} else {

		//현재 경로만 검색
		fileList := FindZipList(sPath)
		for _, file := range fileList {

			sZipFile := fmt.Sprintf("%s\\%s", sPath, file)
			sUnZipFile := fmt.Sprintf("%s\\%s", getDesktopDir(), file[:strings.LastIndex(file, ".")])

			err := UnZipfile(sZipFile, sUnZipFile)
			if err != nil {
				fmt.Println("UnZipfile error")
			} else {
				err = os.Remove(sZipFile)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}

	//pId, pName, err := FindProcessIdByName("Wpfunzip.exe")
	//if err != nil {
	//	fmt.Println(pId)
	//	fmt.Println(pName)
	//}

	hwnd := FindWindowTitle("Wpfunzip")
	if hwnd != 0 {
		SendMessage(hwnd, ZIP_COMPLETEUNZIP, uintptr(0), uintptr(0))
	}

	fmt.Println("complete")
}
