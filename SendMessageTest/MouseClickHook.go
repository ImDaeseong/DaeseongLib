// MouseClickHook
package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

const (
	WH_MOUSE_LL    = 14
	WM_QUIT        = 18
	WM_LBUTTONDOWN = 513
	WM_LBUTTONUP   = 514
	WM_RBUTTONDOWN = 516
	WM_RBUTTONUP   = 517
	WM_USER        = 1024

	PT_GETWINDOWFROMPOINT = WM_USER + 1
)

type POINT struct {
	X, Y int32
}

type MSG struct {
	Hwnd    HWND
	Message uint32
	WParam  uintptr
	LParam  uintptr
	Time    uint32
	Pt      POINT
}

type (
	HANDLE    uintptr
	HHOOK     HANDLE
	HINSTANCE HANDLE
	DWORD     uint32
	LRESULT   uintptr
	LPARAM    uintptr
	WPARAM    uintptr
	HWND      HANDLE
)

type MSLLHOOKSTRUCT struct {
	pt          POINT
	mouseData   uint32
	flags       uint32
	time        uint32
	dwExtraInfo uintptr
}

type HOOKPROC func(int, WPARAM, LPARAM) LRESULT

var (
	user32A   = syscall.NewLazyDLL("user32.dll")
	kernel32A = syscall.NewLazyDLL("kernel32.dll")

	GetModuleHandleWProc = kernel32A.NewProc("GetModuleHandleW")
	ExitProcessProc      = kernel32A.NewProc("ExitProcess")

	GetMessageWProc          = user32A.NewProc("GetMessageW")
	SetWindowsHookExWProc    = user32A.NewProc("SetWindowsHookExW")
	UnhookWindowsHookExProc  = user32A.NewProc("UnhookWindowsHookEx")
	CallNextHookExProc       = user32A.NewProc("CallNextHookEx")
	GetWindowTextLengthWProc = user32A.NewProc("GetWindowTextLengthW")
	GetWindowTextWProc       = user32A.NewProc("GetWindowTextW")
	GetClassNameWProc        = user32A.NewProc("GetClassNameW")

	procSendMessageW = user32A.NewProc("SendMessageW")
	procFindWindowW  = user32A.NewProc("FindWindowW")

	libuser32, _           = syscall.LoadLibrary("user32.dll")
	WindowFromPointProc, _ = syscall.GetProcAddress(libuser32, "WindowFromPoint")
)

func SendMessage(hwnd HWND, msg uint32, wParam uintptr, lParam uintptr) uintptr {

	ret, _, _ := procSendMessageW.Call(uintptr(hwnd), uintptr(msg), wParam, lParam)
	return ret
}

func FindWindow(className string, windowName string) (HWND, unsafe.Pointer, HANDLE) {

	ptrclassName := syscall.StringToUTF16Ptr(className)
	ptrwindowName := syscall.StringToUTF16Ptr(windowName)

	ret, _, _ := procFindWindowW.Call(uintptr(unsafe.Pointer(ptrclassName)), uintptr(unsafe.Pointer(ptrwindowName)))
	return HWND(ret), unsafe.Pointer(ret), HANDLE(ret)
}

var hHook HHOOK

func GetMessage(msg *MSG, hwnd HWND, msgFilterMin, msgFilterMax uint32) int {
	ret, _, _ := GetMessageWProc.Call(uintptr(unsafe.Pointer(msg)), uintptr(hwnd), uintptr(msgFilterMin), uintptr(msgFilterMax))
	return int(ret)
}

func GetModuleHandle(modulename string) HINSTANCE {
	var mn uintptr
	if modulename == "" {
		mn = 0
	} else {
		mn = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(modulename)))
	}
	ret, _, _ := GetModuleHandleWProc.Call(mn)
	return HINSTANCE(ret)
}

func ExitProcess(uExitCode uint32) {
	ExitProcessProc.Call(uintptr(uExitCode))
}

func GetWindowTextLength(hwnd HWND) int {
	ret, _, _ := GetWindowTextLengthWProc.Call(uintptr(hwnd))
	return int(ret)
}

func GetWindowText(hwnd HWND) string {
	textLen := GetWindowTextLength(hwnd) + 1
	buf := make([]uint16, textLen)
	GetWindowTextWProc.Call(uintptr(hwnd), uintptr(unsafe.Pointer(&buf[0])), uintptr(textLen))
	return syscall.UTF16ToString(buf)
}

func GetClassName(hwnd HWND) string {
	b := make([]uint16, 1024)
	p := uintptr(unsafe.Pointer(&b[0]))
	GetClassNameWProc.Call(uintptr(hwnd), p, 1024)
	return string(syscall.UTF16ToString(b))
}

func WindowFromPoint(Point POINT) HWND {
	ret, _, _ := syscall.Syscall(WindowFromPointProc, 2, uintptr(Point.X), uintptr(Point.Y), 0)
	return HWND(ret)
}

func SetWindowsHookEx(idHook int, lpfn HOOKPROC, hMod HINSTANCE, dwThreadId DWORD) HHOOK {
	ret, _, _ := SetWindowsHookExWProc.Call(uintptr(idHook), uintptr(syscall.NewCallback(lpfn)), uintptr(hMod), uintptr(dwThreadId))
	return HHOOK(ret)
}

func UnhookWindowsHookEx(hhk HHOOK) bool {
	ret, _, _ := UnhookWindowsHookExProc.Call(uintptr(hhk))
	return ret != 0
}

func CallNextHookEx(hhk HHOOK, nCode int, wParam WPARAM, lParam LPARAM) LRESULT {
	ret, _, _ := CallNextHookExProc.Call(uintptr(hhk), uintptr(nCode), uintptr(wParam), uintptr(lParam))
	return LRESULT(ret)
}

func HookCallback(nCode int, wParam WPARAM, lParam LPARAM) LRESULT {

	if nCode >= 0 && (wParam == WM_LBUTTONDOWN || wParam == WM_LBUTTONUP) {

		hookStruct := (*MSLLHOOKSTRUCT)(unsafe.Pointer(lParam))

		hwnd := WindowFromPoint(hookStruct.pt)
		if hwnd != 0 {

			hwnd, _, _ := FindWindow("WindowsForms10.Window.8.app.0.bf7771", "msgHook")
			if hwnd != 0 {
				SendMessage(hwnd, PT_GETWINDOWFROMPOINT, uintptr(hookStruct.pt.X), uintptr(hookStruct.pt.Y))
			}

			//fmt.Println(hwnd, "WM_LBUTTONDOWN", hookStruct.pt.X, hookStruct.pt.Y)
			//sTitle := GetWindowText(hwnd)
			//sClass := GetClassName(hwnd)
			//fmt.Println("title:", sTitle, "class:", sClass)
		}

	} else if nCode >= 0 && (wParam == WM_RBUTTONDOWN || wParam == WM_RBUTTONUP) {

		UnhookWindowsHookEx(hHook)
		ExitProcess(0)
		fmt.Println("WM_RBUTTONDOWN")
	}

	return CallNextHookEx(hHook, nCode, wParam, lParam)
}

func main() {

	defer syscall.FreeLibrary(libuser32)

	hinst := GetModuleHandle("")
	hHook = SetWindowsHookEx(WH_MOUSE_LL, HookCallback, hinst, 0)

	var msg MSG
	//for GetMessage(nil, 0, 0, 0) != 0 {
	for GetMessage(&msg, 0, 0, 0) != 0 {
		fmt.Println(msg)
	}
	UnhookWindowsHookEx(hHook)
}
