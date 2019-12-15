package main

import "fmt"

// 结构体
// Go 当中，数组可以存储同一类型的数据，但是，如果是结构体，我们可以给不同项存储不同数据类型
// 相当于 Java 的类，一个类有一到多个属性，每个属性之间可以有不同的数据类型
type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	//println("创建一个结构体Books：", Books{"Go 语言", "www.xxx.com", "Go 教程", 1232})
	bookA := Books{"Go 语言", "www.xxx.com", "Go 教程", 1232}
	bookB := Books{title: "Python 语言", book_id: 4444}
	fmt.Printf("bookA %v\n", bookA)
	fmt.Printf("bookB %v\n", bookB)

	println("结构体指针", &bookA)
	println("结构体也能作为形参")
	printBook(bookA)
}

func printBook(book Books) {
	//println("printBook", book.title, book.author)
	fmt.Printf("%v\n", book)
}
