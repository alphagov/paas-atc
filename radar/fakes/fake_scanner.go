// This file was generated by counterfeiter
package fakes

import (
	"sync"
	"time"

	"github.com/concourse/atc"
	"github.com/concourse/atc/radar"
	"github.com/pivotal-golang/lager"
)

type FakeScanner struct {
	RunStub        func(lager.Logger, string) (time.Duration, error)
	runMutex       sync.RWMutex
	runArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	runReturns struct {
		result1 time.Duration
		result2 error
	}
	ScanStub        func(lager.Logger, string) error
	scanMutex       sync.RWMutex
	scanArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	scanReturns struct {
		result1 error
	}
	ScanFromVersionStub        func(lager.Logger, string, atc.Version) error
	scanFromVersionMutex       sync.RWMutex
	scanFromVersionArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 atc.Version
	}
	scanFromVersionReturns struct {
		result1 error
	}
}

func (fake *FakeScanner) Run(arg1 lager.Logger, arg2 string) (time.Duration, error) {
	fake.runMutex.Lock()
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		return fake.RunStub(arg1, arg2)
	} else {
		return fake.runReturns.result1, fake.runReturns.result2
	}
}

func (fake *FakeScanner) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeScanner) RunArgsForCall(i int) (lager.Logger, string) {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return fake.runArgsForCall[i].arg1, fake.runArgsForCall[i].arg2
}

func (fake *FakeScanner) RunReturns(result1 time.Duration, result2 error) {
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 time.Duration
		result2 error
	}{result1, result2}
}

func (fake *FakeScanner) Scan(arg1 lager.Logger, arg2 string) error {
	fake.scanMutex.Lock()
	fake.scanArgsForCall = append(fake.scanArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.scanMutex.Unlock()
	if fake.ScanStub != nil {
		return fake.ScanStub(arg1, arg2)
	} else {
		return fake.scanReturns.result1
	}
}

func (fake *FakeScanner) ScanCallCount() int {
	fake.scanMutex.RLock()
	defer fake.scanMutex.RUnlock()
	return len(fake.scanArgsForCall)
}

func (fake *FakeScanner) ScanArgsForCall(i int) (lager.Logger, string) {
	fake.scanMutex.RLock()
	defer fake.scanMutex.RUnlock()
	return fake.scanArgsForCall[i].arg1, fake.scanArgsForCall[i].arg2
}

func (fake *FakeScanner) ScanReturns(result1 error) {
	fake.ScanStub = nil
	fake.scanReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeScanner) ScanFromVersion(arg1 lager.Logger, arg2 string, arg3 atc.Version) error {
	fake.scanFromVersionMutex.Lock()
	fake.scanFromVersionArgsForCall = append(fake.scanFromVersionArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 atc.Version
	}{arg1, arg2, arg3})
	fake.scanFromVersionMutex.Unlock()
	if fake.ScanFromVersionStub != nil {
		return fake.ScanFromVersionStub(arg1, arg2, arg3)
	} else {
		return fake.scanFromVersionReturns.result1
	}
}

func (fake *FakeScanner) ScanFromVersionCallCount() int {
	fake.scanFromVersionMutex.RLock()
	defer fake.scanFromVersionMutex.RUnlock()
	return len(fake.scanFromVersionArgsForCall)
}

func (fake *FakeScanner) ScanFromVersionArgsForCall(i int) (lager.Logger, string, atc.Version) {
	fake.scanFromVersionMutex.RLock()
	defer fake.scanFromVersionMutex.RUnlock()
	return fake.scanFromVersionArgsForCall[i].arg1, fake.scanFromVersionArgsForCall[i].arg2, fake.scanFromVersionArgsForCall[i].arg3
}

func (fake *FakeScanner) ScanFromVersionReturns(result1 error) {
	fake.ScanFromVersionStub = nil
	fake.scanFromVersionReturns = struct {
		result1 error
	}{result1}
}

var _ radar.Scanner = new(FakeScanner)
