package dna

import (
	"bio/character"
)

//这个文件主要写了各种基因接口

//一个DNA的总接口
type DNA interface {
	Dominance() bool				//返回是否显性，显性就true
	Character() character.Character			//DNA解析后应该能返回一个特性接口
}


