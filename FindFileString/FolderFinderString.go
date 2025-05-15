package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	sTargetPath = "D:\\mySourceFolder"
	sOutputFile = "outputs.txt"
	sExtension  = []string{".cpp", ".h"}
	hRegex      = regexp.MustCompile(`"(https?://[^"]+)"`)
	sResult     = make(map[string]bool)
)

// 파일 확장자 찾기(.cpp, .h)
func findExtension(sPath string) bool {
	sExt := strings.ToLower(filepath.Ext(sPath))
	for _, sFind := range sExtension {
		if sExt == sFind {
			return true
		}
	}
	return false
}

// 파일 읽어서 문자열 찾기(http,https 포함문자열 검색)
func findLines(sPath string) {

	file, err := os.Open(sPath)
	if err != nil {
		fmt.Printf("파일 열기 실패: %s (%v)\n", sPath, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := hRegex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if len(match) > 1 {
				sResult[match[1]] = true
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("파일 읽기 에러: %s (%v)\n", sPath, err)
	}
}

func findFolder() {

	err := filepath.Walk(sTargetPath, func(sPath string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("파일 처리 중 에러: %v\n", err)
			return nil
		}

		if info.IsDir() {
			return nil
		}

		if findExtension(sPath) {
			findLines(sPath)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("디렉토리 탐색 중 에러 발생: %v\n", err)
		return
	}

	// 결과 저장
	if len(sResult) > 0 {

		var lines []string
		for url := range sResult {
			lines = append(lines, url)
		}

		err := os.WriteFile(sOutputFile, []byte(strings.Join(lines, "\n")), 0644)
		if err != nil {
			fmt.Printf("결과 파일 저장 실패: %v\n", err)
		} else {
			fmt.Printf("총 %d개의 문자열을 %s 파일에 저장했습니다.\n", len(sResult), sOutputFile)
		}
	} else {
		fmt.Println("일치하는 문자열이 없습니다.")
	}
}

// 파일 읽기
func readFile(sPath string) (map[string]bool, error) {

	result := make(map[string]bool)

	file, err := os.Open(sPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		result[line] = true
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

func checkFile(file1, file2 map[string]bool) (file1line, file2line []string) {

	for line := range file1 {

		//file2에 없는것만 설정
		if !file2[line] {
			file1line = append(file1line, line)
		}
	}
	for line := range file2 {

		//file1에 없는것만 설정
		if !file1[line] {
			file2line = append(file2line, line)
		}
	}
	return
}

func compareFile() {

	file1, err := readFile("file1.txt")
	if err != nil {
		return
	}

	file2, err := readFile("file2.txt")
	if err != nil {
		return
	}

	//내용 비교
	file1line, file2line := checkFile(file1, file2)

	//file1, file2 에 동일한 내용인것 찾기
	var commonLines []string
	for line := range file1 {
		if file2[line] {
			commonLines = append(commonLines, line)
		}
	}

	//공통된 줄 출력
	if len(commonLines) > 0 {
		for _, line := range commonLines {
			fmt.Println(line)
		}
	}

	if len(file1line) == 0 && len(file2line) == 0 {
		fmt.Println("두 파일의 내용은 동일합니다.")
	} else {
		fmt.Println("두 파일의 내용이 다릅니다.")

		if len(file1line) > 0 {
			for _, line := range file1line {
				fmt.Println(line)
			}
		}

		if len(file2line) > 0 {
			for _, line := range file2line {
				fmt.Println(line)
			}
		}
	}
}

func main() {

	findFolder()
	//compareFile()
}
