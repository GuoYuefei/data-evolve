package main

import "fmt"

func main() {
	c := NewCC(12)
	c.DoSome()
	c.DoSome1(21)
	//cc := NewCc()
	////虽然能被实例化，但是没有写下面的函数，所以无实际意义
	//cc.DoSome1(12)
	//cc.DoSome()
	var iq Iq = c
	iq.DoSome()
	var icc Icc = c
	icc.DoSome1(22)
}


type Iq interface {
	DoSome()
}

type Icc interface {
	Iq
	DoSome1(i int)
}

type Itt interface {
	Iq
	DoSome2()
}

type Cc struct {
	Icc
}

type CC struct {
	i int
	Cc
}

func (c *CC)DoSome(){
	fmt.Println("i = ",c.i)
}

func (c *CC)DoSome1(i int) {
	fmt.Println("i = ",i)
}

func NewCC(i int) *CC {
	return &CC{i,Cc{}}
}

func NewCc() *Cc {
	return &Cc{}
}
