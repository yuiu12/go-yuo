package main  
import "fmt" 
var defaultli = 10//新加
// func main() {
// 	offset := 5
// 	fmt.Println(offset)
// 	offset = 10 
// 	fmt.Println(offset)
// }

func main() {
	offset := defaultli 
	fmt.Println(offset) 
	offset = offset + defaultli 
	fmt.Println(offset)
}