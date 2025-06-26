// NOTE: This is not thread safe!
package main

import (
	"fmt"
)

var (
	uniqueInstance *Singleton = nil
)

type Singleton struct {
}

func newSingleton() *Singleton {
	fmt.Println("Classic Singleton is created!")

	return &Singleton{}
}

func (s *Singleton) DoJob() {
	fmt.Println("Classic Singleton is doing its job...")
}

func Instance() *Singleton {
	if uniqueInstance == nil {
		uniqueInstance = newSingleton()
	}

	return uniqueInstance
}
