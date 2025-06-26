package main

import (
	"fmt"
	"sync"
)

var (
	uniqueInstance *Singleton = nil
	mu                        = &sync.Mutex{}
)

type Singleton struct {
}

func newSingleton() *Singleton {
	fmt.Println("Thread safe Singleton is created!")

	return &Singleton{}
}

func (s *Singleton) DoJob() {
	fmt.Println("Thread safe Singleton is doing its job...")
}

func Instance() *Singleton {
	mu.Lock()
	defer mu.Unlock()

	if uniqueInstance == nil {
		uniqueInstance = newSingleton()
	}

	return uniqueInstance
}
