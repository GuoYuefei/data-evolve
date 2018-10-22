package chromosome



/**
①染色体与基因的关系：一条染色体上有许多基因，基因在染色体上呈直线排列。
②染色体与DNA的关系：每一条染色体上只有一个DNA分子，染色体是DNA分子的主要载体。
③DNA与基因的关系：每个DNA上有许多基因，基因是有遗传效应的DNA片段。
当然染色体应该是有一条DNA和蛋白质组成
 */

//染色体接口
type Ichromosome interface {
	//染色体可以自我复制
	Copy() Ichromosome
}

//一对染色体的接口 也就是里面应该有两个同源染色体
type Ichromosomes interface {
	//一对染色体应该能分裂成两个染色体
	Divide() (C1,C2 Ichromosome)
	//同源染色体之间应该能进行交叉互换     **这是也是出现生物差异性的一个源头
	//重新返回一对染色体
	Chiasmata()	Ichromosomes
}




