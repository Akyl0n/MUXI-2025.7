package main

import "fmt"

// 判断一个数是否为质数
func judge(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 返回 n 以内的所有质数
func Prime(n int) []int {
	if n <= 2 {
		fmt.Println("n小于2，没有比2更小的质数")
		return nil
	}
	a := []int{}
	for i := 2; i <= n; i++ {
		if judge(i) {
			a = append(a, i)
		}
	}
	return a
}

func main() {
	var n int
	fmt.Print("输入一个n值：")
	fmt.Scan(&n)
	a := Prime(n)
	fmt.Println("质数有：", a)
}