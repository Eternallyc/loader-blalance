package loader_balance

//定义结点
type Node struct {
	address string //地址加端口
	weight  int    //权重
}

func NewNode(address string, weight int) *Node {
	return &Node{
		address: address,
		weight:  weight,
	}
}
func (n Node) GetAddress() string {
	return n.address
}
func (n Node) GetWeight() int {
	return n.weight
}
