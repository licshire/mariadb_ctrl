// Code generated by counterfeiter. DO NOT EDIT.
package seederfakes

import (
	"sync"

	"github.com/cloudfoundry/mariadb_ctrl/mariadb_helper/seeder"
	_ "github.com/go-sql-driver/mysql"
)

type FakeSeeder struct {
	CreateDBIfNeededStub        func() error
	createDBIfNeededMutex       sync.RWMutex
	createDBIfNeededArgsForCall []struct{}
	createDBIfNeededReturns     struct {
		result1 error
	}
	createDBIfNeededReturnsOnCall map[int]struct {
		result1 error
	}
	IsExistingUserStub        func() (bool, error)
	isExistingUserMutex       sync.RWMutex
	isExistingUserArgsForCall []struct{}
	isExistingUserReturns     struct {
		result1 bool
		result2 error
	}
	isExistingUserReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	CreateUserStub        func() error
	createUserMutex       sync.RWMutex
	createUserArgsForCall []struct{}
	createUserReturns     struct {
		result1 error
	}
	createUserReturnsOnCall map[int]struct {
		result1 error
	}
	UpdateUserStub        func() error
	updateUserMutex       sync.RWMutex
	updateUserArgsForCall []struct{}
	updateUserReturns     struct {
		result1 error
	}
	updateUserReturnsOnCall map[int]struct {
		result1 error
	}
	GrantUserPrivilegesStub        func() error
	grantUserPrivilegesMutex       sync.RWMutex
	grantUserPrivilegesArgsForCall []struct{}
	grantUserPrivilegesReturns     struct {
		result1 error
	}
	grantUserPrivilegesReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSeeder) CreateDBIfNeeded() error {
	fake.createDBIfNeededMutex.Lock()
	ret, specificReturn := fake.createDBIfNeededReturnsOnCall[len(fake.createDBIfNeededArgsForCall)]
	fake.createDBIfNeededArgsForCall = append(fake.createDBIfNeededArgsForCall, struct{}{})
	fake.recordInvocation("CreateDBIfNeeded", []interface{}{})
	fake.createDBIfNeededMutex.Unlock()
	if fake.CreateDBIfNeededStub != nil {
		return fake.CreateDBIfNeededStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createDBIfNeededReturns.result1
}

func (fake *FakeSeeder) CreateDBIfNeededCallCount() int {
	fake.createDBIfNeededMutex.RLock()
	defer fake.createDBIfNeededMutex.RUnlock()
	return len(fake.createDBIfNeededArgsForCall)
}

func (fake *FakeSeeder) CreateDBIfNeededReturns(result1 error) {
	fake.CreateDBIfNeededStub = nil
	fake.createDBIfNeededReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSeeder) CreateDBIfNeededReturnsOnCall(i int, result1 error) {
	fake.CreateDBIfNeededStub = nil
	if fake.createDBIfNeededReturnsOnCall == nil {
		fake.createDBIfNeededReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createDBIfNeededReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSeeder) IsExistingUser() (bool, error) {
	fake.isExistingUserMutex.Lock()
	ret, specificReturn := fake.isExistingUserReturnsOnCall[len(fake.isExistingUserArgsForCall)]
	fake.isExistingUserArgsForCall = append(fake.isExistingUserArgsForCall, struct{}{})
	fake.recordInvocation("IsExistingUser", []interface{}{})
	fake.isExistingUserMutex.Unlock()
	if fake.IsExistingUserStub != nil {
		return fake.IsExistingUserStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.isExistingUserReturns.result1, fake.isExistingUserReturns.result2
}

func (fake *FakeSeeder) IsExistingUserCallCount() int {
	fake.isExistingUserMutex.RLock()
	defer fake.isExistingUserMutex.RUnlock()
	return len(fake.isExistingUserArgsForCall)
}

func (fake *FakeSeeder) IsExistingUserReturns(result1 bool, result2 error) {
	fake.IsExistingUserStub = nil
	fake.isExistingUserReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeSeeder) IsExistingUserReturnsOnCall(i int, result1 bool, result2 error) {
	fake.IsExistingUserStub = nil
	if fake.isExistingUserReturnsOnCall == nil {
		fake.isExistingUserReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.isExistingUserReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeSeeder) CreateUser() error {
	fake.createUserMutex.Lock()
	ret, specificReturn := fake.createUserReturnsOnCall[len(fake.createUserArgsForCall)]
	fake.createUserArgsForCall = append(fake.createUserArgsForCall, struct{}{})
	fake.recordInvocation("CreateUser", []interface{}{})
	fake.createUserMutex.Unlock()
	if fake.CreateUserStub != nil {
		return fake.CreateUserStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createUserReturns.result1
}

func (fake *FakeSeeder) CreateUserCallCount() int {
	fake.createUserMutex.RLock()
	defer fake.createUserMutex.RUnlock()
	return len(fake.createUserArgsForCall)
}

func (fake *FakeSeeder) CreateUserReturns(result1 error) {
	fake.CreateUserStub = nil
	fake.createUserReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSeeder) CreateUserReturnsOnCall(i int, result1 error) {
	fake.CreateUserStub = nil
	if fake.createUserReturnsOnCall == nil {
		fake.createUserReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createUserReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSeeder) UpdateUser() error {
	fake.updateUserMutex.Lock()
	ret, specificReturn := fake.updateUserReturnsOnCall[len(fake.updateUserArgsForCall)]
	fake.updateUserArgsForCall = append(fake.updateUserArgsForCall, struct{}{})
	fake.recordInvocation("UpdateUser", []interface{}{})
	fake.updateUserMutex.Unlock()
	if fake.UpdateUserStub != nil {
		return fake.UpdateUserStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.updateUserReturns.result1
}

func (fake *FakeSeeder) UpdateUserCallCount() int {
	fake.updateUserMutex.RLock()
	defer fake.updateUserMutex.RUnlock()
	return len(fake.updateUserArgsForCall)
}

func (fake *FakeSeeder) UpdateUserReturns(result1 error) {
	fake.UpdateUserStub = nil
	fake.updateUserReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSeeder) UpdateUserReturnsOnCall(i int, result1 error) {
	fake.UpdateUserStub = nil
	if fake.updateUserReturnsOnCall == nil {
		fake.updateUserReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateUserReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSeeder) GrantUserPrivileges() error {
	fake.grantUserPrivilegesMutex.Lock()
	ret, specificReturn := fake.grantUserPrivilegesReturnsOnCall[len(fake.grantUserPrivilegesArgsForCall)]
	fake.grantUserPrivilegesArgsForCall = append(fake.grantUserPrivilegesArgsForCall, struct{}{})
	fake.recordInvocation("GrantUserPrivileges", []interface{}{})
	fake.grantUserPrivilegesMutex.Unlock()
	if fake.GrantUserPrivilegesStub != nil {
		return fake.GrantUserPrivilegesStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.grantUserPrivilegesReturns.result1
}

func (fake *FakeSeeder) GrantUserPrivilegesCallCount() int {
	fake.grantUserPrivilegesMutex.RLock()
	defer fake.grantUserPrivilegesMutex.RUnlock()
	return len(fake.grantUserPrivilegesArgsForCall)
}

func (fake *FakeSeeder) GrantUserPrivilegesReturns(result1 error) {
	fake.GrantUserPrivilegesStub = nil
	fake.grantUserPrivilegesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeSeeder) GrantUserPrivilegesReturnsOnCall(i int, result1 error) {
	fake.GrantUserPrivilegesStub = nil
	if fake.grantUserPrivilegesReturnsOnCall == nil {
		fake.grantUserPrivilegesReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.grantUserPrivilegesReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeSeeder) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createDBIfNeededMutex.RLock()
	defer fake.createDBIfNeededMutex.RUnlock()
	fake.isExistingUserMutex.RLock()
	defer fake.isExistingUserMutex.RUnlock()
	fake.createUserMutex.RLock()
	defer fake.createUserMutex.RUnlock()
	fake.updateUserMutex.RLock()
	defer fake.updateUserMutex.RUnlock()
	fake.grantUserPrivilegesMutex.RLock()
	defer fake.grantUserPrivilegesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSeeder) recordInvocation(key string, args []interface{}) {
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

var _ seeder.Seeder = new(FakeSeeder)
