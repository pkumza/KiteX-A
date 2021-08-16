package main

import (
	"context"
	"kitex.service.a/kitex_gen/api"
)

// ServiceAImpl implements the last service interface defined in the IDL.
type ServiceAImpl struct{}

// ServiceA implements the ServiceAImpl interface.
func (s *ServiceAImpl) ServiceA(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	// TODO: Your code here...
	return
}
