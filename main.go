package main

import (
	"context"
	"errors"
	"github.com/charmbracelet/log"
	v1beta1 "github.com/kevinmichaelchen/buf-vanguard-rest/internal/idl/bvr/v1beta1"
	"net"
	"net/http"
	"time"

	connectGo "connectrpc.com/connect"
	"connectrpc.com/grpcreflect"
	"connectrpc.com/vanguard"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	"github.com/kevinmichaelchen/buf-vanguard-rest/internal"
	connectPB "github.com/kevinmichaelchen/buf-vanguard-rest/internal/idl/bvr/v1beta1/bvrv1beta1connect"
)

type fooService struct{}

func (s fooService) CreateFoo(
	_ context.Context,
	req *connectGo.Request[v1beta1.CreateFooRequest],
) (*connectGo.Response[v1beta1.CreateFooResponse], error) {
	log.Info("Creating new Foo...", "name", req.Msg.GetName())

	return connectGo.NewResponse(
		&v1beta1.CreateFooResponse{
			Foo: &v1beta1.Foo{
				Id:   "1",
				Name: req.Msg.GetName(),
			},
		},
	), nil
}

func (s fooService) GetFoo(
	_ context.Context,
	req *connectGo.Request[v1beta1.GetFooRequest],
) (*connectGo.Response[v1beta1.GetFooResponse], error) {
	log.Info("Retrieving Foo...", "id", req.Msg.GetId())

	return connectGo.NewResponse(
		&v1beta1.GetFooResponse{
			Foo: &v1beta1.Foo{
				Id:   "1",
				Name: "Bar",
			},
		},
	), nil
}

func (s fooService) ListFoos(
	context.Context,
	*connectGo.Request[v1beta1.ListFoosRequest],
) (*connectGo.Response[v1beta1.ListFoosResponse], error) {
	return connectGo.NewResponse(
		&v1beta1.ListFoosResponse{
			Foos: []*v1beta1.Foo{
				{
					Id:   "1",
					Name: "Bar",
				},
			},
		},
	), nil
}

func (s fooService) DeleteFoo(
	context.Context,
	*connectGo.Request[v1beta1.DeleteFooRequest],
) (*connectGo.Response[v1beta1.DeleteFooResponse], error) {
	return connectGo.NewResponse(
		&v1beta1.DeleteFooResponse{},
	), nil
}

func main() {
	path, connectHandler := connectPB.NewFooServiceHandler(&fooService{})

	// Wrap handler with Vanguard,
	// so it can accept Connect, gRPC, or gRPC-Web.
	vanguardHandler, err := vanguard.NewTranscoder(
		[]*vanguard.Service{
			vanguard.NewService(
				connectPB.FooServiceName,
				connectHandler,
			),
		},
	)
	if err != nil {
		log.Fatal(err)
	}

	serveMux := http.NewServeMux()

	// Our Connect handler
	serveMux.Handle(path, connectHandler)

	// Similar to above, we trace incoming requests to stdout. That way you can see the
	// original RPC request and the proxied REST request to see Vanguard in action.
	serveMux.Handle("/", internal.TraceHandler(vanguardHandler))

	// We add gRPC reflection support so that you can use tools like `buf curl` or `grpcurl`
	// with this server.
	serveMux.Handle(grpcreflect.NewHandlerV1(grpcreflect.NewStaticReflector(connectPB.FooServiceName)))

	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}

	svr := &http.Server{
		Addr: ":http",
		// We use h2c to support HTTP/2 without TLS (and thus support the gRPC protocol).
		Handler:           h2c.NewHandler(serveMux, &http2.Server{}),
		ReadHeaderTimeout: 15 * time.Second,
	}

	err = svr.Serve(listener)
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal(err)
	}
}
