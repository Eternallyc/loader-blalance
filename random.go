package loader_balance

import (
	"context"
	"math/rand"
)

func init() {
	Register(context.Background(), Random, RandomBalance{})
}

type RandomBalance struct {
}

func (r RandomBalance) GetNode(ctx context.Context, nodes []Node) (Node, error) {
	//随机获得node
	return nodes[rand.Intn(len(nodes))], nil
}
