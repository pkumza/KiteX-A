package main

import (
	"context"
	"github.com/KiteX-A/kitex_gen/api"
)

// HelloImpl implements the last service interface defined in the IDL.
type HelloImpl struct{}

// Echo implements the HelloImpl interface.
func (s *HelloImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	// TODO: Your code here...
	return
}
