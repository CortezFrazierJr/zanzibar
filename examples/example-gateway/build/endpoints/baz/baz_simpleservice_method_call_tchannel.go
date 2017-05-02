// Code generated by zanzibar
// @generated

// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by zanzibar
// @generated

// Package baz is generated code used to handle TChannel calls using Thrift.
package baz

import (
	"context"
	"errors"

	"github.com/uber/zanzibar/examples/example-gateway/build/clients"
	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/thriftrw/wire"
	"go.uber.org/zap"

	endpointsBazBaz "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/baz/baz"
	customBaz "github.com/uber/zanzibar/examples/example-gateway/endpoints/baz"
)

// NewSimpleServiceCallHandler creates a handler to be registered with a thrift server.
func NewSimpleServiceCallHandler(
	gateway *zanzibar.Gateway,
) zanzibar.TChannelHandler {
	return &SimpleServiceCallHandler{
		Clients: gateway.Clients.(*clients.Clients),
		Logger:  gateway.Logger,
	}
}

// SimpleServiceCallHandler is the handler for "SimpleService::Call".
type SimpleServiceCallHandler struct {
	Clients *clients.Clients
	Logger  *zap.Logger
}

// Handle handles RPC call of "SimpleService::Call".
func (h *SimpleServiceCallHandler) Handle(
	ctx context.Context,
	reqHeaders map[string]string,
	wireValue *wire.Value,
) (bool, zanzibar.RWTStruct, map[string]string, error) {
	wfReqHeaders := zanzibar.ServerTChannelHeader(reqHeaders)
	if err := wfReqHeaders.Ensure([]string{"x-uuid", "x-token"}); err != nil {
		return false, nil, nil, err
	}

	var res endpointsBazBaz.SimpleService_Call_Result

	var req endpointsBazBaz.SimpleService_Call_Args
	if err := req.FromWire(*wireValue); err != nil {
		return false, nil, nil, err
	}
	workflow := customBaz.CallEndpoint{
		Clients: h.Clients,
		Logger:  h.Logger,
	}

	wfResHeaders, err := workflow.Handle(ctx, wfReqHeaders, &req)

	if err := wfResHeaders.Ensure([]string{"some-res-header"}); err != nil {
		return false, nil, nil, err
	}

	resHeaders := map[string]string{}
	for _, key := range wfResHeaders.Keys() {
		resHeaders[key], _ = wfResHeaders.Get(key)
	}

	if err != nil {
		switch v := err.(type) {
		case *endpointsBazBaz.AuthErr:
			if v == nil {
				return false, nil, resHeaders, errors.New(
					"Handler for Call returned non-nil error type *endpointsBazBaz.AuthErr but nil value",
				)
			}
			res.AuthErr = v
		default:
			return false, nil, resHeaders, err
		}
	}

	return err == nil, &res, resHeaders, nil
}
