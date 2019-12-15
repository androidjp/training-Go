package main

import "fmt"

// 集合
// 特点： 无序，默认为nil，内部使用hash表实现
// 初始化格式：
/**
声明变量，默认 map 是 nil
var map_variable map[key_data_type]value_data_type
使用 make 函数
map_variable := make(map[key_data_type]value_data_type)
*/

func main() {
	var countryCapitalMap map[string]string /*创建集合 */
	countryCapitalMap = make(map[string]string)

	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap [ "France" ] = "巴黎"
	countryCapitalMap [ "Italy" ] = "罗马"
	countryCapitalMap [ "Japan" ] = "东京"
	countryCapitalMap [ "India " ] = "新德里"

	/*使用键输出地图值 */
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap [country])
	}

	/*查看元素在集合中是否存在 */
	capital, ok := countryCapitalMap [ "American" ] /*如果确定是真实的,则存在,否则不存在 */
	/*fmt.Println(capital) */
	/*fmt.Println(ok) */
	if (ok) {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}

	println("-----------------------------------------")
	println("delete() 函数")
	println("-----------------------------------------")

	/* 创建map */
	countryCapitalMap2 := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}

	fmt.Println("原始地图")

	/* 打印地图 */
	for country := range countryCapitalMap2 {
		fmt.Println(country, "首都是", countryCapitalMap2 [ country ])
	}

	/*删除元素*/ delete(countryCapitalMap2, "France")
	fmt.Println("法国条目被删除")

	fmt.Println("删除元素后地图")

	/*打印地图*/
	for country := range countryCapitalMap2 {
		fmt.Println(country, "首都是", countryCapitalMap2 [ country ])
	}
}
