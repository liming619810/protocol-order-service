// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: merchant.proto

package merchant

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

// Api Endpoints for Merchant service

func NewMerchantEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Merchant service

type MerchantService interface {
	// 新增 收单
	MerchantBalance(ctx context.Context, in *MerchantBalanceRequest, opts ...client.CallOption) (*MerchantBalanceResponse, error)
}

type merchantService struct {
	c    client.Client
	name string
}

func NewMerchantService(name string, c client.Client) MerchantService {
	return &merchantService{
		c:    c,
		name: name,
	}
}

func (c *merchantService) MerchantBalance(ctx context.Context, in *MerchantBalanceRequest, opts ...client.CallOption) (*MerchantBalanceResponse, error) {
	req := c.c.NewRequest(c.name, "Merchant.MerchantBalance", in)
	out := new(MerchantBalanceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Merchant service

type MerchantHandler interface {
	// 新增 收单
	MerchantBalance(context.Context, *MerchantBalanceRequest, *MerchantBalanceResponse) error
}

func RegisterMerchantHandler(s server.Server, hdlr MerchantHandler, opts ...server.HandlerOption) error {
	type merchant interface {
		MerchantBalance(ctx context.Context, in *MerchantBalanceRequest, out *MerchantBalanceResponse) error
	}
	type Merchant struct {
		merchant
	}
	h := &merchantHandler{hdlr}
	return s.Handle(s.NewHandler(&Merchant{h}, opts...))
}

type merchantHandler struct {
	MerchantHandler
}

func (h *merchantHandler) MerchantBalance(ctx context.Context, in *MerchantBalanceRequest, out *MerchantBalanceResponse) error {
	return h.MerchantHandler.MerchantBalance(ctx, in, out)
}
