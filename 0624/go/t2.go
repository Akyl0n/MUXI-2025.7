package main
import (
	"fmt"
	"sync"
)

// 开启两个协程，来交替打印字母(A-Z)和数字(1-26)，最终呈现的效果大概是
// A1B2C3D4....

func main(){
	letterCh := make(chan struct{})
	numCh := make (chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)

	go func (){
		defer wg.Done()
		for i:='A' ; i <= 'Z';i++{
			<-letterCh
			fmt.Print(string(i))
			numCh <- struct{}{}
		}
	}()

	go func (){
		defer wg.Done()
		for i := 1;i <= 26 ; i++{
			<-numCh
			fmt.Print(i)
			if i < 26 {
				letterCh <- struct{}{}
			}
		}
	}()

	letterCh <- struct{}{}
	wg.Wait()
}