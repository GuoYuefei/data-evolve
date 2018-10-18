## 这是一个随机包

生物进化是建造在大量偶然性上的必然结果，所以必须要引入大量的随机函数。系统包中虽然提供了很多随机函数，但是对于我们需要重复使用的代码还是有必要进行一定量的封装的。

[TOC]



#### func PercentTrue(percent float32) bool

返回一个随机布尔量，这个布尔量有percent%的机率是true的。也就是说有1 - percent%的机率是false的。



#### func MillionTrue(million int) bool

返回一个随机布尔量，这个布尔量有百万分之million的可能是true。


#### func BillionTure(billion int) bool

返回一个随机布尔量，这个布尔量有十亿分之billion的可能是true。