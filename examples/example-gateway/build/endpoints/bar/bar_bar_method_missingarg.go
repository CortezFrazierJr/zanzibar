// Code generated by zanzibar
// @generated

// Copyright (c) 2018 Uber Technologies, Inc.
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

package barEndpoint

import (
	"context"

	zanzibar "github.com/uber/zanzibar/runtime"
	"go.uber.org/zap"

	clientsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/bar/bar"
	endpointsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints/bar/bar"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/bar/module"
)

// BarMissingArgHandler is the handler for "/bar/missing-arg-path"
type BarMissingArgHandler struct {
	Clients  *module.ClientDependencies
	endpoint *zanzibar.RouterEndpoint
}

// NewBarMissingArgHandler creates a handler
func NewBarMissingArgHandler(
	g *zanzibar.Gateway,
	deps *module.Dependencies,
) *BarMissingArgHandler {
	handler := &BarMissingArgHandler{
		Clients: deps.Client,
	}
	handler.endpoint = zanzibar.NewRouterEndpoint(
		deps.Default.Logger, deps.Default.Scope,
		"bar", "missingArg",
		handler.HandleRequest,
	)
	return handler
}

// Register adds the http handler to the gateway's http router
func (h *BarMissingArgHandler) Register(g *zanzibar.Gateway) error {
	g.HTTPRouter.Register(
		"GET", "/bar/missing-arg-path",
		h.endpoint,
	)
	// TODO: register should return errors on route conflicts
	return nil
}

// HandleRequest handles "/bar/missing-arg-path".
func (h *BarMissingArgHandler) HandleRequest(
	ctx context.Context,
	req *zanzibar.ServerHTTPRequest,
	res *zanzibar.ServerHTTPResponse,
) {

	workflow := BarMissingArgEndpoint{
		Clients: h.Clients,
		Logger:  req.Logger,
		Request: req,
	}

	response, cliRespHeaders, err := workflow.Handle(ctx, req.Header)
	if err != nil {
		switch errValue := err.(type) {

		case *endpointsBarBar.BarException:
			res.WriteJSON(
				403, cliRespHeaders, errValue,
			)
			return

		default:
			req.Logger.Warn("Workflow for endpoint returned error", zap.Error(errValue))
			res.SendErrorString(500, "Unexpected server error")
			return
		}
	}
	// TODO(jakev): implement writing fields into response headers

	res.WriteJSON(200, cliRespHeaders, response)
}

// BarMissingArgEndpoint calls thrift client Bar.MissingArg
type BarMissingArgEndpoint struct {
	Clients *module.ClientDependencies
	Logger  *zap.Logger
	Request *zanzibar.ServerHTTPRequest
}

// Handle calls thrift client.
func (w BarMissingArgEndpoint) Handle(
	ctx context.Context,
	reqHeaders zanzibar.Header,
) (*endpointsBarBar.BarResponse, zanzibar.Header, error) {

	clientHeaders := map[string]string{}

	clientRespBody, _, err := w.Clients.Bar.MissingArg(
		ctx, clientHeaders,
	)

	if err != nil {
		switch errValue := err.(type) {

		case *clientsBarBar.BarException:
			serverErr := convertMissingArgBarException(
				errValue,
			)
			// TODO(sindelar): Consider returning partial headers

			return nil, nil, serverErr

		default:
			w.Logger.Warn("Could not make client request", zap.Error(errValue))
			// TODO(sindelar): Consider returning partial headers

			return nil, nil, err

		}
	}

	// Filter and map response headers from client to server response.

	// TODO: Add support for TChannel Headers with a switch here
	resHeaders := zanzibar.ServerHTTPHeader{}

	response := convertBarMissingArgClientResponse(clientRespBody)
	return response, resHeaders, nil
}

func convertMissingArgBarException(
	clientError *clientsBarBar.BarException,
) *endpointsBarBar.BarException {
	// TODO: Add error fields mapping here.
	serverError := &endpointsBarBar.BarException{}
	return serverError
}

func convertBarMissingArgClientResponse(in *clientsBarBar.BarResponse) *endpointsBarBar.BarResponse {
	out := &endpointsBarBar.BarResponse{}

	out.StringField = string(in.StringField)
	out.IntWithRange = int32(in.IntWithRange)
	out.IntWithoutRange = int32(in.IntWithoutRange)
	out.MapIntWithRange = make(map[endpointsBarBar.UUID]int32, len(in.MapIntWithRange))
	for key1, value2 := range in.MapIntWithRange {
		out.MapIntWithRange[endpointsBarBar.UUID(key1)] = int32(value2)
	}
	out.MapIntWithoutRange = make(map[string]int32, len(in.MapIntWithoutRange))
	for key3, value4 := range in.MapIntWithoutRange {
		out.MapIntWithoutRange[key3] = int32(value4)
	}
	out.BinaryField = []byte(in.BinaryField)
	var convertBarResponseHelper5 func(in *clientsBarBar.BarResponse) (out *endpointsBarBar.BarResponse)
	convertBarResponseHelper5 = func(in *clientsBarBar.BarResponse) (out *endpointsBarBar.BarResponse) {
		if in != nil {
			out = &endpointsBarBar.BarResponse{}
			out.StringField = string(in.StringField)
			out.IntWithRange = int32(in.IntWithRange)
			out.IntWithoutRange = int32(in.IntWithoutRange)
			out.MapIntWithRange = make(map[endpointsBarBar.UUID]int32, len(in.MapIntWithRange))
			for key6, value7 := range in.MapIntWithRange {
				out.MapIntWithRange[endpointsBarBar.UUID(key6)] = int32(value7)
			}
			out.MapIntWithoutRange = make(map[string]int32, len(in.MapIntWithoutRange))
			for key8, value9 := range in.MapIntWithoutRange {
				out.MapIntWithoutRange[key8] = int32(value9)
			}
			out.BinaryField = []byte(in.BinaryField)
			out.NextResponse = convertBarResponseHelper5(in.NextResponse)
		} else {
			out = nil
		}
		return
	}
	out.NextResponse = convertBarResponseHelper5(in.NextResponse)

	return out
}
