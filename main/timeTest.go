package main

import (
	"fmt"
	"time"
)

func main() {

	//now := time.Now()
	//fmt.Println(now)
	//fmt.Println(now.Year())
	//fmt.Println(now.YearDay())
	//fmt.Println(now.Weekday())
	//
	//fmt.Println(now.Unix())
	//t := time.Unix(now.Unix(),-1)
	//fmt.Println(t)
	//format := now.Format("2006-01-02 15:04:00")
	//fmt.Println(format)
	//
	//s := time.Now().Format("2006/01/02 15:04:00")
	//fmt.Println(s)

	location, err := time.LoadLocation("Asia/Shanghai")
	if nil != err {
		fmt.Println(err)
	}
	fmt.Println(location)
	timea, err := time.ParseInLocation("2006-01-02 15:04:00", time.Now().String(), location)
	if err != nil {
	}
	fmt.Println(timea)
	second := time.Second
	second += time.Duration(25)
	//
	//fmt.Println(now.Month())
	//fmt.Println(now.Minute())
	//fmt.Println(now.Day())
	//
	//
	//fmt.Println(now.Unix())
	//fmt.Println(now.UnixNano())

	//tick := time.Tick(time.Second * 2)
	//i := 0
	//for _ = range tick {
	//	if i == 5 {
	//		break
	//	}
	//	fmt.Println("我砍！！！")
	//	i++
	//}
}
