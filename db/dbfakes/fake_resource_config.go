// Code generated by counterfeiter. DO NOT EDIT.
package dbfakes

import (
	"sync"
	"time"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/atc"
	"github.com/concourse/atc/db"
	"github.com/concourse/atc/db/lock"
)

type FakeResourceConfig struct {
	IDStub        func() int
	iDMutex       sync.RWMutex
	iDArgsForCall []struct{}
	iDReturns     struct {
		result1 int
	}
	iDReturnsOnCall map[int]struct {
		result1 int
	}
	CheckErrorStub        func() error
	checkErrorMutex       sync.RWMutex
	checkErrorArgsForCall []struct{}
	checkErrorReturns     struct {
		result1 error
	}
	checkErrorReturnsOnCall map[int]struct {
		result1 error
	}
	CreatedByResourceCacheStub        func() db.UsedResourceCache
	createdByResourceCacheMutex       sync.RWMutex
	createdByResourceCacheArgsForCall []struct{}
	createdByResourceCacheReturns     struct {
		result1 db.UsedResourceCache
	}
	createdByResourceCacheReturnsOnCall map[int]struct {
		result1 db.UsedResourceCache
	}
	CreatedByBaseResourceTypeStub        func() *db.UsedBaseResourceType
	createdByBaseResourceTypeMutex       sync.RWMutex
	createdByBaseResourceTypeArgsForCall []struct{}
	createdByBaseResourceTypeReturns     struct {
		result1 *db.UsedBaseResourceType
	}
	createdByBaseResourceTypeReturnsOnCall map[int]struct {
		result1 *db.UsedBaseResourceType
	}
	OriginBaseResourceTypeStub        func() *db.UsedBaseResourceType
	originBaseResourceTypeMutex       sync.RWMutex
	originBaseResourceTypeArgsForCall []struct{}
	originBaseResourceTypeReturns     struct {
		result1 *db.UsedBaseResourceType
	}
	originBaseResourceTypeReturnsOnCall map[int]struct {
		result1 *db.UsedBaseResourceType
	}
	AcquireResourceConfigCheckingLockWithIntervalCheckStub        func(logger lager.Logger, interval time.Duration, immediate bool) (lock.Lock, bool, error)
	acquireResourceConfigCheckingLockWithIntervalCheckMutex       sync.RWMutex
	acquireResourceConfigCheckingLockWithIntervalCheckArgsForCall []struct {
		logger    lager.Logger
		interval  time.Duration
		immediate bool
	}
	acquireResourceConfigCheckingLockWithIntervalCheckReturns struct {
		result1 lock.Lock
		result2 bool
		result3 error
	}
	acquireResourceConfigCheckingLockWithIntervalCheckReturnsOnCall map[int]struct {
		result1 lock.Lock
		result2 bool
		result3 error
	}
	LatestVersionStub        func() (db.ResourceConfigVersion, bool, error)
	latestVersionMutex       sync.RWMutex
	latestVersionArgsForCall []struct{}
	latestVersionReturns     struct {
		result1 db.ResourceConfigVersion
		result2 bool
		result3 error
	}
	latestVersionReturnsOnCall map[int]struct {
		result1 db.ResourceConfigVersion
		result2 bool
		result3 error
	}
	SaveVersionStub        func(version atc.Version, metadata db.ResourceConfigMetadataFields) (bool, error)
	saveVersionMutex       sync.RWMutex
	saveVersionArgsForCall []struct {
		version  atc.Version
		metadata db.ResourceConfigMetadataFields
	}
	saveVersionReturns struct {
		result1 bool
		result2 error
	}
	saveVersionReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	SaveVersionsStub        func(versions []atc.Version) error
	saveVersionsMutex       sync.RWMutex
	saveVersionsArgsForCall []struct {
		versions []atc.Version
	}
	saveVersionsReturns struct {
		result1 error
	}
	saveVersionsReturnsOnCall map[int]struct {
		result1 error
	}
	SetCheckErrorStub        func(error) error
	setCheckErrorMutex       sync.RWMutex
	setCheckErrorArgsForCall []struct {
		arg1 error
	}
	setCheckErrorReturns struct {
		result1 error
	}
	setCheckErrorReturnsOnCall map[int]struct {
		result1 error
	}
	FindVersionStub        func(atc.Version) (db.ResourceConfigVersion, bool, error)
	findVersionMutex       sync.RWMutex
	findVersionArgsForCall []struct {
		arg1 atc.Version
	}
	findVersionReturns struct {
		result1 db.ResourceConfigVersion
		result2 bool
		result3 error
	}
	findVersionReturnsOnCall map[int]struct {
		result1 db.ResourceConfigVersion
		result2 bool
		result3 error
	}
	VersionsStub        func(page db.Page) (db.ResourceConfigVersions, db.Pagination, bool, error)
	versionsMutex       sync.RWMutex
	versionsArgsForCall []struct {
		page db.Page
	}
	versionsReturns struct {
		result1 db.ResourceConfigVersions
		result2 db.Pagination
		result3 bool
		result4 error
	}
	versionsReturnsOnCall map[int]struct {
		result1 db.ResourceConfigVersions
		result2 db.Pagination
		result3 bool
		result4 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeResourceConfig) ID() int {
	fake.iDMutex.Lock()
	ret, specificReturn := fake.iDReturnsOnCall[len(fake.iDArgsForCall)]
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct{}{})
	fake.recordInvocation("ID", []interface{}{})
	fake.iDMutex.Unlock()
	if fake.IDStub != nil {
		return fake.IDStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.iDReturns.result1
}

func (fake *FakeResourceConfig) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeResourceConfig) IDReturns(result1 int) {
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeResourceConfig) IDReturnsOnCall(i int, result1 int) {
	fake.IDStub = nil
	if fake.iDReturnsOnCall == nil {
		fake.iDReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.iDReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeResourceConfig) CheckError() error {
	fake.checkErrorMutex.Lock()
	ret, specificReturn := fake.checkErrorReturnsOnCall[len(fake.checkErrorArgsForCall)]
	fake.checkErrorArgsForCall = append(fake.checkErrorArgsForCall, struct{}{})
	fake.recordInvocation("CheckError", []interface{}{})
	fake.checkErrorMutex.Unlock()
	if fake.CheckErrorStub != nil {
		return fake.CheckErrorStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.checkErrorReturns.result1
}

func (fake *FakeResourceConfig) CheckErrorCallCount() int {
	fake.checkErrorMutex.RLock()
	defer fake.checkErrorMutex.RUnlock()
	return len(fake.checkErrorArgsForCall)
}

func (fake *FakeResourceConfig) CheckErrorReturns(result1 error) {
	fake.CheckErrorStub = nil
	fake.checkErrorReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceConfig) CheckErrorReturnsOnCall(i int, result1 error) {
	fake.CheckErrorStub = nil
	if fake.checkErrorReturnsOnCall == nil {
		fake.checkErrorReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.checkErrorReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceConfig) CreatedByResourceCache() db.UsedResourceCache {
	fake.createdByResourceCacheMutex.Lock()
	ret, specificReturn := fake.createdByResourceCacheReturnsOnCall[len(fake.createdByResourceCacheArgsForCall)]
	fake.createdByResourceCacheArgsForCall = append(fake.createdByResourceCacheArgsForCall, struct{}{})
	fake.recordInvocation("CreatedByResourceCache", []interface{}{})
	fake.createdByResourceCacheMutex.Unlock()
	if fake.CreatedByResourceCacheStub != nil {
		return fake.CreatedByResourceCacheStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createdByResourceCacheReturns.result1
}

func (fake *FakeResourceConfig) CreatedByResourceCacheCallCount() int {
	fake.createdByResourceCacheMutex.RLock()
	defer fake.createdByResourceCacheMutex.RUnlock()
	return len(fake.createdByResourceCacheArgsForCall)
}

func (fake *FakeResourceConfig) CreatedByResourceCacheReturns(result1 db.UsedResourceCache) {
	fake.CreatedByResourceCacheStub = nil
	fake.createdByResourceCacheReturns = struct {
		result1 db.UsedResourceCache
	}{result1}
}

func (fake *FakeResourceConfig) CreatedByResourceCacheReturnsOnCall(i int, result1 db.UsedResourceCache) {
	fake.CreatedByResourceCacheStub = nil
	if fake.createdByResourceCacheReturnsOnCall == nil {
		fake.createdByResourceCacheReturnsOnCall = make(map[int]struct {
			result1 db.UsedResourceCache
		})
	}
	fake.createdByResourceCacheReturnsOnCall[i] = struct {
		result1 db.UsedResourceCache
	}{result1}
}

func (fake *FakeResourceConfig) CreatedByBaseResourceType() *db.UsedBaseResourceType {
	fake.createdByBaseResourceTypeMutex.Lock()
	ret, specificReturn := fake.createdByBaseResourceTypeReturnsOnCall[len(fake.createdByBaseResourceTypeArgsForCall)]
	fake.createdByBaseResourceTypeArgsForCall = append(fake.createdByBaseResourceTypeArgsForCall, struct{}{})
	fake.recordInvocation("CreatedByBaseResourceType", []interface{}{})
	fake.createdByBaseResourceTypeMutex.Unlock()
	if fake.CreatedByBaseResourceTypeStub != nil {
		return fake.CreatedByBaseResourceTypeStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createdByBaseResourceTypeReturns.result1
}

func (fake *FakeResourceConfig) CreatedByBaseResourceTypeCallCount() int {
	fake.createdByBaseResourceTypeMutex.RLock()
	defer fake.createdByBaseResourceTypeMutex.RUnlock()
	return len(fake.createdByBaseResourceTypeArgsForCall)
}

func (fake *FakeResourceConfig) CreatedByBaseResourceTypeReturns(result1 *db.UsedBaseResourceType) {
	fake.CreatedByBaseResourceTypeStub = nil
	fake.createdByBaseResourceTypeReturns = struct {
		result1 *db.UsedBaseResourceType
	}{result1}
}

func (fake *FakeResourceConfig) CreatedByBaseResourceTypeReturnsOnCall(i int, result1 *db.UsedBaseResourceType) {
	fake.CreatedByBaseResourceTypeStub = nil
	if fake.createdByBaseResourceTypeReturnsOnCall == nil {
		fake.createdByBaseResourceTypeReturnsOnCall = make(map[int]struct {
			result1 *db.UsedBaseResourceType
		})
	}
	fake.createdByBaseResourceTypeReturnsOnCall[i] = struct {
		result1 *db.UsedBaseResourceType
	}{result1}
}

func (fake *FakeResourceConfig) OriginBaseResourceType() *db.UsedBaseResourceType {
	fake.originBaseResourceTypeMutex.Lock()
	ret, specificReturn := fake.originBaseResourceTypeReturnsOnCall[len(fake.originBaseResourceTypeArgsForCall)]
	fake.originBaseResourceTypeArgsForCall = append(fake.originBaseResourceTypeArgsForCall, struct{}{})
	fake.recordInvocation("OriginBaseResourceType", []interface{}{})
	fake.originBaseResourceTypeMutex.Unlock()
	if fake.OriginBaseResourceTypeStub != nil {
		return fake.OriginBaseResourceTypeStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.originBaseResourceTypeReturns.result1
}

func (fake *FakeResourceConfig) OriginBaseResourceTypeCallCount() int {
	fake.originBaseResourceTypeMutex.RLock()
	defer fake.originBaseResourceTypeMutex.RUnlock()
	return len(fake.originBaseResourceTypeArgsForCall)
}

func (fake *FakeResourceConfig) OriginBaseResourceTypeReturns(result1 *db.UsedBaseResourceType) {
	fake.OriginBaseResourceTypeStub = nil
	fake.originBaseResourceTypeReturns = struct {
		result1 *db.UsedBaseResourceType
	}{result1}
}

func (fake *FakeResourceConfig) OriginBaseResourceTypeReturnsOnCall(i int, result1 *db.UsedBaseResourceType) {
	fake.OriginBaseResourceTypeStub = nil
	if fake.originBaseResourceTypeReturnsOnCall == nil {
		fake.originBaseResourceTypeReturnsOnCall = make(map[int]struct {
			result1 *db.UsedBaseResourceType
		})
	}
	fake.originBaseResourceTypeReturnsOnCall[i] = struct {
		result1 *db.UsedBaseResourceType
	}{result1}
}

func (fake *FakeResourceConfig) AcquireResourceConfigCheckingLockWithIntervalCheck(logger lager.Logger, interval time.Duration, immediate bool) (lock.Lock, bool, error) {
	fake.acquireResourceConfigCheckingLockWithIntervalCheckMutex.Lock()
	ret, specificReturn := fake.acquireResourceConfigCheckingLockWithIntervalCheckReturnsOnCall[len(fake.acquireResourceConfigCheckingLockWithIntervalCheckArgsForCall)]
	fake.acquireResourceConfigCheckingLockWithIntervalCheckArgsForCall = append(fake.acquireResourceConfigCheckingLockWithIntervalCheckArgsForCall, struct {
		logger    lager.Logger
		interval  time.Duration
		immediate bool
	}{logger, interval, immediate})
	fake.recordInvocation("AcquireResourceConfigCheckingLockWithIntervalCheck", []interface{}{logger, interval, immediate})
	fake.acquireResourceConfigCheckingLockWithIntervalCheckMutex.Unlock()
	if fake.AcquireResourceConfigCheckingLockWithIntervalCheckStub != nil {
		return fake.AcquireResourceConfigCheckingLockWithIntervalCheckStub(logger, interval, immediate)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.acquireResourceConfigCheckingLockWithIntervalCheckReturns.result1, fake.acquireResourceConfigCheckingLockWithIntervalCheckReturns.result2, fake.acquireResourceConfigCheckingLockWithIntervalCheckReturns.result3
}

func (fake *FakeResourceConfig) AcquireResourceConfigCheckingLockWithIntervalCheckCallCount() int {
	fake.acquireResourceConfigCheckingLockWithIntervalCheckMutex.RLock()
	defer fake.acquireResourceConfigCheckingLockWithIntervalCheckMutex.RUnlock()
	return len(fake.acquireResourceConfigCheckingLockWithIntervalCheckArgsForCall)
}

func (fake *FakeResourceConfig) AcquireResourceConfigCheckingLockWithIntervalCheckArgsForCall(i int) (lager.Logger, time.Duration, bool) {
	fake.acquireResourceConfigCheckingLockWithIntervalCheckMutex.RLock()
	defer fake.acquireResourceConfigCheckingLockWithIntervalCheckMutex.RUnlock()
	return fake.acquireResourceConfigCheckingLockWithIntervalCheckArgsForCall[i].logger, fake.acquireResourceConfigCheckingLockWithIntervalCheckArgsForCall[i].interval, fake.acquireResourceConfigCheckingLockWithIntervalCheckArgsForCall[i].immediate
}

func (fake *FakeResourceConfig) AcquireResourceConfigCheckingLockWithIntervalCheckReturns(result1 lock.Lock, result2 bool, result3 error) {
	fake.AcquireResourceConfigCheckingLockWithIntervalCheckStub = nil
	fake.acquireResourceConfigCheckingLockWithIntervalCheckReturns = struct {
		result1 lock.Lock
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeResourceConfig) AcquireResourceConfigCheckingLockWithIntervalCheckReturnsOnCall(i int, result1 lock.Lock, result2 bool, result3 error) {
	fake.AcquireResourceConfigCheckingLockWithIntervalCheckStub = nil
	if fake.acquireResourceConfigCheckingLockWithIntervalCheckReturnsOnCall == nil {
		fake.acquireResourceConfigCheckingLockWithIntervalCheckReturnsOnCall = make(map[int]struct {
			result1 lock.Lock
			result2 bool
			result3 error
		})
	}
	fake.acquireResourceConfigCheckingLockWithIntervalCheckReturnsOnCall[i] = struct {
		result1 lock.Lock
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeResourceConfig) LatestVersion() (db.ResourceConfigVersion, bool, error) {
	fake.latestVersionMutex.Lock()
	ret, specificReturn := fake.latestVersionReturnsOnCall[len(fake.latestVersionArgsForCall)]
	fake.latestVersionArgsForCall = append(fake.latestVersionArgsForCall, struct{}{})
	fake.recordInvocation("LatestVersion", []interface{}{})
	fake.latestVersionMutex.Unlock()
	if fake.LatestVersionStub != nil {
		return fake.LatestVersionStub()
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.latestVersionReturns.result1, fake.latestVersionReturns.result2, fake.latestVersionReturns.result3
}

func (fake *FakeResourceConfig) LatestVersionCallCount() int {
	fake.latestVersionMutex.RLock()
	defer fake.latestVersionMutex.RUnlock()
	return len(fake.latestVersionArgsForCall)
}

func (fake *FakeResourceConfig) LatestVersionReturns(result1 db.ResourceConfigVersion, result2 bool, result3 error) {
	fake.LatestVersionStub = nil
	fake.latestVersionReturns = struct {
		result1 db.ResourceConfigVersion
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeResourceConfig) LatestVersionReturnsOnCall(i int, result1 db.ResourceConfigVersion, result2 bool, result3 error) {
	fake.LatestVersionStub = nil
	if fake.latestVersionReturnsOnCall == nil {
		fake.latestVersionReturnsOnCall = make(map[int]struct {
			result1 db.ResourceConfigVersion
			result2 bool
			result3 error
		})
	}
	fake.latestVersionReturnsOnCall[i] = struct {
		result1 db.ResourceConfigVersion
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeResourceConfig) SaveVersion(version atc.Version, metadata db.ResourceConfigMetadataFields) (bool, error) {
	fake.saveVersionMutex.Lock()
	ret, specificReturn := fake.saveVersionReturnsOnCall[len(fake.saveVersionArgsForCall)]
	fake.saveVersionArgsForCall = append(fake.saveVersionArgsForCall, struct {
		version  atc.Version
		metadata db.ResourceConfigMetadataFields
	}{version, metadata})
	fake.recordInvocation("SaveVersion", []interface{}{version, metadata})
	fake.saveVersionMutex.Unlock()
	if fake.SaveVersionStub != nil {
		return fake.SaveVersionStub(version, metadata)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.saveVersionReturns.result1, fake.saveVersionReturns.result2
}

func (fake *FakeResourceConfig) SaveVersionCallCount() int {
	fake.saveVersionMutex.RLock()
	defer fake.saveVersionMutex.RUnlock()
	return len(fake.saveVersionArgsForCall)
}

func (fake *FakeResourceConfig) SaveVersionArgsForCall(i int) (atc.Version, db.ResourceConfigMetadataFields) {
	fake.saveVersionMutex.RLock()
	defer fake.saveVersionMutex.RUnlock()
	return fake.saveVersionArgsForCall[i].version, fake.saveVersionArgsForCall[i].metadata
}

func (fake *FakeResourceConfig) SaveVersionReturns(result1 bool, result2 error) {
	fake.SaveVersionStub = nil
	fake.saveVersionReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceConfig) SaveVersionReturnsOnCall(i int, result1 bool, result2 error) {
	fake.SaveVersionStub = nil
	if fake.saveVersionReturnsOnCall == nil {
		fake.saveVersionReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.saveVersionReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeResourceConfig) SaveVersions(versions []atc.Version) error {
	var versionsCopy []atc.Version
	if versions != nil {
		versionsCopy = make([]atc.Version, len(versions))
		copy(versionsCopy, versions)
	}
	fake.saveVersionsMutex.Lock()
	ret, specificReturn := fake.saveVersionsReturnsOnCall[len(fake.saveVersionsArgsForCall)]
	fake.saveVersionsArgsForCall = append(fake.saveVersionsArgsForCall, struct {
		versions []atc.Version
	}{versionsCopy})
	fake.recordInvocation("SaveVersions", []interface{}{versionsCopy})
	fake.saveVersionsMutex.Unlock()
	if fake.SaveVersionsStub != nil {
		return fake.SaveVersionsStub(versions)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.saveVersionsReturns.result1
}

func (fake *FakeResourceConfig) SaveVersionsCallCount() int {
	fake.saveVersionsMutex.RLock()
	defer fake.saveVersionsMutex.RUnlock()
	return len(fake.saveVersionsArgsForCall)
}

func (fake *FakeResourceConfig) SaveVersionsArgsForCall(i int) []atc.Version {
	fake.saveVersionsMutex.RLock()
	defer fake.saveVersionsMutex.RUnlock()
	return fake.saveVersionsArgsForCall[i].versions
}

func (fake *FakeResourceConfig) SaveVersionsReturns(result1 error) {
	fake.SaveVersionsStub = nil
	fake.saveVersionsReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceConfig) SaveVersionsReturnsOnCall(i int, result1 error) {
	fake.SaveVersionsStub = nil
	if fake.saveVersionsReturnsOnCall == nil {
		fake.saveVersionsReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.saveVersionsReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceConfig) SetCheckError(arg1 error) error {
	fake.setCheckErrorMutex.Lock()
	ret, specificReturn := fake.setCheckErrorReturnsOnCall[len(fake.setCheckErrorArgsForCall)]
	fake.setCheckErrorArgsForCall = append(fake.setCheckErrorArgsForCall, struct {
		arg1 error
	}{arg1})
	fake.recordInvocation("SetCheckError", []interface{}{arg1})
	fake.setCheckErrorMutex.Unlock()
	if fake.SetCheckErrorStub != nil {
		return fake.SetCheckErrorStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.setCheckErrorReturns.result1
}

func (fake *FakeResourceConfig) SetCheckErrorCallCount() int {
	fake.setCheckErrorMutex.RLock()
	defer fake.setCheckErrorMutex.RUnlock()
	return len(fake.setCheckErrorArgsForCall)
}

func (fake *FakeResourceConfig) SetCheckErrorArgsForCall(i int) error {
	fake.setCheckErrorMutex.RLock()
	defer fake.setCheckErrorMutex.RUnlock()
	return fake.setCheckErrorArgsForCall[i].arg1
}

func (fake *FakeResourceConfig) SetCheckErrorReturns(result1 error) {
	fake.SetCheckErrorStub = nil
	fake.setCheckErrorReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceConfig) SetCheckErrorReturnsOnCall(i int, result1 error) {
	fake.SetCheckErrorStub = nil
	if fake.setCheckErrorReturnsOnCall == nil {
		fake.setCheckErrorReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setCheckErrorReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeResourceConfig) FindVersion(arg1 atc.Version) (db.ResourceConfigVersion, bool, error) {
	fake.findVersionMutex.Lock()
	ret, specificReturn := fake.findVersionReturnsOnCall[len(fake.findVersionArgsForCall)]
	fake.findVersionArgsForCall = append(fake.findVersionArgsForCall, struct {
		arg1 atc.Version
	}{arg1})
	fake.recordInvocation("FindVersion", []interface{}{arg1})
	fake.findVersionMutex.Unlock()
	if fake.FindVersionStub != nil {
		return fake.FindVersionStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.findVersionReturns.result1, fake.findVersionReturns.result2, fake.findVersionReturns.result3
}

func (fake *FakeResourceConfig) FindVersionCallCount() int {
	fake.findVersionMutex.RLock()
	defer fake.findVersionMutex.RUnlock()
	return len(fake.findVersionArgsForCall)
}

func (fake *FakeResourceConfig) FindVersionArgsForCall(i int) atc.Version {
	fake.findVersionMutex.RLock()
	defer fake.findVersionMutex.RUnlock()
	return fake.findVersionArgsForCall[i].arg1
}

func (fake *FakeResourceConfig) FindVersionReturns(result1 db.ResourceConfigVersion, result2 bool, result3 error) {
	fake.FindVersionStub = nil
	fake.findVersionReturns = struct {
		result1 db.ResourceConfigVersion
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeResourceConfig) FindVersionReturnsOnCall(i int, result1 db.ResourceConfigVersion, result2 bool, result3 error) {
	fake.FindVersionStub = nil
	if fake.findVersionReturnsOnCall == nil {
		fake.findVersionReturnsOnCall = make(map[int]struct {
			result1 db.ResourceConfigVersion
			result2 bool
			result3 error
		})
	}
	fake.findVersionReturnsOnCall[i] = struct {
		result1 db.ResourceConfigVersion
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeResourceConfig) Versions(page db.Page) (db.ResourceConfigVersions, db.Pagination, bool, error) {
	fake.versionsMutex.Lock()
	ret, specificReturn := fake.versionsReturnsOnCall[len(fake.versionsArgsForCall)]
	fake.versionsArgsForCall = append(fake.versionsArgsForCall, struct {
		page db.Page
	}{page})
	fake.recordInvocation("Versions", []interface{}{page})
	fake.versionsMutex.Unlock()
	if fake.VersionsStub != nil {
		return fake.VersionsStub(page)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3, ret.result4
	}
	return fake.versionsReturns.result1, fake.versionsReturns.result2, fake.versionsReturns.result3, fake.versionsReturns.result4
}

func (fake *FakeResourceConfig) VersionsCallCount() int {
	fake.versionsMutex.RLock()
	defer fake.versionsMutex.RUnlock()
	return len(fake.versionsArgsForCall)
}

func (fake *FakeResourceConfig) VersionsArgsForCall(i int) db.Page {
	fake.versionsMutex.RLock()
	defer fake.versionsMutex.RUnlock()
	return fake.versionsArgsForCall[i].page
}

func (fake *FakeResourceConfig) VersionsReturns(result1 db.ResourceConfigVersions, result2 db.Pagination, result3 bool, result4 error) {
	fake.VersionsStub = nil
	fake.versionsReturns = struct {
		result1 db.ResourceConfigVersions
		result2 db.Pagination
		result3 bool
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeResourceConfig) VersionsReturnsOnCall(i int, result1 db.ResourceConfigVersions, result2 db.Pagination, result3 bool, result4 error) {
	fake.VersionsStub = nil
	if fake.versionsReturnsOnCall == nil {
		fake.versionsReturnsOnCall = make(map[int]struct {
			result1 db.ResourceConfigVersions
			result2 db.Pagination
			result3 bool
			result4 error
		})
	}
	fake.versionsReturnsOnCall[i] = struct {
		result1 db.ResourceConfigVersions
		result2 db.Pagination
		result3 bool
		result4 error
	}{result1, result2, result3, result4}
}

func (fake *FakeResourceConfig) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.checkErrorMutex.RLock()
	defer fake.checkErrorMutex.RUnlock()
	fake.createdByResourceCacheMutex.RLock()
	defer fake.createdByResourceCacheMutex.RUnlock()
	fake.createdByBaseResourceTypeMutex.RLock()
	defer fake.createdByBaseResourceTypeMutex.RUnlock()
	fake.originBaseResourceTypeMutex.RLock()
	defer fake.originBaseResourceTypeMutex.RUnlock()
	fake.acquireResourceConfigCheckingLockWithIntervalCheckMutex.RLock()
	defer fake.acquireResourceConfigCheckingLockWithIntervalCheckMutex.RUnlock()
	fake.latestVersionMutex.RLock()
	defer fake.latestVersionMutex.RUnlock()
	fake.saveVersionMutex.RLock()
	defer fake.saveVersionMutex.RUnlock()
	fake.saveVersionsMutex.RLock()
	defer fake.saveVersionsMutex.RUnlock()
	fake.setCheckErrorMutex.RLock()
	defer fake.setCheckErrorMutex.RUnlock()
	fake.findVersionMutex.RLock()
	defer fake.findVersionMutex.RUnlock()
	fake.versionsMutex.RLock()
	defer fake.versionsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeResourceConfig) recordInvocation(key string, args []interface{}) {
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

var _ db.ResourceConfig = new(FakeResourceConfig)
