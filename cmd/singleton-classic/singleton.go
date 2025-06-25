package main

// NOTE: This is not thread safe!

var (
	uniqueInstance *Singleton = nil
)

type Singleton struct {
}

func newSingleton() *Singleton {
	return &Singleton{}
}

// other useful methods here
func (s *Singleton) GetDescription() string {
	return "I'm a classic Singleton!"
}

func GetInstance() *Singleton {
	if uniqueInstance == nil {
		uniqueInstance = newSingleton()
	}

	return uniqueInstance
}
