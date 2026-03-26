package proxy

import "fmt"

type Proxy interface {
	Do()
}

type Work struct {}

func (Work) Do() {
	fmt.Println("do something")
}

type ProxyClass struct {
	work Work
}

func (p ProxyClass) Do() {
	
	// 在调用真实对象工作之前的处理逻辑, 如 检查缓存, 权限校验, 实例化对象等
	fmt.Println("代理之前的处理逻辑...")

	p.work.Do()

	// 在调用真实对象工作之后的处理逻辑, 如缓存结果, 对结果进行处理等
	fmt.Println("代理之后的处理逻辑...")

}
