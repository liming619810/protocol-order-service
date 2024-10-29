// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: statement.proto

package statement

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

// Api Endpoints for Statement service

func NewStatementEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Statement service

type StatementService interface {
	// 新增 收单
	PageQuery(ctx context.Context, in *StatementPageRequest, opts ...client.CallOption) (*StatementPageResponse, error)
}

type statementService struct {
	c    client.Client
	name string
}

func NewStatementService(name string, c client.Client) StatementService {
	return &statementService{
		c:    c,
		name: name,
	}
}

func (c *statementService) PageQuery(ctx context.Context, in *StatementPageRequest, opts ...client.CallOption) (*StatementPageResponse, error) {
	req := c.c.NewRequest(c.name, "Statement.PageQuery", in)
	out := new(StatementPageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Statement service

type StatementHandler interface {
	// 新增 收单
	PageQuery(context.Context, *StatementPageRequest, *StatementPageResponse) error
}

func RegisterStatementHandler(s server.Server, hdlr StatementHandler, opts ...server.HandlerOption) error {
	type statement interface {
		PageQuery(ctx context.Context, in *StatementPageRequest, out *StatementPageResponse) error
	}
	type Statement struct {
		statement
	}
	h := &statementHandler{hdlr}
	return s.Handle(s.NewHandler(&Statement{h}, opts...))
}

type statementHandler struct {
	StatementHandler
}

func (h *statementHandler) PageQuery(ctx context.Context, in *StatementPageRequest, out *StatementPageResponse) error {
	return h.StatementHandler.PageQuery(ctx, in, out)
}
