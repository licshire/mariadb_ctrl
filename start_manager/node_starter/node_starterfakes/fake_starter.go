// This file was generated by counterfeiter
package node_starterfakes

import (
	"os/exec"
	"sync"

	"github.com/cloudfoundry/mariadb_ctrl/start_manager/node_starter"
)

type FakeStarter struct {
	StartNodeFromStateStub        func(string) (string, error)
	startNodeFromStateMutex       sync.RWMutex
	startNodeFromStateArgsForCall []struct {
		arg1 string
	}
	startNodeFromStateReturns struct {
		result1 string
		result2 error
	}
	startNodeFromStateReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetMysqlCmdStub        func() (*exec.Cmd, error)
	getMysqlCmdMutex       sync.RWMutex
	getMysqlCmdArgsForCall []struct{}
	getMysqlCmdReturns     struct {
		result1 *exec.Cmd
		result2 error
	}
	getMysqlCmdReturnsOnCall map[int]struct {
		result1 *exec.Cmd
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStarter) StartNodeFromState(arg1 string) (string, error) {
	fake.startNodeFromStateMutex.Lock()
	ret, specificReturn := fake.startNodeFromStateReturnsOnCall[len(fake.startNodeFromStateArgsForCall)]
	fake.startNodeFromStateArgsForCall = append(fake.startNodeFromStateArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("StartNodeFromState", []interface{}{arg1})
	fake.startNodeFromStateMutex.Unlock()
	if fake.StartNodeFromStateStub != nil {
		return fake.StartNodeFromStateStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.startNodeFromStateReturns.result1, fake.startNodeFromStateReturns.result2
}

func (fake *FakeStarter) StartNodeFromStateCallCount() int {
	fake.startNodeFromStateMutex.RLock()
	defer fake.startNodeFromStateMutex.RUnlock()
	return len(fake.startNodeFromStateArgsForCall)
}

func (fake *FakeStarter) StartNodeFromStateArgsForCall(i int) string {
	fake.startNodeFromStateMutex.RLock()
	defer fake.startNodeFromStateMutex.RUnlock()
	return fake.startNodeFromStateArgsForCall[i].arg1
}

func (fake *FakeStarter) StartNodeFromStateReturns(result1 string, result2 error) {
	fake.StartNodeFromStateStub = nil
	fake.startNodeFromStateReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeStarter) StartNodeFromStateReturnsOnCall(i int, result1 string, result2 error) {
	fake.StartNodeFromStateStub = nil
	if fake.startNodeFromStateReturnsOnCall == nil {
		fake.startNodeFromStateReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.startNodeFromStateReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeStarter) GetMysqlCmd() (*exec.Cmd, error) {
	fake.getMysqlCmdMutex.Lock()
	ret, specificReturn := fake.getMysqlCmdReturnsOnCall[len(fake.getMysqlCmdArgsForCall)]
	fake.getMysqlCmdArgsForCall = append(fake.getMysqlCmdArgsForCall, struct{}{})
	fake.recordInvocation("GetMysqlCmd", []interface{}{})
	fake.getMysqlCmdMutex.Unlock()
	if fake.GetMysqlCmdStub != nil {
		return fake.GetMysqlCmdStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getMysqlCmdReturns.result1, fake.getMysqlCmdReturns.result2
}

func (fake *FakeStarter) GetMysqlCmdCallCount() int {
	fake.getMysqlCmdMutex.RLock()
	defer fake.getMysqlCmdMutex.RUnlock()
	return len(fake.getMysqlCmdArgsForCall)
}

func (fake *FakeStarter) GetMysqlCmdReturns(result1 *exec.Cmd, result2 error) {
	fake.GetMysqlCmdStub = nil
	fake.getMysqlCmdReturns = struct {
		result1 *exec.Cmd
		result2 error
	}{result1, result2}
}

func (fake *FakeStarter) GetMysqlCmdReturnsOnCall(i int, result1 *exec.Cmd, result2 error) {
	fake.GetMysqlCmdStub = nil
	if fake.getMysqlCmdReturnsOnCall == nil {
		fake.getMysqlCmdReturnsOnCall = make(map[int]struct {
			result1 *exec.Cmd
			result2 error
		})
	}
	fake.getMysqlCmdReturnsOnCall[i] = struct {
		result1 *exec.Cmd
		result2 error
	}{result1, result2}
}

func (fake *FakeStarter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.startNodeFromStateMutex.RLock()
	defer fake.startNodeFromStateMutex.RUnlock()
	fake.getMysqlCmdMutex.RLock()
	defer fake.getMysqlCmdMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeStarter) recordInvocation(key string, args []interface{}) {
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

var _ node_starter.Starter = new(FakeStarter)
