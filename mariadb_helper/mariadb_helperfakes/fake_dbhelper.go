// This file was generated by counterfeiter
package mariadb_helperfakes

import (
	"os/exec"
	"sync"

	"github.com/cloudfoundry/mariadb_ctrl/mariadb_helper"
)

type FakeDBHelper struct {
	StartMysqldInModeStub        func(command string) error
	startMysqldInModeMutex       sync.RWMutex
	startMysqldInModeArgsForCall []struct {
		command string
	}
	startMysqldInModeReturns struct {
		result1 error
	}
	startMysqldInModeReturnsOnCall map[int]struct {
		result1 error
	}
	StartMysqldInJoinStub        func() (*exec.Cmd, error)
	startMysqldInJoinMutex       sync.RWMutex
	startMysqldInJoinArgsForCall []struct{}
	startMysqldInJoinReturns     struct {
		result1 *exec.Cmd
		result2 error
	}
	startMysqldInJoinReturnsOnCall map[int]struct {
		result1 *exec.Cmd
		result2 error
	}
	StartMysqldInBootstrapStub        func() (*exec.Cmd, error)
	startMysqldInBootstrapMutex       sync.RWMutex
	startMysqldInBootstrapArgsForCall []struct{}
	startMysqldInBootstrapReturns     struct {
		result1 *exec.Cmd
		result2 error
	}
	startMysqldInBootstrapReturnsOnCall map[int]struct {
		result1 *exec.Cmd
		result2 error
	}
	StopMysqldStub        func() error
	stopMysqldMutex       sync.RWMutex
	stopMysqldArgsForCall []struct{}
	stopMysqldReturns     struct {
		result1 error
	}
	stopMysqldReturnsOnCall map[int]struct {
		result1 error
	}
	StopStandaloneMysqldStub        func() error
	stopStandaloneMysqldMutex       sync.RWMutex
	stopStandaloneMysqldArgsForCall []struct{}
	stopStandaloneMysqldReturns     struct {
		result1 error
	}
	stopStandaloneMysqldReturnsOnCall map[int]struct {
		result1 error
	}
	UpgradeStub        func() (output string, err error)
	upgradeMutex       sync.RWMutex
	upgradeArgsForCall []struct{}
	upgradeReturns     struct {
		result1 string
		result2 error
	}
	upgradeReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	IsDatabaseReachableStub        func() bool
	isDatabaseReachableMutex       sync.RWMutex
	isDatabaseReachableArgsForCall []struct{}
	isDatabaseReachableReturns     struct {
		result1 bool
	}
	isDatabaseReachableReturnsOnCall map[int]struct {
		result1 bool
	}
	IsProcessRunningStub        func() bool
	isProcessRunningMutex       sync.RWMutex
	isProcessRunningArgsForCall []struct{}
	isProcessRunningReturns     struct {
		result1 bool
	}
	isProcessRunningReturnsOnCall map[int]struct {
		result1 bool
	}
	SeedStub        func() error
	seedMutex       sync.RWMutex
	seedArgsForCall []struct{}
	seedReturns     struct {
		result1 error
	}
	seedReturnsOnCall map[int]struct {
		result1 error
	}
	ManageReadOnlyUserStub        func() error
	manageReadOnlyUserMutex       sync.RWMutex
	manageReadOnlyUserArgsForCall []struct{}
	manageReadOnlyUserReturns     struct {
		result1 error
	}
	manageReadOnlyUserReturnsOnCall map[int]struct {
		result1 error
	}
	RunPostStartSQLStub        func() error
	runPostStartSQLMutex       sync.RWMutex
	runPostStartSQLArgsForCall []struct{}
	runPostStartSQLReturns     struct {
		result1 error
	}
	runPostStartSQLReturnsOnCall map[int]struct {
		result1 error
	}
	TestDatabaseCleanupStub        func() error
	testDatabaseCleanupMutex       sync.RWMutex
	testDatabaseCleanupArgsForCall []struct{}
	testDatabaseCleanupReturns     struct {
		result1 error
	}
	testDatabaseCleanupReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDBHelper) StartMysqldInMode(command string) error {
	fake.startMysqldInModeMutex.Lock()
	ret, specificReturn := fake.startMysqldInModeReturnsOnCall[len(fake.startMysqldInModeArgsForCall)]
	fake.startMysqldInModeArgsForCall = append(fake.startMysqldInModeArgsForCall, struct {
		command string
	}{command})
	fake.recordInvocation("StartMysqldInMode", []interface{}{command})
	fake.startMysqldInModeMutex.Unlock()
	if fake.StartMysqldInModeStub != nil {
		return fake.StartMysqldInModeStub(command)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.startMysqldInModeReturns.result1
}

func (fake *FakeDBHelper) StartMysqldInModeCallCount() int {
	fake.startMysqldInModeMutex.RLock()
	defer fake.startMysqldInModeMutex.RUnlock()
	return len(fake.startMysqldInModeArgsForCall)
}

func (fake *FakeDBHelper) StartMysqldInModeArgsForCall(i int) string {
	fake.startMysqldInModeMutex.RLock()
	defer fake.startMysqldInModeMutex.RUnlock()
	return fake.startMysqldInModeArgsForCall[i].command
}

func (fake *FakeDBHelper) StartMysqldInModeReturns(result1 error) {
	fake.StartMysqldInModeStub = nil
	fake.startMysqldInModeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDBHelper) StartMysqldInModeReturnsOnCall(i int, result1 error) {
	fake.StartMysqldInModeStub = nil
	if fake.startMysqldInModeReturnsOnCall == nil {
		fake.startMysqldInModeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.startMysqldInModeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDBHelper) StartMysqldInJoin() (*exec.Cmd, error) {
	fake.startMysqldInJoinMutex.Lock()
	ret, specificReturn := fake.startMysqldInJoinReturnsOnCall[len(fake.startMysqldInJoinArgsForCall)]
	fake.startMysqldInJoinArgsForCall = append(fake.startMysqldInJoinArgsForCall, struct{}{})
	fake.recordInvocation("StartMysqldInJoin", []interface{}{})
	fake.startMysqldInJoinMutex.Unlock()
	if fake.StartMysqldInJoinStub != nil {
		return fake.StartMysqldInJoinStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.startMysqldInJoinReturns.result1, fake.startMysqldInJoinReturns.result2
}

func (fake *FakeDBHelper) StartMysqldInJoinCallCount() int {
	fake.startMysqldInJoinMutex.RLock()
	defer fake.startMysqldInJoinMutex.RUnlock()
	return len(fake.startMysqldInJoinArgsForCall)
}

func (fake *FakeDBHelper) StartMysqldInJoinReturns(result1 *exec.Cmd, result2 error) {
	fake.StartMysqldInJoinStub = nil
	fake.startMysqldInJoinReturns = struct {
		result1 *exec.Cmd
		result2 error
	}{result1, result2}
}

func (fake *FakeDBHelper) StartMysqldInJoinReturnsOnCall(i int, result1 *exec.Cmd, result2 error) {
	fake.StartMysqldInJoinStub = nil
	if fake.startMysqldInJoinReturnsOnCall == nil {
		fake.startMysqldInJoinReturnsOnCall = make(map[int]struct {
			result1 *exec.Cmd
			result2 error
		})
	}
	fake.startMysqldInJoinReturnsOnCall[i] = struct {
		result1 *exec.Cmd
		result2 error
	}{result1, result2}
}

func (fake *FakeDBHelper) StartMysqldInBootstrap() (*exec.Cmd, error) {
	fake.startMysqldInBootstrapMutex.Lock()
	ret, specificReturn := fake.startMysqldInBootstrapReturnsOnCall[len(fake.startMysqldInBootstrapArgsForCall)]
	fake.startMysqldInBootstrapArgsForCall = append(fake.startMysqldInBootstrapArgsForCall, struct{}{})
	fake.recordInvocation("StartMysqldInBootstrap", []interface{}{})
	fake.startMysqldInBootstrapMutex.Unlock()
	if fake.StartMysqldInBootstrapStub != nil {
		return fake.StartMysqldInBootstrapStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.startMysqldInBootstrapReturns.result1, fake.startMysqldInBootstrapReturns.result2
}

func (fake *FakeDBHelper) StartMysqldInBootstrapCallCount() int {
	fake.startMysqldInBootstrapMutex.RLock()
	defer fake.startMysqldInBootstrapMutex.RUnlock()
	return len(fake.startMysqldInBootstrapArgsForCall)
}

func (fake *FakeDBHelper) StartMysqldInBootstrapReturns(result1 *exec.Cmd, result2 error) {
	fake.StartMysqldInBootstrapStub = nil
	fake.startMysqldInBootstrapReturns = struct {
		result1 *exec.Cmd
		result2 error
	}{result1, result2}
}

func (fake *FakeDBHelper) StartMysqldInBootstrapReturnsOnCall(i int, result1 *exec.Cmd, result2 error) {
	fake.StartMysqldInBootstrapStub = nil
	if fake.startMysqldInBootstrapReturnsOnCall == nil {
		fake.startMysqldInBootstrapReturnsOnCall = make(map[int]struct {
			result1 *exec.Cmd
			result2 error
		})
	}
	fake.startMysqldInBootstrapReturnsOnCall[i] = struct {
		result1 *exec.Cmd
		result2 error
	}{result1, result2}
}

func (fake *FakeDBHelper) StopMysqld() error {
	fake.stopMysqldMutex.Lock()
	ret, specificReturn := fake.stopMysqldReturnsOnCall[len(fake.stopMysqldArgsForCall)]
	fake.stopMysqldArgsForCall = append(fake.stopMysqldArgsForCall, struct{}{})
	fake.recordInvocation("StopMysqld", []interface{}{})
	fake.stopMysqldMutex.Unlock()
	if fake.StopMysqldStub != nil {
		return fake.StopMysqldStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.stopMysqldReturns.result1
}

func (fake *FakeDBHelper) StopMysqldCallCount() int {
	fake.stopMysqldMutex.RLock()
	defer fake.stopMysqldMutex.RUnlock()
	return len(fake.stopMysqldArgsForCall)
}

func (fake *FakeDBHelper) StopMysqldReturns(result1 error) {
	fake.StopMysqldStub = nil
	fake.stopMysqldReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDBHelper) StopMysqldReturnsOnCall(i int, result1 error) {
	fake.StopMysqldStub = nil
	if fake.stopMysqldReturnsOnCall == nil {
		fake.stopMysqldReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.stopMysqldReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDBHelper) StopStandaloneMysqld() error {
	fake.stopStandaloneMysqldMutex.Lock()
	ret, specificReturn := fake.stopStandaloneMysqldReturnsOnCall[len(fake.stopStandaloneMysqldArgsForCall)]
	fake.stopStandaloneMysqldArgsForCall = append(fake.stopStandaloneMysqldArgsForCall, struct{}{})
	fake.recordInvocation("StopStandaloneMysqld", []interface{}{})
	fake.stopStandaloneMysqldMutex.Unlock()
	if fake.StopStandaloneMysqldStub != nil {
		return fake.StopStandaloneMysqldStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.stopStandaloneMysqldReturns.result1
}

func (fake *FakeDBHelper) StopStandaloneMysqldCallCount() int {
	fake.stopStandaloneMysqldMutex.RLock()
	defer fake.stopStandaloneMysqldMutex.RUnlock()
	return len(fake.stopStandaloneMysqldArgsForCall)
}

func (fake *FakeDBHelper) StopStandaloneMysqldReturns(result1 error) {
	fake.StopStandaloneMysqldStub = nil
	fake.stopStandaloneMysqldReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDBHelper) StopStandaloneMysqldReturnsOnCall(i int, result1 error) {
	fake.StopStandaloneMysqldStub = nil
	if fake.stopStandaloneMysqldReturnsOnCall == nil {
		fake.stopStandaloneMysqldReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.stopStandaloneMysqldReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDBHelper) Upgrade() (output string, err error) {
	fake.upgradeMutex.Lock()
	ret, specificReturn := fake.upgradeReturnsOnCall[len(fake.upgradeArgsForCall)]
	fake.upgradeArgsForCall = append(fake.upgradeArgsForCall, struct{}{})
	fake.recordInvocation("Upgrade", []interface{}{})
	fake.upgradeMutex.Unlock()
	if fake.UpgradeStub != nil {
		return fake.UpgradeStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.upgradeReturns.result1, fake.upgradeReturns.result2
}

func (fake *FakeDBHelper) UpgradeCallCount() int {
	fake.upgradeMutex.RLock()
	defer fake.upgradeMutex.RUnlock()
	return len(fake.upgradeArgsForCall)
}

func (fake *FakeDBHelper) UpgradeReturns(result1 string, result2 error) {
	fake.UpgradeStub = nil
	fake.upgradeReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeDBHelper) UpgradeReturnsOnCall(i int, result1 string, result2 error) {
	fake.UpgradeStub = nil
	if fake.upgradeReturnsOnCall == nil {
		fake.upgradeReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.upgradeReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeDBHelper) IsDatabaseReachable() bool {
	fake.isDatabaseReachableMutex.Lock()
	ret, specificReturn := fake.isDatabaseReachableReturnsOnCall[len(fake.isDatabaseReachableArgsForCall)]
	fake.isDatabaseReachableArgsForCall = append(fake.isDatabaseReachableArgsForCall, struct{}{})
	fake.recordInvocation("IsDatabaseReachable", []interface{}{})
	fake.isDatabaseReachableMutex.Unlock()
	if fake.IsDatabaseReachableStub != nil {
		return fake.IsDatabaseReachableStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.isDatabaseReachableReturns.result1
}

func (fake *FakeDBHelper) IsDatabaseReachableCallCount() int {
	fake.isDatabaseReachableMutex.RLock()
	defer fake.isDatabaseReachableMutex.RUnlock()
	return len(fake.isDatabaseReachableArgsForCall)
}

func (fake *FakeDBHelper) IsDatabaseReachableReturns(result1 bool) {
	fake.IsDatabaseReachableStub = nil
	fake.isDatabaseReachableReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeDBHelper) IsDatabaseReachableReturnsOnCall(i int, result1 bool) {
	fake.IsDatabaseReachableStub = nil
	if fake.isDatabaseReachableReturnsOnCall == nil {
		fake.isDatabaseReachableReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isDatabaseReachableReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeDBHelper) IsProcessRunning() bool {
	fake.isProcessRunningMutex.Lock()
	ret, specificReturn := fake.isProcessRunningReturnsOnCall[len(fake.isProcessRunningArgsForCall)]
	fake.isProcessRunningArgsForCall = append(fake.isProcessRunningArgsForCall, struct{}{})
	fake.recordInvocation("IsProcessRunning", []interface{}{})
	fake.isProcessRunningMutex.Unlock()
	if fake.IsProcessRunningStub != nil {
		return fake.IsProcessRunningStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.isProcessRunningReturns.result1
}

func (fake *FakeDBHelper) IsProcessRunningCallCount() int {
	fake.isProcessRunningMutex.RLock()
	defer fake.isProcessRunningMutex.RUnlock()
	return len(fake.isProcessRunningArgsForCall)
}

func (fake *FakeDBHelper) IsProcessRunningReturns(result1 bool) {
	fake.IsProcessRunningStub = nil
	fake.isProcessRunningReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeDBHelper) IsProcessRunningReturnsOnCall(i int, result1 bool) {
	fake.IsProcessRunningStub = nil
	if fake.isProcessRunningReturnsOnCall == nil {
		fake.isProcessRunningReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isProcessRunningReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeDBHelper) Seed() error {
	fake.seedMutex.Lock()
	ret, specificReturn := fake.seedReturnsOnCall[len(fake.seedArgsForCall)]
	fake.seedArgsForCall = append(fake.seedArgsForCall, struct{}{})
	fake.recordInvocation("Seed", []interface{}{})
	fake.seedMutex.Unlock()
	if fake.SeedStub != nil {
		return fake.SeedStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.seedReturns.result1
}

func (fake *FakeDBHelper) SeedCallCount() int {
	fake.seedMutex.RLock()
	defer fake.seedMutex.RUnlock()
	return len(fake.seedArgsForCall)
}

func (fake *FakeDBHelper) SeedReturns(result1 error) {
	fake.SeedStub = nil
	fake.seedReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDBHelper) SeedReturnsOnCall(i int, result1 error) {
	fake.SeedStub = nil
	if fake.seedReturnsOnCall == nil {
		fake.seedReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.seedReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDBHelper) ManageReadOnlyUser() error {
	fake.manageReadOnlyUserMutex.Lock()
	ret, specificReturn := fake.manageReadOnlyUserReturnsOnCall[len(fake.manageReadOnlyUserArgsForCall)]
	fake.manageReadOnlyUserArgsForCall = append(fake.manageReadOnlyUserArgsForCall, struct{}{})
	fake.recordInvocation("ManageReadOnlyUser", []interface{}{})
	fake.manageReadOnlyUserMutex.Unlock()
	if fake.ManageReadOnlyUserStub != nil {
		return fake.ManageReadOnlyUserStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.manageReadOnlyUserReturns.result1
}

func (fake *FakeDBHelper) ManageReadOnlyUserCallCount() int {
	fake.manageReadOnlyUserMutex.RLock()
	defer fake.manageReadOnlyUserMutex.RUnlock()
	return len(fake.manageReadOnlyUserArgsForCall)
}

func (fake *FakeDBHelper) ManageReadOnlyUserReturns(result1 error) {
	fake.ManageReadOnlyUserStub = nil
	fake.manageReadOnlyUserReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDBHelper) ManageReadOnlyUserReturnsOnCall(i int, result1 error) {
	fake.ManageReadOnlyUserStub = nil
	if fake.manageReadOnlyUserReturnsOnCall == nil {
		fake.manageReadOnlyUserReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.manageReadOnlyUserReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDBHelper) RunPostStartSQL() error {
	fake.runPostStartSQLMutex.Lock()
	ret, specificReturn := fake.runPostStartSQLReturnsOnCall[len(fake.runPostStartSQLArgsForCall)]
	fake.runPostStartSQLArgsForCall = append(fake.runPostStartSQLArgsForCall, struct{}{})
	fake.recordInvocation("RunPostStartSQL", []interface{}{})
	fake.runPostStartSQLMutex.Unlock()
	if fake.RunPostStartSQLStub != nil {
		return fake.RunPostStartSQLStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.runPostStartSQLReturns.result1
}

func (fake *FakeDBHelper) RunPostStartSQLCallCount() int {
	fake.runPostStartSQLMutex.RLock()
	defer fake.runPostStartSQLMutex.RUnlock()
	return len(fake.runPostStartSQLArgsForCall)
}

func (fake *FakeDBHelper) RunPostStartSQLReturns(result1 error) {
	fake.RunPostStartSQLStub = nil
	fake.runPostStartSQLReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDBHelper) RunPostStartSQLReturnsOnCall(i int, result1 error) {
	fake.RunPostStartSQLStub = nil
	if fake.runPostStartSQLReturnsOnCall == nil {
		fake.runPostStartSQLReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.runPostStartSQLReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDBHelper) TestDatabaseCleanup() error {
	fake.testDatabaseCleanupMutex.Lock()
	ret, specificReturn := fake.testDatabaseCleanupReturnsOnCall[len(fake.testDatabaseCleanupArgsForCall)]
	fake.testDatabaseCleanupArgsForCall = append(fake.testDatabaseCleanupArgsForCall, struct{}{})
	fake.recordInvocation("TestDatabaseCleanup", []interface{}{})
	fake.testDatabaseCleanupMutex.Unlock()
	if fake.TestDatabaseCleanupStub != nil {
		return fake.TestDatabaseCleanupStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.testDatabaseCleanupReturns.result1
}

func (fake *FakeDBHelper) TestDatabaseCleanupCallCount() int {
	fake.testDatabaseCleanupMutex.RLock()
	defer fake.testDatabaseCleanupMutex.RUnlock()
	return len(fake.testDatabaseCleanupArgsForCall)
}

func (fake *FakeDBHelper) TestDatabaseCleanupReturns(result1 error) {
	fake.TestDatabaseCleanupStub = nil
	fake.testDatabaseCleanupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDBHelper) TestDatabaseCleanupReturnsOnCall(i int, result1 error) {
	fake.TestDatabaseCleanupStub = nil
	if fake.testDatabaseCleanupReturnsOnCall == nil {
		fake.testDatabaseCleanupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.testDatabaseCleanupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDBHelper) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.startMysqldInModeMutex.RLock()
	defer fake.startMysqldInModeMutex.RUnlock()
	fake.startMysqldInJoinMutex.RLock()
	defer fake.startMysqldInJoinMutex.RUnlock()
	fake.startMysqldInBootstrapMutex.RLock()
	defer fake.startMysqldInBootstrapMutex.RUnlock()
	fake.stopMysqldMutex.RLock()
	defer fake.stopMysqldMutex.RUnlock()
	fake.stopStandaloneMysqldMutex.RLock()
	defer fake.stopStandaloneMysqldMutex.RUnlock()
	fake.upgradeMutex.RLock()
	defer fake.upgradeMutex.RUnlock()
	fake.isDatabaseReachableMutex.RLock()
	defer fake.isDatabaseReachableMutex.RUnlock()
	fake.isProcessRunningMutex.RLock()
	defer fake.isProcessRunningMutex.RUnlock()
	fake.seedMutex.RLock()
	defer fake.seedMutex.RUnlock()
	fake.manageReadOnlyUserMutex.RLock()
	defer fake.manageReadOnlyUserMutex.RUnlock()
	fake.runPostStartSQLMutex.RLock()
	defer fake.runPostStartSQLMutex.RUnlock()
	fake.testDatabaseCleanupMutex.RLock()
	defer fake.testDatabaseCleanupMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeDBHelper) recordInvocation(key string, args []interface{}) {
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

var _ mariadb_helper.DBHelper = new(FakeDBHelper)
