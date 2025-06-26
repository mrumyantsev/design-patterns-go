package main

import "sync"

var (
	uniqueInstance *Singleton = nil
	mu                        = &sync.Mutex{}
)

type Singleton struct {
	// other useful instance variables here
}

func newSingleton() *Singleton {
	return &Singleton{}
}

// other useful methods here
func (s *Singleton) GetDescription() string {
	return "I'm a thread safe Singleton!"
}

func GetInstance() *Singleton {
	mu.Lock()
	defer mu.Unlock()

	if uniqueInstance == nil {
		uniqueInstance = newSingleton()
	}

	return uniqueInstance
}
