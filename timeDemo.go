package main

import (
	"fmt"
	"time"
)

func timestampDemo2(timestamp int64) {
	timeObj := time.Unix(timestamp, 0) //将时间戳转为时间格式
	fmt.Println(timeObj)
	year := timeObj.Year()     //年
	month := timeObj.Month()   //月
	day := timeObj.Day()       //日
	hour := timeObj.Hour()     //小时
	minute := timeObj.Minute() //分钟
	second := timeObj.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

func tick() {
	tick := time.Tick(time.Second * 2) //定义一个2秒执行一次的定时器,tick本质上是一个channel
	i := 0
	for _ = range tick {
		i++
		if i > 3 {
			break
		}
		fmt.Println("每隔2秒打印一次")
	}
}

func formatDemo() {
	now := time.Now()
	// 格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	// 12小时制
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}

func standard() {
	now := time.Now()
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	i, month2, day2 := now.Date()
	weekday := now.Weekday()
	year2, week := now.ISOWeek()
	timestamp1 := now.Unix()           //时间戳
	timestamp2 := now.UnixNano() / 1e6 //毫秒时间戳
	fmt.Println(now)
	fmt.Println(year)
	fmt.Println(month)
	fmt.Println(day)
	fmt.Println(hour)
	fmt.Println(minute)
	fmt.Println(second)
	fmt.Println(weekday)
	fmt.Println(timestamp1)
	fmt.Println(timestamp2)
	fmt.Println()
	fmt.Printf("i=%v;month2=%v;day2=%v", i, month2, day2)
	fmt.Println()
	fmt.Printf("year2=%v;week=%v", year2, week)
	fmt.Println()

}

func parseIn() {
	now := time.Now()
	fmt.Println(now)
	// 加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2019/08/04 14:15:20", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Printf("%T", timeObj)
	fmt.Println()
	fmt.Println(timeObj.Sub(now))
}

func main() {
	//基础方法
	fmt.Println("基础方法")
	standard()
	fmt.Println()

	//放时间戳戳转换时间
	fmt.Println("放时间戳戳转换时间")
	timestampDemo2(1514736000)
	fmt.Println()

	//时间格式化,时间类型有一个自带的方法Format进行格式化，需要注意的是Go语言中格式化时间模板不是常见的Y-m-d H:M:S
	// 而是使用Go的诞生时间2006年1月2号15点04分（记忆口诀为2006 1 2 3 4）。也许这就是技术人员的浪漫吧。
	fmt.Println("时间格式化")
	formatDemo()
	fmt.Println()

	//解析字符串
	fmt.Println("解析字符串")
	parseIn()
	fmt.Println()

	//定时器
	fmt.Println("定时器")
	tick()
	fmt.Println()
}
