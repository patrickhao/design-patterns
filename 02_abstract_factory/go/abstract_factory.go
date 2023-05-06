package main

import "fmt"

// 合理划分产品等级结构和产品族
// 划分后产品等级结构不要改变，产品簇可做增加
// 否则不适合使用抽象工厂模式

// --- 抽象层 ---
type AbstractApple interface {
	ShowApple()
}

type AbstractBanana interface {
	ShowBanana()
}

type AbstractFactory interface {
	CreateApple() AbstractApple
	CreateBanana() AbstractBanana
}

// --- 实现层 ---
// 产品簇可以横向开辟，不影响产品
// 针对产品簇进行添加，符合开闭原则
// 如果涉及到产品等级结构修改，则不符合开闭原则，而且这种情况下也不适合使用抽象工厂方法

// 中国产品簇
type ChinaApple struct{}

func (ca *ChinaApple) ShowApple() {
	fmt.Println("china apple")
}

type ChinaBanana struct{}

func (cb *ChinaBanana) ShowBanana() {
	fmt.Println("china banana")
}

type ChinaFactory struct{}

func (cf *ChinaFactory) CreateApple() AbstractApple {
	var apple AbstractApple
	apple = new(ChinaApple)
	return apple
}

func (cf *ChinaFactory) CreateBanana() AbstractBanana {
	var banana AbstractBanana
	banana = new(ChinaBanana)
	return banana
}

// 日本产品簇
type JapanApple struct{}

func (ja *JapanApple) ShowApple() {
	fmt.Println("japan apple")
}

type JapanBanana struct{}

func (jb *JapanBanana) ShowBanana() {
	fmt.Println("japan banana")
}

type JapanFactory struct{}

func (jf *JapanFactory) CreateApple() AbstractApple {
	var apple AbstractApple
	apple = new(JapanApple)
	return apple
}

func (jf *JapanFactory) CreateBanana() AbstractBanana {
	var banana AbstractBanana
	banana = new(JapanBanana)
	return banana
}

// --- 逻辑层 ---
func main() {
	cFac := new(ChinaFactory)

	var chinaApple AbstractApple
	chinaApple = cFac.CreateApple()
	chinaApple.ShowApple()

	jFac := new(JapanFactory)

	var japanApple AbstractApple
	japanApple = jFac.CreateApple()
	japanApple.ShowApple()
}
