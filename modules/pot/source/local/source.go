package local

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/forbole/juno/v4/node/local"
	pottypes "github.com/stratosnet/stratos-chain/x/pot/types"

	potsource "github.com/forbole/bdjuno/v3/modules/pot/source"
)

var (
	_ potsource.Source = &Source{}
)

// Source implements govsource.Source by using a local node
type Source struct {
	*local.Source
	q pottypes.QueryServer
}

// NewSource returns a new Source instance
func NewSource(source *local.Source, potKeeper pottypes.QueryServer) *Source {
	return &Source{
		Source: source,
		q:      potKeeper,
	}
}

// Metrics implements potsource.Source
func (s Source) Metrics() (pottypes.Metrics, error) {
	ctx, err := s.LoadHeight(0)
	if err != nil {
		return pottypes.Metrics{}, fmt.Errorf("error while loading height: %s", err)
	}

	res, err := s.q.Metrics(sdk.WrapSDKContext(ctx), &pottypes.QueryMetricsRequest{})
	if err != nil {
		return pottypes.Metrics{}, err
	}

	return res.Metrics, nil
}
