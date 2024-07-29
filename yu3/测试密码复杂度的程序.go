package main 
import (
	"fmt"
	"unicode"
)
func passwordChecker(pw string) bool {
	pwR := []rune(pw) 
	if len(pwR) < 8 {
		return false
	}
	hasUpper := false 
	hasLower := false 
	hasNumber := false 
	hasSymbol := false 
	for _, v := range pwR {
		//大写
		if unicode.IsUpper(v) {
			hasUpper = true
		}
		//小写
		if unicode.IsLower(v) {
			hasLower = true
		}
		if unicode.IsNumber(v) {
			hasNumber = true
		}//标点符号或符号
		if unicode.IsPunct(v) || unicode.IsSymbol(v) {
			hasSymbol = true
		}
	}
	return hasUpper && hasLower && hasNumber && hasSymbol
}
func main() {
	if passwordChecker("") {
		fmt.Println("password good")
	}else {
		fmt.Println("password bad")
	}
	if passwordChecker("This!I5A") {
		fmt.Println("password good")
	}else {
		fmt.Println("password bad")
	}

}