package main

import (
	"fmt"
	"sync"
	"time"
)

// *-- understand goroutines --*

// func printNumbers(){
// 	for i:=0; i<5; i++{
// 		fmt.Println(i)
// 		time.Sleep(500 * time.Millisecond)
// 	}
// }

// func main(){
// 	go printNumbers()

// 	fmt.Println("Hello from main")
// 	time.Sleep(3 * time.Second) // time to complete the goroutine printnumbers

// }

// *-- Learn chanels --*

// func Generate(data chan<- int){
// 	for i:=0; i<15; i++{
// 		data<-i
// 	}
// 	close(data)
// }

// func main(){
// 	ch := make(chan int)

// 	go Generate(ch)

// 	for i := range ch{
// 		fmt.Println(i)
// 	}

// }

// *-- Buffered channel --*
// Buffered channels allow sending multiple values without blocking immediately
// it stores limited number of values
// sending to a full channel or recieving from a empty channel blocks execution

// func main(){
// 	ch := make(chan int, 2)

// 	ch<-100
// 	ch<-200
// 	fmt.Println(<-ch, <-ch)
// }

// *-- synchronize with sync.WaitGroup --*
// func work(wg *sync.WaitGroup){
// 	defer wg.Done()
// 	fmt.Println("started work")
// 	fmt.Println("completed work")
// }
// func main(){

// 	var wg sync.WaitGroup

// 	for i:=0; i<10; i++{
// 		wg.Add(1)
// 		go work(&wg)

// 	}
// 	wg.Wait()
// 	fmt.Println("succesfully completed")
// }

// *-- Avoid race condition using sync.Mutex --*
// var counter=0
// func increment(mutex *sync.Mutex, wg *sync.WaitGroup){
// 	mutex.Lock()
// 	defer wg.Done()
// 	defer mutex.Unlock()
// 	counter++
// 	fmt.Println("started and completed work")
// }
// func main(){

// 	var wg sync.WaitGroup
// 	var mutex sync.Mutex

// 	for i:=0; i<5; i++{
// 		wg.Add(1)
// 		go increment(&mutex, &wg)

// 	}
// 	wg.Wait()
// 	fmt.Println("Completed : ", counter)
// }

// *-- Use select for multipluxing channels --*

// func main(){
// 	ch1 := make(chan string)
// 	ch2 := make(chan string)

// 	go func ()  {
// 		time.Sleep(1 * time.Second)
// 		ch1 <- "hello to ch1"
// 	}()

// 	go func ()  {
// 		time.Sleep(2 * time.Second)
// 		ch2 <- "hello to ch2"
// 	}()

// 	for i:=0; i<2; i++{
// 		select{
// 		case msg := <-ch1:
// 			fmt.Println("recieved from ch1 :", msg)
// 		case msg := <-ch2:
// 			fmt.Println("recieved from ch2 :", msg)
// 		}
// 	}
// }

// *--Handle timeouts with select and time.After --*

// func main(){
// 	ch := make(chan int)

// 	go func ()  {
// 		time.Sleep(3*time.Second)
// 		ch<-1001
// 	}()

// 	select{
// 	case num :=<-ch:
// 		fmt.Println("recieved from channel :", num)
// 	case <-time.After(4*time.Second):
// 		fmt.Println("Timeout error")
// 	}
// }

// *-- Web Scraper Simulation --*

func fetchURL(id int, url string, wg *sync.WaitGroup, result chan<- string){
	defer wg.Done()
	time.Sleep(time.Duration(id)*time.Second)
	result <- fmt.Sprintf("fetched %s by worker %d", url, id)
}

func main(){
	urls := []string{"http://example.com", "http://golang.org", "http://github.com"}
	results := make(chan string, len(urls))
	var wg sync.WaitGroup

	for i, url := range urls{
		wg.Add(1)
		go fetchURL(i+1, url, &wg, results)

	}
	wg.Wait()
	close(results)

	for result := range results{
		fmt.Println(result)
	}
}

