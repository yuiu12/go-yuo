// 定义一个包含未排序数字的切片。
// 将此切片打印到控制台。
// 使用交换对值进行排序。
// 完成后，将现在排序的数字打印到控制台。
package main 
import "fmt"
func main() {
	nums := []int{5,8,2,4,0,1,3,7,9,6}
	fmt.Println("Before:",nums)
	for stops := true; stops;{
		stops = false 
		for i := 1;i < len(nums);i++{
			if nums[i-1] > nums[i]{
				nums[i],nums[i - 1] = nums[i -1],nums[i]
				stops = true
			}
		}
	}
	fmt.Println("After :", nums)
}