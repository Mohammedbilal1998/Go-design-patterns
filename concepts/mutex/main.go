package main

import(
	"fmt"
	"sync"
)

type Container struct {
	Mu sync.Mutex
	Counters map[string]int
}

func (c *Container) Inc(name string) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.Counters[name]++
}

func NewContainer(counters map[string]int) *Container{
	return &Container{Counters: counters}
}

func main(){
	var wg sync.WaitGroup
	c := NewContainer(map[string]int{})

	doInc := func(name string, n int){
		defer wg.Done()
		for i:=0; i<n; i++ {
			c.Inc(name)
		}
	}
	wg.Add(3)
	go doInc("a", 1000)
	go doInc("b", 1000)
	go doInc("a", 1000)

	wg.Wait()
	fmt.Println(c)

}