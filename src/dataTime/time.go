package dataTime

/**
注意 本包是模拟时间，和现实中的时间有出入
 */

import (
	"strconv"
)




type Time struct {
	//毫秒作为本模拟时间的最小时间单位
	//这个值是从执行时间开始计算
	Millisecond Duration
}

func NewTime(MS Duration) *Time {
	return &Time{MS}
}

//这个时间是一个虚拟时间，虚拟时间可能在现实时间上可能分布不均匀，这是有cpu调度决定的
//这个函数会返回一个从执行该程序开始到现在的一个虚拟时间
//当然不一定是从这次执行开始，将来程序会支持断点运行
//func Now() *Time {
//
//}


// A Duration represents the elapsed time between two instants
// as an int64 millisecond count. The representation limits the
type Duration int64



const (
	//tenth of a second
	Millisecond Duration = 1
	Second	= 1000 * Millisecond
	Minute	= 60 * Second
	Hour	= 60 * Minute
	Day 	= 24 * Hour
	Week	= 7 * Day
	Month	= 30 * Day
	Year	= 12 * Month
	MY 		= 1e6 * Year					//百万年
)

func (t *Time)Sec() Duration {
	return t.Millisecond / Second
}

func (t *Time)Minute() Duration {
	return t.Millisecond / Minute
}

func (t *Time)Hour() Duration {
	return t.Millisecond / Hour
}

func (t *Time)Day() Duration {
	return t.Millisecond / Day
}

func (t *Time)Week() Duration {
	return t.Millisecond / Week
}

func (t *Time)Month() Duration {
	return t.Millisecond / Month
}

func (t *Time)Year() Duration {
	return t.Millisecond / Year
}

//百万年 单位MY
func (t *Time)MillYear() Duration {
	return t.Millisecond / MY
}

//为了输出方便，实现接口
func (t *Time)String() (timeStr string) {

	timeStr = timeStr + strconv.Itoa(int(t.MillYear())) + "-" + strconv.Itoa(int(t.Year() % 1e6)) + "-" +
		strconv.Itoa(int(t.Month() % 1e6 % 12)) + "-" + strconv.Itoa(int(t.Day() % 12e6 % 30)) + " " +
		strconv.Itoa(int(t.Hour() % 36e7 % 24)) + ":" + strconv.Itoa(int(t.Minute() % 864e7 % 60)) + ":" +
		strconv.Itoa(int(t.Sec() % 5184e8 % 60)) + ":" + strconv.Itoa(int(t.Millisecond % 1000))
	return
}












