package creational

import (
	"fmt"
	"sync"
)

type Singleton struct {
}

var locker = &sync.Mutex{}
var instance *Singleton

func GetInstance() *Singleton {
	locker.Lock()
	defer locker.Unlock()

	if instance == nil {
		instance = &Singleton{}
		fmt.Println("creating instance")
	}

	fmt.Println("instance created")

	return instance
}
