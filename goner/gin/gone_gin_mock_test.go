// Code generated by MockGen. DO NOT EDIT.
// Source: ../../gin_interface.go

// Package gin is a generated GoMock package.
package gin

import (
	goneMock "github.com/gone-io/gone"
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	gone "github.com/gone-io/gone"
)

// MockIRoutes is a mock of IRoutes interface.
type MockIRoutes struct {
	goneMock.Flag
	ctrl     *gomock.Controller
	recorder *MockIRoutesMockRecorder
}

// MockIRoutesMockRecorder is the mock recorder for MockIRoutes.
type MockIRoutesMockRecorder struct {
	mock *MockIRoutes
}

// NewMockIRoutes creates a new mock instance.
func NewMockIRoutes(ctrl *gomock.Controller) *MockIRoutes {
	mock := &MockIRoutes{ctrl: ctrl}
	mock.recorder = &MockIRoutesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRoutes) EXPECT() *MockIRoutesMockRecorder {
	return m.recorder
}

// Any mocks base method.
func (m *MockIRoutes) Any(arg0 string, arg1 ...gone.HandlerFunc) gone.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Any", varargs...)
	ret0, _ := ret[0].(gone.IRoutes)
	return ret0
}

// Any indicates an expected call of Any.
func (mr *MockIRoutesMockRecorder) Any(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Any", reflect.TypeOf((*MockIRoutes)(nil).Any), varargs...)
}

// DELETE mocks base method.
func (m *MockIRoutes) DELETE(arg0 string, arg1 ...gone.HandlerFunc) gone.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DELETE", varargs...)
	ret0, _ := ret[0].(gone.IRoutes)
	return ret0
}

// DELETE indicates an expected call of DELETE.
func (mr *MockIRoutesMockRecorder) DELETE(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DELETE", reflect.TypeOf((*MockIRoutes)(nil).DELETE), varargs...)
}

// GET mocks base method.
func (m *MockIRoutes) GET(arg0 string, arg1 ...gone.HandlerFunc) gone.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GET", varargs...)
	ret0, _ := ret[0].(gone.IRoutes)
	return ret0
}

// GET indicates an expected call of GET.
func (mr *MockIRoutesMockRecorder) GET(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GET", reflect.TypeOf((*MockIRoutes)(nil).GET), varargs...)
}

// HEAD mocks base method.
func (m *MockIRoutes) HEAD(arg0 string, arg1 ...gone.HandlerFunc) gone.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HEAD", varargs...)
	ret0, _ := ret[0].(gone.IRoutes)
	return ret0
}

// HEAD indicates an expected call of HEAD.
func (mr *MockIRoutesMockRecorder) HEAD(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HEAD", reflect.TypeOf((*MockIRoutes)(nil).HEAD), varargs...)
}

// Handle mocks base method.
func (m *MockIRoutes) Handle(arg0, arg1 string, arg2 ...gone.HandlerFunc) gone.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Handle", varargs...)
	ret0, _ := ret[0].(gone.IRoutes)
	return ret0
}

// Handle indicates an expected call of Handle.
func (mr *MockIRoutesMockRecorder) Handle(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockIRoutes)(nil).Handle), varargs...)
}

// OPTIONS mocks base method.
func (m *MockIRoutes) OPTIONS(arg0 string, arg1 ...gone.HandlerFunc) gone.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "OPTIONS", varargs...)
	ret0, _ := ret[0].(gone.IRoutes)
	return ret0
}

// OPTIONS indicates an expected call of OPTIONS.
func (mr *MockIRoutesMockRecorder) OPTIONS(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OPTIONS", reflect.TypeOf((*MockIRoutes)(nil).OPTIONS), varargs...)
}

// PATCH mocks base method.
func (m *MockIRoutes) PATCH(arg0 string, arg1 ...gone.HandlerFunc) gone.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PATCH", varargs...)
	ret0, _ := ret[0].(gone.IRoutes)
	return ret0
}

// PATCH indicates an expected call of PATCH.
func (mr *MockIRoutesMockRecorder) PATCH(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PATCH", reflect.TypeOf((*MockIRoutes)(nil).PATCH), varargs...)
}

// POST mocks base method.
func (m *MockIRoutes) POST(arg0 string, arg1 ...gone.HandlerFunc) gone.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "POST", varargs...)
	ret0, _ := ret[0].(gone.IRoutes)
	return ret0
}

// POST indicates an expected call of POST.
func (mr *MockIRoutesMockRecorder) POST(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "POST", reflect.TypeOf((*MockIRoutes)(nil).POST), varargs...)
}

// PUT mocks base method.
func (m *MockIRoutes) PUT(arg0 string, arg1 ...gone.HandlerFunc) gone.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PUT", varargs...)
	ret0, _ := ret[0].(gone.IRoutes)
	return ret0
}

// PUT indicates an expected call of PUT.
func (mr *MockIRoutesMockRecorder) PUT(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PUT", reflect.TypeOf((*MockIRoutes)(nil).PUT), varargs...)
}

// Use mocks base method.
func (m *MockIRoutes) Use(arg0 ...gone.HandlerFunc) gone.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Use", varargs...)
	ret0, _ := ret[0].(gone.IRoutes)
	return ret0
}

// Use indicates an expected call of Use.
func (mr *MockIRoutesMockRecorder) Use(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Use", reflect.TypeOf((*MockIRoutes)(nil).Use), arg0...)
}

// MockIRouter is a mock of IRouter interface.
type MockIRouter struct {
	goneMock.Flag
	ctrl     *gomock.Controller
	recorder *MockIRouterMockRecorder
}

// MockIRouterMockRecorder is the mock recorder for MockIRouter.
type MockIRouterMockRecorder struct {
	mock *MockIRouter
}

// NewMockIRouter creates a new mock instance.
func NewMockIRouter(ctrl *gomock.Controller) *MockIRouter {
	mock := &MockIRouter{ctrl: ctrl}
	mock.recorder = &MockIRouterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIRouter) EXPECT() *MockIRouterMockRecorder {
	return m.recorder
}

// Any mocks base method.
func (m *MockIRouter) Any(arg0 string, arg1 ...gone.HandlerFunc) gone.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Any", varargs...)
	ret0, _ := ret[0].(gone.IRoutes)
	return ret0
}

// Any indicates an expected call of Any.
func (mr *MockIRouterMockRecorder) Any(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Any", reflect.TypeOf((*MockIRouter)(nil).Any), varargs...)
}

// DELETE mocks base method.
func (m *MockIRouter) DELETE(arg0 string, arg1 ...gone.HandlerFunc) gone.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DELETE", varargs...)
	ret0, _ := ret[0].(gone.IRoutes)
	return ret0
}

// DELETE indicates an expected call of DELETE.
func (mr *MockIRouterMockRecorder) DELETE(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DELETE", reflect.TypeOf((*MockIRouter)(nil).DELETE), varargs...)
}

// GET mocks base method.
func (m *MockIRouter) GET(arg0 string, arg1 ...gone.HandlerFunc) gone.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GET", varargs...)
	ret0, _ := ret[0].(gone.IRoutes)
	return ret0
}

// GET indicates an expected call of GET.
func (mr *MockIRouterMockRecorder) GET(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GET", reflect.TypeOf((*MockIRouter)(nil).GET), varargs...)
}

// GetGinRouter mocks base method.
func (m *MockIRouter) GetGinRouter() gin.IRouter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGinRouter")
	ret0, _ := ret[0].(gin.IRouter)
	return ret0
}

// GetGinRouter indicates an expected call of GetGinRouter.
func (mr *MockIRouterMockRecorder) GetGinRouter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGinRouter", reflect.TypeOf((*MockIRouter)(nil).GetGinRouter))
}

// Group mocks base method.
func (m *MockIRouter) Group(arg0 string, arg1 ...gone.HandlerFunc) gone.RouteGroup {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Group", varargs...)
	ret0, _ := ret[0].(gone.RouteGroup)
	return ret0
}

// Group indicates an expected call of Group.
func (mr *MockIRouterMockRecorder) Group(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Group", reflect.TypeOf((*MockIRouter)(nil).Group), varargs...)
}

// HEAD mocks base method.
func (m *MockIRouter) HEAD(arg0 string, arg1 ...gone.HandlerFunc) gone.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HEAD", varargs...)
	ret0, _ := ret[0].(gone.IRoutes)
	return ret0
}

// HEAD indicates an expected call of HEAD.
func (mr *MockIRouterMockRecorder) HEAD(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HEAD", reflect.TypeOf((*MockIRouter)(nil).HEAD), varargs...)
}

// Handle mocks base method.
func (m *MockIRouter) Handle(arg0, arg1 string, arg2 ...gone.HandlerFunc) gone.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Handle", varargs...)
	ret0, _ := ret[0].(gone.IRoutes)
	return ret0
}

// Handle indicates an expected call of Handle.
func (mr *MockIRouterMockRecorder) Handle(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockIRouter)(nil).Handle), varargs...)
}

// LoadHTMLGlob mocks base method.
func (m *MockIRouter) LoadHTMLGlob(pattern string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LoadHTMLGlob", pattern)
}

// LoadHTMLGlob indicates an expected call of LoadHTMLGlob.
func (mr *MockIRouterMockRecorder) LoadHTMLGlob(pattern interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadHTMLGlob", reflect.TypeOf((*MockIRouter)(nil).LoadHTMLGlob), pattern)
}

// OPTIONS mocks base method.
func (m *MockIRouter) OPTIONS(arg0 string, arg1 ...gone.HandlerFunc) gone.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "OPTIONS", varargs...)
	ret0, _ := ret[0].(gone.IRoutes)
	return ret0
}

// OPTIONS indicates an expected call of OPTIONS.
func (mr *MockIRouterMockRecorder) OPTIONS(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OPTIONS", reflect.TypeOf((*MockIRouter)(nil).OPTIONS), varargs...)
}

// PATCH mocks base method.
func (m *MockIRouter) PATCH(arg0 string, arg1 ...gone.HandlerFunc) gone.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PATCH", varargs...)
	ret0, _ := ret[0].(gone.IRoutes)
	return ret0
}

// PATCH indicates an expected call of PATCH.
func (mr *MockIRouterMockRecorder) PATCH(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PATCH", reflect.TypeOf((*MockIRouter)(nil).PATCH), varargs...)
}

// POST mocks base method.
func (m *MockIRouter) POST(arg0 string, arg1 ...gone.HandlerFunc) gone.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "POST", varargs...)
	ret0, _ := ret[0].(gone.IRoutes)
	return ret0
}

// POST indicates an expected call of POST.
func (mr *MockIRouterMockRecorder) POST(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "POST", reflect.TypeOf((*MockIRouter)(nil).POST), varargs...)
}

// PUT mocks base method.
func (m *MockIRouter) PUT(arg0 string, arg1 ...gone.HandlerFunc) gone.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PUT", varargs...)
	ret0, _ := ret[0].(gone.IRoutes)
	return ret0
}

// PUT indicates an expected call of PUT.
func (mr *MockIRouterMockRecorder) PUT(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PUT", reflect.TypeOf((*MockIRouter)(nil).PUT), varargs...)
}

// Use mocks base method.
func (m *MockIRouter) Use(arg0 ...gone.HandlerFunc) gone.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Use", varargs...)
	ret0, _ := ret[0].(gone.IRoutes)
	return ret0
}

// Use indicates an expected call of Use.
func (mr *MockIRouterMockRecorder) Use(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Use", reflect.TypeOf((*MockIRouter)(nil).Use), arg0...)
}

// MockRouteGroup is a mock of RouteGroup interface.
type MockRouteGroup struct {
	goneMock.Flag
	ctrl     *gomock.Controller
	recorder *MockRouteGroupMockRecorder
}

// MockRouteGroupMockRecorder is the mock recorder for MockRouteGroup.
type MockRouteGroupMockRecorder struct {
	mock *MockRouteGroup
}

// NewMockRouteGroup creates a new mock instance.
func NewMockRouteGroup(ctrl *gomock.Controller) *MockRouteGroup {
	mock := &MockRouteGroup{ctrl: ctrl}
	mock.recorder = &MockRouteGroupMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRouteGroup) EXPECT() *MockRouteGroupMockRecorder {
	return m.recorder
}

// Any mocks base method.
func (m *MockRouteGroup) Any(arg0 string, arg1 ...gone.HandlerFunc) gone.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Any", varargs...)
	ret0, _ := ret[0].(gone.IRoutes)
	return ret0
}

// Any indicates an expected call of Any.
func (mr *MockRouteGroupMockRecorder) Any(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Any", reflect.TypeOf((*MockRouteGroup)(nil).Any), varargs...)
}

// DELETE mocks base method.
func (m *MockRouteGroup) DELETE(arg0 string, arg1 ...gone.HandlerFunc) gone.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DELETE", varargs...)
	ret0, _ := ret[0].(gone.IRoutes)
	return ret0
}

// DELETE indicates an expected call of DELETE.
func (mr *MockRouteGroupMockRecorder) DELETE(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DELETE", reflect.TypeOf((*MockRouteGroup)(nil).DELETE), varargs...)
}

// GET mocks base method.
func (m *MockRouteGroup) GET(arg0 string, arg1 ...gone.HandlerFunc) gone.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GET", varargs...)
	ret0, _ := ret[0].(gone.IRoutes)
	return ret0
}

// GET indicates an expected call of GET.
func (mr *MockRouteGroupMockRecorder) GET(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GET", reflect.TypeOf((*MockRouteGroup)(nil).GET), varargs...)
}

// GetGinRouter mocks base method.
func (m *MockRouteGroup) GetGinRouter() gin.IRouter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGinRouter")
	ret0, _ := ret[0].(gin.IRouter)
	return ret0
}

// GetGinRouter indicates an expected call of GetGinRouter.
func (mr *MockRouteGroupMockRecorder) GetGinRouter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGinRouter", reflect.TypeOf((*MockRouteGroup)(nil).GetGinRouter))
}

// Group mocks base method.
func (m *MockRouteGroup) Group(arg0 string, arg1 ...gone.HandlerFunc) gone.RouteGroup {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Group", varargs...)
	ret0, _ := ret[0].(gone.RouteGroup)
	return ret0
}

// Group indicates an expected call of Group.
func (mr *MockRouteGroupMockRecorder) Group(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Group", reflect.TypeOf((*MockRouteGroup)(nil).Group), varargs...)
}

// HEAD mocks base method.
func (m *MockRouteGroup) HEAD(arg0 string, arg1 ...gone.HandlerFunc) gone.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "HEAD", varargs...)
	ret0, _ := ret[0].(gone.IRoutes)
	return ret0
}

// HEAD indicates an expected call of HEAD.
func (mr *MockRouteGroupMockRecorder) HEAD(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HEAD", reflect.TypeOf((*MockRouteGroup)(nil).HEAD), varargs...)
}

// Handle mocks base method.
func (m *MockRouteGroup) Handle(arg0, arg1 string, arg2 ...gone.HandlerFunc) gone.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Handle", varargs...)
	ret0, _ := ret[0].(gone.IRoutes)
	return ret0
}

// Handle indicates an expected call of Handle.
func (mr *MockRouteGroupMockRecorder) Handle(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Handle", reflect.TypeOf((*MockRouteGroup)(nil).Handle), varargs...)
}

// LoadHTMLGlob mocks base method.
func (m *MockRouteGroup) LoadHTMLGlob(pattern string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "LoadHTMLGlob", pattern)
}

// LoadHTMLGlob indicates an expected call of LoadHTMLGlob.
func (mr *MockRouteGroupMockRecorder) LoadHTMLGlob(pattern interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadHTMLGlob", reflect.TypeOf((*MockRouteGroup)(nil).LoadHTMLGlob), pattern)
}

// OPTIONS mocks base method.
func (m *MockRouteGroup) OPTIONS(arg0 string, arg1 ...gone.HandlerFunc) gone.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "OPTIONS", varargs...)
	ret0, _ := ret[0].(gone.IRoutes)
	return ret0
}

// OPTIONS indicates an expected call of OPTIONS.
func (mr *MockRouteGroupMockRecorder) OPTIONS(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OPTIONS", reflect.TypeOf((*MockRouteGroup)(nil).OPTIONS), varargs...)
}

// PATCH mocks base method.
func (m *MockRouteGroup) PATCH(arg0 string, arg1 ...gone.HandlerFunc) gone.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PATCH", varargs...)
	ret0, _ := ret[0].(gone.IRoutes)
	return ret0
}

// PATCH indicates an expected call of PATCH.
func (mr *MockRouteGroupMockRecorder) PATCH(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PATCH", reflect.TypeOf((*MockRouteGroup)(nil).PATCH), varargs...)
}

// POST mocks base method.
func (m *MockRouteGroup) POST(arg0 string, arg1 ...gone.HandlerFunc) gone.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "POST", varargs...)
	ret0, _ := ret[0].(gone.IRoutes)
	return ret0
}

// POST indicates an expected call of POST.
func (mr *MockRouteGroupMockRecorder) POST(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "POST", reflect.TypeOf((*MockRouteGroup)(nil).POST), varargs...)
}

// PUT mocks base method.
func (m *MockRouteGroup) PUT(arg0 string, arg1 ...gone.HandlerFunc) gone.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "PUT", varargs...)
	ret0, _ := ret[0].(gone.IRoutes)
	return ret0
}

// PUT indicates an expected call of PUT.
func (mr *MockRouteGroupMockRecorder) PUT(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PUT", reflect.TypeOf((*MockRouteGroup)(nil).PUT), varargs...)
}

// Use mocks base method.
func (m *MockRouteGroup) Use(arg0 ...gone.HandlerFunc) gone.IRoutes {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Use", varargs...)
	ret0, _ := ret[0].(gone.IRoutes)
	return ret0
}

// Use indicates an expected call of Use.
func (mr *MockRouteGroupMockRecorder) Use(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Use", reflect.TypeOf((*MockRouteGroup)(nil).Use), arg0...)
}
