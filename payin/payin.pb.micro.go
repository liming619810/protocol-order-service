// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: payin.proto

package payin

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

// Api Endpoints for Payin service

func NewPayinEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Payin service

type PayinService interface {
	// 新增 收单
	OrderAndPay(ctx context.Context, in *PayinRequest, opts ...client.CallOption) (*PayinResponse, error)
	OrderQuery(ctx context.Context, in *PayinQueryRequest, opts ...client.CallOption) (*PayinQueryResponse, error)
	PayinDetail(ctx context.Context, in *PayinDetailRequest, opts ...client.CallOption) (*PayinDetailResponse, error)
	OrderQueryPage(ctx context.Context, in *PayinQueryPageRequest, opts ...client.CallOption) (*PayinQueryPageResponse, error)
}

type payinService struct {
	c    client.Client
	name string
}

func NewPayinService(name string, c client.Client) PayinService {
	return &payinService{
		c:    c,
		name: name,
	}
}

func (c *payinService) OrderAndPay(ctx context.Context, in *PayinRequest, opts ...client.CallOption) (*PayinResponse, error) {
	req := c.c.NewRequest(c.name, "Payin.OrderAndPay", in)
	out := new(PayinResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payinService) OrderQuery(ctx context.Context, in *PayinQueryRequest, opts ...client.CallOption) (*PayinQueryResponse, error) {
	req := c.c.NewRequest(c.name, "Payin.OrderQuery", in)
	out := new(PayinQueryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payinService) PayinDetail(ctx context.Context, in *PayinDetailRequest, opts ...client.CallOption) (*PayinDetailResponse, error) {
	req := c.c.NewRequest(c.name, "Payin.PayinDetail", in)
	out := new(PayinDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payinService) OrderQueryPage(ctx context.Context, in *PayinQueryPageRequest, opts ...client.CallOption) (*PayinQueryPageResponse, error) {
	req := c.c.NewRequest(c.name, "Payin.OrderQueryPage", in)
	out := new(PayinQueryPageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Payin service

type PayinHandler interface {
	// 新增 收单
	OrderAndPay(context.Context, *PayinRequest, *PayinResponse) error
	OrderQuery(context.Context, *PayinQueryRequest, *PayinQueryResponse) error
	PayinDetail(context.Context, *PayinDetailRequest, *PayinDetailResponse) error
	OrderQueryPage(context.Context, *PayinQueryPageRequest, *PayinQueryPageResponse) error
}

func RegisterPayinHandler(s server.Server, hdlr PayinHandler, opts ...server.HandlerOption) error {
	type payin interface {
		OrderAndPay(ctx context.Context, in *PayinRequest, out *PayinResponse) error
		OrderQuery(ctx context.Context, in *PayinQueryRequest, out *PayinQueryResponse) error
		PayinDetail(ctx context.Context, in *PayinDetailRequest, out *PayinDetailResponse) error
		OrderQueryPage(ctx context.Context, in *PayinQueryPageRequest, out *PayinQueryPageResponse) error
	}
	type Payin struct {
		payin
	}
	h := &payinHandler{hdlr}
	return s.Handle(s.NewHandler(&Payin{h}, opts...))
}

type payinHandler struct {
	PayinHandler
}

func (h *payinHandler) OrderAndPay(ctx context.Context, in *PayinRequest, out *PayinResponse) error {
	return h.PayinHandler.OrderAndPay(ctx, in, out)
}

func (h *payinHandler) OrderQuery(ctx context.Context, in *PayinQueryRequest, out *PayinQueryResponse) error {
	return h.PayinHandler.OrderQuery(ctx, in, out)
}

func (h *payinHandler) PayinDetail(ctx context.Context, in *PayinDetailRequest, out *PayinDetailResponse) error {
	return h.PayinHandler.PayinDetail(ctx, in, out)
}

func (h *payinHandler) OrderQueryPage(ctx context.Context, in *PayinQueryPageRequest, out *PayinQueryPageResponse) error {
	return h.PayinHandler.OrderQueryPage(ctx, in, out)
}
