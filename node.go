package loader_balance

//定义结点
type Node struct {
	ip   string
	port int
}

func NewNode(ip string, port int) *Node {
	return &Node{
		ip:   ip,
		port: port,
	}
}
func (n Node) GetIp() string {
	return n.ip
}
func (n Node) GetPort() int {
	return n.port
}
