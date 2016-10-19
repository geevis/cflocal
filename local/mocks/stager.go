// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/sclevine/cflocal/plugin (interfaces: Stager)

package mocks

import (
	gomock "github.com/golang/mock/gomock"
	app "github.com/sclevine/cflocal/app"
	io "io"
)

// Mock of Stager interface
type MockStager struct {
	ctrl     *gomock.Controller
	recorder *_MockStagerRecorder
}

// Recorder for MockStager (not exported)
type _MockStagerRecorder struct {
	mock *MockStager
}

func NewMockStager(ctrl *gomock.Controller) *MockStager {
	mock := &MockStager{ctrl: ctrl}
	mock.recorder = &_MockStagerRecorder{mock}
	return mock
}

func (_m *MockStager) EXPECT() *_MockStagerRecorder {
	return _m.recorder
}

func (_m *MockStager) Launcher() (io.ReadCloser, int64, error) {
	ret := _m.ctrl.Call(_m, "Launcher")
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockStagerRecorder) Launcher() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Launcher")
}

func (_m *MockStager) Stage(_param0 string, _param1 app.Colorizer, _param2 *app.StageConfig) (io.ReadCloser, int64, error) {
	ret := _m.ctrl.Call(_m, "Stage", _param0, _param1, _param2)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockStagerRecorder) Stage(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stage", arg0, arg1, arg2)
}
