package main

import (
	"fmt"
	"time"
)

// 时间 与 字符串的 转换与格式化
func main() {
	// ----------------------------------------------------
	// time 转 string
	// ----------------------------------------------------
	fmt.Println("------------------------------time 转 string")
	now := time.Now()
	fmt.Println(now.Format("2006-01-01"))     // YYYY-DD-MM
	fmt.Println(now.Format(time.RFC822))      // 02 Nov 20 20:04 CST
	fmt.Println(now.Format(time.Kitchen))     // 8:04PM
	fmt.Println(now.Format(time.UnixDate))    // Mon Nov  2 20:04:02 CST 2020
	var timeLayoutStr = "2006-01-02 15:04:05" //go中的时间格式化必须是这个时间
	//var timeLayoutStr = 2006/01/02 03:04:05 //合法, 格式可以改变
	//var timeLayoutStr = 2019/01/02 15:04:05 //不合法, 时间必须是2016年1月2号这个时间
	fmt.Println(now.Format(timeLayoutStr))

	// ----------------------------------------------------
	// string 转 time
	// ----------------------------------------------------
	fmt.Println("------------------------------string 转 time")

	t := time.Now() //当前时间
	t.Unix()        //时间戳

	ts := t.Format(timeLayoutStr) //time转string
	fmt.Println("time转string：", ts)
	st, _ := time.Parse(timeLayoutStr, ts) //string转time
	fmt.Println("string转time：", st)

	//在go中, 可以格式化一个带前后缀的时间字符串
	prefixTStr := "PREFIX-- 2019-01-01 -TEST- 10:31:12 --SUFFIX"       //带前后缀的时间字符串
	preTimeLayoutStr := "PREFIX-- 2006-01-02 -TEST- 15:04:05 --SUFFIX" //需要转换的时间格式, 格式和前后缀需要一致, 这种写法的限制很大, 但一些特殊场景可以用到
	prefixTime, _ := time.Parse(preTimeLayoutStr, prefixTStr)
	fmt.Println("带有前后缀的时间字符串，转换得到的时间：", prefixTime)

	//时间加减 time.ParseDuration()
	// such as "300ms", "-1.5h" or "2h45m".
	// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	at, _ := time.ParseDuration("2h") //2个小时后的时间, 负数就是之前的时间
	fmt.Println("两小时后的时间：", (t.Add(at)).Format(timeLayoutStr))

	//两个时间差
	sub := t.Sub(prefixTime)
	fmt.Println("两个时间差：", sub.Seconds()) //秒,  sub.Minutes()分钟,  sub.Hours()小时...
}
