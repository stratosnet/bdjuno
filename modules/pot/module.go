package pot

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/forbole/callisto/v4/database"
	"github.com/forbole/callisto/v4/modules/pot/source"

	"github.com/forbole/juno/v5/modules"
	junomessages "github.com/forbole/juno/v5/modules/messages"
)

var (
	_ modules.Module = &Module{}
)

// Module represents the x/bank module
type Module struct {
	cdc codec.Codec
	db  *database.Db

	messageParser junomessages.MessageAddressesParser
	keeper        source.Source
}

// NewModule returns a new Module instance
func NewModule(
	messageParser junomessages.MessageAddressesParser, keeper source.Source, cdc codec.Codec, db *database.Db,
) *Module {
	return &Module{
		cdc:           cdc,
		db:            db,
		messageParser: messageParser,
		keeper:        keeper,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "pot"
}
