//类型推断
package main  
import "math/rand" 

func main() {
	var seed int64 = 1234456789 
	rand.NewSource(seed)
}