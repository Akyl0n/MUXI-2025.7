package main
import "fmt"

// 给你两个按 非递减顺序 排列的整数数组 nums1 和 nums2 ，
// 另有两个整数 m 和 n ，
// 分别表示 nums1和 nums2 中的元素数目，请你 合并 nums2 到 nums1 中，
// 使合并后的数组同样按 非递减顺序 排列，
// 最后返回合并后的数组

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	result := []int{}
	i := 0
	j := 0
	for i < m && j < n {
		if nums1[i] < nums2[j] {
			result = append(result, nums1[i])
			i++
		} else {
			result = append(result, nums2[j])
			j++
		}
	}
	
	if i == m && j < n{
		for ; j < n ; j++{
			result = append(result , nums2[j])
		}
	}

	if i < m && j == n{
		for ; i < m ; i++{
			result = append(result , nums1[i])
		}
	}
	return result
}

func main() {
	nums1 := []int{1, 1, 2, 3, 3, 4, 5, 6, 7, 7, 8}
	nums2 := []int{3, 4, 5, 5, 5, 6, 7, 8, 9, 9, 10, 22, 33, 34, 56}
	fmt.Println(merge(nums1, len(nums1), nums2, len(nums2)))
}