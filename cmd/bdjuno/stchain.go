package main

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	junomessages "github.com/forbole/juno/v4/modules/messages"
	pottypes "github.com/stratosnet/stratos-chain/x/pot/types"
	//profilestypes "github.com/desmos-labs/desmos/v2/x/profiles/types"
	registertypes "github.com/stratosnet/stratos-chain/x/register/types"
	sdstypes "github.com/stratosnet/stratos-chain/x/sds/types"
)

// stchainMessageAddressesParser represents a parser able to get the addresses of the involved
// account from a stchain message
var stchainMessageAddressesParser = junomessages.JoinMessageParsers(
	registerMessageAddressesParser,
	potMessageAddressesParser,
	sdsMessageAddressesParser,
)

// registerMessageAddressesParser represents a MessageAddressesParser for the x/register module
func registerMessageAddressesParser(_ codec.Codec, cosmosMsg sdk.Msg) ([]string, error) {
	switch msg := cosmosMsg.(type) {

	case *registertypes.MsgCreateResourceNode:
		return []string{msg.NetworkAddress, msg.OwnerAddress}, nil

	case *registertypes.MsgCreateMetaNode:
		return []string{msg.NetworkAddress, msg.OwnerAddress}, nil

	case *registertypes.MsgRemoveResourceNode:
		return []string{msg.ResourceNodeAddress, msg.OwnerAddress}, nil

	case *registertypes.MsgRemoveMetaNode:
		return []string{msg.MetaNodeAddress, msg.OwnerAddress}, nil

	case *registertypes.MsgUpdateResourceNode:
		return []string{msg.NetworkAddress, msg.OwnerAddress}, nil

	case *registertypes.MsgUpdateMetaNode:
		return []string{msg.NetworkAddress, msg.OwnerAddress}, nil

	case *registertypes.MsgUpdateResourceNodeStake:
		return []string{msg.NetworkAddress, msg.OwnerAddress}, nil

	case *registertypes.MsgUpdateMetaNodeStake:
		return []string{msg.NetworkAddress, msg.OwnerAddress}, nil

	case *registertypes.MsgMetaNodeRegistrationVote:
		return []string{msg.CandidateNetworkAddress, msg.CandidateOwnerAddress,
			msg.VoterNetworkAddress, msg.VoterOwnerAddress}, nil
	}

	return nil, junomessages.MessageNotSupported(cosmosMsg)
}

// profilesMessageAddressesParser represents a MessageAddressesParser for the x/pot module
func potMessageAddressesParser(_ codec.Codec, cosmosMsg sdk.Msg) ([]string, error) {
	switch msg := cosmosMsg.(type) {

	case *pottypes.MsgVolumeReport:
		return []string{msg.Reporter, msg.ReporterOwner}, nil

	case *pottypes.MsgWithdraw:
		return []string{msg.WalletAddress, msg.TargetAddress}, nil

	case *pottypes.MsgFoundationDeposit:
		return []string{msg.From}, nil

	case *pottypes.MsgSlashingResourceNode:
		return []string{msg.NetworkAddress, msg.WalletAddress}, nil
	}

	return nil, junomessages.MessageNotSupported(cosmosMsg)
}

// sdsMessageAddressesParser represents a MessageAddressesParser for the x/sds module
func sdsMessageAddressesParser(_ codec.Codec, cosmosMsg sdk.Msg) ([]string, error) {
	switch msg := cosmosMsg.(type) {

	case *sdstypes.MsgFileUpload:
		return []string{msg.From, msg.Reporter, msg.Uploader}, nil

	case *sdstypes.MsgPrepay:
		return []string{msg.Sender}, nil
	}

	return nil, junomessages.MessageNotSupported(cosmosMsg)
}
