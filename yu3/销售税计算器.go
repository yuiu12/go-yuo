package main  
import "fmt" 
func main() {
	yueu := .0 
	//cake 
	yueu += sale(.99,.075)
	//milk 
	yueu += sale(2.75,.015)
	//butter 
	yueu += sale(.87,.02)
	fmt.Println("Sales Tax Total:",yueu)
}
func sale(yuer float64,yuer1 float64) float64 {
	return yuer * yuer1
}