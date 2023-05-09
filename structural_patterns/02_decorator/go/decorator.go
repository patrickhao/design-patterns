package main

import "fmt"

// --- 抽象层 ---
type IPizza interface {
	getPrice() int
}

// 装饰器，本来应该是接口，但是go的结构中不能有成员，因此这里实现为struct
type Decorator struct {
	pizza IPizza
}

// 需要子类去实现
func (d *Decorator) getPrice() int { return 0 }

// --- 实现层 ---
type VeggeMania struct{}

func (p *VeggeMania) getPrice() int {
	return 15
}

// 继承装饰器接口，装饰器也继承自IPizza，因此也可以使用IPizza调用装饰器的方法
type TomatoTopping struct {
	Decorator
}

func NewTomatoTopping(pizza IPizza) IPizza {
	return &TomatoTopping{Decorator{pizza: pizza}}
}

// 土豆装饰器在方法中增加额外的功能，这里是添加土豆的价格
func (c *TomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 7
}

// 芝士装饰器
type CheeseTopping struct {
	Decorator
}

func NewCheeseTopping(pizza IPizza) IPizza {
	return &CheeseTopping{Decorator{pizza: pizza}}
}

// 增加芝士的价格
func (c *CheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 9
}

// --- 逻辑层 ---
func main() {
	pizza := &VeggeMania{}

	fmt.Println("pizza price", pizza.getPrice())

	// 加入土豆装饰器
	var pizzaWithTomato IPizza
	pizzaWithTomato = NewTomatoTopping(pizza)

	fmt.Println("pizza with tomato price", pizzaWithTomato.getPrice())

	// 加入芝士装饰器
	var pizzaWithCheese IPizza
	pizzaWithCheese = NewCheeseTopping(pizza)

	fmt.Println("pizza with cheese price", pizzaWithCheese.getPrice())

	// 装饰器与代理模式的不同，装饰器可以随意组合，因为各种装饰器和产品都是同一个接口，可以叠加装饰器，功能也叠加
	var pizzaWithCheeseAndTomato IPizza
	// 传入的是带芝士的装饰器，最后得到的是带芝士和土豆的
	pizzaWithCheeseAndTomato = NewTomatoTopping(pizzaWithCheese)

	fmt.Println("pizza with cheese and tomato price", pizzaWithCheeseAndTomato.getPrice())
}
