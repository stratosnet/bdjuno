package source

import pottypes "github.com/stratosnet/stratos-chain/x/pot/types"

type Source interface {
	Metrics() (pottypes.Metrics, error)
}
