// Code generated by counterfeiter. DO NOT EDIT.
package aviatorfakes

import (
	"sync"

	"github.com/JulzDiverse/aviator"
)

type FakeModifier struct {
	ModifyStub        func([]byte, aviator.Modify) ([]byte, error)
	modifyMutex       sync.RWMutex
	modifyArgsForCall []struct {
		arg1 []byte
		arg2 aviator.Modify
	}
	modifyReturns struct {
		result1 []byte
		result2 error
	}
	modifyReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeModifier) Modify(arg1 []byte, arg2 aviator.Modify) ([]byte, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.modifyMutex.Lock()
	ret, specificReturn := fake.modifyReturnsOnCall[len(fake.modifyArgsForCall)]
	fake.modifyArgsForCall = append(fake.modifyArgsForCall, struct {
		arg1 []byte
		arg2 aviator.Modify
	}{arg1Copy, arg2})
	fake.recordInvocation("Modify", []interface{}{arg1Copy, arg2})
	fake.modifyMutex.Unlock()
	if fake.ModifyStub != nil {
		return fake.ModifyStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.modifyReturns.result1, fake.modifyReturns.result2
}

func (fake *FakeModifier) ModifyCallCount() int {
	fake.modifyMutex.RLock()
	defer fake.modifyMutex.RUnlock()
	return len(fake.modifyArgsForCall)
}

func (fake *FakeModifier) ModifyArgsForCall(i int) ([]byte, aviator.Modify) {
	fake.modifyMutex.RLock()
	defer fake.modifyMutex.RUnlock()
	return fake.modifyArgsForCall[i].arg1, fake.modifyArgsForCall[i].arg2
}

func (fake *FakeModifier) ModifyReturns(result1 []byte, result2 error) {
	fake.ModifyStub = nil
	fake.modifyReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeModifier) ModifyReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.ModifyStub = nil
	if fake.modifyReturnsOnCall == nil {
		fake.modifyReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.modifyReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeModifier) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.modifyMutex.RLock()
	defer fake.modifyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeModifier) recordInvocation(key string, args []interface{}) {
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

var _ aviator.Modifier = new(FakeModifier)