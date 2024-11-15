package main

import(
	"fmt"
)

func main(){
	jobs := make(chan int, 100)
	result := make(chan int, 100)

	for i:=1; i<3; i++{
		go func(w int){
			for job := range jobs{
				fmt.Printf("worker %d Performing job %d \n", w,  job)
				result<-job*2
			}
			close(result)
		}(i)
	}

	for i:= 0; i< 15; i++{
		jobs<-i
	}
	close(jobs)

	for r := range result{
		fmt.Printf("Got job result %d\n", r)
	}

}