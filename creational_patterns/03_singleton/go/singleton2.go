package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type singleton struct{}

func (s *singleton) Show() {
	fmt.Println("some function")
}

var instance *singleton

// 标记，原子操作
var initialized uint32

// 锁
var lock sync.Mutex

// func GetInstance() *singleton {
// 	// 懒汉式，在第一次调用时才实例化类
//
// 	// 问题，两个goroutine并发时，可能此时instance都为空，创建了两个实例
// 	// 因此需要加锁，增加互斥锁
// 	lock.Lock()
// 	defer lock.Unlock()
//
// 	if instance == nil {
// 		instance = new(singleton)
// 		// 设置标记为1
// 		atomic.StoreUint32(&initialized, 1)
// 	}
//
// 	return instance
// }

func GetInstance() *singleton {
	// 懒汉式，在第一次调用时才实例化类
	// 防止每次GetInstance时都加锁导致的性能问题
	// 默认值是0，第一次不会进该if
	// 如果值为1，则表示该实例已经有了，则无需加锁申请，直接返回皆可
	// 原子操作只读，无需加锁，性能更好
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	// initialized值为0，当前还没有该实例，为了防止多线程同时申请时的问题
	// 加锁之后再申请，即通过维护initialized，只在实例没有初始化的时候加锁保护
	// 已经初始化之后直接返回即可
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		instance = new(singleton)
		// 设置标记为1
		atomic.StoreUint32(&initialized, 1)
	}

	return instance
}

// 这里是为了实现方便，实际上单例放在单独的包中，只暴露出GetInstance接口
func main() {
	s1 := GetInstance()
	s1.Show()

	s2 := GetInstance()
	s2.Show()
}
