package main

import "fmt"

type Goods struct {
	Kind string // 商品的种类
	Fact bool   // 商品的真伪
}

// --- 抽象层 ---
type Shopping interface {
	Buy(goods *Goods)
}

// --- 实现层 ---
type KoreaShpping struct{}

func (ks *KoreaShpping) Buy(goods *Goods) {
	fmt.Println("go shopping in korea, buy", goods.Kind)
}

type AmericanShopping struct{}

func (as *AmericanShopping) Buy(goods *Goods) {
	fmt.Println("go shopping in american, buy", goods.Kind)
}

// 符合开闭
// 这里面向同一个接口来实现代理，即OverseasProxy也是Shopping类型
// 因此调用是和直接调用Shopping中方法一致，只是提供了额外的功能
type OverseasProxy struct {
	shopping Shopping
}

func (op *OverseasProxy) distinguish(goods *Goods) bool {
	// 1 辨别真伪
	fmt.Println("distinguish authenticity", goods.Kind)
	if goods.Fact == false {
		fmt.Println("false")
	}

	return goods.Fact
}

func (op *OverseasProxy) check(goods *Goods) {
	fmt.Println("customs inspection", goods.Kind)
}

func (op *OverseasProxy) Buy(goods *Goods) {
	// 1 辨别真伪
	if op.distinguish(goods) {
		// 2 调用具体要被代理的购物方式的Buy()方法
		op.shopping.Buy(goods)
		// 3 海关检查
		op.check(goods)
	}
}

// 创建一个代理，并配置被代理的主题
func NewProxy(shopping Shopping) Shopping {
	return &OverseasProxy{shopping}
}

// --- 逻辑层 ---
func main() {
	g1 := Goods{
		Kind: "apple",
		Fact: true,
	}

	g2 := Goods{
		Kind: "abibas",
		Fact: false,
	}

	var kShopping Shopping
	kShopping = new(KoreaShpping)

	var proxy Shopping
	proxy = NewProxy(kShopping)

	proxy.Buy(&g1)
	proxy.Buy(&g2)
}
