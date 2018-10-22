package bio

import "environment/material"


//应该是可以把生物看作一种material.Imaterial


type Ilife interface {
	//生物都有摄取外界必要物质的能力，在高级生物上大多以吃喝为主
	Intake(IM material.Imaterial)
	//一个繁衍能力
	Breeding() Ilife

	//繁殖条件的判断
	//满足条件就允许繁殖
	CanBreed() bool


	//繁殖周期的判断
	//微生物可以用繁殖周期去判定是否可以繁殖，而动、植物则需要达到成年状态后才能有一个繁殖周期
	//同时满足繁殖条件和繁殖周期为true是才可以繁殖
	//本函数的周期数据提供者应该是各个实现类型的状态量
	BreedCycle() bool

}

//微生物接口
type Imicroorganism interface {
	Ilife
}

type ICell interface {
	Ilife
}


type Ianimal interface {
	Ilife
	Eat(ilife Ilife)						//这个eat过程其实包含kill life的过程
	Drink(IM material.Imaterial)			//这个形参应该以后会变化
	Growing()					//成长可能要开一个goroutine
}


type Iplant interface {
	Ilife
	Growing()					//植物也需要一个成长过程
	//光合作用，这个过程是一个复杂的 过程，设计到很多方面
	//这个过程设计到氧气、二氧化碳的变化、还设计到阳光的摄取以及植物本身光合作用的各项能力
	Photosynthesis()
}