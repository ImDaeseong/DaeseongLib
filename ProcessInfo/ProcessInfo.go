// ProcessInfo
package DaeseongLib

import (
	"errors"
	_ "fmt"
	"strings"
	"syscall"
	"unsafe"
)

type (
	HANDLE uintptr
	HWND   HANDLE
)

const (
	WM_CLOSE = 16
	WM_QUIT  = 18
)

const (
	SW_HIDE = 0
	SW_SHOW = 5
)

var (
	user32   = syscall.NewLazyDLL("user32.dll")
	shell32  = syscall.NewLazyDLL("shell32.dll")
	kernel32 = syscall.NewLazyDLL("kernel32.dll")

	procFindWindowW  = user32.NewProc("FindWindowW")
	procSendMessageW = user32.NewProc("SendMessageW")
	procPostMessageW = user32.NewProc("PostMessageW")

	procShellExecuteW = shell32.NewProc("ShellExecuteW")

	procCreateToolhelp32Snapshot = kernel32.NewProc("CreateToolhelp32Snapshot")
	procProcess32First           = kernel32.NewProc("Process32FirstW")
	procProcess32Next            = kernel32.NewProc("Process32NextW")
	procCloseHandle              = kernel32.NewProc("CloseHandle")
	procOpenProcess              = kernel32.NewProc("OpenProcess")
	procTerminateProcess         = kernel32.NewProc("TerminateProcess")
)

func FindWindow(className string, windowName string) (HWND, unsafe.Pointer, HANDLE) {

	ptrclassName := syscall.StringToUTF16Ptr(className)
	ptrwindowName := syscall.StringToUTF16Ptr(windowName)

	ret, _, _ := procFindWindowW.Call(uintptr(unsafe.Pointer(ptrclassName)), uintptr(unsafe.Pointer(ptrwindowName)))
	return HWND(ret), unsafe.Pointer(ret), HANDLE(ret)
}

func SendMessage(hwnd HWND, msg uint32, wParam uintptr, lParam uintptr) uintptr {

	ret, _, _ := procSendMessageW.Call(uintptr(hwnd), uintptr(msg), wParam, lParam)
	return ret
}

func PostMessage(hwnd HWND, msg uint32, wParam uintptr, lParam uintptr) bool {

	ret, _, _ := procPostMessageW.Call(uintptr(hwnd), uintptr(msg), wParam, lParam)
	return ret != 0
}

func ShellExecute(hwnd HWND, lpOperation string, lpFile string, lpParameters string, lpDirectory string, nShowCmd int) error {

	var ptrlpOperation uintptr
	var ptrlpFile uintptr
	var ptrlpParameters uintptr
	var ptrlpDirectory uintptr

	if len(lpOperation) != 0 {
		ptrlpOperation = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpOperation)))
	} else {
		ptrlpOperation = uintptr(0)
	}

	if len(lpFile) != 0 {
		ptrlpFile = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpFile)))
	} else {
		ptrlpFile = uintptr(0)
	}

	if len(lpParameters) != 0 {
		ptrlpParameters = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpParameters)))
	} else {
		ptrlpParameters = uintptr(0)
	}

	if len(lpDirectory) != 0 {
		ptrlpDirectory = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(lpDirectory)))
	} else {
		ptrlpDirectory = uintptr(0)
	}

	ret, _, _ := procShellExecuteW.Call(uintptr(hwnd), ptrlpOperation, ptrlpFile, ptrlpParameters, ptrlpDirectory, uintptr(nShowCmd))

	errMsg := ""
	if ret != 0 && ret <= 32 {
		errMsg = "error"
	} else {
		return nil
	}
	return errors.New(errMsg)
}

func TerminateProcess(pid int) bool {

	handle, _, _ := procOpenProcess.Call(syscall.PROCESS_TERMINATE, uintptr(0), uintptr(pid))
	if handle < 0 {
		return false
	}
	defer procCloseHandle.Call(handle)

	ret, _, _ := procTerminateProcess.Call(handle, uintptr(0))
	if ret != 1 {
		return false
	}
	return true
}

func GetPID(szFileName string) int {

	var PID int

	snapshot, _, _ := procCreateToolhelp32Snapshot.Call(syscall.TH32CS_SNAPPROCESS, 0)
	if snapshot < 0 {
		return 0
	}
	defer procCloseHandle.Call(snapshot)

	var entry syscall.ProcessEntry32
	entry.Size = uint32(unsafe.Sizeof(entry))
	ret, _, _ := procProcess32First.Call(snapshot, uintptr(unsafe.Pointer(&entry)))
	if ret < 0 {
		return 0
	}

	for {

		ExeName := strings.ToLower(syscall.UTF16ToString(entry.ExeFile[:]))
		if ExeName == szFileName {
			PID = int(entry.ProcessID)
			//fmt.Printf("%d %d %s \n", entry.ProcessID, entry.ParentProcessID, ExeName)
			break
		}

		ret, _, _ := procProcess32Next.Call(snapshot, uintptr(unsafe.Pointer(&entry)))
		if ret == 0 {
			break
		}
	}

	return PID
}

/*
func f1() {

	hwnd, hPointer, hHandle := FindWindow("Progman", "Program Manager")
	if hHandle == 0 {
		println("nil")
	} else {
		println("hwnd")
	}

	if hwnd == 0 {
		println("nil")
	} else {
		println("hwnd")
	}

	if hPointer == nil {
		println("nil")
	} else {
		println("hwnd")
	}
}

func f2() {
	hwnd, _, _ := FindWindow("#32770", "editing mp3")
	if hwnd != 0 {
		PostMessage(hwnd, WM_QUIT, uintptr(0), uintptr(0))
		PostMessage(hwnd, WM_CLOSE, uintptr(0), uintptr(0))
	}
}

func f3() {

	var url string
	url = "http://www.naver.com"
	if len(url) != 0 {
		ShellExecute(0, "open", url, "", "", SW_SHOW)
	}

	var exePath string
	exePath = "C:\\Windows\\System32\\notepad.exe"
	if len(exePath) != 0 {
		ShellExecute(0, "open", exePath, "", "", SW_SHOW)
	}
}

func f4() {
	pid := GetPID("mp3tag.exe")
	TerminateProcess(pid)
}

func main() {
	f4()
}
*/
