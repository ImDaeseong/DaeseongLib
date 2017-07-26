// StringTestInfo
package DaeseongLib

import (
	"fmt"
	"strings"
	"unicode"
)

//소문자로
func ToLower() {
	fmt.Println(strings.ToLower("AbC")) //abc
}

//대문자로
func ToUpper() {
	fmt.Println(strings.ToUpper("aBc")) //ABC
}

func ToLowerSpecial() {
	var SC unicode.SpecialCase
	fmt.Println(strings.ToLowerSpecial(SC, "Gopher")) //gopher
}

func ToUpperSpecial() {
	var SC unicode.SpecialCase
	fmt.Println(strings.ToUpperSpecial(SC, "Gopher")) //GOPHER
}

//앞에서 문자 검색후 제거
func TrimPrefix() {
	//yes
	//var text = "Gopher test"
	//text = strings.TrimPrefix(text, "Gopher") //test

	//no
	//var text = "Gopher test"
	//text = strings.TrimPrefix(text, "gopher") //Gopher test

	var text = "고퍼 테스트"
	text = strings.TrimPrefix(text, "고퍼") //테스트
	fmt.Println(text)
}

//뒤에서 문자 검색후 제거
func TrimSuffix() {
	//yes
	//var text = "Gopher test"
	//text = strings.TrimSuffix(text, "test") //Gopher

	//no
	//var text = "Gopher test"
	//text = strings.TrimSuffix(text, "tesT") //Gopher test

	var text = "고퍼 테스트"
	text = strings.TrimSuffix(text, "테스트") //고퍼
	fmt.Println(text)
}

func TrimSpace() {
	var text = "   \t\n  고퍼 테스트 \n\t\r\n "
	text = strings.TrimSpace(text) //고퍼 테스트
	fmt.Println(text)
}

func TrimLeft() {
	var text = " !!! 고퍼! 테스트! !!! "
	fmt.Println(strings.TrimLeft(text, "! ")) //고퍼! 테스트! !!!
	fmt.Println(strings.TrimLeft(text, "!"))  //!!! 고퍼! 테스트! !!!
}

func TrimRight() {
	var text = " !!! 고퍼! 테스트! !!! "
	fmt.Println(strings.TrimRight(text, "! ")) //!!! 고퍼! 테스트
	fmt.Println(strings.TrimRight(text, " !")) //!!! 고퍼! 테스트
}

func Trim() {
	var text = " !!! 고퍼! 테스트! !!! "
	fmt.Println(strings.Trim(text, "! ")) //고퍼! 테스트
	fmt.Println(strings.Trim(text, " !")) //고퍼! 테스트
}

//단어의 첫글자만 대문자
func Tilte() {
	var text = "gopher test"         //고퍼 테스트"
	fmt.Println(strings.Title(text)) //Gopher Test
}

//모든 글자 대문자
func ToTitle() {
	fmt.Println(strings.ToTitle("gopher test")) //GOPHER TEST
	fmt.Println(strings.ToTitle("고퍼 테스트"))      //고퍼 테스트
}

//모든 글자 대문자
func ToTitleSpecial() {
	var SC unicode.SpecialCase
	fmt.Println(strings.ToTitleSpecial(SC, "gopher test")) //GOPHER TEST
}

func Repeat() {
	fmt.Println("고퍼" + strings.Repeat("테스트", 2)) // 고퍼테스트테스트
}

func Replace() {
	fmt.Println(strings.Replace("golang golang golang", "lang", "pher", 2))     // gopher gopher golang
	fmt.Println(strings.Replace("golang golang golang", "golang", "goher", -1)) // goher goher goher
	fmt.Println(strings.Replace("고퍼 테스트 고퍼 테스트 고퍼 테스트", "테스트", "", -1))         //고퍼  고퍼  고퍼
}

func Count() {
	fmt.Println(strings.Count("gopher test", "t")) // 2
	fmt.Println(strings.Count("gopher test", ""))  // 12
	fmt.Println(strings.Count("고퍼 테스트", ""))       // 7
	fmt.Println(strings.Count("고퍼 테스트", "고퍼"))     // 1
}

func Contains() {
	fmt.Println(strings.Contains("gopher test", "gopher")) // true
	fmt.Println(strings.Contains("고퍼 테스트", "테스트"))         // true
	fmt.Println(strings.Contains("gopher test", ""))       // true
	fmt.Println(strings.Contains("", ""))                  // true
}

func Join() {
	text := []string{"gopher", "test", "!!"}
	fmt.Println(strings.Join(text, ", ")) // gopher, test, !!
	fmt.Println(strings.Join(text, " "))  // gopher test !!
	fmt.Println(strings.Join(text, ""))   // gophertest!!
}

func Index() {
	fmt.Println(strings.Index("gopher test", "te")) // 7
	fmt.Println(strings.Index("gopher test", "TE")) // -1
}

func LastIndex() {
	fmt.Println(strings.LastIndex("gopher test", "go")) // 0
	fmt.Println(strings.LastIndex("gopher test", "GO")) // -1
}

func Split() {
	fmt.Println(strings.Split("a,b,c", ","))                              // ["a" "b" "c"]
	fmt.Println(strings.Split("a gopher a gopher a gopher gopher", "a ")) // ["" "gopher " "gopher " "gopher gopher"]
	fmt.Println(strings.Split(" abc ", ""))                               // [" " "a" "b" "c" " "]
	fmt.Println(strings.Split("", "gopher test"))                         // [""]
}

/*
func f1() {
	ToLower()
	ToUpper()
	ToLowerSpecial()
	ToUpperSpecial()
	TrimPrefix()
	TrimSuffix()
	TrimSpace()
	TrimLeft()
	TrimRight()
	Trim()
	Tilte()
	ToTitle()
	ToTitleSpecial()
	Repeat()
	Replace()
	Count()
	Contains()
	Join()
	Index()
	LastIndex()
	Split()
}
func main() {
	f1()
}
*/
