package main

import "fmt"

// 降低Cooker中业务逻辑的复杂度，其只用专注于业务，与场景中的业务场景分离
// 并且可以在调用者中将命令组合起来，形成宏命令，完成更复杂的业务逻辑
type Cooker struct {
}

func (c *Cooker) MakeChuan() {
	fmt.Println("烤肉师傅烤了羊肉串")
}

func (c *Cooker) MakeChicken() {
	fmt.Println("烤肉师傅烤了鸡肉串")
}

// 抽象的命令
type Command interface {
	Make()
}

type CommandCookChuan struct {
	cooker *Cooker
}

func (cmd *CommandCookChuan) Make() {
	cmd.cooker.MakeChuan()
}

type CommandCookChicken struct {
	cooker *Cooker
}

func (cmd *CommandCookChicken) Make() {
	cmd.cooker.MakeChicken()
}

// 命令的调用者
type Waiter struct {
	CmdList []Command
}

func (w Waiter) Notify() {
	if w.CmdList == nil {
		return
	}

	// 命令调用者批量执行方法，Cooker与具体业务场景解耦，由调用者批量执行
	for _, cmd := range w.CmdList {
		cmd.Make()
	}
}

func main() {
	cooker := new(Cooker)

	cmdChuan := CommandCookChuan{cooker: cooker}
	cmdChicken := CommandCookChicken{cooker: cooker}

	w := new(Waiter)

	// 批量发送订单，其调用Cooker执行
	// 因为这里是接口，因此传入指针而不是具体的类的实例
	// 这里多态的实现方式和c++中一致，面向指针的多态而不是具体实例
	w.CmdList = append(w.CmdList, &cmdChuan)
	w.CmdList = append(w.CmdList, &cmdChicken)

	w.Notify()
}
