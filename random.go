package loader_balance

/**
https://blog.csdn.net/baidu_17508977/article/details/81233415?utm_source=blogxgwz0
*/
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
	randInt := rand.Intn(len(nodes))
	return nodes[randInt], nil
}
