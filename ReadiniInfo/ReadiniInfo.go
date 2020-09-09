// ReadiniInfo
package DaeseongLib

import (
	"bufio"
	_ "fmt"
	"os"
	"strconv"
	"strings"
)

type ReadiniInfo struct {
	data map[string]string
}

func newReadiniInfo() *ReadiniInfo {
	ret := new(ReadiniInfo)
	ret.data = make(map[string]string)
	return ret
}

func LoadiniConfig(sPath string) *ReadiniInfo {

	ret := newReadiniInfo()

	f, err := os.OpenFile(sPath, os.O_RDONLY, 0644)
	if err != nil {
		return nil
	}

	reader := bufio.NewScanner(f)
	for reader.Scan() {
		line := reader.Text()
		line = strings.TrimSpace(line)
		//fmt.Printf("%s \n", line)

		if strings.Contains(line, "=") {
			parts := strings.SplitN(line, "=", 2)
			//fmt.Printf("%s %s\n", parts[0], parts[1])
			ret.data[parts[0]] = parts[1]
		}
	}
	return ret
}

func (ini *ReadiniInfo) get(sKey string) string {
	if val, ok := ini.data[sKey]; ok {
		return val
	}
	return ""
}

func (ini *ReadiniInfo) GetString(sKey string) string {

	val := ini.get(sKey)
	if val == "" {
		return ""
	}

	ret := strings.Replace(val, "\\n", "\n", -1)
	return ret
}

func (ini *ReadiniInfo) GetInt(sKey string) int {

	val := ini.get(sKey)
	if val == "" {
		return 0
	}

	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0
	}

	return int(i)
}

/*
func fun1() {

	//readini.cfg style
	//
	// [GameInfo]
	// GameFileCount=3
	// File0=a.dll
	// File1=b.ocx
	// File2=c.exe
	//
	ini := LoadiniConfig("readini.cfg")
	count := ini.GetInt("GameFileCount")
	for i := 0; i < count; i++ {

		key := fmt.Sprintf("File%d", i)
		Value := ini.GetString(key)
		fmt.Printf("%s \n", Value)
	}
}

func main() {
	fun1()
}
*/
