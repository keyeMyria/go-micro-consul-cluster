// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/hotel.proto

/*
Package hotel is a generated protocol buffer package.

It is generated from these files:
	proto/hotel.proto

It has these top-level messages:
	Reqeust
	Response
*/
package hotel

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go-micro-consul-cluster/srv/profile/proto"
import _ "go-micro-consul-cluster/srv/rate/proto"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Hotel service

type HotelService interface {
	Rates(ctx context.Context, in *Reqeust, opts ...client.CallOption) (*Response, error)
}

type hotelService struct {
	c    client.Client
	name string
}

func NewHotelService(name string, c client.Client) HotelService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "hotel"
	}
	return &hotelService{
		c:    c,
		name: name,
	}
}

func (c *hotelService) Rates(ctx context.Context, in *Reqeust, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Hotel.Rates", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Hotel service

type HotelHandler interface {
	Rates(context.Context, *Reqeust, *Response) error
}

func RegisterHotelHandler(s server.Server, hdlr HotelHandler, opts ...server.HandlerOption) {
	type hotel interface {
		Rates(ctx context.Context, in *Reqeust, out *Response) error
	}
	type Hotel struct {
		hotel
	}
	h := &hotelHandler{hdlr}
	s.Handle(s.NewHandler(&Hotel{h}, opts...))
}

type hotelHandler struct {
	HotelHandler
}

func (h *hotelHandler) Rates(ctx context.Context, in *Reqeust, out *Response) error {
	return h.HotelHandler.Rates(ctx, in, out)
}
