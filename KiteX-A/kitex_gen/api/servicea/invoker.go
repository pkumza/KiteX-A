// Code generated by Kitex v0.0.3. DO NOT EDIT.

package servicea

import (
	"github.com/Duslia997/KiteX-A/KiteX-A/kitex_gen/api"
	"github.com/cloudwego/kitex/server"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler api.ServiceA, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
