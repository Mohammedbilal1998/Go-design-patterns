package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type Job struct {
	ID     int
	Input  int
	Result int
}

type WorkerPool struct {
	numWorkers int
	jobs       chan Job
	results    chan Job
	done       chan struct{}
	wg         sync.WaitGroup
	ctx        context.Context
	cancel     context.CancelFunc
}

func NewWorkerPool(numWorkers int) *WorkerPool {
	ctx, cancel := context.WithCancel(context.Background())

	return &WorkerPool{
		numWorkers: numWorkers,
		jobs: make(chan Job),
		results: make(chan Job),
		done: make(chan struct{}),
		ctx: ctx,
		cancel: cancel,
	}
}

func (wp *WorkerPool) Start(){
	for i:=0; i<wp.numWorkers; i++{
		wp.wg.Add(1)
		go wp.Worker(i)
	}
	go wp.CollectResults()

}

func (wp *WorkerPool) Worker(i int){
	defer wp.wg.Done()

	for{
		select{
		case <-wp.ctx.Done():
			fmt.Printf("Worker %d shutting down!\n", i)
			return
		case job, ok := <-wp.jobs:
			if !ok{
				return
			}
			fmt.Printf("Worker %d started job %v\n", i, job)
			time.Sleep(time.Second)
			fmt.Printf("Worker %d finished job %v\n", i, job)

			select{
			case wp.results <-job:
			case <-wp.ctx.Done():
				return
			}
		}
	}
}

func (wp *WorkerPool) SubmitJob(jobInput int){
	select{
	case <-wp.ctx.Done():
		return
	case wp.jobs <- Job{ID: jobInput, Input: jobInput}:
	}
}

func (wp *WorkerPool) CollectResults() {
	for{
		select{
		case <-wp.ctx.Done():
			return
		case result, ok := <- wp.results:
			if !ok{
				return
			}
			fmt.Printf("Result Recieved %d\n", result)
		}
	}
}

func (wp *WorkerPool) Stop(){
	fmt.Println("Initiating shutdown...")
	wp.cancel()
	close(wp.jobs)
	wp.wg.Wait()
	close(wp.results)
	wp.done <- struct{}{} 
	fmt.Println("Shutdown completed")
}

func main(){
	wp := NewWorkerPool(3)

	wp.Start()

	go func() {
		for i := 0; i < 5; i++ {
			wp.SubmitJob(i) 
		}
	}()

	go func() {
		<-wp.done
		fmt.Println("Received done signal, worker pool has shut down")
	}()

	time.Sleep(2 * time.Second)
	wp.Stop()
	
	// Wait a moment to see the done signal handling
	time.Sleep(1 * time.Second)
	
}