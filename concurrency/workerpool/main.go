package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	var wg sync.WaitGroup
	jobs := make(chan int)
	result := make(chan int)
	now := time.Now()
	for i:=1; i<=3; i++{
		wg.Add(1)
		go func(w int){
			defer wg.Done()
			for job := range jobs{
				fmt.Printf("worker %d started job %d \n", w,  job)
				time.Sleep(time.Second)
				fmt.Printf("worker %d finished job %d \n", w,  job)
				result<-job*2
			}
		}(i)
	}

	go func() {
		for i:= 0; i< 15; i++{
			jobs<-i
		}
		close(jobs)
	}()
	
	go func ()  {
		wg.Wait()
		close(result)
	}()

	for r := range result{
		fmt.Printf("result %d\n", r)
	}
	fmt.Println("Total time taken: ", time.Since(now))

}