// SteamgameJson
package DaeseongLib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type SteamgameItem struct {
	Img      string
	Name     string
	System   string
	Package  string
	Appstore string
	Activity string
	Version  string
}

var (
	Item []SteamgameItem
)

func GetReadSteamGameInfo(sPath string) bool {

	file, err := ioutil.ReadFile(sPath)
	if err != nil {
		return false
	}

	err = json.Unmarshal(file, &Item)
	if err != nil {
		return false
	}

	for i := range Item {
		fmt.Printf("%s %s \n", Item[i].Name, Item[i].Activity)
	}
	return true
}

func FindGameInfo(sName string) []string {
	var output []string
	for i := range Item {

		if Item[i].Name == sName {

			output = append(output, Item[i].Img)
			output = append(output, Item[i].Name)
			output = append(output, Item[i].System)
			output = append(output, Item[i].Package)
			output = append(output, Item[i].Appstore)
			output = append(output, Item[i].Activity)
			output = append(output, Item[i].Version)
			break
		}
	}
	return output
}

func IsCheckJson(sPath string) bool {
	if filepath.Ext(sPath) == ".json" {
		return true
	}
	return false
}

func GetCurrentTime() string {
	now := time.Now()
	return fmt.Sprintf("%d/%d/%d", now.Year(), now.Month(), now.Day())
}

func createFile(sPath string) *os.File {
	f, err := os.Create(sPath)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File, sContent string) {
	fmt.Fprintln(f, sContent)
}

func closeFile(f *os.File) {
	f.Close()
}

/*
func f1(filename string) {

	if !IsCheckJson(filename) {
		return
	}

	if GetReadSteamGameInfo(filename) {
		gameinfo := FindGameInfo("Help")
		fmt.Printf("%s", gameinfo)
	}
}

func f2(filename string) {

	GetReadSteamGameInfo(filename)

	file, err := os.Create(filename)
	if err != nil {
	}
	defer file.Close()

	file.WriteString("[")
	for i := range Item {
		setmap := map[string]string{
			"img":      Item[i].Img,
			"name":     Item[i].Name,
			"system":   Item[i].System,
			"package":  Item[i].Package,
			"appstore": Item[i].Appstore,
			"activity": Item[i].Activity,
			"version":  Item[i].Version}
		ret, _ := json.Marshal(setmap)
		file.WriteString(string(ret))
	}
	file.WriteString("]")
}

func f3(filename string) {

	GetReadSteamGameInfo(filename)

	file := createFile(filename)
	writeFile(file, "[")

	//	var sVal string
	//	for i := 0; i < len(Item); i++ {
	//		if i == (len(Item) - 1) {
	//			sVal = fmt.Sprintf("{\"img\": \"%s\", \"name\": \"%s\", \"system\": \"%s\", \"package\": \"%s\", \"appstore\": \"%s\", \"activity\": \"%s\", \"version\": \"%s\"}",
	//				Item[i].Img, Item[i].Name, Item[i].System, Item[i].Package, Item[i].Appstore, Item[i].Activity, Item[i].Version)
	//		} else {
	//			sVal = fmt.Sprintf("{\"img\": \"%s\", \"name\": \"%s\", \"system\": \"%s\", \"package\": \"%s\", \"appstore\": \"%s\", \"activity\": \"%s\", \"version\": \"%s\"},",
	//				Item[i].Img, Item[i].Name, Item[i].System, Item[i].Package, Item[i].Appstore, Item[i].Activity, Item[i].Version)
	//		}
	//		writeFile(file, sVal)
	//	}

	for i := range Item {
		sVal := fmt.Sprintf("{\"img\": \"%s\", \"name\": \"%s\", \"system\": \"%s\", \"package\": \"%s\", \"appstore\": \"%s\", \"activity\": \"%s\", \"version\": \"%s\"},",
			Item[i].Img, Item[i].Name, Item[i].System, Item[i].Package, Item[i].Appstore, Item[i].Activity, Item[i].Version)
		writeFile(file, sVal)
	}

	sVal := fmt.Sprintf("{\"img\": \"%s\", \"name\": \"%s\", \"system\": \"%s\", \"package\": \"%s\", \"appstore\": \"%s\", \"activity\": \"%s\", \"version\": \"%s\"}",
		"Img", "Name", "System", "Package", "Appstore", "Activity", "Version")
	writeFile(file, sVal)

	writeFile(file, "]")
	defer closeFile(file)
}

func f4() {

	var apps map[string]SteamgameItem
	apps = make(map[string]SteamgameItem)

	apps["1"] = SteamgameItem{"img1", "name1", "system1", "package1", "appstore1", "activity1", "version1"}
	apps["2"] = SteamgameItem{"img2", "name2", "system2", "package2", "appstore2", "activity2", "version2"}

	for key, val := range apps {
		fmt.Println(key, val)
	}

	delete(apps, "2")

	item, result := apps["2"]
	if result {
		fmt.Println("found: " + item.Name)
	} else {
		fmt.Println("not found")
	}
}

//apps.json style
//[{"img": "com.bluestacks.help.com.bluestacks.help.HelpActivity.png", "name": "Help", "system": "0", "package": "com.bluestacks.help", "appstore": "no", "activity": "com.bluestacks.help.HelpActivity", "version": "Unknown"}]

func main() {
	//f3("apps.json")
}
*/
