package singleton

import (
	"sync"
)

// var lock = &sync.Mutex{}

// type Single struct {
// }

// var singleInstance *Single

// func GetInstance() *Single {
// 	if singleInstance == nil {
// 		lock.Lock()
// 		defer lock.Unlock()
// 		if singleInstance == nil {
// 			fmt.Println("Creating single instance now.")
// 			singleInstance = &Single{}
// 		} else {
// 			fmt.Println("Single instance already created.")
// 		}
// 	} else {
// 		fmt.Println("Single instance already created.")
// 	}

// 	return singleInstance
// }

//Singleton 是单例模式类
type Singleton struct{}

var singleton *Singleton
var once sync.Once

//GetInstance 用于获取单例模式对象
func GetInstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{}
	})

	return singleton
}
