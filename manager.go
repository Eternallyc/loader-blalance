package loader_balance

import "context"

var mgr = LoaderBalance{BalanceMap: map[BalanceType]Balance{}}

//负载均衡管理器
type LoaderBalance struct {
	BalanceMap map[BalanceType]Balance
}

func (l LoaderBalance) registerBalance(_ context.Context, typ BalanceType, balance Balance) {
	l.BalanceMap[typ] = balance
}

//注册负载均衡算法
func Register(ctx context.Context, typ BalanceType, balance Balance) {
	mgr.registerBalance(ctx, typ, balance)
}

func BalanceGetNode(ctx context.Context, typ BalanceType, nodes []Node) (Node, error) {
	if len(nodes) == 0 {
		return Node{}, ErrMissingNode
	}
	b, ok := mgr.BalanceMap[typ]
	if !ok {
		return Node{}, ErrBalanceTypeNotExist
	}
	if node, err := b.GetNode(ctx, nodes); err != nil {
		return Node{}, err
	} else {
		return node, nil
	}
}
