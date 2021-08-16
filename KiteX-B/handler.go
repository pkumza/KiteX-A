package main

import (
	"context"
	"github.com/Duslia997/KiteX-A/KiteX-B/kitex_gen/api"
)

// ServiceBImpl implements the last service interface defined in the IDL.
type ServiceBImpl struct{}

// ServiceB implements the ServiceBImpl interface.
func (s *ServiceBImpl) ServiceB(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	resp = api.NewResponse()
	resp.SetMessage(req.GetMessage() + "_ServiceB")

	return
}
