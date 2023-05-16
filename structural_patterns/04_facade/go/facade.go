package main

import "fmt"

type TV struct{}

func (t *TV) On() {
	fmt.Println("TV on")
}

type Light struct{}

func (l *Light) On() {
	fmt.Println("Light on")
}

type Xbox struct {
}

func (x *Xbox) On() {
	fmt.Println("XBox on")
}

// 通过外观模式，将各种功能独立的模块组合起来，提供方法整合各个独立的模块
type Facade struct {
	tv    TV
	light Light
	xbox  Xbox
}

func (f *Facade) DoMovie() {
	f.tv.On()
	f.light.On()
}

func (f *Facade) DoGame() {
	f.tv.On()
	f.light.On()
	f.xbox.On()
}

func main() {
	facade := Facade{}

	facade.DoMovie()

	fmt.Println("---------------------")

	facade.DoGame()
}
