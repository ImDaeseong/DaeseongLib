// SteamgameInfo
package DaeseongLib

import (
	"fmt"
	"internal/syscall/windows/registry"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

type SteamItem struct {
	APPID      string
	INSTALLDIR string
}

func GetSteamPath(skey string) (string, error) {
	key, err := registry.OpenKey(registry.CURRENT_USER, "Software\\Valve\\Steam", registry.QUERY_VALUE)
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

func StartSteam(strUserName string, strPassword string) bool {

	strPath, _ := GetSteamPath("SteamPath")
	SteamRunApp := fmt.Sprintf("%s\\Steam.exe", strPath)
	_, err := os.Stat(SteamRunApp)
	if err != nil {
		return false
	}

	attr := &os.ProcAttr{Files: []*os.File{nil, nil, nil}}
	proc, runerr := os.StartProcess(SteamRunApp, []string{SteamRunApp, "-login", strUserName, strPassword}, attr)
	if runerr != nil {
		log.Fatalln(runerr)
	}
	log.Println(proc.Pid)

	return true
}

func StopSteam() bool {
	strPath, _ := GetSteamPath("SteamPath")
	SteamRunApp := fmt.Sprintf("%s\\Steam.exe", strPath)
	sParam := fmt.Sprintf("-shutdown")
	_, err := os.Stat(SteamRunApp)
	if err != nil {
		return false
	}
	exec.Command(SteamRunApp, sParam).Run()
	return true
}

func ReadSteamGameInfo(strPath string) (strAppid string, strInstallDir string) {

	var sAppid, sInstallDir string = "", ""

	bytes, err := ioutil.ReadFile(strPath)
	if err != nil {
		println(err.Error())
		return "", ""
	}

	lines := strings.Split(string(bytes), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		//fmt.Println(line)

		if strings.Contains(line, "appid") {
			replacer := strings.NewReplacer("appid", "", "\"", "", "\r", "", "\n", "", "\t", "")
			line = replacer.Replace(line)
			sAppid = line

		} else if strings.Contains(line, "installdir") {
			replacer := strings.NewReplacer("installdir", "", "\"", "", "\r", "", "\n", "", "\t", "")
			line = replacer.Replace(line)
			sInstallDir = line
		}
	}

	if sAppid != "" && sInstallDir != "" {
		return sAppid, sInstallDir
	}

	return "", ""
}

func GetGameInstallList() []SteamItem {

	var results []SteamItem
	var ary SteamItem

	strPath, _ := GetSteamPath("SteamPath")
	strGamePath := fmt.Sprintf("%s\\steamapps", strPath)

	files, err := ioutil.ReadDir(strGamePath)
	if err != nil {
		log.Fatalln(err)
	}

	for _, file := range files {

		if file.IsDir() {
			continue
		}

		if !strings.Contains(file.Name(), ".acf") {
			fmt.Printf("File '%s' is not a acf file. \n", file.Name())
			continue
		}

		filePath := path.Join(strGamePath, file.Name())
		if IsExistFile(filePath) {
			strAppid, strInstallDir := ReadSteamGameInfo(filePath)
			if strAppid != "" && strInstallDir != "" {
				ary.APPID = strAppid
				ary.INSTALLDIR = strInstallDir
				results = append(results, ary)
			}
		}
	}
	return results
}

func IsExistDir(sPath string) bool {
	fileinfo, err := os.Stat(sPath)
	if err == nil {
		if fileinfo.IsDir() {
			return true
		}
	}
	return false
}

func IsExistFile(sPath string) bool {
	fileinfo, err := os.Stat(sPath)
	if err == nil {
		if fileinfo.IsDir() {
			return false
		}
		return true
	}
	return os.IsExist(err)
}

/*
func f1() {
	var results = GetGameInstallList()
	for _, Apps := range results {

		strPath, _ := GetSteamPath("SteamPath")
		installPath := fmt.Sprintf("%s\\steamapps\\common\\%s", strPath, Apps.INSTALLDIR)
		if IsExistDir(installPath) {
			fmt.Printf("appid:%s  installPath:%s\r\n", Apps.APPID, installPath)
		}
	}
}

func f2() {
	StartSteam("id", "password")
}

func main() {
	f1()
}
*/
