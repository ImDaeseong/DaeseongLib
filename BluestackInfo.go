// BluestackInfo
package DaeseongLib

import (
	"fmt"
	"internal/syscall/windows/registry"
	"io/ioutil"
	"log"
	"os"
	_ "os/exec"
	"strings"
)

type bluestackApp struct {
	NAME     string
	PACKAGE  string
	ACTIVITY string
	VERSION  string
	APPSTORE string
	SYSTEM   string
	IMG      string
}

func GetBluestackkey(skey string) (int, error) {

	key, err := registry.OpenKey(registry.LOCAL_MACHINE, "SOFTWARE\\Wow6432Node\\BlueStacks\\Guests\\Android\\FrameBuffer\\0", registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer key.Close()

	val, _, err := key.GetIntegerValue(skey)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	return int(val), nil
}

func SetBluestackkey(skey string, nValue int) bool {

	key, err := registry.OpenKey(registry.LOCAL_MACHINE, "SOFTWARE\\Wow6432Node\\BlueStacks\\Guests\\Android\\FrameBuffer\\0", registry.SET_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer key.Close()

	dValue := uint32(nValue)
	err1 := key.SetDWordValue(skey, dValue)
	if err1 != nil {
		log.Fatal(err1)
		return false
	}
	return true
}

func SetBluestackkeyIntFromString(skey string, nValue int) bool {

	key, err := registry.OpenKey(registry.LOCAL_MACHINE, "SOFTWARE\\Wow6432Node\\BlueStacks\\Guests\\Android\\FrameBuffer\\0", registry.SET_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer key.Close()

	sValue := fmt.Sprintf("%d", nValue)
	err1 := key.SetStringValue(skey, sValue)
	if err1 != nil {
		log.Fatal(err1)
		return false
	}
	return true
}

func GetBluestackMemory() (int, error) {

	key, err := registry.OpenKey(registry.LOCAL_MACHINE, "SOFTWARE\\BlueStacks\\Guests\\Android\\", registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer key.Close()

	val, _, err := key.GetIntegerValue("Memory")
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	return int(val), nil
}

func SetBluestackMemory(nValue int) bool {

	key, err := registry.OpenKey(registry.LOCAL_MACHINE, "SOFTWARE\\BlueStacks\\Guests\\Android\\", registry.SET_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer key.Close()

	dValue := uint32(nValue)
	err1 := key.SetDWordValue("Memory", dValue)
	if err1 != nil {
		log.Fatal(err1)
		return false
	}
	return true
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

func HDRunAPP() bool {

	InstallDir, _ := GetBluestackDataInfo("InstallDir")
	sHDRunApp := fmt.Sprintf("%s\\HD-RunApp.exe", InstallDir)

	_, err := os.Stat(sHDRunApp)
	if err != nil {
		return false
	}

	attr := &os.ProcAttr{Files: []*os.File{nil, nil, nil}}
	proc, runerr := os.StartProcess(sHDRunApp, []string{}, attr)
	if runerr != nil {
		log.Fatalln(runerr)
	}

	log.Println(proc.Pid)

	return true
}

func HDRunAPPParam(strpackage string, stractivity string) bool {

	InstallDir, _ := GetBluestackDataInfo("InstallDir")
	sHDRunApp := fmt.Sprintf("%s\\HD-RunApp.exe", InstallDir)

	_, err := os.Stat(sHDRunApp)
	if err != nil {
		return false
	}

	attr := &os.ProcAttr{Files: []*os.File{nil, nil, nil}}
	proc, runerr := os.StartProcess(sHDRunApp, []string{sHDRunApp, "--p", strpackage, "-a", stractivity}, attr)
	if runerr != nil {
		log.Fatalln(runerr)
	}
	log.Println(proc.Pid)

	return true
}

func GetBluestackGameInfo() []bluestackApp {

	var results []bluestackApp

	DataDir, _ := GetBluestackDataInfo("DataDir")
	strPath := fmt.Sprintf("%s\\UserData\\Gadget\\apps.json", DataDir)

	_, err := os.Stat(strPath)
	if err != nil {
		log.Fatal(err)
		return results
	}

	bytes, err := ioutil.ReadFile(strPath)
	if err != nil {
		log.Fatal(err)
		return results
	}

	sContent := string(bytes)
	if sContent == "" {
		return results
	}

	pkgsJson := []string{}
	for _, sRead := range strings.Split(sContent, `{`) {
		if strings.Contains(sRead, `}`) {
			replacer := strings.NewReplacer("}", "", "]", "")
			sRead := replacer.Replace(sRead)
			pkgsJson = append(pkgsJson, sRead)
		}
	}
	//fmt.Println(strings.Join(pkgsJson, "\n"))
	//for _, element := range pkgsJson {
	//	fmt.Println(element)
	//}

	for i := 0; i < len(pkgsJson); i++ {

		var ary bluestackApp

		rows := strings.Split(pkgsJson[i], ",")
		for _, line := range rows {

			line := strings.Trim(line, " ")
			if line == "" {
				continue
			}

			values := strings.Split(line, ":")
			replacer := strings.NewReplacer("\"", "", "\r", "", "\n", "", "\t", "")
			name := replacer.Replace(values[0])
			key := replacer.Replace(values[1])
			name = strings.Trim(name, " ")
			key = strings.Trim(key, " ")
			//fmt.Println(name + "\n" + key)

			if name == "name" {
				ary.NAME = key
			} else if name == "package" {
				ary.PACKAGE = key
			} else if name == "activity" {
				ary.ACTIVITY = key
			} else if name == "version" {
				ary.VERSION = key
			} else if name == "appstore" {
				ary.APPSTORE = key
			} else if name == "system" {
				ary.SYSTEM = key
			} else if name == "img" {
				ary.IMG = key
			}
		}

		//fmt.Println(pkgsJson[i])
		results = append(results, ary)
	}

	//fmt.Println(results)
	return results
}

/*
func f1() {
	FullScreen, _ := GetBluestackkey("FullScreen")
	GuestWidth, _ := GetBluestackkey("GuestWidth")
	WindowWidth, _ := GetBluestackkey("WindowWidth")
	GuestHeight, _ := GetBluestackkey("GuestHeight")
	WindowHeight, _ := GetBluestackkey("WindowHeight")
	Memory, _ := GetBluestackMemory()
	fmt.Printf("FullScreen:%d, GuestWidth:%d, WindowWidth:%d, GuestHeight:%d, WindowHeight:%d, Memory:%d\n", FullScreen, GuestWidth, WindowWidth, GuestHeight, WindowHeight, Memory)
}

func f2() {
	HDRunAPP()
}

func f3() {
	HDRunAPPParam("com.netmarble.lineageII", "com.epicgames.ue4.SplashActivity")
}

func f4() {
	SetBluestackkey("FullScreen", 0)
	SetBluestackMemory(1024)

	FullScreen, _ := GetBluestackkey("FullScreen")
	Memory, _ := GetBluestackMemory()
	fmt.Printf("FullScreen:%d, Memory:%d\n", FullScreen, Memory)
}

func f5() {
	var results = GetBluestackGameInfo()
	for _, Apps := range results {
		fmt.Printf("name:%s  package:%s  activity:%s\r\n", Apps.NAME, Apps.PACKAGE, Apps.ACTIVITY)
	}
}

func main() {
	f3()
}
*/
