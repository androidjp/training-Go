package main

// 定义接口
type Phone interface {
	call()
}

type NokiaPhone struct {
}
// 实现call() 方法
func (nokiaPhone NokiaPhone) call() {
	println("I am Nokia")
}

type IPhone struct {
}

// 实现call() 方法
func (iPhone IPhone) call() {
	println("I am iPhone")
}

func main() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}
