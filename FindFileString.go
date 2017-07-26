// FindFileString
package DaeseongLib

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
)

var FileItem map[string]string

func flags() (string, string) {

	spath := flag.String("path", "C:\\Go\\src\\DaeseongLib\\lib", "Search Path")
	skey := flag.String("key", "Split", "Search word")
	flag.Parse()

	return *spath, *skey
}

func Right(sString string, nCount int) string {

	if nCount < 0 {
		return sString
	}

	nLength := len(sString)
	if nLength <= nCount {
		return sString
	}
	sString = string(sString[nLength-nCount : nLength])
	return sString
}

func Left(sString string, nCount int) string {

	if nCount < 0 {
		return sString
	}

	nLength := len(sString)
	if nLength <= nCount {
		return sString
	}
	sString = string(sString[:nCount])
	return sString
}

func ReadLines(spath string) ([]string, error) {

	file, err := os.Open(spath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func SearchFiles(spath string, bRecursive bool) {

	FileItem = make(map[string]string)

	files, err := ioutil.ReadDir(spath)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		if file.IsDir() {

			if bRecursive {
				SearchFiles(spath+file.Name()+string(os.PathSeparator), true)
			}

		} else {

			content, err := ioutil.ReadFile(spath + string(os.PathSeparator) + file.Name())
			if err == nil {

				FileItem[file.Name()] = string(content)
			}
		}
	}
}

func SearchSourceFile() {

	path, key := flags()

	if Right(path, 1) != "/" {
		path += string(os.PathSeparator)
	}

	SearchFiles(path, true)

	for filename, content := range FileItem {

		if strings.Contains(content, key) {
			fmt.Println("Contains:", filename)
		}
	}
}

func FindFileWords(sPath, sKey string) {

	var wg sync.WaitGroup

	files, _ := ioutil.ReadDir(sPath)

	for _, file := range files {

		if file.IsDir() {

			sSubPath := filepath.Join(sPath, file.Name())

			FindFileWords(sSubPath, sKey)

		} else {

			sFullPath := filepath.Join(sPath, file.Name())

			wg.Add(1)
			go func() {
				defer wg.Done()

				pattern := regexp.MustCompile(sKey)
				lines, err := ReadLines(sFullPath)
				if err == nil {

					for _, line := range lines {

						if line != "" {

							data := pattern.FindAllStringSubmatch(line, -1)
							if len(data) != 0 {
								fmt.Println(sFullPath, line)
							}
						}

					}
				}

			}()
		}
	}
	wg.Wait()
}

/*
func f1() {

	//test1
	//if len(os.Args) < 2 {
	//	panic("panic")
	//}
	//ProgramName := os.Args[0:1]
	//Arg1 := os.Args[1:2]
	//Arg2 := os.Args[2:3]
	//allArgs := os.Args[1:]
	//fmt.Println(ProgramName, Arg1, Arg2, allArgs)

	//test2
	//flagstring := flag.String("sValue", "", "command line string")
	//flagint := flag.Int("nValue", 0, "command line int")
	//flagbool := flag.Bool("Ok", false, "command line bool")
	//flag.Parse()
	//fmt.Println(*flagstring, *flagint, *flagbool)

	//test3
	//lines, err := ReadLines("C:\\Go\\src\\DaeseongLib\\lib\\YoutubeInfo.go")
	//if err == nil {
	//	for _, line := range lines {
	//		index := strings.Index(line, "Split")
	//		if index > 0 {
	//			fmt.Println(line)
	//		}
	//	}
	//}

	//test4
	//lines, err = ReadLines("C:\\Go\\src\\DaeseongLib\\lib\\YoutubeInfo.go")
	//if err == nil {
	//
	//	pattern := regexp.MustCompile("Split")
	//	for _, line := range lines {
	//
	//		data := pattern.FindAllStringSubmatch(line, -1)
	//		if len(data) != 0 {
	//			fmt.Println(line)
	//		}
	//	}
	//}
}

func f2() {

	SearchSourceFile()
}

func f3() {

	sPath, sKey := flags()
	FindFileWords(sPath, sKey)
}

func main() {

	f3()
}
*/
