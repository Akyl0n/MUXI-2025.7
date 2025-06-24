package main
import "fmt"

// 建一颗二叉树
type TreeNode struct{
	Num int
	Lchild *TreeNode
	Rchild *TreeNode
}

// 传入数组，返回数的根节点
// 如果是空节点，其值为-1
func BuildTree(nums []int) *TreeNode{
	if len(nums) == 0 || nums[0] == -1{
		return nil
	}
	root := &TreeNode{Num:nums[0]} // 给根赋初值
	queue := []*TreeNode{root} // 层次遍历辅助队列

	i := 1
	for i < len(nums){
		node := queue[0] // 当前节点
		queue = queue[1:] // 把当前这个节点给出掉，开始给它配左右孩子

		// Lchild
		if i<len(nums) && nums[i] != -1{
			node.Lchild = &TreeNode{Num:nums[i]}
			queue = append(queue , node.Lchild)
		}
		i++

		// Rchild
		if i<len(nums) && nums[i] != -1{
			node.Rchild = &TreeNode{Num:nums[i]}
			queue = append(queue , node.Rchild)
		}
		i++
	}
	return root
}

func main (){
	nums := []int{1,2,3,-1,4,5,-1}
	root :=BuildTree(nums)
	fmt.Print("根节点值：",root.Num)
}