package loader_balance

import "errors"

var (
	ErrBalanceTypeNotExist = errors.New("balance type not exist") //负载均衡策略不存在
	ErrMissingNode         = errors.New("missing node")           //没有传node
)
