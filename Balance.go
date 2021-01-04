package loader_balance

import "context"

//负载均衡算法接口
type Balance interface {
	//根据具体的算法来获得 用语法糖可以使客户端调用此方法时不用多传参数
	GetNode(ctx context.Context, nodes []Node) (Node, error)
}
