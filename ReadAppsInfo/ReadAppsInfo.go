// ReadAppsInfo
package main

import (
	_ "bytes"
	"encoding/json"
	"fmt"
	"internal/syscall/windows/registry"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"unsafe"
)

var (
	kernel32              = syscall.NewLazyDLL("kernel32.dll")
	GetModuleFileNameProc = kernel32.NewProc("GetModuleFileNameW")
)

type ItemApp struct {
	Img      string
	Name     string
	System   string
	sPackage string
	Appstore string
	Activity string
	Version  string
}

var (
	Item []ItemApp
)

func IsFile(f string) bool {
	fi, err := os.Stat(f)
	return err == nil && !fi.IsDir()
}

func GetModulePath() string {

	var wpath [syscall.MAX_PATH]uint16
	r1, _, _ := GetModuleFileNameProc.Call(0, uintptr(unsafe.Pointer(&wpath[0])), uintptr(len(wpath)))
	if r1 == 0 {
		return ""
	}
	return filepath.Dir(syscall.UTF16ToString(wpath[:]))
}

func GetBluestackDataInfo(skey string) (string, error) {

	key, err := registry.OpenKey(registry.LOCAL_MACHINE, "SOFTWARE\\BlueStacks", registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer key.Close()

	val, _, err := key.GetStringValue(skey)
	if err != nil {
		log.Fatal(err)
		return "", err
	}

	return val, nil
}

func main() {

	DataDir, _ := GetBluestackDataInfo("DataDir")
	sPath := fmt.Sprintf("%s\\UserData\\Gadget\\apps.json", DataDir)

	data, err := ioutil.ReadFile(sPath)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = json.Unmarshal(data, &Item)
	if err != nil {
		return
	}

	for i := range Item {
		sIconName := Item[i].Img
		sIconName = sIconName[:strings.LastIndex(sIconName, ".")]
		sIconPath := fmt.Sprintf("%s\\UserData\\Library\\Icons\\%s.ico", DataDir, sIconName)

		if IsFile(sIconPath) {
			sExePath := fmt.Sprintf("%s\\DaeseongLib.exe", GetModulePath())
			sLnkPath := fmt.Sprintf("%s\\%s.lnk", GetModulePath(), Item[i].Name)
			sArg := fmt.Sprintf("%s;%s;%s", Item[i].Name, Item[i].sPackage, Item[i].Activity)
			sExeName := fmt.Sprintf("%s", Item[i].Name)

			//go 에서 CoCreateInstance 콜 하는 방법은?
			//CreateShortCuttor(sExePath, sLnkPath, sArg, sExeName, sIconPath)

			sParam := fmt.Sprintf("%s|%s|%s|%s|%s", sExePath, sLnkPath, sArg, sExeName, sIconPath)
			sRunPath := "C:\\Go\\src\\DaeseongLib\\CreateShortCutBluestackIcon.exe"
			cmd := exec.Command("cmd", "/c", sRunPath, sParam)
			cmd.Run()
		}
	}
	fmt.Println("complete")
}
