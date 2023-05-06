package main

import "fmt"

// 简单工厂模式 + 开闭原则 = 工厂模式
// 将工厂类抽象出来，各个工厂专注于创建某个类
// --- 抽象层 ---
type Fruit interface {
	Show()
}

type AbstractFactory interface {
	CreateFruit() Fruit // 返回的是接口，指向具体水果实例
}

// --- 实现层 ---
type Apple struct {
	Fruit
}

func (apple *Apple) Show() {
	fmt.Println("this is apple")
}

type Banana struct {
	Fruit
}

func (banana *Banana) Show() {
	fmt.Println("this is banana")
}

type AppleFactory struct {
	AbstractFactory
}

func (f *AppleFactory) CreateFruit() Fruit {
	var fruit Fruit
	fruit = new(Apple)

	return fruit
}

type BananaFactory struct {
	AbstractFactory
}

func (f *BananaFactory) CreateFruit() Fruit {
	var fruit Fruit
	fruit = new(Banana)

	return fruit
}

// --- 业务逻辑层 ---
func main() {
	// 面向抽象接口编程，使用接口而不是具体的类
	var appleFac AbstractFactory
	appleFac = new(AppleFactory)

	var apple Fruit
	apple = appleFac.CreateFruit()

	apple.Show()

	// 另一种更简单的写法
	var bananaFac AbstractFactory
	// 这里的bananaFac只能是AbstractFactory，其指向的类是其具体实现
	bananaFac = new(BananaFactory)

	banana := bananaFac.CreateFruit()
	banana.Show()
}
