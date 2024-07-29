package main  
import "fmt" 

func main() {
	//创建变量保存总数，账单上的项目，客户购买了两件价值13美元的项目，乘法
	var total float64 = 2 * 13 
	fmt.Println(total)
	//四件价值2.25美元的商品，乘法项目的总数，前一个总值相加，分配回总数
	total = total + (4 * 2.25) 
	fmt.Println("sub :",total)
	//获得5美元折扣，总数减去5美元
	total = total - 5 
	fmt.Println("Sub :",total) 
	//乘法计算10%的提示
	tip := total * 0.1 
	fmt.Println("Tip :",tip) 
	total = total + tip 
	fmt.Println("Total:",total)
	split := total / 2 
	fmt.Println("Split:",split) 
	//计算客户是否获得奖励
	visitCount := 24 
	visitCount = visitCount + 1 
	//除以5美元的任何余数
	remainder := visitCount % 5 
	//每5次获得奖励，余数为0，一次访问，运算符检查余数是否为0
	if remainder == 0 {
		fmt.Println("With this visit, you've earned a reward.")
	}
}