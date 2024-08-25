package playground

import (
	"fmt"
	"sync"
	"testing"
)

func TestGoroutine(t *testing.T){
	oneChan := make(chan int,50)
	twoChan := make(chan int,50)
	
	lock := sync.Mutex{}
	cond:=sync.NewCond(&lock)

	go func(){
		for i:=0;i<100;i++{
			if i%2==0{
				twoChan<-i
			}else{
				oneChan<-i
			}
		}
	}()

	go func () {
		for {
			cond.Signal()
			fmt.Printf("goroutine 1 is printing %d\n",<-oneChan)
			cond.Wait()
		}
	}()
	go func () {
		for {
			cond.Signal()
			fmt.Printf("goroutine 2 is printing %d\n",<-twoChan)
			cond.Wait()
		}
	}()

	fmt.Println("finish")
}

