package bio

import "time"

//这个文件主要包含各基本生物类型



//生命基本类型
type Life struct {
	//应该用各种共同属性，想写寿命来着。。。嗯。。发现需要为这个系统写一个时间系统
	Ilife
}



//微生物基本类型
type Microorganism struct {
	Imicroorganism
}

//植物基本类型
type Plant struct {
	Iplant
}

//动物基本类型
type Animal struct {
	Ianimal
	e time.Time
}






