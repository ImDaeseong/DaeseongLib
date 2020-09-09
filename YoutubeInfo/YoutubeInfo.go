// YoutubeInfo
package DaeseongLib

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	_ "regexp"
	"strconv"
	"strings"
)

type ProgressInfo struct {
	io.Reader
	current int64
	title   string
	length  int64
}

var (
	sSavePath string = "d:\\Daeseong\\"
)

func (pr *ProgressInfo) Read(p []byte) (int, error) {
	var (
		n   int
		err error
	)

	if n, err = pr.Reader.Read(p); err == nil {
		pr.current += int64(n)
		var bytesToMB float64 = 1048576

		fmt.Printf(
			"%s   %vMB / %vMB   %v%%\r",
			pr.title,
			fmt.Sprintf("%.2f", float64(pr.current)/bytesToMB),
			fmt.Sprintf("%.2f", float64(pr.length)/bytesToMB),
			int(float64(pr.current)/float64(pr.length)*float64(100)+1),
		)
	}
	return n, err
}

func GetBodyData(sUrl string) (body []byte, err error) {

	res, err := http.Get(sUrl)
	if err != nil {
		return
	}
	defer res.Body.Close()

	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}
	return
}

func GetVideoList(sUrl string) (idList []string, err error) {

	body, err := GetBodyData(sUrl)
	if err != nil {
		return
	}

	byTag := []byte("class=\"yt-lockup yt-lockup-tile yt-lockup-video vve-check clearfix\" data-context-item-id=\"")
	nIndex := 0
	for {

		offset := bytes.Index(body[nIndex:], byTag)
		if offset < 0 {
			return
		}
		nIndex += offset + len(byTag)

		offset = bytes.Index(body[nIndex:], []byte("\""))
		if offset < 0 {
			return
		}
		end := nIndex + offset
		idList = append(idList, string(body[nIndex:end]))
	}
	return
}

func GetDownloadUrlInfo(sVID string) (string, string) {

	sUrl := fmt.Sprintf("%s%s", "http://www.youtube.com/get_video_info?video_id=", sVID)
	body, err := GetBodyData(sUrl)
	if err != nil {
		return "", ""
	}

	link, err := url.ParseQuery(string(body))
	if err != nil {
		return "", ""
	}

	if link.Get("errorcode") != "" || link.Get("status") == "fail" {
		return "", ""
	}

	//fmt.Println(link.Get("title"))
	//fmt.Println(link.Get("author"))
	//fmt.Println(link.Get("keywords"))
	//fmt.Println(link.Get("thumbnail_url"))
	//fmt.Println(link.Get("view_count"))

	var CheckUrl string
	//var Urls []string
	urlList := strings.Split(link.Get("url_encoded_fmt_stream_map"), ",")
	for _, streams := range urlList {
		u, err := url.ParseQuery(streams)
		if err != nil {
			break
		}

		for _, url := range u["url"] {
			//Urls = append(Urls, url)
			CheckUrl = url
		}
	}

	if CheckUrl == "" {
		return "", ""
	}

	Replacer := strings.NewReplacer(
		"\\", "",
		"/", "",
		":", "",
		"*", "",
		"?", "",
		"\"", "",
		"<", "",
		">", "",
		"|", "",
	)
	sTitle := Replacer.Replace(link.Get("title"))
	if sTitle == "" {
		return "", ""
	}

	m, err := url.ParseQuery(link["url_encoded_fmt_stream_map"][0])
	if err != nil {
		return "", ""
	}

	sDownloadUrl := m["url"][0]
	if sUrl == "" {
		return "", ""
	}

	return sDownloadUrl, sTitle
}

func GetFileSize(sPath string) (int, error) {

	info, err := os.Stat(sPath)
	if err != nil {
		return 0, err
	}
	return int(info.Size()), nil
}

func CreateMp4File(sPath, sFileName string) *os.File {

	// Spath 끝에 \\ 있는지 확인해서 없으면 \\ 붙여줌
	if strings.LastIndex(sPath, string(filepath.Separator)) != len(sPath)-1 {
		sPath = sPath + string(filepath.Separator)
	}

	sFullPath := sPath + sFileName + ".mp4"

	if _, err := os.Stat(sFullPath); !os.IsNotExist(err) {
		os.Remove(sFullPath)
	}

	outfile, _ := os.Create(sFullPath)
	return outfile
}

func GetSearchYoutubeFile(sKey string) []string {

	UrlList := []string{}

	sSearchUrl := "https://www.youtube.com/results?q=" + sKey
	videoIdlst, _ := GetVideoList(sSearchUrl)

	for _, idV := range videoIdlst {
		UrlList = append(UrlList, idV)
	}
	return UrlList
}

func DownLoadFile(sUrl, sTitle string) (int, error) {

	_, err := os.Stat(sSavePath)
	if os.IsNotExist(err) {
		os.MkdirAll(sSavePath, 0744)
	}

	out := CreateMp4File(sSavePath, sTitle)
	defer out.Close()

	resp, err := http.Get(sUrl)
	defer resp.Body.Close()
	if err != nil {
		return 0, err
	}

	_, err = io.Copy(out, &ProgressInfo{
		Reader: resp.Body,
		title:  sTitle,
		length: resp.ContentLength,
	})

	filesize, _ := GetFileSize(sSavePath)

	return filesize, nil
}

func DownLoadFileA(sUrl, sTitle string) (int, error) {

	_, err := os.Stat(sSavePath)
	if os.IsNotExist(err) {
		os.MkdirAll(sSavePath, 0744)
	}

	out := CreateMp4File(sSavePath, sTitle)
	defer out.Close()

	resp, err := http.Get(sUrl)
	defer resp.Body.Close()
	if err != nil {
		return 0, err
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return 0, err
	}

	filesize, _ := GetFileSize(sSavePath)

	return filesize, nil
}

func GetUrlList(sVID string) ([]string, string) {

	sUrl := fmt.Sprintf("%s%s", "http://www.youtube.com/get_video_info?video_id=", sVID)
	body, err := GetBodyData(sUrl)
	if err != nil {
		return nil, ""
	}

	link, err := url.ParseQuery(string(body))
	if err != nil {
		return nil, ""
	}

	if link.Get("errorcode") != "" || link.Get("status") == "fail" {
		return nil, ""
	}

	//fmt.Println(link.Get("title"))
	//fmt.Println(link.Get("author"))
	//fmt.Println(link.Get("keywords"))
	//fmt.Println(link.Get("thumbnail_url"))
	//fmt.Println(link.Get("view_count"))

	var Urls []string
	urlList := strings.Split(link.Get("url_encoded_fmt_stream_map"), ",")
	for _, streams := range urlList {
		u, err := url.ParseQuery(streams)
		if err != nil {
			break
		}

		for _, url := range u["url"] {
			Urls = append(Urls, url)
		}
	}

	Replacer := strings.NewReplacer(
		"\\", "",
		"/", "",
		":", "",
		"*", "",
		"?", "",
		"\"", "",
		"<", "",
		">", "",
		"|", "",
	)
	sTitle := Replacer.Replace(link.Get("title"))

	return Urls, sTitle
}

func YoutubeDownload(sVid string) bool {

	sUrl, sTitle := GetDownloadUrlInfo(sVid)

	filesize, err := DownLoadFile(sUrl, sTitle)
	if err != nil {
		return false
	}

	if filesize == 0 {
		return false
	}
	return true
}

func YoutubeDownloadList(sSearchKey string) bool {

	var bDown bool = false

	vlst := GetSearchYoutubeFile(sSearchKey)
	for i, vID := range vlst {

		sUrl, sTitle := GetDownloadUrlInfo(vID)

		if sUrl != "" && sTitle != "" {

			filesize, err := DownLoadFile(sUrl, sTitle)
			if err == nil {
			} else if filesize == 0 {
				fmt.Println(i, "YoutubeDownloadList failed")
			}
		}
	}
	return bDown
}

func GetYoutubeUrlInfo(sVid string) (map[int]string, string, bool) {

	sUrl := "http://www.youtube.com/get_video_info?hl=en_US&el=detailpage&video_id="
	res, err := http.Get(fmt.Sprintf("%s%s", sUrl, sVid))
	defer res.Body.Close()

	if err != nil {
		return nil, "", false
	}

	body, err := ioutil.ReadAll(res.Body)

	videoInfo, err := url.ParseQuery(string(body))
	if err != nil {
		return nil, "", false
	}

	if videoInfo.Get("errorcode") != "" || videoInfo.Get("status") == "fail" {
		return nil, "", false
	}

	Replacer := strings.NewReplacer(
		"\\", "",
		"/", "",
		":", "",
		"*", "",
		"?", "",
		"\"", "",
		"<", "",
		">", "",
		"|", "",
	)
	sTitle := Replacer.Replace(videoInfo.Get("title"))
	if sTitle == "" {
		return nil, "", false
	}

	UrlList := make(map[int]string, 0)
	for _, v := range strings.Split(videoInfo["url_encoded_fmt_stream_map"][0], ",") {
		Values, err := url.ParseQuery(v)
		if err != nil {
			continue
		}

		itag, err := strconv.Atoi(Values["itag"][0])
		if err != nil {
			continue
		}

		url := Values["url"][0]
		if sig, ok := Values["sig"]; ok {
			url += "&signature=" + sig[0]
		}
		UrlList[itag] = url
	}

	return UrlList, sTitle, true
}

/*
func f1() {

	sVideoID := "9h0Arg_-380"

	OK := YoutubeDownload(sVideoID)
	if OK == true {
		fmt.Println("YoutubeDownload Successed")
	} else {

		fmt.Println("ReTry YoutubeDownload")

		urlList, sTitle := GetUrlList(sVideoID)
		for i, sUrl := range urlList {

			filesize, err := DownLoadFile(sUrl, sTitle)

			if filesize == 0 {
				fmt.Println(i, "YoutubeDownload failed")
			} else {
				if err == nil {
					fmt.Println(i, "YoutubeDownload Successed")
				} else {
					fmt.Println(i, "YoutubeDownload failed")
				}
			}
		}
	}

	fileList := FindFileList("C:\\test")
	for _, f := range fileList {
		fmt.Println(f)
	}
}

func f2() {

	sVideoID := "9h0Arg_-380"

	UrlList, sTitle, Ok := GetYoutubeUrlInfo(sVideoID)
	if Ok == true {
		for i, url := range UrlList {

			fmt.Println(i, sTitle, url)

			filesize, err := DownLoadFile(url, sTitle)
			if filesize == 0 {
				fmt.Println(i, "YoutubeDownload failed")
			} else {
				if err == nil {
					fmt.Println(i, "YoutubeDownload Successed")
					break
				} else {
					fmt.Println(i, "YoutubeDownload failed")
				}
			}
		}
	}
}

func f3() {

	sKeyword := "adele hello" //"surat cinta untuk starla" //"너의 이름은"
	sKeyword = strings.Replace(sKeyword, " ", "+", -1)

	Ok := YoutubeDownloadList(sKeyword)
	if Ok == true {
		fmt.Println("YoutubeDownloadList Successed")
	} else {
		fmt.Println("YoutubeDownloadList failed")
	}
}

func f4() {

	sKeyword := "adele hello"
	sKeyword = strings.Replace(sKeyword, " ", "+", -1)

	vlst := GetSearchYoutubeFile(sKeyword)
	for _, vID := range vlst {

		UrlList, sTitle, Ok := GetYoutubeUrlInfo(vID)
		if Ok == true {
			for i, url := range UrlList {

				filesize, err := DownLoadFile(url, sTitle)
				if filesize == 0 {
					fmt.Println(i, "YoutubeDownload failed")
				} else {
					if err == nil {
						fmt.Println(i, "YoutubeDownload Successed")
						break
					} else {
						fmt.Println(i, "YoutubeDownload failed")
					}
				}
			}
		}
	}
}

func main() {
	f4()
}
*/
