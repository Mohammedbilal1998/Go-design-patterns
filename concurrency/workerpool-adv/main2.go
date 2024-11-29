package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Worker1Pool struct{
	wg sync.WaitGroup
	numWorkers int
	input chan int
	resutl chan int
	ctx  context.Context
	cancel context.CancelFunc
	done chan struct{}
}

func NewWorker1Pool(numberOfWorkers int) *Worker1Pool{
	ctx, cancle := context.WithCancel(context.Background())

	return &Worker1Pool{
		wg: sync.WaitGroup{},
		numWorkers: numberOfWorkers,
		input: make(chan int),
		resutl: make(chan int),
		ctx: ctx,
		cancel: cancle,
		done: make(chan struct{}),
	}
}

func (w *Worker1Pool) Start(){
	for i:=0; i<w.numWorkers; i++{
		w.wg.Add(1)
		go w.Worker(i)
	}
	go w.FetchResults()
}

func (w *Worker1Pool) Worker(i int){
	defer w.wg.Done()
	for{
		select{
		case <-w.ctx.Done():
			fmt.Printf("worker %d is shuttingdown\n", i)
			return
		case msg, ok := <-w.input:
			if !ok{
				return
			}
			fmt.Printf("Worker %d is working\n", i)
			w.resutl <- msg * msg
		}
	}

}

func (w *Worker1Pool) AssignWork(i int){
	select{
	case <-w.ctx.Done():
		return
	case w.input<-i:
	}
}

func (w *Worker1Pool) FetchResults(){
	for {
		select{
		case <-w.ctx.Done():
			return
		case res, ok := <-w.resutl:
			if !ok{
				return
			}
			fmt.Printf("Recieved result %d \n", res)
		}
	}
}

func (w *Worker1Pool) Stop(){
	fmt.Println("Initiating shutdown.....")
	w.cancel()
	close(w.input)
	w.wg.Wait()
	close(w.resutl)
	w.done<-struct{}{}
	fmt.Println("shutdown completed..")
}

func main(){
	wp := NewWorker1Pool(3)
	wp.Start()

	//assign work
	go func(){
		for i:=0; i<20; i++{
			wp.AssignWork(i)
		}
	}()

	go func ()  {
		<-wp.done
		fmt.Println("Received done signal, worker pool has shut down")
	}()


	time.Sleep(5*time.Second)
	wp.Stop()

	time.Sleep(2*time.Second)

}