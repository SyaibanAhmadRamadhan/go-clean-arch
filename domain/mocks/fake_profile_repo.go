// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"context"
	"database/sql"
	"sync"

	"github.com/SyaibanAhmadRamadhan/go-clean-arch/domain/model"
	"github.com/SyaibanAhmadRamadhan/go-clean-arch/domain/repository"
)

type FakeProfileRepo struct {
	CloseConnStub        func()
	closeConnMutex       sync.RWMutex
	closeConnArgsForCall []struct {
	}
	EndTxStub        func(error) error
	endTxMutex       sync.RWMutex
	endTxArgsForCall []struct {
		arg1 error
	}
	endTxReturns struct {
		result1 error
	}
	endTxReturnsOnCall map[int]struct {
		result1 error
	}
	GetConnStub        func() (*sql.Conn, error)
	getConnMutex       sync.RWMutex
	getConnArgsForCall []struct {
	}
	getConnReturns struct {
		result1 *sql.Conn
		result2 error
	}
	getConnReturnsOnCall map[int]struct {
		result1 *sql.Conn
		result2 error
	}
	GetProfileByIDStub        func(context.Context, string) (model.Profile, error)
	getProfileByIDMutex       sync.RWMutex
	getProfileByIDArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getProfileByIDReturns struct {
		result1 model.Profile
		result2 error
	}
	getProfileByIDReturnsOnCall map[int]struct {
		result1 model.Profile
		result2 error
	}
	GetProfileByUserIDStub        func(context.Context, string) (model.Profile, error)
	getProfileByUserIDMutex       sync.RWMutex
	getProfileByUserIDArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	getProfileByUserIDReturns struct {
		result1 model.Profile
		result2 error
	}
	getProfileByUserIDReturnsOnCall map[int]struct {
		result1 model.Profile
		result2 error
	}
	GetTxStub        func() (*sql.Tx, error)
	getTxMutex       sync.RWMutex
	getTxArgsForCall []struct {
	}
	getTxReturns struct {
		result1 *sql.Tx
		result2 error
	}
	getTxReturnsOnCall map[int]struct {
		result1 *sql.Tx
		result2 error
	}
	OpenConnStub        func(context.Context) error
	openConnMutex       sync.RWMutex
	openConnArgsForCall []struct {
		arg1 context.Context
	}
	openConnReturns struct {
		result1 error
	}
	openConnReturnsOnCall map[int]struct {
		result1 error
	}
	StartTxStub        func(context.Context, *sql.TxOptions) error
	startTxMutex       sync.RWMutex
	startTxArgsForCall []struct {
		arg1 context.Context
		arg2 *sql.TxOptions
	}
	startTxReturns struct {
		result1 error
	}
	startTxReturnsOnCall map[int]struct {
		result1 error
	}
	StoreProfileStub        func(context.Context, model.Profile) (model.Profile, error)
	storeProfileMutex       sync.RWMutex
	storeProfileArgsForCall []struct {
		arg1 context.Context
		arg2 model.Profile
	}
	storeProfileReturns struct {
		result1 model.Profile
		result2 error
	}
	storeProfileReturnsOnCall map[int]struct {
		result1 model.Profile
		result2 error
	}
	UpdateProfileStub        func(context.Context, model.Profile) (model.Profile, error)
	updateProfileMutex       sync.RWMutex
	updateProfileArgsForCall []struct {
		arg1 context.Context
		arg2 model.Profile
	}
	updateProfileReturns struct {
		result1 model.Profile
		result2 error
	}
	updateProfileReturnsOnCall map[int]struct {
		result1 model.Profile
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeProfileRepo) CloseConn() {
	fake.closeConnMutex.Lock()
	fake.closeConnArgsForCall = append(fake.closeConnArgsForCall, struct {
	}{})
	stub := fake.CloseConnStub
	fake.recordInvocation("CloseConn", []interface{}{})
	fake.closeConnMutex.Unlock()
	if stub != nil {
		fake.CloseConnStub()
	}
}

func (fake *FakeProfileRepo) CloseConnCallCount() int {
	fake.closeConnMutex.RLock()
	defer fake.closeConnMutex.RUnlock()
	return len(fake.closeConnArgsForCall)
}

func (fake *FakeProfileRepo) CloseConnCalls(stub func()) {
	fake.closeConnMutex.Lock()
	defer fake.closeConnMutex.Unlock()
	fake.CloseConnStub = stub
}

func (fake *FakeProfileRepo) EndTx(arg1 error) error {
	fake.endTxMutex.Lock()
	ret, specificReturn := fake.endTxReturnsOnCall[len(fake.endTxArgsForCall)]
	fake.endTxArgsForCall = append(fake.endTxArgsForCall, struct {
		arg1 error
	}{arg1})
	stub := fake.EndTxStub
	fakeReturns := fake.endTxReturns
	fake.recordInvocation("EndTx", []interface{}{arg1})
	fake.endTxMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeProfileRepo) EndTxCallCount() int {
	fake.endTxMutex.RLock()
	defer fake.endTxMutex.RUnlock()
	return len(fake.endTxArgsForCall)
}

func (fake *FakeProfileRepo) EndTxCalls(stub func(error) error) {
	fake.endTxMutex.Lock()
	defer fake.endTxMutex.Unlock()
	fake.EndTxStub = stub
}

func (fake *FakeProfileRepo) EndTxArgsForCall(i int) error {
	fake.endTxMutex.RLock()
	defer fake.endTxMutex.RUnlock()
	argsForCall := fake.endTxArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeProfileRepo) EndTxReturns(result1 error) {
	fake.endTxMutex.Lock()
	defer fake.endTxMutex.Unlock()
	fake.EndTxStub = nil
	fake.endTxReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeProfileRepo) EndTxReturnsOnCall(i int, result1 error) {
	fake.endTxMutex.Lock()
	defer fake.endTxMutex.Unlock()
	fake.EndTxStub = nil
	if fake.endTxReturnsOnCall == nil {
		fake.endTxReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.endTxReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeProfileRepo) GetConn() (*sql.Conn, error) {
	fake.getConnMutex.Lock()
	ret, specificReturn := fake.getConnReturnsOnCall[len(fake.getConnArgsForCall)]
	fake.getConnArgsForCall = append(fake.getConnArgsForCall, struct {
	}{})
	stub := fake.GetConnStub
	fakeReturns := fake.getConnReturns
	fake.recordInvocation("GetConn", []interface{}{})
	fake.getConnMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeProfileRepo) GetConnCallCount() int {
	fake.getConnMutex.RLock()
	defer fake.getConnMutex.RUnlock()
	return len(fake.getConnArgsForCall)
}

func (fake *FakeProfileRepo) GetConnCalls(stub func() (*sql.Conn, error)) {
	fake.getConnMutex.Lock()
	defer fake.getConnMutex.Unlock()
	fake.GetConnStub = stub
}

func (fake *FakeProfileRepo) GetConnReturns(result1 *sql.Conn, result2 error) {
	fake.getConnMutex.Lock()
	defer fake.getConnMutex.Unlock()
	fake.GetConnStub = nil
	fake.getConnReturns = struct {
		result1 *sql.Conn
		result2 error
	}{result1, result2}
}

func (fake *FakeProfileRepo) GetConnReturnsOnCall(i int, result1 *sql.Conn, result2 error) {
	fake.getConnMutex.Lock()
	defer fake.getConnMutex.Unlock()
	fake.GetConnStub = nil
	if fake.getConnReturnsOnCall == nil {
		fake.getConnReturnsOnCall = make(map[int]struct {
			result1 *sql.Conn
			result2 error
		})
	}
	fake.getConnReturnsOnCall[i] = struct {
		result1 *sql.Conn
		result2 error
	}{result1, result2}
}

func (fake *FakeProfileRepo) GetProfileByID(arg1 context.Context, arg2 string) (model.Profile, error) {
	fake.getProfileByIDMutex.Lock()
	ret, specificReturn := fake.getProfileByIDReturnsOnCall[len(fake.getProfileByIDArgsForCall)]
	fake.getProfileByIDArgsForCall = append(fake.getProfileByIDArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.GetProfileByIDStub
	fakeReturns := fake.getProfileByIDReturns
	fake.recordInvocation("GetProfileByID", []interface{}{arg1, arg2})
	fake.getProfileByIDMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeProfileRepo) GetProfileByIDCallCount() int {
	fake.getProfileByIDMutex.RLock()
	defer fake.getProfileByIDMutex.RUnlock()
	return len(fake.getProfileByIDArgsForCall)
}

func (fake *FakeProfileRepo) GetProfileByIDCalls(stub func(context.Context, string) (model.Profile, error)) {
	fake.getProfileByIDMutex.Lock()
	defer fake.getProfileByIDMutex.Unlock()
	fake.GetProfileByIDStub = stub
}

func (fake *FakeProfileRepo) GetProfileByIDArgsForCall(i int) (context.Context, string) {
	fake.getProfileByIDMutex.RLock()
	defer fake.getProfileByIDMutex.RUnlock()
	argsForCall := fake.getProfileByIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeProfileRepo) GetProfileByIDReturns(result1 model.Profile, result2 error) {
	fake.getProfileByIDMutex.Lock()
	defer fake.getProfileByIDMutex.Unlock()
	fake.GetProfileByIDStub = nil
	fake.getProfileByIDReturns = struct {
		result1 model.Profile
		result2 error
	}{result1, result2}
}

func (fake *FakeProfileRepo) GetProfileByIDReturnsOnCall(i int, result1 model.Profile, result2 error) {
	fake.getProfileByIDMutex.Lock()
	defer fake.getProfileByIDMutex.Unlock()
	fake.GetProfileByIDStub = nil
	if fake.getProfileByIDReturnsOnCall == nil {
		fake.getProfileByIDReturnsOnCall = make(map[int]struct {
			result1 model.Profile
			result2 error
		})
	}
	fake.getProfileByIDReturnsOnCall[i] = struct {
		result1 model.Profile
		result2 error
	}{result1, result2}
}

func (fake *FakeProfileRepo) GetProfileByUserID(arg1 context.Context, arg2 string) (model.Profile, error) {
	fake.getProfileByUserIDMutex.Lock()
	ret, specificReturn := fake.getProfileByUserIDReturnsOnCall[len(fake.getProfileByUserIDArgsForCall)]
	fake.getProfileByUserIDArgsForCall = append(fake.getProfileByUserIDArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.GetProfileByUserIDStub
	fakeReturns := fake.getProfileByUserIDReturns
	fake.recordInvocation("GetProfileByUserID", []interface{}{arg1, arg2})
	fake.getProfileByUserIDMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeProfileRepo) GetProfileByUserIDCallCount() int {
	fake.getProfileByUserIDMutex.RLock()
	defer fake.getProfileByUserIDMutex.RUnlock()
	return len(fake.getProfileByUserIDArgsForCall)
}

func (fake *FakeProfileRepo) GetProfileByUserIDCalls(stub func(context.Context, string) (model.Profile, error)) {
	fake.getProfileByUserIDMutex.Lock()
	defer fake.getProfileByUserIDMutex.Unlock()
	fake.GetProfileByUserIDStub = stub
}

func (fake *FakeProfileRepo) GetProfileByUserIDArgsForCall(i int) (context.Context, string) {
	fake.getProfileByUserIDMutex.RLock()
	defer fake.getProfileByUserIDMutex.RUnlock()
	argsForCall := fake.getProfileByUserIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeProfileRepo) GetProfileByUserIDReturns(result1 model.Profile, result2 error) {
	fake.getProfileByUserIDMutex.Lock()
	defer fake.getProfileByUserIDMutex.Unlock()
	fake.GetProfileByUserIDStub = nil
	fake.getProfileByUserIDReturns = struct {
		result1 model.Profile
		result2 error
	}{result1, result2}
}

func (fake *FakeProfileRepo) GetProfileByUserIDReturnsOnCall(i int, result1 model.Profile, result2 error) {
	fake.getProfileByUserIDMutex.Lock()
	defer fake.getProfileByUserIDMutex.Unlock()
	fake.GetProfileByUserIDStub = nil
	if fake.getProfileByUserIDReturnsOnCall == nil {
		fake.getProfileByUserIDReturnsOnCall = make(map[int]struct {
			result1 model.Profile
			result2 error
		})
	}
	fake.getProfileByUserIDReturnsOnCall[i] = struct {
		result1 model.Profile
		result2 error
	}{result1, result2}
}

func (fake *FakeProfileRepo) GetTx() (*sql.Tx, error) {
	fake.getTxMutex.Lock()
	ret, specificReturn := fake.getTxReturnsOnCall[len(fake.getTxArgsForCall)]
	fake.getTxArgsForCall = append(fake.getTxArgsForCall, struct {
	}{})
	stub := fake.GetTxStub
	fakeReturns := fake.getTxReturns
	fake.recordInvocation("GetTx", []interface{}{})
	fake.getTxMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeProfileRepo) GetTxCallCount() int {
	fake.getTxMutex.RLock()
	defer fake.getTxMutex.RUnlock()
	return len(fake.getTxArgsForCall)
}

func (fake *FakeProfileRepo) GetTxCalls(stub func() (*sql.Tx, error)) {
	fake.getTxMutex.Lock()
	defer fake.getTxMutex.Unlock()
	fake.GetTxStub = stub
}

func (fake *FakeProfileRepo) GetTxReturns(result1 *sql.Tx, result2 error) {
	fake.getTxMutex.Lock()
	defer fake.getTxMutex.Unlock()
	fake.GetTxStub = nil
	fake.getTxReturns = struct {
		result1 *sql.Tx
		result2 error
	}{result1, result2}
}

func (fake *FakeProfileRepo) GetTxReturnsOnCall(i int, result1 *sql.Tx, result2 error) {
	fake.getTxMutex.Lock()
	defer fake.getTxMutex.Unlock()
	fake.GetTxStub = nil
	if fake.getTxReturnsOnCall == nil {
		fake.getTxReturnsOnCall = make(map[int]struct {
			result1 *sql.Tx
			result2 error
		})
	}
	fake.getTxReturnsOnCall[i] = struct {
		result1 *sql.Tx
		result2 error
	}{result1, result2}
}

func (fake *FakeProfileRepo) OpenConn(arg1 context.Context) error {
	fake.openConnMutex.Lock()
	ret, specificReturn := fake.openConnReturnsOnCall[len(fake.openConnArgsForCall)]
	fake.openConnArgsForCall = append(fake.openConnArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.OpenConnStub
	fakeReturns := fake.openConnReturns
	fake.recordInvocation("OpenConn", []interface{}{arg1})
	fake.openConnMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeProfileRepo) OpenConnCallCount() int {
	fake.openConnMutex.RLock()
	defer fake.openConnMutex.RUnlock()
	return len(fake.openConnArgsForCall)
}

func (fake *FakeProfileRepo) OpenConnCalls(stub func(context.Context) error) {
	fake.openConnMutex.Lock()
	defer fake.openConnMutex.Unlock()
	fake.OpenConnStub = stub
}

func (fake *FakeProfileRepo) OpenConnArgsForCall(i int) context.Context {
	fake.openConnMutex.RLock()
	defer fake.openConnMutex.RUnlock()
	argsForCall := fake.openConnArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeProfileRepo) OpenConnReturns(result1 error) {
	fake.openConnMutex.Lock()
	defer fake.openConnMutex.Unlock()
	fake.OpenConnStub = nil
	fake.openConnReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeProfileRepo) OpenConnReturnsOnCall(i int, result1 error) {
	fake.openConnMutex.Lock()
	defer fake.openConnMutex.Unlock()
	fake.OpenConnStub = nil
	if fake.openConnReturnsOnCall == nil {
		fake.openConnReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.openConnReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeProfileRepo) StartTx(arg1 context.Context, arg2 *sql.TxOptions) error {
	fake.startTxMutex.Lock()
	ret, specificReturn := fake.startTxReturnsOnCall[len(fake.startTxArgsForCall)]
	fake.startTxArgsForCall = append(fake.startTxArgsForCall, struct {
		arg1 context.Context
		arg2 *sql.TxOptions
	}{arg1, arg2})
	stub := fake.StartTxStub
	fakeReturns := fake.startTxReturns
	fake.recordInvocation("StartTx", []interface{}{arg1, arg2})
	fake.startTxMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeProfileRepo) StartTxCallCount() int {
	fake.startTxMutex.RLock()
	defer fake.startTxMutex.RUnlock()
	return len(fake.startTxArgsForCall)
}

func (fake *FakeProfileRepo) StartTxCalls(stub func(context.Context, *sql.TxOptions) error) {
	fake.startTxMutex.Lock()
	defer fake.startTxMutex.Unlock()
	fake.StartTxStub = stub
}

func (fake *FakeProfileRepo) StartTxArgsForCall(i int) (context.Context, *sql.TxOptions) {
	fake.startTxMutex.RLock()
	defer fake.startTxMutex.RUnlock()
	argsForCall := fake.startTxArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeProfileRepo) StartTxReturns(result1 error) {
	fake.startTxMutex.Lock()
	defer fake.startTxMutex.Unlock()
	fake.StartTxStub = nil
	fake.startTxReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeProfileRepo) StartTxReturnsOnCall(i int, result1 error) {
	fake.startTxMutex.Lock()
	defer fake.startTxMutex.Unlock()
	fake.StartTxStub = nil
	if fake.startTxReturnsOnCall == nil {
		fake.startTxReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.startTxReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeProfileRepo) StoreProfile(arg1 context.Context, arg2 model.Profile) (model.Profile, error) {
	fake.storeProfileMutex.Lock()
	ret, specificReturn := fake.storeProfileReturnsOnCall[len(fake.storeProfileArgsForCall)]
	fake.storeProfileArgsForCall = append(fake.storeProfileArgsForCall, struct {
		arg1 context.Context
		arg2 model.Profile
	}{arg1, arg2})
	stub := fake.StoreProfileStub
	fakeReturns := fake.storeProfileReturns
	fake.recordInvocation("StoreProfile", []interface{}{arg1, arg2})
	fake.storeProfileMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeProfileRepo) StoreProfileCallCount() int {
	fake.storeProfileMutex.RLock()
	defer fake.storeProfileMutex.RUnlock()
	return len(fake.storeProfileArgsForCall)
}

func (fake *FakeProfileRepo) StoreProfileCalls(stub func(context.Context, model.Profile) (model.Profile, error)) {
	fake.storeProfileMutex.Lock()
	defer fake.storeProfileMutex.Unlock()
	fake.StoreProfileStub = stub
}

func (fake *FakeProfileRepo) StoreProfileArgsForCall(i int) (context.Context, model.Profile) {
	fake.storeProfileMutex.RLock()
	defer fake.storeProfileMutex.RUnlock()
	argsForCall := fake.storeProfileArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeProfileRepo) StoreProfileReturns(result1 model.Profile, result2 error) {
	fake.storeProfileMutex.Lock()
	defer fake.storeProfileMutex.Unlock()
	fake.StoreProfileStub = nil
	fake.storeProfileReturns = struct {
		result1 model.Profile
		result2 error
	}{result1, result2}
}

func (fake *FakeProfileRepo) StoreProfileReturnsOnCall(i int, result1 model.Profile, result2 error) {
	fake.storeProfileMutex.Lock()
	defer fake.storeProfileMutex.Unlock()
	fake.StoreProfileStub = nil
	if fake.storeProfileReturnsOnCall == nil {
		fake.storeProfileReturnsOnCall = make(map[int]struct {
			result1 model.Profile
			result2 error
		})
	}
	fake.storeProfileReturnsOnCall[i] = struct {
		result1 model.Profile
		result2 error
	}{result1, result2}
}

func (fake *FakeProfileRepo) UpdateProfile(arg1 context.Context, arg2 model.Profile) (model.Profile, error) {
	fake.updateProfileMutex.Lock()
	ret, specificReturn := fake.updateProfileReturnsOnCall[len(fake.updateProfileArgsForCall)]
	fake.updateProfileArgsForCall = append(fake.updateProfileArgsForCall, struct {
		arg1 context.Context
		arg2 model.Profile
	}{arg1, arg2})
	stub := fake.UpdateProfileStub
	fakeReturns := fake.updateProfileReturns
	fake.recordInvocation("UpdateProfile", []interface{}{arg1, arg2})
	fake.updateProfileMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeProfileRepo) UpdateProfileCallCount() int {
	fake.updateProfileMutex.RLock()
	defer fake.updateProfileMutex.RUnlock()
	return len(fake.updateProfileArgsForCall)
}

func (fake *FakeProfileRepo) UpdateProfileCalls(stub func(context.Context, model.Profile) (model.Profile, error)) {
	fake.updateProfileMutex.Lock()
	defer fake.updateProfileMutex.Unlock()
	fake.UpdateProfileStub = stub
}

func (fake *FakeProfileRepo) UpdateProfileArgsForCall(i int) (context.Context, model.Profile) {
	fake.updateProfileMutex.RLock()
	defer fake.updateProfileMutex.RUnlock()
	argsForCall := fake.updateProfileArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeProfileRepo) UpdateProfileReturns(result1 model.Profile, result2 error) {
	fake.updateProfileMutex.Lock()
	defer fake.updateProfileMutex.Unlock()
	fake.UpdateProfileStub = nil
	fake.updateProfileReturns = struct {
		result1 model.Profile
		result2 error
	}{result1, result2}
}

func (fake *FakeProfileRepo) UpdateProfileReturnsOnCall(i int, result1 model.Profile, result2 error) {
	fake.updateProfileMutex.Lock()
	defer fake.updateProfileMutex.Unlock()
	fake.UpdateProfileStub = nil
	if fake.updateProfileReturnsOnCall == nil {
		fake.updateProfileReturnsOnCall = make(map[int]struct {
			result1 model.Profile
			result2 error
		})
	}
	fake.updateProfileReturnsOnCall[i] = struct {
		result1 model.Profile
		result2 error
	}{result1, result2}
}

func (fake *FakeProfileRepo) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeConnMutex.RLock()
	defer fake.closeConnMutex.RUnlock()
	fake.endTxMutex.RLock()
	defer fake.endTxMutex.RUnlock()
	fake.getConnMutex.RLock()
	defer fake.getConnMutex.RUnlock()
	fake.getProfileByIDMutex.RLock()
	defer fake.getProfileByIDMutex.RUnlock()
	fake.getProfileByUserIDMutex.RLock()
	defer fake.getProfileByUserIDMutex.RUnlock()
	fake.getTxMutex.RLock()
	defer fake.getTxMutex.RUnlock()
	fake.openConnMutex.RLock()
	defer fake.openConnMutex.RUnlock()
	fake.startTxMutex.RLock()
	defer fake.startTxMutex.RUnlock()
	fake.storeProfileMutex.RLock()
	defer fake.storeProfileMutex.RUnlock()
	fake.updateProfileMutex.RLock()
	defer fake.updateProfileMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeProfileRepo) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ repository.ProfileRepo = new(FakeProfileRepo)
