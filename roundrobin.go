package loader_balance

import (
	"context"
)

func init() {
	Register(context.Background(), Random, RandomBalance{})
}

type RoundRobinBalance struct {
	curIndex int
}

func (r RoundRobinBalance) GetNode(ctx context.Context, nodes []Node) (Node, error) {

	length := len(nodes)
	if r.curIndex >= length {
		r.curIndex = 0
	}
	node := nodes[r.curIndex]
	r.curIndex++
	//随机获得node
	return node, nil
}
