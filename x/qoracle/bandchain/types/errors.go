package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// IBC transfer sentinel errors
var (
	ErrDisabled = sdkerrors.Register(SubModuleName, 2, "bandchain oracle module is disabled")
)
