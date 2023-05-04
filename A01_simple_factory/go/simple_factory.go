package main

import "fmt"

// --- 抽象层 ---
type Fruit interface {
	Show()
}

// --- 实现层 ---
type Apple struct {
	Fruit // go中无需显示继承接口，实现相应方法即可，这里是为了展示其中的继承关系
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

type Factory struct {
}

// 简单工厂的问题，该方法仍然是耦合的
// 且违背开闭原则，增加新的水果类需要修改当前代码，而无法通过新增代码解决问题
func (f *Factory) CreateFruit(kind string) Fruit {
	// 父类指针指向子类对象，为了实现多态
	// 这里可以直接返回，会进行隐式类型转换，这里是为了体现多态的实现方式
	var fruit Fruit

	if kind == "apple" {
		fruit = new(Apple)
	} else if kind == "banana" {
		fruit = new(Banana)
	}

	return fruit
}

// --- 业务逻辑层 ---
func main() {
	factory := Factory{}

	apple := factory.CreateFruit("apple")
	apple.Show()

	banana := factory.CreateFruit("banana")
	banana.Show()
}
