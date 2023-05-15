package main

import "fmt"

// 适配的目标
type V5 interface {
	Use5V()
}

type Phone struct {
	v V5
}

// 适配器继承V5，因此可以提供5V电压，其本身对220V电压进行转换，对外提供5V电压
func NewPhone(v V5) *Phone {
	return &Phone{v}
}

func (p *Phone) Charge() {
	fmt.Println("charge phone")
	p.v.Use5V()
}

// 被适配的角色，适配者
type V220 struct{}

func (v *V220) Use220V() {
	fmt.Println("use 220v")
}

// 电源适配器
type Adapter struct {
	V5
	v220 *V220
}

func (a *Adapter) Use5V() {
	fmt.Println("use 220v charge phone")

	// 调用适配者方法，对外提供5V电压
	a.v220.Use220V()
}

func NewAdapter(v220 *V220) *Adapter {
	return &Adapter{v220: v220}
}

func main() {
	phone := NewPhone(NewAdapter(new(V220)))

	phone.Charge()
}
