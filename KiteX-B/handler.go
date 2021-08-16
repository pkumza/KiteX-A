package main

import (
	"context"
	"kitex.service.b/kitex_gen/api"
)

// ServiceBImpl implements the last service interface defined in the IDL.
type ServiceBImpl struct{}

// ServiceB implements the ServiceBImpl interface.
func (s *ServiceBImpl) ServiceB(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	// TODO: Your code here...
	return
}
