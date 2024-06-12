package remote

import (
	"github.com/forbole/juno/v5/node/remote"
	pottypes "github.com/stratosnet/stratos-chain/x/pot/types"

	potsource "github.com/forbole/callisto/v4/modules/pot/source"
)

var (
	_ potsource.Source = &Source{}
)

// Source implements govsource.Source using a remote node
type Source struct {
	*remote.Source
	potClient pottypes.QueryClient
}

// NewSource returns a new Source implementation
func NewSource(source *remote.Source, potClient pottypes.QueryClient) *Source {
	return &Source{
		Source:    source,
		potClient: potClient,
	}
}

// Metrics implements potsource.Source
func (s Source) Metrics() (pottypes.Metrics, error) {
	res, err := s.potClient.Metrics(
		remote.GetHeightRequestContext(s.Ctx, 0),
		&pottypes.QueryMetricsRequest{},
	)
	if err != nil {
		return pottypes.Metrics{}, err
	}

	return res.Metrics, err
}
