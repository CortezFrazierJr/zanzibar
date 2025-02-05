// Code generated by protoc-gen-yarpc-go. DO NOT EDIT.
// source: clients/mirror/mirror.proto

package mirror

import (
	"context"
	"io/ioutil"
	"reflect"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"go.uber.org/fx"
	"go.uber.org/yarpc"
	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/api/x/restriction"
	"go.uber.org/yarpc/encoding/protobuf"
	"go.uber.org/yarpc/encoding/protobuf/reflection"
)

var _ = ioutil.NopCloser

// MirrorYARPCClient is the YARPC client-side interface for the Mirror service.
type MirrorYARPCClient interface {
	Mirror(context.Context, *Request, ...yarpc.CallOption) (*Response, error)
}

func newMirrorYARPCClient(clientConfig transport.ClientConfig, anyResolver jsonpb.AnyResolver, options ...protobuf.ClientOption) MirrorYARPCClient {
	return &_MirrorYARPCCaller{protobuf.NewStreamClient(
		protobuf.ClientParams{
			ServiceName:  "mirror.Mirror",
			ClientConfig: clientConfig,
			AnyResolver:  anyResolver,
			Options:      options,
		},
	)}
}

// NewMirrorYARPCClient builds a new YARPC client for the Mirror service.
func NewMirrorYARPCClient(clientConfig transport.ClientConfig, options ...protobuf.ClientOption) MirrorYARPCClient {
	return newMirrorYARPCClient(clientConfig, nil, options...)
}

// MirrorYARPCServer is the YARPC server-side interface for the Mirror service.
type MirrorYARPCServer interface {
	Mirror(context.Context, *Request) (*Response, error)
}

type buildMirrorYARPCProceduresParams struct {
	Server      MirrorYARPCServer
	AnyResolver jsonpb.AnyResolver
}

func buildMirrorYARPCProcedures(params buildMirrorYARPCProceduresParams) []transport.Procedure {
	handler := &_MirrorYARPCHandler{params.Server}
	return protobuf.BuildProcedures(
		protobuf.BuildProceduresParams{
			ServiceName: "mirror.Mirror",
			UnaryHandlerParams: []protobuf.BuildProceduresUnaryHandlerParams{
				{
					MethodName: "Mirror",
					Handler: protobuf.NewUnaryHandler(
						protobuf.UnaryHandlerParams{
							Handle:      handler.Mirror,
							NewRequest:  newMirrorServiceMirrorYARPCRequest,
							AnyResolver: params.AnyResolver,
						},
					),
				},
			},
			OnewayHandlerParams: []protobuf.BuildProceduresOnewayHandlerParams{},
			StreamHandlerParams: []protobuf.BuildProceduresStreamHandlerParams{},
		},
	)
}

// BuildMirrorYARPCProcedures prepares an implementation of the Mirror service for YARPC registration.
func BuildMirrorYARPCProcedures(server MirrorYARPCServer) []transport.Procedure {
	return buildMirrorYARPCProcedures(buildMirrorYARPCProceduresParams{Server: server})
}

// FxMirrorYARPCClientParams defines the input
// for NewFxMirrorYARPCClient. It provides the
// paramaters to get a MirrorYARPCClient in an
// Fx application.
type FxMirrorYARPCClientParams struct {
	fx.In

	Provider    yarpc.ClientConfig
	AnyResolver jsonpb.AnyResolver  `name:"yarpcfx" optional:"true"`
	Restriction restriction.Checker `optional:"true"`
}

// FxMirrorYARPCClientResult defines the output
// of NewFxMirrorYARPCClient. It provides a
// MirrorYARPCClient to an Fx application.
type FxMirrorYARPCClientResult struct {
	fx.Out

	Client MirrorYARPCClient

	// We are using an fx.Out struct here instead of just returning a client
	// so that we can add more values or add named versions of the client in
	// the future without breaking any existing code.
}

// NewFxMirrorYARPCClient provides a MirrorYARPCClient
// to an Fx application using the given name for routing.
//
//  fx.Provide(
//    mirror.NewFxMirrorYARPCClient("service-name"),
//    ...
//  )
func NewFxMirrorYARPCClient(name string, options ...protobuf.ClientOption) interface{} {
	return func(params FxMirrorYARPCClientParams) FxMirrorYARPCClientResult {
		cc := params.Provider.ClientConfig(name)

		if params.Restriction != nil {
			if namer, ok := cc.GetUnaryOutbound().(transport.Namer); ok {
				if err := params.Restriction.Check(protobuf.Encoding, namer.TransportName()); err != nil {
					panic(err.Error())
				}
			}
		}

		return FxMirrorYARPCClientResult{
			Client: newMirrorYARPCClient(cc, params.AnyResolver, options...),
		}
	}
}

// FxMirrorYARPCProceduresParams defines the input
// for NewFxMirrorYARPCProcedures. It provides the
// paramaters to get MirrorYARPCServer procedures in an
// Fx application.
type FxMirrorYARPCProceduresParams struct {
	fx.In

	Server      MirrorYARPCServer
	AnyResolver jsonpb.AnyResolver `name:"yarpcfx" optional:"true"`
}

// FxMirrorYARPCProceduresResult defines the output
// of NewFxMirrorYARPCProcedures. It provides
// MirrorYARPCServer procedures to an Fx application.
//
// The procedures are provided to the "yarpcfx" value group.
// Dig 1.2 or newer must be used for this feature to work.
type FxMirrorYARPCProceduresResult struct {
	fx.Out

	Procedures     []transport.Procedure `group:"yarpcfx"`
	ReflectionMeta reflection.ServerMeta `group:"yarpcfx"`
}

// NewFxMirrorYARPCProcedures provides MirrorYARPCServer procedures to an Fx application.
// It expects a MirrorYARPCServer to be present in the container.
//
//  fx.Provide(
//    mirror.NewFxMirrorYARPCProcedures(),
//    ...
//  )
func NewFxMirrorYARPCProcedures() interface{} {
	return func(params FxMirrorYARPCProceduresParams) FxMirrorYARPCProceduresResult {
		return FxMirrorYARPCProceduresResult{
			Procedures: buildMirrorYARPCProcedures(buildMirrorYARPCProceduresParams{
				Server:      params.Server,
				AnyResolver: params.AnyResolver,
			}),
			ReflectionMeta: reflection.ServerMeta{
				ServiceName:     "mirror.Mirror",
				FileDescriptors: yarpcFileDescriptorClosure735477c7170897b9,
			},
		}
	}
}

type _MirrorYARPCCaller struct {
	streamClient protobuf.StreamClient
}

func (c *_MirrorYARPCCaller) Mirror(ctx context.Context, request *Request, options ...yarpc.CallOption) (*Response, error) {
	responseMessage, err := c.streamClient.Call(ctx, "Mirror", request, newMirrorServiceMirrorYARPCResponse, options...)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*Response)
	if !ok {
		return nil, protobuf.CastError(emptyMirrorServiceMirrorYARPCResponse, responseMessage)
	}
	return response, err
}

type _MirrorYARPCHandler struct {
	server MirrorYARPCServer
}

func (h *_MirrorYARPCHandler) Mirror(ctx context.Context, requestMessage proto.Message) (proto.Message, error) {
	var request *Request
	var ok bool
	if requestMessage != nil {
		request, ok = requestMessage.(*Request)
		if !ok {
			return nil, protobuf.CastError(emptyMirrorServiceMirrorYARPCRequest, requestMessage)
		}
	}
	response, err := h.server.Mirror(ctx, request)
	if response == nil {
		return nil, err
	}
	return response, err
}

func newMirrorServiceMirrorYARPCRequest() proto.Message {
	return &Request{}
}

func newMirrorServiceMirrorYARPCResponse() proto.Message {
	return &Response{}
}

var (
	emptyMirrorServiceMirrorYARPCRequest  = &Request{}
	emptyMirrorServiceMirrorYARPCResponse = &Response{}
)

// MirrorInternalYARPCClient is the YARPC client-side interface for the MirrorInternal service.
type MirrorInternalYARPCClient interface {
	Mirror(context.Context, *InternalRequest, ...yarpc.CallOption) (*InternalResponse, error)
}

func newMirrorInternalYARPCClient(clientConfig transport.ClientConfig, anyResolver jsonpb.AnyResolver, options ...protobuf.ClientOption) MirrorInternalYARPCClient {
	return &_MirrorInternalYARPCCaller{protobuf.NewStreamClient(
		protobuf.ClientParams{
			ServiceName:  "mirror.MirrorInternal",
			ClientConfig: clientConfig,
			AnyResolver:  anyResolver,
			Options:      options,
		},
	)}
}

// NewMirrorInternalYARPCClient builds a new YARPC client for the MirrorInternal service.
func NewMirrorInternalYARPCClient(clientConfig transport.ClientConfig, options ...protobuf.ClientOption) MirrorInternalYARPCClient {
	return newMirrorInternalYARPCClient(clientConfig, nil, options...)
}

// MirrorInternalYARPCServer is the YARPC server-side interface for the MirrorInternal service.
type MirrorInternalYARPCServer interface {
	Mirror(context.Context, *InternalRequest) (*InternalResponse, error)
}

type buildMirrorInternalYARPCProceduresParams struct {
	Server      MirrorInternalYARPCServer
	AnyResolver jsonpb.AnyResolver
}

func buildMirrorInternalYARPCProcedures(params buildMirrorInternalYARPCProceduresParams) []transport.Procedure {
	handler := &_MirrorInternalYARPCHandler{params.Server}
	return protobuf.BuildProcedures(
		protobuf.BuildProceduresParams{
			ServiceName: "mirror.MirrorInternal",
			UnaryHandlerParams: []protobuf.BuildProceduresUnaryHandlerParams{
				{
					MethodName: "Mirror",
					Handler: protobuf.NewUnaryHandler(
						protobuf.UnaryHandlerParams{
							Handle:      handler.Mirror,
							NewRequest:  newMirrorInternalServiceMirrorYARPCRequest,
							AnyResolver: params.AnyResolver,
						},
					),
				},
			},
			OnewayHandlerParams: []protobuf.BuildProceduresOnewayHandlerParams{},
			StreamHandlerParams: []protobuf.BuildProceduresStreamHandlerParams{},
		},
	)
}

// BuildMirrorInternalYARPCProcedures prepares an implementation of the MirrorInternal service for YARPC registration.
func BuildMirrorInternalYARPCProcedures(server MirrorInternalYARPCServer) []transport.Procedure {
	return buildMirrorInternalYARPCProcedures(buildMirrorInternalYARPCProceduresParams{Server: server})
}

// FxMirrorInternalYARPCClientParams defines the input
// for NewFxMirrorInternalYARPCClient. It provides the
// paramaters to get a MirrorInternalYARPCClient in an
// Fx application.
type FxMirrorInternalYARPCClientParams struct {
	fx.In

	Provider    yarpc.ClientConfig
	AnyResolver jsonpb.AnyResolver  `name:"yarpcfx" optional:"true"`
	Restriction restriction.Checker `optional:"true"`
}

// FxMirrorInternalYARPCClientResult defines the output
// of NewFxMirrorInternalYARPCClient. It provides a
// MirrorInternalYARPCClient to an Fx application.
type FxMirrorInternalYARPCClientResult struct {
	fx.Out

	Client MirrorInternalYARPCClient

	// We are using an fx.Out struct here instead of just returning a client
	// so that we can add more values or add named versions of the client in
	// the future without breaking any existing code.
}

// NewFxMirrorInternalYARPCClient provides a MirrorInternalYARPCClient
// to an Fx application using the given name for routing.
//
//  fx.Provide(
//    mirror.NewFxMirrorInternalYARPCClient("service-name"),
//    ...
//  )
func NewFxMirrorInternalYARPCClient(name string, options ...protobuf.ClientOption) interface{} {
	return func(params FxMirrorInternalYARPCClientParams) FxMirrorInternalYARPCClientResult {
		cc := params.Provider.ClientConfig(name)

		if params.Restriction != nil {
			if namer, ok := cc.GetUnaryOutbound().(transport.Namer); ok {
				if err := params.Restriction.Check(protobuf.Encoding, namer.TransportName()); err != nil {
					panic(err.Error())
				}
			}
		}

		return FxMirrorInternalYARPCClientResult{
			Client: newMirrorInternalYARPCClient(cc, params.AnyResolver, options...),
		}
	}
}

// FxMirrorInternalYARPCProceduresParams defines the input
// for NewFxMirrorInternalYARPCProcedures. It provides the
// paramaters to get MirrorInternalYARPCServer procedures in an
// Fx application.
type FxMirrorInternalYARPCProceduresParams struct {
	fx.In

	Server      MirrorInternalYARPCServer
	AnyResolver jsonpb.AnyResolver `name:"yarpcfx" optional:"true"`
}

// FxMirrorInternalYARPCProceduresResult defines the output
// of NewFxMirrorInternalYARPCProcedures. It provides
// MirrorInternalYARPCServer procedures to an Fx application.
//
// The procedures are provided to the "yarpcfx" value group.
// Dig 1.2 or newer must be used for this feature to work.
type FxMirrorInternalYARPCProceduresResult struct {
	fx.Out

	Procedures     []transport.Procedure `group:"yarpcfx"`
	ReflectionMeta reflection.ServerMeta `group:"yarpcfx"`
}

// NewFxMirrorInternalYARPCProcedures provides MirrorInternalYARPCServer procedures to an Fx application.
// It expects a MirrorInternalYARPCServer to be present in the container.
//
//  fx.Provide(
//    mirror.NewFxMirrorInternalYARPCProcedures(),
//    ...
//  )
func NewFxMirrorInternalYARPCProcedures() interface{} {
	return func(params FxMirrorInternalYARPCProceduresParams) FxMirrorInternalYARPCProceduresResult {
		return FxMirrorInternalYARPCProceduresResult{
			Procedures: buildMirrorInternalYARPCProcedures(buildMirrorInternalYARPCProceduresParams{
				Server:      params.Server,
				AnyResolver: params.AnyResolver,
			}),
			ReflectionMeta: reflection.ServerMeta{
				ServiceName:     "mirror.MirrorInternal",
				FileDescriptors: yarpcFileDescriptorClosure735477c7170897b9,
			},
		}
	}
}

type _MirrorInternalYARPCCaller struct {
	streamClient protobuf.StreamClient
}

func (c *_MirrorInternalYARPCCaller) Mirror(ctx context.Context, request *InternalRequest, options ...yarpc.CallOption) (*InternalResponse, error) {
	responseMessage, err := c.streamClient.Call(ctx, "Mirror", request, newMirrorInternalServiceMirrorYARPCResponse, options...)
	if responseMessage == nil {
		return nil, err
	}
	response, ok := responseMessage.(*InternalResponse)
	if !ok {
		return nil, protobuf.CastError(emptyMirrorInternalServiceMirrorYARPCResponse, responseMessage)
	}
	return response, err
}

type _MirrorInternalYARPCHandler struct {
	server MirrorInternalYARPCServer
}

func (h *_MirrorInternalYARPCHandler) Mirror(ctx context.Context, requestMessage proto.Message) (proto.Message, error) {
	var request *InternalRequest
	var ok bool
	if requestMessage != nil {
		request, ok = requestMessage.(*InternalRequest)
		if !ok {
			return nil, protobuf.CastError(emptyMirrorInternalServiceMirrorYARPCRequest, requestMessage)
		}
	}
	response, err := h.server.Mirror(ctx, request)
	if response == nil {
		return nil, err
	}
	return response, err
}

func newMirrorInternalServiceMirrorYARPCRequest() proto.Message {
	return &InternalRequest{}
}

func newMirrorInternalServiceMirrorYARPCResponse() proto.Message {
	return &InternalResponse{}
}

var (
	emptyMirrorInternalServiceMirrorYARPCRequest  = &InternalRequest{}
	emptyMirrorInternalServiceMirrorYARPCResponse = &InternalResponse{}
)

var yarpcFileDescriptorClosure735477c7170897b9 = [][]byte{
	// clients/mirror/mirror.proto
	[]byte{
		0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4e, 0xce, 0xc9, 0x4c,
		0xcd, 0x2b, 0x29, 0xd6, 0xcf, 0xcd, 0x2c, 0x2a, 0xca, 0x2f, 0x82, 0x52, 0x7a, 0x05, 0x45, 0xf9,
		0x25, 0xf9, 0x42, 0x6c, 0x10, 0x9e, 0x92, 0x32, 0x17, 0x7b, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71,
		0x89, 0x90, 0x04, 0x17, 0x7b, 0x6e, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0xaa, 0x04, 0xa3, 0x02, 0xa3,
		0x06, 0x67, 0x10, 0x8c, 0xab, 0xa4, 0xc2, 0xc5, 0x11, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c,
		0x8a, 0x47, 0x95, 0x36, 0x17, 0xbf, 0x67, 0x5e, 0x49, 0x6a, 0x51, 0x5e, 0x62, 0x0e, 0x61, 0x23,
		0x75, 0xb8, 0x04, 0x10, 0x8a, 0x09, 0x19, 0x6d, 0x64, 0xca, 0xc5, 0xe6, 0x0b, 0x76, 0xaf, 0x90,
		0x36, 0x9c, 0xc5, 0xaf, 0x07, 0xf5, 0x10, 0xd4, 0x32, 0x29, 0x01, 0x84, 0x00, 0xc4, 0x40, 0x23,
		0x5f, 0x2e, 0x3e, 0x88, 0x62, 0x98, 0x55, 0x42, 0xd6, 0x70, 0xed, 0xe2, 0x30, 0xd5, 0x68, 0x6e,
		0x96, 0x92, 0xc0, 0x94, 0x80, 0x18, 0x97, 0xc4, 0x06, 0x0e, 0x3a, 0x63, 0x40, 0x00, 0x00, 0x00,
		0xff, 0xff, 0xfa, 0x4b, 0x9b, 0xc3, 0x59, 0x01, 0x00, 0x00,
	},
}

func init() {
	yarpc.RegisterClientBuilder(
		func(clientConfig transport.ClientConfig, structField reflect.StructField) MirrorYARPCClient {
			return NewMirrorYARPCClient(clientConfig, protobuf.ClientBuilderOptions(clientConfig, structField)...)
		},
	)
	yarpc.RegisterClientBuilder(
		func(clientConfig transport.ClientConfig, structField reflect.StructField) MirrorInternalYARPCClient {
			return NewMirrorInternalYARPCClient(clientConfig, protobuf.ClientBuilderOptions(clientConfig, structField)...)
		},
	)
}
