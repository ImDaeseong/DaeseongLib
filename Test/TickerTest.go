// TickerTest
package DaeseongLib

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

func IsProcessRunning(NameList ...string) bool {

	if len(NameList) == 0 {
		return false
	}

	cmd := exec.Command("tasklist.exe", "/fo", "csv", "/nh")
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(out)
		return false
	}

	for _, val := range NameList {
		if bytes.Contains(out, []byte(val)) {
			return true
		}
	}
	return false
}

func getgoName(sPath string) string {
	filename := strings.Split(sPath, "\\")
	return filename[len(filename)-1]
}

func getgoFolder(sPath string) string {
	filename := strings.Split(sPath, "\\")
	return filename[len(filename)-2]
}

func getDrives() (drives []string) {
	for _, val := range "ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		_, err := os.Open(string(val) + ":\\")
		if err == nil {
			drives = append(drives, string(val))
		}
	}
	return
}

func checkUrl(sUrl string) bool {
	res, err := http.Get(sUrl)

	if err != nil {
		return false
	}

	if res.StatusCode != http.StatusOK {
		return false
	}
	return true
}

func getStateUrl1() {

	//0.5초마다 url 체크, ok시 종료
	Ticker := time.NewTicker(time.Millisecond * 500)

	for {
		select {
		case <-Ticker.C:

			go func() {

				bCheck := checkUrl("https://play.google.com/store/apps/collection/topselling_free")
				if bCheck == true {
					fmt.Println(time.Now(), "http.StatusOK")
					Ticker.Stop()
				} else {
					fmt.Println(time.Now(), "http.StatusFailed")
				}

			}()
		}
	}
}

func getStateUrl2() {

	//1초마다 url 체크, ok시 종료
	Ticker := time.NewTicker(time.Millisecond * 1000)

	go func() {
		for {
			<-Ticker.C

			bCheck := checkUrl("https://play.google.com/store/apps/collection/topselling_free")
			if bCheck == true {
				fmt.Println(time.Now(), "http.StatusOK")
			} else {
				fmt.Println(time.Now(), "http.StatusFailed")
			}
		}
	}()

	//5초 timer 종료
	time.Sleep(time.Millisecond * 5000)
	Ticker.Stop()
}

func getStateUrl3() {

	//1초마다 url 체크, ok시 종료
	CheckHttp := func() {

		Ticker := time.NewTicker(time.Millisecond * 1000)

		for {

			bCheck := checkUrl("https://play.google.com/store/apps/collection/topselling_free")
			if bCheck == true {
				fmt.Println(time.Now(), "http.StatusOK")
				Ticker.Stop()
			} else {
				fmt.Println(time.Now(), "http.StatusFailed")
			}

			<-Ticker.C
		}
	}

	go CheckHttp()

	select {}
}

func getStateUrl4() {

	Ticker1 := time.NewTicker(time.Millisecond * 5000)
	Ticker2 := time.NewTicker(time.Millisecond * 5000)
	Ticker3 := time.NewTicker(time.Millisecond * 5000)
	Ticker4 := time.NewTicker(time.Millisecond * 5000)
	defer Ticker1.Stop()
	defer Ticker2.Stop()
	defer Ticker3.Stop()
	defer Ticker4.Stop()

	bDoneChan := make(chan bool)

	//20초까지만 실행
	go func() {
		time.Sleep(time.Second * 20)
		bDoneChan <- true
	}()

	for {
		select {
		case <-Ticker1.C:

			bCheck := checkUrl("https://play.google.com/store/apps/collection/topselling_free")
			if bCheck == true {
				fmt.Println(time.Now(), "http.StatusOK")
			} else {
				fmt.Println(time.Now(), "http.StatusFailed")
			}

		case <-Ticker2.C:

			bCheck := checkUrl("https://play.google.com/store/apps/collection/topselling_paid")
			if bCheck == true {
				fmt.Println(time.Now(), "http.StatusOK")
			} else {
				fmt.Println(time.Now(), "http.StatusFailed")
			}

		case <-Ticker3.C:

			bCheck := checkUrl("https://play.google.com/store/apps/collection/topgrossing")
			if bCheck == true {
				fmt.Println(time.Now(), "http.StatusOK")
			} else {
				fmt.Println(time.Now(), "http.StatusFailed")
			}

		case <-Ticker4.C:

			bCheck := checkUrl("https://play.google.com/store/apps/category/GAME/collection/topselling_free")
			if bCheck == true {
				fmt.Println(time.Now(), "http.StatusOK")
			} else {
				fmt.Println(time.Now(), "http.StatusFailed")
			}

		case <-bDoneChan:
			fmt.Println("Complete")
			return
		}
	}
}

func getStateUrl5() {

	StopTicker := time.After(time.Millisecond * 5000)
	Ticker := time.NewTicker(time.Millisecond * 1000)
	defer Ticker.Stop()

	for {
		select {
		case <-Ticker.C:

			drives := getDrives()
			for _, val := range drives {

				sFile := fmt.Sprintf("%s:\\Go\\src\\Daeseonglib\\TickerTest.go", val)
				_, err := os.Stat(sFile)
				if err == nil {
					fmt.Println("drive:", val, "filename:", getgoName(sFile), "foldername:", getgoFolder(sFile))
				}
			}

		case <-StopTicker:
			return
		}
	}
}

func getStateUrl6() {

	//run process
	isRunning := IsProcessRunning("cedt.exe", "cedt.exe", "cedt.exe")
	if !isRunning {
		cmd := exec.Command("cmd", "/C", "C:\\Program Files (x86)\\Crimson Editor\\cedt.exe")
		if err := cmd.Run(); err != nil {
			exec.Command("cmd", "/c", "start", "http://www.naver.com").Start()
		}
	}

	//10초마다 체크
	Ticker := time.NewTicker(10 * time.Second)
	defer Ticker.Stop()

	for {
		select {
		case <-Ticker.C:

			isRunning := IsProcessRunning("cedt.exe")
			if isRunning {
				fmt.Println("isRunning")

			} else {
				fmt.Println("isStoping")
			}
		}
	}
}

/*
func main() {

	//getStateUrl1()

	//getStateUrl2()

	//getStateUrl3()

	//getStateUrl4()

	//getStateUrl5()

	getStateUrl6()
}
*/
