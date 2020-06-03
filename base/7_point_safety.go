package main

/**
关于Go指针：只要还有引用在引用这个指针，那么，指针就不会被垃圾回收
Go语言的内存回收机制规定，只要有一个指针指向引用一个变量，那么这个变量就不会给释放，因此，在Go语言中返回函数参数或临时变量，是安全的。

go 对局部变量的生命周期不做假设，而是根据是否被引用了 来决定对象 在堆 or 栈 上创建。此过程被称为 内存 escape。
*/
type Person struct {
	name string
}

func NewPerson(name string) *Person {
	return &Person{name: name}
}

func main() {
	xiaoMingMemoryAddress := NewPerson("小明")
	println(xiaoMingMemoryAddress)         // 查看指针的值：小明对象存储的位置
	println(&xiaoMingMemoryAddress)        // 查看指针对象本身的存储位置
	println((*xiaoMingMemoryAddress).name) // 查看指针指向的 小明对象的name
}
