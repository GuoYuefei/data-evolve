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
	return &SystemClock{*dataTime.NewTime(0),R50}
}

//私有变量now的访问器
func (SC SystemClock)Now() *dataTime.Time {
	return &SC.now
}

//给Relative赋值的常量选择
const (
	UNIT time.Duration = 1
	R50 = UNIT << iota					//这里的iota=1
	R25
	R12
	R6
	R3
	R1 						//其实这个值是512近似500
	R0_8
	R0_4
	R0_2
	R0_1
)

//使用goroutine，需要开启一个计时系统
//当然最后这个函数可以控制虚拟时间行走的快慢
//快慢控制取决于结构体中Relative的值
//快慢的控制是不精确的，因为这需要考虑到系统调度问题
//虽然精度不高但是不同的参数也是有区别的 这里的时间流速参数相当于只是一个挡位而已
func (SC *SystemClock)Walk() {
	go func() {
		for {

			time.Sleep(SC.Relative * time.Millisecond) //现实N毫秒时间
			SC.now.TenSec++                            //虚拟时间加0.1秒
		}
	}()
}


