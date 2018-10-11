// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: protos/example/generated.proto

/*
Package example is a generated protocol buffer package.

It is generated from these files:
	protos/example/generated.proto

It has these top-level messages:
	Category
	MyDuration
	MyTime
	Price
	Product
	GetAlphaTimeRequest
	GetDurationForLengthRequest
	GetDurationForLengthCtxRequest
	GetOmegaTimeRequest
	GetPhoneRequest
	RandomCategoryRequest
	RandomNumberRequest
	RandomNumberResponse
*/
package example

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
//import _ "google/protobuf"
import example_categories "github.com/kennyzhu/go-os/dbservice/tools/example/protos/example/categories"
//import _ "google/protobuf"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = example_categories.CategoryOptions{}

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ExampleService service

type ExampleService interface {
	GetAlphaTime(ctx context.Context, in *GetAlphaTimeRequest, opts ...client.CallOption) (*MyTime, error)
	GetDurationForLength(ctx context.Context, in *GetDurationForLengthRequest, opts ...client.CallOption) (*MyDuration, error)
	GetDurationForLengthCtx(ctx context.Context, in *GetDurationForLengthCtxRequest, opts ...client.CallOption) (*MyDuration, error)
	GetOmegaTime(ctx context.Context, in *GetOmegaTimeRequest, opts ...client.CallOption) (*MyTime, error)
	GetPhone(ctx context.Context, in *GetPhoneRequest, opts ...client.CallOption) (*Product, error)
	RandomCategory(ctx context.Context, in *RandomCategoryRequest, opts ...client.CallOption) (*example_categories.CategoryOptions, error)
	RandomNumber(ctx context.Context, in *RandomNumberRequest, opts ...client.CallOption) (*RandomNumberResponse, error)
}

type exampleService struct {
	c    client.Client
	name string
}

func NewExampleService(name string, c client.Client) ExampleService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "example"
	}
	return &exampleService{
		c:    c,
		name: name,
	}
}

func (c *exampleService) GetAlphaTime(ctx context.Context, in *GetAlphaTimeRequest, opts ...client.CallOption) (*MyTime, error) {
	req := c.c.NewRequest(c.name, "ExampleService.GetAlphaTime", in)
	out := new(MyTime)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleService) GetDurationForLength(ctx context.Context, in *GetDurationForLengthRequest, opts ...client.CallOption) (*MyDuration, error) {
	req := c.c.NewRequest(c.name, "ExampleService.GetDurationForLength", in)
	out := new(MyDuration)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleService) GetDurationForLengthCtx(ctx context.Context, in *GetDurationForLengthCtxRequest, opts ...client.CallOption) (*MyDuration, error) {
	req := c.c.NewRequest(c.name, "ExampleService.GetDurationForLengthCtx", in)
	out := new(MyDuration)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleService) GetOmegaTime(ctx context.Context, in *GetOmegaTimeRequest, opts ...client.CallOption) (*MyTime, error) {
	req := c.c.NewRequest(c.name, "ExampleService.GetOmegaTime", in)
	out := new(MyTime)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleService) GetPhone(ctx context.Context, in *GetPhoneRequest, opts ...client.CallOption) (*Product, error) {
	req := c.c.NewRequest(c.name, "ExampleService.GetPhone", in)
	out := new(Product)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleService) RandomCategory(ctx context.Context, in *RandomCategoryRequest, opts ...client.CallOption) (*example_categories.CategoryOptions, error) {
	req := c.c.NewRequest(c.name, "ExampleService.RandomCategory", in)
	out := new(example_categories.CategoryOptions)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exampleService) RandomNumber(ctx context.Context, in *RandomNumberRequest, opts ...client.CallOption) (*RandomNumberResponse, error) {
	req := c.c.NewRequest(c.name, "ExampleService.RandomNumber", in)
	out := new(RandomNumberResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ExampleService service

type ExampleServiceHandler interface {
	GetAlphaTime(context.Context, *GetAlphaTimeRequest, *MyTime) error
	GetDurationForLength(context.Context, *GetDurationForLengthRequest, *MyDuration) error
	GetDurationForLengthCtx(context.Context, *GetDurationForLengthCtxRequest, *MyDuration) error
	GetOmegaTime(context.Context, *GetOmegaTimeRequest, *MyTime) error
	GetPhone(context.Context, *GetPhoneRequest, *Product) error
	RandomCategory(context.Context, *RandomCategoryRequest, *example_categories.CategoryOptions) error
	RandomNumber(context.Context, *RandomNumberRequest, *RandomNumberResponse) error
}

func RegisterExampleServiceHandler(s server.Server, hdlr ExampleServiceHandler, opts ...server.HandlerOption) {
	type exampleService interface {
		GetAlphaTime(ctx context.Context, in *GetAlphaTimeRequest, out *MyTime) error
		GetDurationForLength(ctx context.Context, in *GetDurationForLengthRequest, out *MyDuration) error
		GetDurationForLengthCtx(ctx context.Context, in *GetDurationForLengthCtxRequest, out *MyDuration) error
		GetOmegaTime(ctx context.Context, in *GetOmegaTimeRequest, out *MyTime) error
		GetPhone(ctx context.Context, in *GetPhoneRequest, out *Product) error
		RandomCategory(ctx context.Context, in *RandomCategoryRequest, out *example_categories.CategoryOptions) error
		RandomNumber(ctx context.Context, in *RandomNumberRequest, out *RandomNumberResponse) error
	}
	type ExampleService struct {
		exampleService
	}
	h := &exampleServiceHandler{hdlr}
	s.Handle(s.NewHandler(&ExampleService{h}, opts...))
}

type exampleServiceHandler struct {
	ExampleServiceHandler
}

func (h *exampleServiceHandler) GetAlphaTime(ctx context.Context, in *GetAlphaTimeRequest, out *MyTime) error {
	return h.ExampleServiceHandler.GetAlphaTime(ctx, in, out)
}

func (h *exampleServiceHandler) GetDurationForLength(ctx context.Context, in *GetDurationForLengthRequest, out *MyDuration) error {
	return h.ExampleServiceHandler.GetDurationForLength(ctx, in, out)
}

func (h *exampleServiceHandler) GetDurationForLengthCtx(ctx context.Context, in *GetDurationForLengthCtxRequest, out *MyDuration) error {
	return h.ExampleServiceHandler.GetDurationForLengthCtx(ctx, in, out)
}

func (h *exampleServiceHandler) GetOmegaTime(ctx context.Context, in *GetOmegaTimeRequest, out *MyTime) error {
	return h.ExampleServiceHandler.GetOmegaTime(ctx, in, out)
}

func (h *exampleServiceHandler) GetPhone(ctx context.Context, in *GetPhoneRequest, out *Product) error {
	return h.ExampleServiceHandler.GetPhone(ctx, in, out)
}

func (h *exampleServiceHandler) RandomCategory(ctx context.Context, in *RandomCategoryRequest, out *example_categories.CategoryOptions) error {
	return h.ExampleServiceHandler.RandomCategory(ctx, in, out)
}

func (h *exampleServiceHandler) RandomNumber(ctx context.Context, in *RandomNumberRequest, out *RandomNumberResponse) error {
	return h.ExampleServiceHandler.RandomNumber(ctx, in, out)
}
