package random

import (
	"math/rand"
	"time"
)

//该函数有percent的值决定有多大机率返回true
//当输入为percent = 10.5时，返回值有10.5%的机率返回true
func PercentTrue(percent float32) bool {
	rand.Seed(time.Now().UnixNano())				//设置随机种子，当随机种子不变时因为随机算法相同所以得到的数其实也是相同的
	total := rand.Float32() * 100
	if total < percent {
		return true
	}else {
		return false
	}
}

//百万分之一是这个函数的最小概率值
func MillionTrue(million int) bool {
	return PercentTrue(1e-4 * float32(million))
}

//十亿分之一*billion
func BillionTure(billion int) bool {
	return PercentTrue(1e-7 * float32(billion))
}



