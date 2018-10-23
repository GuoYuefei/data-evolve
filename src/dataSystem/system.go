package dataSystem

import (
	"dataTime"
	"time"
)

/**
作为时间系统，需要有一个变量去记录当前时间
还必须有一个函数去计时，goroutine
 */

//系统时钟结构体
//需要对dataTime.Time进行再次封装
type SystemClock struct {
	//小写禁止访问 嗯读取应该要写一个读取器
	now dataTime.Time

	//模拟操作系统提供整个系统时间已经产生自己的时序
	//这个属性是用来提供模拟时间与外界时间的快慢比例
	//当该值为1000是 模拟：现实比例为1:1
	//2000 => 2：1
	//500 => 1:2
	//赋值时可以使用2 * dataSystem.UNIT
	Relative time.Duration

}

//默认构造
//now从零计时
//Relative默认是R125
func NewSystemClock() *SystemClock {
	return &SystemClock{*dataTime.NewTime(0),R250}
}

//私有变量now的访问器
func (SC SystemClock)Now() *dataTime.Time {
	return &SC.now
}

//给Relative赋值的常量选择
const (
	UNIT time.Duration = 1
	R1000 = UNIT << iota
	R500
	R250
	R125
	R63
	R31
	R17
	R8
	R4
	R2
	R1									//其实这个值是1024近似1000
	R0_5
	R0_25
	R0_125
)

//使用goroutine，需要开启一个计时系统
//当然最后这个函数可以控制虚拟时间行走的快慢
//快慢控制取决于结构体中Relative的值
func (SC *SystemClock)Walk() {
	go func() {
		for {

			time.Sleep(SC.Relative * time.Microsecond)				//现实N微秒时间
			SC.now.Millisecond++						//虚拟时间加1毫秒
		}
	}()
}


