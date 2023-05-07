package main

import (
	"fmt"
	"sync"
)

type singleton struct{}

func (s *singleton) Show() {
	fmt.Println("some function")
}

var once sync.Once

var instance *singleton

func GetInstance() *singleton {
	// go中的once保证只执行一次，并且是线程安全的，底层实现加了锁
	// 这里可以看底层实现，与singleton2的实现相似
	once.Do(func() {
		instance = new(singleton)
	})

	return instance
}

// 这里是为了实现方便，实际上单例放在单独的包中，只暴露出GetInstance接口
func main() {
	s1 := GetInstance()
	s1.Show()

	s2 := GetInstance()
	s2.Show()
}
