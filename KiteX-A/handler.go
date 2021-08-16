package main

import (
	"context"
	"github.com/Duslia997/KiteX-A/KiteX-A/kitex_gen/api"
	"github.com/Duslia997/KiteX-A/KiteX-A/rpc"
	apiB "github.com/Duslia997/KiteX-A/KiteX-B/kitex_gen/api"
)

// ServiceAImpl implements the last service interface defined in the IDL.
type ServiceAImpl struct{}

// ServiceA implements the ServiceAImpl interface.
func (s *ServiceAImpl) ServiceA(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	reqB := apiB.NewRequest()
	reqB.Message = req.Message

	respB, err := rpc.ServerBClient.ServiceB(ctx, reqB)
	if err != nil {
		return nil, err
	}

	resp = api.NewResponse()
	resp.SetMessage(respB.GetMessage() + "_ServiceA")
	return
}
