//rune是一种具有足够存储空间来存储单个 UTF-8 多字节字符的类型。字符串文字使用 UTF-8 进行编码。
package main 
import "fmt" 
func main() {
	username := "Sir_King_Über"
	// runes := []rune(username)
	for i := 0;i < len(username);i++ {
		fmt.Print(username[i]," ")
		// fmt.Print(string(runes[i]," "))
	}
}