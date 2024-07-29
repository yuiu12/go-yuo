package main 
import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	//随机数生成器播种
	rand.Seed(time.Now().UnixNano()) 
	//生成介于0之间的随机数
	r := rand.Intn(5) + 1 
	stars := strings.Repeat("*", r) 
	fmt.Println(stars) //带有星号的字符串打印到控制台，尾部添加一个新行字符
}
