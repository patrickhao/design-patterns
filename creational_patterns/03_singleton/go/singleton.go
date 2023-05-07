package main

import "fmt"

/*
* 三个要点
* 某个类只能有一个实例
* 它必须自行创建这个实例
* 必须自行向整个系统提供这个实例
*
* 保证一个类永远只有一个对象，这个对象还能被系统其他模块使用
*
* 什么时候使用，全局只能有一个实例，例如写文件时，一个时刻只能有一个线程写，不能同时写
 */

// 1 保证这个类非公有化，外界不同通过访问这个类直接创建一个对象
//   其他语言中可以将构造函数私有化实现
//   go中将类名称首字母小写，变成非公有访问，通过将单例设计成一个单独的包，包内可以访问

type singleton struct{}

func (s *singleton) Show() {
	fmt.Println("some function")
}

// 这里是饿汉式实现，不管需不需要用到，都在最开始实例化了类，会有一些多余的内存占用
var instance *singleton = new(singleton)

func GetInstance() *singleton {
	return instance
}

// 这里是为了实现方便，实际上单例放在单独的包中，只暴露出GetInstance接口
func main() {
	s1 := GetInstance()
	s1.Show()

	s2 := GetInstance()
	s2.Show()
}
