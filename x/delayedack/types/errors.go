package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrCanOnlyUpdatePendingPacket = sdkerrors.Register(ModuleName, 1, "can only update pending packet")
	ErrRollappPacketDoesNotExist  = sdkerrors.Register(ModuleName, 2, "rollapp packet does not exist")
	ErrInvalidEIBCFee             = sdkerrors.Register(ModuleName, 3, "invalid eibc fee")
)
