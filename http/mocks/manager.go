// Code generated by mockery v1.0.0
package mocks

import errors "github.com/stone-payments/goclienttools/errors"
import http "github.com/stone-payments/goclienttools/http"
import mock "github.com/stretchr/testify/mock"

// Manager is an autogenerated mock type for the Manager type
type Manager struct {
	mock.Mock
}

// BuildURL provides a mock function with given fields: endpoint, params
func (_m *Manager) BuildURL(endpoint string, params ...interface{}) string {
	var _ca []interface{}
	_ca = append(_ca, endpoint)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, ...interface{}) string); ok {
		r0 = rf(endpoint, params...)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// Request provides a mock function with given fields: method, url, options, interceptors
func (_m *Manager) Request(method http.RequestMethod, url string, options http.Options, interceptors ...http.Interceptor) (http.Response, errors.Error) {
	_va := make([]interface{}, len(interceptors))
	for _i := range interceptors {
		_va[_i] = interceptors[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, method, url, options)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 http.Response
	if rf, ok := ret.Get(0).(func(http.RequestMethod, string, http.Options, ...http.Interceptor) http.Response); ok {
		r0 = rf(method, url, options, interceptors...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(http.Response)
		}
	}

	var r1 errors.Error
	if rf, ok := ret.Get(1).(func(http.RequestMethod, string, http.Options, ...http.Interceptor) errors.Error); ok {
		r1 = rf(method, url, options, interceptors...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.Error)
		}
	}

	return r0, r1
}
