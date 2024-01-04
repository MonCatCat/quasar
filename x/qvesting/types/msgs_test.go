package types

import (
	"testing"

	"github.com/MonCatCat/quasar/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateVestingAccount_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateVestingAccount
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateVestingAccount{
				FromAddress: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateVestingAccount{
				FromAddress: sample.AccAddress().String(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
