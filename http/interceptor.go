package http

import "github.com/stone-payments/goclienttools/errors"

//Interceptor exposes interceptor methods
type Interceptor interface {
	OnRequest(Options) Options
	OnResponse(Response) (Response, errors.Error)
}

type interceptor struct {
	onRequestCallback  func(Options) Options
	onResponseCallback func(Response) (Response, errors.Error)
}

//OnRequestCallback define request callback interceptor assignature
type OnRequestCallback func(Options) Options

//OnResponseCallback define response callback interceptor assignature
type OnResponseCallback func(Response) (Response, errors.Error)

//NewInterceptor constructs interceptor
func NewInterceptor(onRequestCallback OnRequestCallback, onResponseCallback OnResponseCallback) Interceptor {
	return &interceptor{
		onRequestCallback,
		onResponseCallback,
	}
}

func (p *interceptor) OnRequest(request Options) Options {
	if p.onRequestCallback != nil {
		return p.onRequestCallback(request)
	}

	return request
}

func (p *interceptor) OnResponse(response Response) (Response, errors.Error) {
	if p.onResponseCallback != nil {
		return p.onResponseCallback(response)
	}

	return response, nil
}
