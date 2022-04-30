// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/postAvatar.proto

package postAvatar

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

// Api Endpoints for PostAvatar service

func NewPostAvatarEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for PostAvatar service

type PostAvatarService interface {
	Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error)
	ClientStream(ctx context.Context, opts ...client.CallOption) (PostAvatar_ClientStreamService, error)
	ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...client.CallOption) (PostAvatar_ServerStreamService, error)
	BidiStream(ctx context.Context, opts ...client.CallOption) (PostAvatar_BidiStreamService, error)
}

type postAvatarService struct {
	c    client.Client
	name string
}

func NewPostAvatarService(name string, c client.Client) PostAvatarService {
	return &postAvatarService{
		c:    c,
		name: name,
	}
}

func (c *postAvatarService) Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error) {
	req := c.c.NewRequest(c.name, "PostAvatar.Call", in)
	out := new(CallResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *postAvatarService) ClientStream(ctx context.Context, opts ...client.CallOption) (PostAvatar_ClientStreamService, error) {
	req := c.c.NewRequest(c.name, "PostAvatar.ClientStream", &ClientStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &postAvatarServiceClientStream{stream}, nil
}

type PostAvatar_ClientStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Send(*ClientStreamRequest) error
}

type postAvatarServiceClientStream struct {
	stream client.Stream
}

func (x *postAvatarServiceClientStream) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *postAvatarServiceClientStream) Close() error {
	return x.stream.Close()
}

func (x *postAvatarServiceClientStream) Context() context.Context {
	return x.stream.Context()
}

func (x *postAvatarServiceClientStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *postAvatarServiceClientStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *postAvatarServiceClientStream) Send(m *ClientStreamRequest) error {
	return x.stream.Send(m)
}

func (c *postAvatarService) ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...client.CallOption) (PostAvatar_ServerStreamService, error) {
	req := c.c.NewRequest(c.name, "PostAvatar.ServerStream", &ServerStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &postAvatarServiceServerStream{stream}, nil
}

type PostAvatar_ServerStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Recv() (*ServerStreamResponse, error)
}

type postAvatarServiceServerStream struct {
	stream client.Stream
}

func (x *postAvatarServiceServerStream) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *postAvatarServiceServerStream) Close() error {
	return x.stream.Close()
}

func (x *postAvatarServiceServerStream) Context() context.Context {
	return x.stream.Context()
}

func (x *postAvatarServiceServerStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *postAvatarServiceServerStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *postAvatarServiceServerStream) Recv() (*ServerStreamResponse, error) {
	m := new(ServerStreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *postAvatarService) BidiStream(ctx context.Context, opts ...client.CallOption) (PostAvatar_BidiStreamService, error) {
	req := c.c.NewRequest(c.name, "PostAvatar.BidiStream", &BidiStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &postAvatarServiceBidiStream{stream}, nil
}

type PostAvatar_BidiStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Send(*BidiStreamRequest) error
	Recv() (*BidiStreamResponse, error)
}

type postAvatarServiceBidiStream struct {
	stream client.Stream
}

func (x *postAvatarServiceBidiStream) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *postAvatarServiceBidiStream) Close() error {
	return x.stream.Close()
}

func (x *postAvatarServiceBidiStream) Context() context.Context {
	return x.stream.Context()
}

func (x *postAvatarServiceBidiStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *postAvatarServiceBidiStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *postAvatarServiceBidiStream) Send(m *BidiStreamRequest) error {
	return x.stream.Send(m)
}

func (x *postAvatarServiceBidiStream) Recv() (*BidiStreamResponse, error) {
	m := new(BidiStreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for PostAvatar service

type PostAvatarHandler interface {
	Call(context.Context, *CallRequest, *CallResponse) error
	ClientStream(context.Context, PostAvatar_ClientStreamStream) error
	ServerStream(context.Context, *ServerStreamRequest, PostAvatar_ServerStreamStream) error
	BidiStream(context.Context, PostAvatar_BidiStreamStream) error
}

func RegisterPostAvatarHandler(s server.Server, hdlr PostAvatarHandler, opts ...server.HandlerOption) error {
	type postAvatar interface {
		Call(ctx context.Context, in *CallRequest, out *CallResponse) error
		ClientStream(ctx context.Context, stream server.Stream) error
		ServerStream(ctx context.Context, stream server.Stream) error
		BidiStream(ctx context.Context, stream server.Stream) error
	}
	type PostAvatar struct {
		postAvatar
	}
	h := &postAvatarHandler{hdlr}
	return s.Handle(s.NewHandler(&PostAvatar{h}, opts...))
}

type postAvatarHandler struct {
	PostAvatarHandler
}

func (h *postAvatarHandler) Call(ctx context.Context, in *CallRequest, out *CallResponse) error {
	return h.PostAvatarHandler.Call(ctx, in, out)
}

func (h *postAvatarHandler) ClientStream(ctx context.Context, stream server.Stream) error {
	return h.PostAvatarHandler.ClientStream(ctx, &postAvatarClientStreamStream{stream})
}

type PostAvatar_ClientStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*ClientStreamRequest, error)
}

type postAvatarClientStreamStream struct {
	stream server.Stream
}

func (x *postAvatarClientStreamStream) Close() error {
	return x.stream.Close()
}

func (x *postAvatarClientStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *postAvatarClientStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *postAvatarClientStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *postAvatarClientStreamStream) Recv() (*ClientStreamRequest, error) {
	m := new(ClientStreamRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *postAvatarHandler) ServerStream(ctx context.Context, stream server.Stream) error {
	m := new(ServerStreamRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.PostAvatarHandler.ServerStream(ctx, m, &postAvatarServerStreamStream{stream})
}

type PostAvatar_ServerStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ServerStreamResponse) error
}

type postAvatarServerStreamStream struct {
	stream server.Stream
}

func (x *postAvatarServerStreamStream) Close() error {
	return x.stream.Close()
}

func (x *postAvatarServerStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *postAvatarServerStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *postAvatarServerStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *postAvatarServerStreamStream) Send(m *ServerStreamResponse) error {
	return x.stream.Send(m)
}

func (h *postAvatarHandler) BidiStream(ctx context.Context, stream server.Stream) error {
	return h.PostAvatarHandler.BidiStream(ctx, &postAvatarBidiStreamStream{stream})
}

type PostAvatar_BidiStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*BidiStreamResponse) error
	Recv() (*BidiStreamRequest, error)
}

type postAvatarBidiStreamStream struct {
	stream server.Stream
}

func (x *postAvatarBidiStreamStream) Close() error {
	return x.stream.Close()
}

func (x *postAvatarBidiStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *postAvatarBidiStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *postAvatarBidiStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *postAvatarBidiStreamStream) Send(m *BidiStreamResponse) error {
	return x.stream.Send(m)
}

func (x *postAvatarBidiStreamStream) Recv() (*BidiStreamRequest, error) {
	m := new(BidiStreamRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
