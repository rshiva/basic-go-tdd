package main

import "sync"

//Use channels when passing ownership of data 
//Use mutexes for managing state


type Counter struct {
	//A Mutex must not be copied after first use.
	//Mutex allows us to add locks to our data
	mu sync.Mutex
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc(){
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int{
	return c.value

}