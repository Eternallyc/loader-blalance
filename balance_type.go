package loader_balance

//定义枚举
type BalanceType = string

var (
	Random     BalanceType = "Random"     //随机
	RoundRobin BalanceType = "RoundRobin" //轮询
	WeightRand BalanceType = "WeightRand" //加权随机
)
