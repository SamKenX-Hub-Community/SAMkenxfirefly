// Code generated by mockery v2.15.0. DO NOT EDIT.

package txcommonmocks

import (
	context "context"

	fftypes "github.com/hyperledger/firefly-common/pkg/fftypes"
	core "github.com/hyperledger/firefly/pkg/core"

	mock "github.com/stretchr/testify/mock"
)

// Helper is an autogenerated mock type for the Helper type
type Helper struct {
	mock.Mock
}

// AddBlockchainTX provides a mock function with given fields: ctx, tx, blockchainTXID
func (_m *Helper) AddBlockchainTX(ctx context.Context, tx *core.Transaction, blockchainTXID string) error {
	ret := _m.Called(ctx, tx, blockchainTXID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *core.Transaction, string) error); ok {
		r0 = rf(ctx, tx, blockchainTXID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBlockchainEventByIDCached provides a mock function with given fields: ctx, id
func (_m *Helper) GetBlockchainEventByIDCached(ctx context.Context, id *fftypes.UUID) (*core.BlockchainEvent, error) {
	ret := _m.Called(ctx, id)

	var r0 *core.BlockchainEvent
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID) *core.BlockchainEvent); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.BlockchainEvent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTransactionByIDCached provides a mock function with given fields: ctx, id
func (_m *Helper) GetTransactionByIDCached(ctx context.Context, id *fftypes.UUID) (*core.Transaction, error) {
	ret := _m.Called(ctx, id)

	var r0 *core.Transaction
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID) *core.Transaction); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.Transaction)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertOrGetBlockchainEvent provides a mock function with given fields: ctx, event
func (_m *Helper) InsertOrGetBlockchainEvent(ctx context.Context, event *core.BlockchainEvent) (*core.BlockchainEvent, error) {
	ret := _m.Called(ctx, event)

	var r0 *core.BlockchainEvent
	if rf, ok := ret.Get(0).(func(context.Context, *core.BlockchainEvent) *core.BlockchainEvent); ok {
		r0 = rf(ctx, event)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*core.BlockchainEvent)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *core.BlockchainEvent) error); ok {
		r1 = rf(ctx, event)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PersistTransaction provides a mock function with given fields: ctx, id, txType, blockchainTXID
func (_m *Helper) PersistTransaction(ctx context.Context, id *fftypes.UUID, txType fftypes.FFEnum, blockchainTXID string) (bool, error) {
	ret := _m.Called(ctx, id, txType, blockchainTXID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *fftypes.UUID, fftypes.FFEnum, string) bool); ok {
		r0 = rf(ctx, id, txType, blockchainTXID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *fftypes.UUID, fftypes.FFEnum, string) error); ok {
		r1 = rf(ctx, id, txType, blockchainTXID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SubmitNewTransaction provides a mock function with given fields: ctx, txType, idempotencyKey
func (_m *Helper) SubmitNewTransaction(ctx context.Context, txType fftypes.FFEnum, idempotencyKey core.IdempotencyKey) (*fftypes.UUID, error) {
	ret := _m.Called(ctx, txType, idempotencyKey)

	var r0 *fftypes.UUID
	if rf, ok := ret.Get(0).(func(context.Context, fftypes.FFEnum, core.IdempotencyKey) *fftypes.UUID); ok {
		r0 = rf(ctx, txType, idempotencyKey)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*fftypes.UUID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, fftypes.FFEnum, core.IdempotencyKey) error); ok {
		r1 = rf(ctx, txType, idempotencyKey)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewHelper interface {
	mock.TestingT
	Cleanup(func())
}

// NewHelper creates a new instance of Helper. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewHelper(t mockConstructorTestingTNewHelper) *Helper {
	mock := &Helper{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
