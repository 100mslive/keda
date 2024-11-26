// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/scalers/scaler.go
//
// Generated by this command:
//
//	mockgen -destination=pkg/mock/mock_scaler/mock_scaler.go -package=mock_scalers -source=pkg/scalers/scaler.go
//

// Package mock_scalers is a generated GoMock package.
package mock_scalers

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	v2 "k8s.io/api/autoscaling/v2"
	external_metrics "k8s.io/metrics/pkg/apis/external_metrics"
)

// MockScaler is a mock of Scaler interface.
type MockScaler struct {
	ctrl     *gomock.Controller
	recorder *MockScalerMockRecorder
	isgomock struct{}
}

// MockScalerMockRecorder is the mock recorder for MockScaler.
type MockScalerMockRecorder struct {
	mock *MockScaler
}

// NewMockScaler creates a new mock instance.
func NewMockScaler(ctrl *gomock.Controller) *MockScaler {
	mock := &MockScaler{ctrl: ctrl}
	mock.recorder = &MockScalerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScaler) EXPECT() *MockScalerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockScaler) Close(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockScalerMockRecorder) Close(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockScaler)(nil).Close), ctx)
}

// GetMetricSpecForScaling mocks base method.
func (m *MockScaler) GetMetricSpecForScaling(ctx context.Context) []v2.MetricSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetricSpecForScaling", ctx)
	ret0, _ := ret[0].([]v2.MetricSpec)
	return ret0
}

// GetMetricSpecForScaling indicates an expected call of GetMetricSpecForScaling.
func (mr *MockScalerMockRecorder) GetMetricSpecForScaling(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetricSpecForScaling", reflect.TypeOf((*MockScaler)(nil).GetMetricSpecForScaling), ctx)
}

// GetMetricsAndActivity mocks base method.
func (m *MockScaler) GetMetricsAndActivity(ctx context.Context, metricName string) ([]external_metrics.ExternalMetricValue, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetricsAndActivity", ctx, metricName)
	ret0, _ := ret[0].([]external_metrics.ExternalMetricValue)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMetricsAndActivity indicates an expected call of GetMetricsAndActivity.
func (mr *MockScalerMockRecorder) GetMetricsAndActivity(ctx, metricName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetricsAndActivity", reflect.TypeOf((*MockScaler)(nil).GetMetricsAndActivity), ctx, metricName)
}

// MockPushScaler is a mock of PushScaler interface.
type MockPushScaler struct {
	ctrl     *gomock.Controller
	recorder *MockPushScalerMockRecorder
	isgomock struct{}
}

// MockPushScalerMockRecorder is the mock recorder for MockPushScaler.
type MockPushScalerMockRecorder struct {
	mock *MockPushScaler
}

// NewMockPushScaler creates a new mock instance.
func NewMockPushScaler(ctrl *gomock.Controller) *MockPushScaler {
	mock := &MockPushScaler{ctrl: ctrl}
	mock.recorder = &MockPushScalerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPushScaler) EXPECT() *MockPushScalerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockPushScaler) Close(ctx context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockPushScalerMockRecorder) Close(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockPushScaler)(nil).Close), ctx)
}

// GetMetricSpecForScaling mocks base method.
func (m *MockPushScaler) GetMetricSpecForScaling(ctx context.Context) []v2.MetricSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetricSpecForScaling", ctx)
	ret0, _ := ret[0].([]v2.MetricSpec)
	return ret0
}

// GetMetricSpecForScaling indicates an expected call of GetMetricSpecForScaling.
func (mr *MockPushScalerMockRecorder) GetMetricSpecForScaling(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetricSpecForScaling", reflect.TypeOf((*MockPushScaler)(nil).GetMetricSpecForScaling), ctx)
}

// GetMetricsAndActivity mocks base method.
func (m *MockPushScaler) GetMetricsAndActivity(ctx context.Context, metricName string) ([]external_metrics.ExternalMetricValue, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetricsAndActivity", ctx, metricName)
	ret0, _ := ret[0].([]external_metrics.ExternalMetricValue)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMetricsAndActivity indicates an expected call of GetMetricsAndActivity.
func (mr *MockPushScalerMockRecorder) GetMetricsAndActivity(ctx, metricName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetricsAndActivity", reflect.TypeOf((*MockPushScaler)(nil).GetMetricsAndActivity), ctx, metricName)
}

// Run mocks base method.
func (m *MockPushScaler) Run(ctx context.Context, active chan<- bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Run", ctx, active)
}

// Run indicates an expected call of Run.
func (mr *MockPushScalerMockRecorder) Run(ctx, active any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockPushScaler)(nil).Run), ctx, active)
}
