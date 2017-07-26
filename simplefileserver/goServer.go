// goServer
package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"
)

func getPageInfo() (int, string) {

	nPort := flag.Int("port", 8080, "port")
	sDir := flag.String("dir", "webfile", "directory")
	flag.Parse()

	return *nPort, *sDir
}

func main() {

	nPort, sDir := getPageInfo()

	err := os.Mkdir(sDir, 0777)
	if err != nil {
		log.Println(err)
	}

	//log.Println("파일 서버가 시작되었습니다.(http://127.0.0.1:8080/)")
	//http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("webfile"))))

	log.Println("파일 서버가 시작되었습니다.(http://127.0.0.1:8080/Daeseong/)")
	http.Handle("/Daeseong/", http.StripPrefix("/Daeseong/", http.FileServer(http.Dir(sDir))))

	err = http.ListenAndServe(":"+strconv.Itoa(nPort), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
