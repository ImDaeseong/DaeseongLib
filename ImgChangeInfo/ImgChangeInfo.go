// ImgChangeInfo
package DaeseongLib

import (
	_ "fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"unsafe"
)

var (
	kernel32B             = syscall.NewLazyDLL("kernel32.dll")
	GetModuleFileNameProc = kernel32B.NewProc("GetModuleFileNameW")
)

func GetModuleFileName() string {

	var wpath [syscall.MAX_PATH]uint16
	r1, _, _ := GetModuleFileNameProc.Call(0, uintptr(unsafe.Pointer(&wpath[0])), uintptr(len(wpath)))
	if r1 == 0 {
		return ""
	}
	return syscall.UTF16ToString(wpath[:])
}

func GetModulePath() string {

	var wpath [syscall.MAX_PATH]uint16
	r1, _, _ := GetModuleFileNameProc.Call(0, uintptr(unsafe.Pointer(&wpath[0])), uintptr(len(wpath)))
	if r1 == 0 {
		return ""
	}
	return filepath.Dir(syscall.UTF16ToString(wpath[:]))
}

func CreateFolder() {

	var sCurPath, sCheckPath string

	sCurPath = GetModulePath()

	sCheckPath = filepath.Join(sCurPath, "PNG")
	_, err := os.Stat(sCheckPath)
	if os.IsNotExist(err) {
		os.MkdirAll(sCheckPath, 0744)
	}

	sCheckPath = filepath.Join(sCurPath, "JPG")
	_, err = os.Stat(sCheckPath)
	if os.IsNotExist(err) {
		os.MkdirAll(sCheckPath, 0744)
	}

	sCheckPath = filepath.Join(sCurPath, "GIF")
	_, err = os.Stat(sCheckPath)
	if os.IsNotExist(err) {
		os.MkdirAll(sCheckPath, 0744)
	}
}

func GetFilePath(sFilename string) string {

	return filepath.Dir(sFilename)
}

func GetFileName(sFilename string) string {

	return filepath.Base(sFilename)
}

var (
	allimgList = []string{}
)

func FindimageList(sPath string) {

	dirlst, err := ioutil.ReadDir(sPath)
	if err != nil {
		panic(err)
	}

	for _, filelst := range dirlst {

		sDir := filepath.Join(sPath, filelst.Name())
		if filelst.IsDir() {

			FindimageList(sDir)

		} else {

			sExt := strings.ToUpper(filepath.Ext(filelst.Name()))

			if strings.HasSuffix(sExt, ".GIF") || strings.HasSuffix(sExt, ".PNG") || strings.HasSuffix(sExt, ".JPG") {

				allimgList = append(allimgList, sDir)
			}
		}
	}
}

func ImageSize(sPath string) (int, int) {

	file, err := os.Open(sPath)
	if err != nil {
		return 0, 0
	}
	defer file.Close()

	img, _, err := image.DecodeConfig(file)
	if err != nil {
		return 0, 0
	}

	return img.Width, img.Height
}

func LoadImage(sPath string) (image.Image, string, error) {

	file, err := os.Open(sPath)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	img, format, err := image.Decode(file)
	if err != nil {
		return nil, "", err
	}
	return img, format, err
}

func SaveToFileImage(sPath string, nImage int) bool {

	var sSavedPath string

	filename := GetFileName(sPath)
	onlyFilename := filename[:strings.LastIndex(filename, ".")]
	//fmt.Println(onlyFilename)

	if nImage == 1 {
		sSavedPath = filepath.Join(GetModulePath(), "PNG", onlyFilename+".png")
	} else if nImage == 2 {
		sSavedPath = filepath.Join(GetModulePath(), "JPG", onlyFilename+".jpg")
	} else if nImage == 3 {
		sSavedPath = filepath.Join(GetModulePath(), "GIF", onlyFilename+".gif")
	}
	//fmt.Println(sSavedPath)

	readfile, err := os.Open(sPath)
	if err != nil {
		return false
	}
	defer readfile.Close()

	readimg, _, err := image.Decode(readfile)
	if err != nil {
		return false
	}

	Savefile, err := os.Create(sSavedPath)
	if err != nil {
		return false
	}
	defer Savefile.Close()

	if nImage == 1 {

		err = png.Encode(Savefile, readimg)

	} else if nImage == 2 {

		var opt jpeg.Options
		opt.Quality = 75
		err = jpeg.Encode(Savefile, readimg, &opt)

	} else if nImage == 3 {

		var opt gif.Options
		opt.NumColors = 256
		err = gif.Encode(Savefile, readimg, &opt)

	}

	return true
}

/*
func f1() {

	dir1, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dir1)

	dir2, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(dir2)

	dir3 := GetModulePath()
	fmt.Println(dir3)
}

func f2() {

	sFolderPath := "D:\\image"
	FindimageList(sFolderPath)

	for _, file := range allimgList {

		imgwidht, imgheight := ImageSize(file)
		fmt.Println("width:", imgwidht, "height:", imgheight)

		_, format, _ := LoadImage(file)
		fmt.Println(format)
	}
}

func f3() {

	CreateFolder()

	sFolderPath := "D:\\image"
	FindimageList(sFolderPath)

	for _, file := range allimgList {

		SaveToFileImage(file, 1)
		SaveToFileImage(file, 2)
		SaveToFileImage(file, 3)
	}
}

func main() {

	f3()
}
*/
