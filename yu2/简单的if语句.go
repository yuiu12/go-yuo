//我们将定义一个将被硬编码的值，但在实际应用程序中，这可能是用户输入。然后，我们将在变量上使用运算符（也称为模数表达式）检查该值是奇数还是偶数。模量为您提供了除法后的剩余量。我们将使用模数除以 2 后得到余数。如果我们得到 0 的余数，那么我们知道这个数字是偶数。如果余数是 1，那么我们知道这个数字是奇数。
package main  
import  "fmt" 

func main() {
	//初始值的变量，设置为5，也可以是偶数
	input := 5 
	//使用模数表达式的语句
	if input%2 == 0 {
		fmt.Println(input,"is even")
	}
	if input%2 == 1 {
		fmt.Println(input,"is odd")
	}
	
}