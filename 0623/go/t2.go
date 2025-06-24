package main
import "fmt"

//给你一组数，返回去重后的数组
func Deduplicate(nums []int) []int{
	seen := make(map[int]bool)
	rezult := []int{}

	for _,num := range nums{
		if !seen[num]{
			seen[num] = true
			rezult = append(rezult , num)
		}
	}
	return rezult
}

func main(){
	nums :=[]int{1,2,2,2,3,3,3,4,4,5,1,1,3,6,7,8,9,0}
	fmt.Println(Deduplicate(nums))
}