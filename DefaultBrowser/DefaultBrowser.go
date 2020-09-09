// DefaultBrowser
package DaeseongLib

import (
	"fmt"
	"os/exec"
)

func Webbrowser(sUrl string) (err error) {
	_, err = exec.Command("cmd", "/c", "start", sUrl).Output()
	return err
}

func main() {
	err := Webbrowser("https://time.navyism.com/?host=http://naver.com")
	if err != nil {
		fmt.Println(err)
	}
}
