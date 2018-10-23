package dataTime

/**
注意 本包是模拟时间，和现实中的时间有出入
 */

import (
	"strconv"
)




type Time struct {
	//10秒作为本模拟时间的最小时间单位
	//这个值是从执行时间开始计算
	TenSec Duration
}

func NewTime(TS Duration) *Time {
	return &Time{TS}
}

//这个时间是一个虚拟时间，虚拟时间可能在现实时间上可能分布不均匀，这是有cpu调度决定的
//这个函数会返回一个从执行该程序开始到现在的一个虚拟时间
//当然不一定是从这次执行开始，将来程序会支持断点运行
//func Now() *Time {
//
//}


// A Duration represents the elapsed time between two instants
//不直接用time中的包是因为在print是time包中的Duration会被格式化，而这个格式化不适合虚拟时间
type Duration int64



const (
	//0.1s
	TenSecond Duration = 1
	Second 			= 10*TenSecond
	Minute          = 60 * Second
	Hour            = 60 * Minute
	Day             = 24 * Hour
	Week            = 7 * Day
	Month           = 30 * Day
	Year            = 12 * Month
	MY              = 1e6 * Year					//百万年
)

func (t *Time) Second() Duration {
	return t.TenSec / Second
}

func (t *Time)Minute() Duration {
	return t.TenSec / Minute
}

func (t *Time)Hour() Duration {
	return t.TenSec / Hour
}

func (t *Time)Day() Duration {
	return t.TenSec / Day
}

func (t *Time)Week() Duration {
	return t.TenSec / Week
}

func (t *Time)Month() Duration {
	return t.TenSec / Month
}

func (t *Time)Year() Duration {
	return t.TenSec / Year
}

//百万年 单位MY
func (t *Time)MillYear() Duration {
	return t.TenSec / MY
}

//为了输出方便，实现接口
func (t *Time)String() (timeStr string) {

	timeStr = timeStr + strconv.Itoa(int(t.MillYear())) + "-" + strconv.Itoa(int(t.Year() % 1e6)) + "-" +
		strconv.Itoa(int(t.Month() % 1e6 % 12)) + "-" + strconv.Itoa(int(t.Day() % 12e6 % 30)) + " " +
		strconv.Itoa(int(t.Hour() % 36e7 % 24)) + ":" + strconv.Itoa(int(t.Minute() % 864e7 % 60)) + ":" +
		strconv.Itoa(int(t.Second() % 5184e8 % 60)) + "." + strconv.Itoa(int(t.TenSec % 10))
	return
}












