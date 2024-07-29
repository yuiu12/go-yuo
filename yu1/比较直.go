// ==：如果两个值相同，则为 True
// !=：如果两个值不相同，则为 True
// <：如果左侧值小于右侧值，则为 True
// <=：如果左侧值小于或等于右侧值，则为 True
// >：如果左侧值大于右侧值，则为 True
// >=：如果左侧值大于或等于右侧值，则为 True 
// &&：如果 left 和 right 值均为 true，则为 True
// ||：如果左边和右边的值都为真，则为 True
// !：此运算符仅适用于单个值，如果值为 false，则结果为 true
package main  
import "fmt" 
func main() {
	visits := 15
	fmt.Println("First visit   :",visits == 1) 
	fmt.Println("Return visit   :",visits != 1) 
	fmt.Println("Silver member :",visits >= 10 && visits < 21) 
	fmt.Println("Gold member   :", visits > 20 && visits <= 30)
	fmt.Println("Platinum member :", visits > 30)

}