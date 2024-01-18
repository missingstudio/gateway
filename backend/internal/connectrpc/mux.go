package connectrpc

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"connectrpc.com/connect"
	"connectrpc.com/grpchealth"
	"connectrpc.com/grpcreflect"
	"connectrpc.com/validate"
	greetingv1 "github.com/missingstudio/protos/pkg/greeting/v1"
	"github.com/missingstudio/protos/pkg/greeting/v1/greetingv1connect"
)

type Deps struct{}

func NewConnectMux(d Deps) (*http.ServeMux, error) {
	mux := http.NewServeMux()

	validateInterceptor, err := validate.NewInterceptor()
	if err != nil {
		return nil, fmt.Errorf("validate interceptor not created: %w", err)
	}

	compress1KB := connect.WithCompressMinBytes(1024)
	mux.Handle(greetingv1connect.NewGreetServiceHandler(
		&GreetServer{},
		compress1KB,
		connect.WithInterceptors(validateInterceptor),
	))

	mux.Handle(grpchealth.NewHandler(
		grpchealth.NewStaticChecker(greetingv1connect.GreetServiceName),
		compress1KB,
	))

	reflector := grpcreflect.NewStaticReflector(greetingv1connect.GreetServiceName)
	mux.Handle(grpcreflect.NewHandlerV1(reflector, compress1KB))
	mux.Handle(grpcreflect.NewHandlerV1Alpha(reflector, compress1KB))

	return mux, nil
}

type GreetServer struct{}

func (s *GreetServer) Greet(
	ctx context.Context,
	req *connect.Request[greetingv1.GreetRequest],
) (*connect.Response[greetingv1.GreetResponse], error) {
	log.Println("Request headers: ", req.Header())

	res := connect.NewResponse(&greetingv1.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})
	return res, nil
}
