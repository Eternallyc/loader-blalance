package loader_balance

import (
	"context"
	"math/rand"
)

func init() {
	Register(context.Background(), WeightRand, RoundRobinBalance{})
}

type WeightRandBalance struct {
}

func (r WeightRandBalance) GetNode(ctx context.Context, nodes []Node) (Node, error) {
	weightSum := 0
	for _, i2 := range nodes {
		weightSum += i2.weight
	}
	randInt := rand.Intn(weightSum)
	sum := 0
	for i, i2 := range nodes {
		sum += i2.weight
		if sum >= randInt {
			return nodes[i], nil
		}
	}

	return nodes[0], nil
}
