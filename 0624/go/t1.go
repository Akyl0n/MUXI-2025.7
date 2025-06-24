package main
import (
	"fmt"
	"sync"
)

// 要求开启50个协程去并发地将一个变量从0增加到5000，
// 也就是说每个协程需要加给这个数加100，
// 但是每个协程每次只能给这个变量加1

func counter (count *int ,id int , wg *sync.WaitGroup , mu *sync.Mutex){
	defer wg.Done()
	for i := 1 ; i <= 100 ; i++{
		mu.Lock()
		(*count)++
		mu.Unlock()
	}
}

func main(){
	var wg sync.WaitGroup
	var mu sync.Mutex
	count := 0
	for i:= 1 ;i <= 50 ; i++{
		wg.Add(1)
		go counter(&count , i , &wg , &mu)
	}
	wg.Wait()
	fmt.Print(count)
}