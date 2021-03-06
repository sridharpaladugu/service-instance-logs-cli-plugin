// Code generated by counterfeiter. DO NOT EDIT.
package logclientfakes

import (
	"sync"

	"github.com/cloudfoundry/noaa/consumer"
	"github.com/cloudfoundry/sonde-go/events"
	"github.com/pivotal-cf/service-instance-logs-cli-plugin/logclient"
)

type FakeConsumer struct {
	SetRecentPathBuilderStub        func(b consumer.RecentPathBuilder)
	setRecentPathBuilderMutex       sync.RWMutex
	setRecentPathBuilderArgsForCall []struct {
		b consumer.RecentPathBuilder
	}
	SetStreamPathBuilderStub        func(b consumer.StreamPathBuilder)
	setStreamPathBuilderMutex       sync.RWMutex
	setStreamPathBuilderArgsForCall []struct {
		b consumer.StreamPathBuilder
	}
	SetDebugPrinterStub        func(debugPrinter consumer.DebugPrinter)
	setDebugPrinterMutex       sync.RWMutex
	setDebugPrinterArgsForCall []struct {
		debugPrinter consumer.DebugPrinter
	}
	RecentLogsStub        func(appGuid string, authToken string) ([]*events.LogMessage, error)
	recentLogsMutex       sync.RWMutex
	recentLogsArgsForCall []struct {
		appGuid   string
		authToken string
	}
	recentLogsReturns struct {
		result1 []*events.LogMessage
		result2 error
	}
	recentLogsReturnsOnCall map[int]struct {
		result1 []*events.LogMessage
		result2 error
	}
	TailingLogsStub        func(appGuid, authToken string) (<-chan *events.LogMessage, <-chan error)
	tailingLogsMutex       sync.RWMutex
	tailingLogsArgsForCall []struct {
		appGuid   string
		authToken string
	}
	tailingLogsReturns struct {
		result1 <-chan *events.LogMessage
		result2 <-chan error
	}
	tailingLogsReturnsOnCall map[int]struct {
		result1 <-chan *events.LogMessage
		result2 <-chan error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeConsumer) SetRecentPathBuilder(b consumer.RecentPathBuilder) {
	fake.setRecentPathBuilderMutex.Lock()
	fake.setRecentPathBuilderArgsForCall = append(fake.setRecentPathBuilderArgsForCall, struct {
		b consumer.RecentPathBuilder
	}{b})
	fake.recordInvocation("SetRecentPathBuilder", []interface{}{b})
	fake.setRecentPathBuilderMutex.Unlock()
	if fake.SetRecentPathBuilderStub != nil {
		fake.SetRecentPathBuilderStub(b)
	}
}

func (fake *FakeConsumer) SetRecentPathBuilderCallCount() int {
	fake.setRecentPathBuilderMutex.RLock()
	defer fake.setRecentPathBuilderMutex.RUnlock()
	return len(fake.setRecentPathBuilderArgsForCall)
}

func (fake *FakeConsumer) SetRecentPathBuilderArgsForCall(i int) consumer.RecentPathBuilder {
	fake.setRecentPathBuilderMutex.RLock()
	defer fake.setRecentPathBuilderMutex.RUnlock()
	return fake.setRecentPathBuilderArgsForCall[i].b
}

func (fake *FakeConsumer) SetStreamPathBuilder(b consumer.StreamPathBuilder) {
	fake.setStreamPathBuilderMutex.Lock()
	fake.setStreamPathBuilderArgsForCall = append(fake.setStreamPathBuilderArgsForCall, struct {
		b consumer.StreamPathBuilder
	}{b})
	fake.recordInvocation("SetStreamPathBuilder", []interface{}{b})
	fake.setStreamPathBuilderMutex.Unlock()
	if fake.SetStreamPathBuilderStub != nil {
		fake.SetStreamPathBuilderStub(b)
	}
}

func (fake *FakeConsumer) SetStreamPathBuilderCallCount() int {
	fake.setStreamPathBuilderMutex.RLock()
	defer fake.setStreamPathBuilderMutex.RUnlock()
	return len(fake.setStreamPathBuilderArgsForCall)
}

func (fake *FakeConsumer) SetStreamPathBuilderArgsForCall(i int) consumer.StreamPathBuilder {
	fake.setStreamPathBuilderMutex.RLock()
	defer fake.setStreamPathBuilderMutex.RUnlock()
	return fake.setStreamPathBuilderArgsForCall[i].b
}

func (fake *FakeConsumer) SetDebugPrinter(debugPrinter consumer.DebugPrinter) {
	fake.setDebugPrinterMutex.Lock()
	fake.setDebugPrinterArgsForCall = append(fake.setDebugPrinterArgsForCall, struct {
		debugPrinter consumer.DebugPrinter
	}{debugPrinter})
	fake.recordInvocation("SetDebugPrinter", []interface{}{debugPrinter})
	fake.setDebugPrinterMutex.Unlock()
	if fake.SetDebugPrinterStub != nil {
		fake.SetDebugPrinterStub(debugPrinter)
	}
}

func (fake *FakeConsumer) SetDebugPrinterCallCount() int {
	fake.setDebugPrinterMutex.RLock()
	defer fake.setDebugPrinterMutex.RUnlock()
	return len(fake.setDebugPrinterArgsForCall)
}

func (fake *FakeConsumer) SetDebugPrinterArgsForCall(i int) consumer.DebugPrinter {
	fake.setDebugPrinterMutex.RLock()
	defer fake.setDebugPrinterMutex.RUnlock()
	return fake.setDebugPrinterArgsForCall[i].debugPrinter
}

func (fake *FakeConsumer) RecentLogs(appGuid string, authToken string) ([]*events.LogMessage, error) {
	fake.recentLogsMutex.Lock()
	ret, specificReturn := fake.recentLogsReturnsOnCall[len(fake.recentLogsArgsForCall)]
	fake.recentLogsArgsForCall = append(fake.recentLogsArgsForCall, struct {
		appGuid   string
		authToken string
	}{appGuid, authToken})
	fake.recordInvocation("RecentLogs", []interface{}{appGuid, authToken})
	fake.recentLogsMutex.Unlock()
	if fake.RecentLogsStub != nil {
		return fake.RecentLogsStub(appGuid, authToken)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.recentLogsReturns.result1, fake.recentLogsReturns.result2
}

func (fake *FakeConsumer) RecentLogsCallCount() int {
	fake.recentLogsMutex.RLock()
	defer fake.recentLogsMutex.RUnlock()
	return len(fake.recentLogsArgsForCall)
}

func (fake *FakeConsumer) RecentLogsArgsForCall(i int) (string, string) {
	fake.recentLogsMutex.RLock()
	defer fake.recentLogsMutex.RUnlock()
	return fake.recentLogsArgsForCall[i].appGuid, fake.recentLogsArgsForCall[i].authToken
}

func (fake *FakeConsumer) RecentLogsReturns(result1 []*events.LogMessage, result2 error) {
	fake.RecentLogsStub = nil
	fake.recentLogsReturns = struct {
		result1 []*events.LogMessage
		result2 error
	}{result1, result2}
}

func (fake *FakeConsumer) RecentLogsReturnsOnCall(i int, result1 []*events.LogMessage, result2 error) {
	fake.RecentLogsStub = nil
	if fake.recentLogsReturnsOnCall == nil {
		fake.recentLogsReturnsOnCall = make(map[int]struct {
			result1 []*events.LogMessage
			result2 error
		})
	}
	fake.recentLogsReturnsOnCall[i] = struct {
		result1 []*events.LogMessage
		result2 error
	}{result1, result2}
}

func (fake *FakeConsumer) TailingLogs(appGuid string, authToken string) (<-chan *events.LogMessage, <-chan error) {
	fake.tailingLogsMutex.Lock()
	ret, specificReturn := fake.tailingLogsReturnsOnCall[len(fake.tailingLogsArgsForCall)]
	fake.tailingLogsArgsForCall = append(fake.tailingLogsArgsForCall, struct {
		appGuid   string
		authToken string
	}{appGuid, authToken})
	fake.recordInvocation("TailingLogs", []interface{}{appGuid, authToken})
	fake.tailingLogsMutex.Unlock()
	if fake.TailingLogsStub != nil {
		return fake.TailingLogsStub(appGuid, authToken)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.tailingLogsReturns.result1, fake.tailingLogsReturns.result2
}

func (fake *FakeConsumer) TailingLogsCallCount() int {
	fake.tailingLogsMutex.RLock()
	defer fake.tailingLogsMutex.RUnlock()
	return len(fake.tailingLogsArgsForCall)
}

func (fake *FakeConsumer) TailingLogsArgsForCall(i int) (string, string) {
	fake.tailingLogsMutex.RLock()
	defer fake.tailingLogsMutex.RUnlock()
	return fake.tailingLogsArgsForCall[i].appGuid, fake.tailingLogsArgsForCall[i].authToken
}

func (fake *FakeConsumer) TailingLogsReturns(result1 <-chan *events.LogMessage, result2 <-chan error) {
	fake.TailingLogsStub = nil
	fake.tailingLogsReturns = struct {
		result1 <-chan *events.LogMessage
		result2 <-chan error
	}{result1, result2}
}

func (fake *FakeConsumer) TailingLogsReturnsOnCall(i int, result1 <-chan *events.LogMessage, result2 <-chan error) {
	fake.TailingLogsStub = nil
	if fake.tailingLogsReturnsOnCall == nil {
		fake.tailingLogsReturnsOnCall = make(map[int]struct {
			result1 <-chan *events.LogMessage
			result2 <-chan error
		})
	}
	fake.tailingLogsReturnsOnCall[i] = struct {
		result1 <-chan *events.LogMessage
		result2 <-chan error
	}{result1, result2}
}

func (fake *FakeConsumer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.setRecentPathBuilderMutex.RLock()
	defer fake.setRecentPathBuilderMutex.RUnlock()
	fake.setStreamPathBuilderMutex.RLock()
	defer fake.setStreamPathBuilderMutex.RUnlock()
	fake.setDebugPrinterMutex.RLock()
	defer fake.setDebugPrinterMutex.RUnlock()
	fake.recentLogsMutex.RLock()
	defer fake.recentLogsMutex.RUnlock()
	fake.tailingLogsMutex.RLock()
	defer fake.tailingLogsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeConsumer) recordInvocation(key string, args []interface{}) {
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

var _ logclient.Consumer = new(FakeConsumer)
