// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: payinchannel.proto

package payinchannel

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for PayinChannel service

func NewPayinChannelEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for PayinChannel service

type PayinChannelService interface {
	QueryToExecute(ctx context.Context, in *PayinQueryToExecuteRequest, opts ...client.CallOption) (*PayinQueryToExecuteResponse, error)
	UpdatePayinStatus(ctx context.Context, in *UpdateStatusRequest, opts ...client.CallOption) (*BaseResponse, error)
}

type payinChannelService struct {
	c    client.Client
	name string
}

func NewPayinChannelService(name string, c client.Client) PayinChannelService {
	return &payinChannelService{
		c:    c,
		name: name,
	}
}

func (c *payinChannelService) QueryToExecute(ctx context.Context, in *PayinQueryToExecuteRequest, opts ...client.CallOption) (*PayinQueryToExecuteResponse, error) {
	req := c.c.NewRequest(c.name, "PayinChannel.QueryToExecute", in)
	out := new(PayinQueryToExecuteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payinChannelService) UpdatePayinStatus(ctx context.Context, in *UpdateStatusRequest, opts ...client.CallOption) (*BaseResponse, error) {
	req := c.c.NewRequest(c.name, "PayinChannel.UpdatePayinStatus", in)
	out := new(BaseResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PayinChannel service

type PayinChannelHandler interface {
	QueryToExecute(context.Context, *PayinQueryToExecuteRequest, *PayinQueryToExecuteResponse) error
	UpdatePayinStatus(context.Context, *UpdateStatusRequest, *BaseResponse) error
}

func RegisterPayinChannelHandler(s server.Server, hdlr PayinChannelHandler, opts ...server.HandlerOption) error {
	type payinChannel interface {
		QueryToExecute(ctx context.Context, in *PayinQueryToExecuteRequest, out *PayinQueryToExecuteResponse) error
		UpdatePayinStatus(ctx context.Context, in *UpdateStatusRequest, out *BaseResponse) error
	}
	type PayinChannel struct {
		payinChannel
	}
	h := &payinChannelHandler{hdlr}
	return s.Handle(s.NewHandler(&PayinChannel{h}, opts...))
}

type payinChannelHandler struct {
	PayinChannelHandler
}

func (h *payinChannelHandler) QueryToExecute(ctx context.Context, in *PayinQueryToExecuteRequest, out *PayinQueryToExecuteResponse) error {
	return h.PayinChannelHandler.QueryToExecute(ctx, in, out)
}

func (h *payinChannelHandler) UpdatePayinStatus(ctx context.Context, in *UpdateStatusRequest, out *BaseResponse) error {
	return h.PayinChannelHandler.UpdatePayinStatus(ctx, in, out)
}
