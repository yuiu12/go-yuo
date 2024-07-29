package main 
import (
	"fmt"
	"runtime"
)
func main() {
	var list []int 
	for i := 0; i< 10000000;i++ {
		list = append(list,100)
	}
	//存储内存统计信息
	var m runtime.MemStats
	//获取当前内存分配统计信息
	runtime.ReadMemStats(&m) //字节转换为兆字节
	fmt.Printf("TotalAlloc (Heap) = %v MiB\n",m.TotalAlloc/1024/1024)
}