// Code generated by MockGen. DO NOT EDIT.
// Source: interface.go
//
// Generated by this command:
//
//	mockgen -package=gin -self_package=github.com/gone-io/gone/goner/gin -source=interface.go
//

// Package gin is a generated GoMock package.
package gin

import (
	goneMock "github.com/gone-io/gone"
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "go.uber.org/mock/gomock"
)

// MockController is a mock of Controller interface.
type MockController struct {
	goneMock.Flag
	ctrl     *gomock.Controller
	recorder *MockControllerMockRecorder
	isgomock struct{}
}

// MockControllerMockRecorder is the mock recorder for MockController.
type MockControllerMockRecorder struct {
	mock *MockController
}

// NewMockController creates a new mock instance.
func NewMockController(ctrl *gomock.Controller) *MockController {
	mock := &MockController{ctrl: ctrl}
	mock.recorder = &MockControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockController) EXPECT() *MockControllerMockRecorder {
	return m.recorder
}

// Mount mocks base method.
func (m *MockController) Mount() MountError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mount")
	ret0, _ := ret[0].(MountError)
	return ret0
}

// Mount indicates an expected call of Mount.
func (mr *MockControllerMockRecorder) Mount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mount", reflect.TypeOf((*MockController)(nil).Mount))
}

// MockHandleProxyToGin is a mock of HandleProxyToGin interface.
type MockHandleProxyToGin struct {
	goneMock.Flag
	ctrl     *gomock.Controller
	recorder *MockHandleProxyToGinMockRecorder
	isgomock struct{}
}

// MockHandleProxyToGinMockRecorder is the mock recorder for MockHandleProxyToGin.
type MockHandleProxyToGinMockRecorder struct {
	mock *MockHandleProxyToGin
}

// NewMockHandleProxyToGin creates a new mock instance.
func NewMockHandleProxyToGin(ctrl *gomock.Controller) *MockHandleProxyToGin {
	mock := &MockHandleProxyToGin{ctrl: ctrl}
	mock.recorder = &MockHandleProxyToGinMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHandleProxyToGin) EXPECT() *MockHandleProxyToGinMockRecorder {
	return m.recorder
}

// Proxy mocks base method.
func (m *MockHandleProxyToGin) Proxy(handler ...HandlerFunc) []gin.HandlerFunc {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range handler {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Proxy", varargs...)
	ret0, _ := ret[0].([]gin.HandlerFunc)
	return ret0
}

// Proxy indicates an expected call of Proxy.
func (mr *MockHandleProxyToGinMockRecorder) Proxy(handler ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Proxy", reflect.TypeOf((*MockHandleProxyToGin)(nil).Proxy), handler...)
}

// ProxyForMiddleware mocks base method.
func (m *MockHandleProxyToGin) ProxyForMiddleware(handlers ...HandlerFunc) []gin.HandlerFunc {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range handlers {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ProxyForMiddleware", varargs...)
	ret0, _ := ret[0].([]gin.HandlerFunc)
	return ret0
}

// ProxyForMiddleware indicates an expected call of ProxyForMiddleware.
func (mr *MockHandleProxyToGinMockRecorder) ProxyForMiddleware(handlers ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProxyForMiddleware", reflect.TypeOf((*MockHandleProxyToGin)(nil).ProxyForMiddleware), handlers...)
}

// MockXContext is a mock of XContext interface.
type MockXContext struct {
	goneMock.Flag
	ctrl     *gomock.Controller
	recorder *MockXContextMockRecorder
	isgomock struct{}
}

// MockXContextMockRecorder is the mock recorder for MockXContext.
type MockXContextMockRecorder struct {
	mock *MockXContext
}

// NewMockXContext creates a new mock instance.
func NewMockXContext(ctrl *gomock.Controller) *MockXContext {
	mock := &MockXContext{ctrl: ctrl}
	mock.recorder = &MockXContextMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockXContext) EXPECT() *MockXContextMockRecorder {
	return m.recorder
}

// Abort mocks base method.
func (m *MockXContext) Abort() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Abort")
}

// Abort indicates an expected call of Abort.
func (mr *MockXContextMockRecorder) Abort() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Abort", reflect.TypeOf((*MockXContext)(nil).Abort))
}

// JSON mocks base method.
func (m *MockXContext) JSON(code int, obj any) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "JSON", code, obj)
}

// JSON indicates an expected call of JSON.
func (mr *MockXContextMockRecorder) JSON(code, obj any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JSON", reflect.TypeOf((*MockXContext)(nil).JSON), code, obj)
}

// String mocks base method.
func (m *MockXContext) String(code int, format string, values ...any) {
	m.ctrl.T.Helper()
	varargs := []any{code, format}
	for _, a := range values {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "String", varargs...)
}

// String indicates an expected call of String.
func (mr *MockXContextMockRecorder) String(code, format any, values ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{code, format}, values...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockXContext)(nil).String), varargs...)
}

// MockResponser is a mock of Responser interface.
type MockResponser struct {
	goneMock.Flag
	ctrl     *gomock.Controller
	recorder *MockResponserMockRecorder
	isgomock struct{}
}

// MockResponserMockRecorder is the mock recorder for MockResponser.
type MockResponserMockRecorder struct {
	mock *MockResponser
}

// NewMockResponser creates a new mock instance.
func NewMockResponser(ctrl *gomock.Controller) *MockResponser {
	mock := &MockResponser{ctrl: ctrl}
	mock.recorder = &MockResponserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResponser) EXPECT() *MockResponserMockRecorder {
	return m.recorder
}

// Failed mocks base method.
func (m *MockResponser) Failed(ctx XContext, err error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Failed", ctx, err)
}

// Failed indicates an expected call of Failed.
func (mr *MockResponserMockRecorder) Failed(ctx, err any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Failed", reflect.TypeOf((*MockResponser)(nil).Failed), ctx, err)
}

// ProcessResults mocks base method.
func (m *MockResponser) ProcessResults(context XContext, writer gin.ResponseWriter, last bool, funcName string, results ...any) {
	m.ctrl.T.Helper()
	varargs := []any{context, writer, last, funcName}
	for _, a := range results {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "ProcessResults", varargs...)
}

// ProcessResults indicates an expected call of ProcessResults.
func (mr *MockResponserMockRecorder) ProcessResults(context, writer, last, funcName any, results ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{context, writer, last, funcName}, results...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessResults", reflect.TypeOf((*MockResponser)(nil).ProcessResults), varargs...)
}

// Success mocks base method.
func (m *MockResponser) Success(ctx XContext, data any) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Success", ctx, data)
}

// Success indicates an expected call of Success.
func (mr *MockResponserMockRecorder) Success(ctx, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Success", reflect.TypeOf((*MockResponser)(nil).Success), ctx, data)
}

// MockHttInjector is a mock of HttInjector interface.
type MockHttInjector struct {
	goneMock.Flag
	ctrl     *gomock.Controller
	recorder *MockHttInjectorMockRecorder
	isgomock struct{}
}

// MockHttInjectorMockRecorder is the mock recorder for MockHttInjector.
type MockHttInjectorMockRecorder struct {
	mock *MockHttInjector
}

// NewMockHttInjector creates a new mock instance.
func NewMockHttInjector(ctrl *gomock.Controller) *MockHttInjector {
	mock := &MockHttInjector{ctrl: ctrl}
	mock.recorder = &MockHttInjectorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHttInjector) EXPECT() *MockHttInjectorMockRecorder {
	return m.recorder
}

// BindFuncs mocks base method.
func (m *MockHttInjector) BindFuncs() BindStructFunc {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BindFuncs")
	ret0, _ := ret[0].(BindStructFunc)
	return ret0
}

// BindFuncs indicates an expected call of BindFuncs.
func (mr *MockHttInjectorMockRecorder) BindFuncs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BindFuncs", reflect.TypeOf((*MockHttInjector)(nil).BindFuncs))
}

// StartBindFuncs mocks base method.
func (m *MockHttInjector) StartBindFuncs() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartBindFuncs")
}

// StartBindFuncs indicates an expected call of StartBindFuncs.
func (mr *MockHttInjectorMockRecorder) StartBindFuncs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartBindFuncs", reflect.TypeOf((*MockHttInjector)(nil).StartBindFuncs))
}

// MockMiddleware is a mock of Middleware interface.
type MockMiddleware struct {
	goneMock.Flag
	ctrl     *gomock.Controller
	recorder *MockMiddlewareMockRecorder
	isgomock struct{}
}

// MockMiddlewareMockRecorder is the mock recorder for MockMiddleware.
type MockMiddlewareMockRecorder struct {
	mock *MockMiddleware
}

// NewMockMiddleware creates a new mock instance.
func NewMockMiddleware(ctrl *gomock.Controller) *MockMiddleware {
	mock := &MockMiddleware{ctrl: ctrl}
	mock.recorder = &MockMiddlewareMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMiddleware) EXPECT() *MockMiddlewareMockRecorder {
	return m.recorder
}

// Process mocks base method.
func (m *MockMiddleware) Process(ctx *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Process", ctx)
}

// Process indicates an expected call of Process.
func (mr *MockMiddlewareMockRecorder) Process(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Process", reflect.TypeOf((*MockMiddleware)(nil).Process), ctx)
}
