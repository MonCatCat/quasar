// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/MonCatCat/quasar/x/intergamm/types (interfaces: IBCTransferKeeper)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	types "github.com/cosmos/cosmos-sdk/types"
	types0 "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	types1 "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	gomock "github.com/golang/mock/gomock"
	bytes "github.com/tendermint/tendermint/libs/bytes"
)

// MockIBCTransferKeeper is a mock of IBCTransferKeeper interface.
type MockIBCTransferKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockIBCTransferKeeperMockRecorder
}

// MockIBCTransferKeeperMockRecorder is the mock recorder for MockIBCTransferKeeper.
type MockIBCTransferKeeperMockRecorder struct {
	mock *MockIBCTransferKeeper
}

// NewMockIBCTransferKeeper creates a new mock instance.
func NewMockIBCTransferKeeper(ctrl *gomock.Controller) *MockIBCTransferKeeper {
	mock := &MockIBCTransferKeeper{ctrl: ctrl}
	mock.recorder = &MockIBCTransferKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIBCTransferKeeper) EXPECT() *MockIBCTransferKeeperMockRecorder {
	return m.recorder
}

// GetDenomTrace mocks base method.
func (m *MockIBCTransferKeeper) GetDenomTrace(arg0 types.Context, arg1 bytes.HexBytes) (types0.DenomTrace, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDenomTrace", arg0, arg1)
	ret0, _ := ret[0].(types0.DenomTrace)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetDenomTrace indicates an expected call of GetDenomTrace.
func (mr *MockIBCTransferKeeperMockRecorder) GetDenomTrace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDenomTrace", reflect.TypeOf((*MockIBCTransferKeeper)(nil).GetDenomTrace), arg0, arg1)
}

// SendTransfer mocks base method.
func (m *MockIBCTransferKeeper) SendTransfer(arg0 types.Context, arg1, arg2 string, arg3 types.Coin, arg4 types.AccAddress, arg5 string, arg6 types1.Height, arg7 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendTransfer", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendTransfer indicates an expected call of SendTransfer.
func (mr *MockIBCTransferKeeperMockRecorder) SendTransfer(arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTransfer", reflect.TypeOf((*MockIBCTransferKeeper)(nil).SendTransfer), arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7)
}
