// Code generated by counterfeiter. DO NOT EDIT.
package workerfakes

import (
	"io"
	"sync"

	"github.com/concourse/atc/db"
	"github.com/concourse/atc/worker"
)

type FakeImageFetchingDelegate struct {
	StdoutStub        func() io.Writer
	stdoutMutex       sync.RWMutex
	stdoutArgsForCall []struct{}
	stdoutReturns     struct {
		result1 io.Writer
	}
	stdoutReturnsOnCall map[int]struct {
		result1 io.Writer
	}
	StderrStub        func() io.Writer
	stderrMutex       sync.RWMutex
	stderrArgsForCall []struct{}
	stderrReturns     struct {
		result1 io.Writer
	}
	stderrReturnsOnCall map[int]struct {
		result1 io.Writer
	}
	ImageVersionDeterminedStub        func(db.UsedResourceCache) error
	imageVersionDeterminedMutex       sync.RWMutex
	imageVersionDeterminedArgsForCall []struct {
		arg1 db.UsedResourceCache
	}
	imageVersionDeterminedReturns struct {
		result1 error
	}
	imageVersionDeterminedReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImageFetchingDelegate) Stdout() io.Writer {
	fake.stdoutMutex.Lock()
	ret, specificReturn := fake.stdoutReturnsOnCall[len(fake.stdoutArgsForCall)]
	fake.stdoutArgsForCall = append(fake.stdoutArgsForCall, struct{}{})
	fake.recordInvocation("Stdout", []interface{}{})
	fake.stdoutMutex.Unlock()
	if fake.StdoutStub != nil {
		return fake.StdoutStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.stdoutReturns.result1
}

func (fake *FakeImageFetchingDelegate) StdoutCallCount() int {
	fake.stdoutMutex.RLock()
	defer fake.stdoutMutex.RUnlock()
	return len(fake.stdoutArgsForCall)
}

func (fake *FakeImageFetchingDelegate) StdoutReturns(result1 io.Writer) {
	fake.StdoutStub = nil
	fake.stdoutReturns = struct {
		result1 io.Writer
	}{result1}
}

func (fake *FakeImageFetchingDelegate) StdoutReturnsOnCall(i int, result1 io.Writer) {
	fake.StdoutStub = nil
	if fake.stdoutReturnsOnCall == nil {
		fake.stdoutReturnsOnCall = make(map[int]struct {
			result1 io.Writer
		})
	}
	fake.stdoutReturnsOnCall[i] = struct {
		result1 io.Writer
	}{result1}
}

func (fake *FakeImageFetchingDelegate) Stderr() io.Writer {
	fake.stderrMutex.Lock()
	ret, specificReturn := fake.stderrReturnsOnCall[len(fake.stderrArgsForCall)]
	fake.stderrArgsForCall = append(fake.stderrArgsForCall, struct{}{})
	fake.recordInvocation("Stderr", []interface{}{})
	fake.stderrMutex.Unlock()
	if fake.StderrStub != nil {
		return fake.StderrStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.stderrReturns.result1
}

func (fake *FakeImageFetchingDelegate) StderrCallCount() int {
	fake.stderrMutex.RLock()
	defer fake.stderrMutex.RUnlock()
	return len(fake.stderrArgsForCall)
}

func (fake *FakeImageFetchingDelegate) StderrReturns(result1 io.Writer) {
	fake.StderrStub = nil
	fake.stderrReturns = struct {
		result1 io.Writer
	}{result1}
}

func (fake *FakeImageFetchingDelegate) StderrReturnsOnCall(i int, result1 io.Writer) {
	fake.StderrStub = nil
	if fake.stderrReturnsOnCall == nil {
		fake.stderrReturnsOnCall = make(map[int]struct {
			result1 io.Writer
		})
	}
	fake.stderrReturnsOnCall[i] = struct {
		result1 io.Writer
	}{result1}
}

func (fake *FakeImageFetchingDelegate) ImageVersionDetermined(arg1 db.UsedResourceCache) error {
	fake.imageVersionDeterminedMutex.Lock()
	ret, specificReturn := fake.imageVersionDeterminedReturnsOnCall[len(fake.imageVersionDeterminedArgsForCall)]
	fake.imageVersionDeterminedArgsForCall = append(fake.imageVersionDeterminedArgsForCall, struct {
		arg1 db.UsedResourceCache
	}{arg1})
	fake.recordInvocation("ImageVersionDetermined", []interface{}{arg1})
	fake.imageVersionDeterminedMutex.Unlock()
	if fake.ImageVersionDeterminedStub != nil {
		return fake.ImageVersionDeterminedStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.imageVersionDeterminedReturns.result1
}

func (fake *FakeImageFetchingDelegate) ImageVersionDeterminedCallCount() int {
	fake.imageVersionDeterminedMutex.RLock()
	defer fake.imageVersionDeterminedMutex.RUnlock()
	return len(fake.imageVersionDeterminedArgsForCall)
}

func (fake *FakeImageFetchingDelegate) ImageVersionDeterminedArgsForCall(i int) db.UsedResourceCache {
	fake.imageVersionDeterminedMutex.RLock()
	defer fake.imageVersionDeterminedMutex.RUnlock()
	return fake.imageVersionDeterminedArgsForCall[i].arg1
}

func (fake *FakeImageFetchingDelegate) ImageVersionDeterminedReturns(result1 error) {
	fake.ImageVersionDeterminedStub = nil
	fake.imageVersionDeterminedReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageFetchingDelegate) ImageVersionDeterminedReturnsOnCall(i int, result1 error) {
	fake.ImageVersionDeterminedStub = nil
	if fake.imageVersionDeterminedReturnsOnCall == nil {
		fake.imageVersionDeterminedReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.imageVersionDeterminedReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageFetchingDelegate) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.stdoutMutex.RLock()
	defer fake.stdoutMutex.RUnlock()
	fake.stderrMutex.RLock()
	defer fake.stderrMutex.RUnlock()
	fake.imageVersionDeterminedMutex.RLock()
	defer fake.imageVersionDeterminedMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeImageFetchingDelegate) recordInvocation(key string, args []interface{}) {
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

var _ worker.ImageFetchingDelegate = new(FakeImageFetchingDelegate)
