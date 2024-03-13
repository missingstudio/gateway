// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: llm/v1/service.proto

package llmv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/missingstudio/ai/protos/pkg/llm/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// LLMServiceName is the fully-qualified name of the LLMService service.
	LLMServiceName = "llm.v1.LLMService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// LLMServiceChatCompletionsProcedure is the fully-qualified name of the LLMService's
	// ChatCompletions RPC.
	LLMServiceChatCompletionsProcedure = "/llm.v1.LLMService/ChatCompletions"
	// LLMServiceStreamChatCompletionsProcedure is the fully-qualified name of the LLMService's
	// StreamChatCompletions RPC.
	LLMServiceStreamChatCompletionsProcedure = "/llm.v1.LLMService/StreamChatCompletions"
	// LLMServiceListModelsProcedure is the fully-qualified name of the LLMService's ListModels RPC.
	LLMServiceListModelsProcedure = "/llm.v1.LLMService/ListModels"
	// LLMServiceListProvidersProcedure is the fully-qualified name of the LLMService's ListProviders
	// RPC.
	LLMServiceListProvidersProcedure = "/llm.v1.LLMService/ListProviders"
	// LLMServiceCreateProviderProcedure is the fully-qualified name of the LLMService's CreateProvider
	// RPC.
	LLMServiceCreateProviderProcedure = "/llm.v1.LLMService/CreateProvider"
	// LLMServiceGetProviderProcedure is the fully-qualified name of the LLMService's GetProvider RPC.
	LLMServiceGetProviderProcedure = "/llm.v1.LLMService/GetProvider"
	// LLMServiceUpsertProviderProcedure is the fully-qualified name of the LLMService's UpsertProvider
	// RPC.
	LLMServiceUpsertProviderProcedure = "/llm.v1.LLMService/UpsertProvider"
	// LLMServiceDeleteProviderProcedure is the fully-qualified name of the LLMService's DeleteProvider
	// RPC.
	LLMServiceDeleteProviderProcedure = "/llm.v1.LLMService/DeleteProvider"
	// LLMServiceGetProviderConfigProcedure is the fully-qualified name of the LLMService's
	// GetProviderConfig RPC.
	LLMServiceGetProviderConfigProcedure = "/llm.v1.LLMService/GetProviderConfig"
	// LLMServiceListAPIKeysProcedure is the fully-qualified name of the LLMService's ListAPIKeys RPC.
	LLMServiceListAPIKeysProcedure = "/llm.v1.LLMService/ListAPIKeys"
	// LLMServiceCreateAPIKeyProcedure is the fully-qualified name of the LLMService's CreateAPIKey RPC.
	LLMServiceCreateAPIKeyProcedure = "/llm.v1.LLMService/CreateAPIKey"
	// LLMServiceGetAPIKeyProcedure is the fully-qualified name of the LLMService's GetAPIKey RPC.
	LLMServiceGetAPIKeyProcedure = "/llm.v1.LLMService/GetAPIKey"
	// LLMServiceUpdateAPIKeyProcedure is the fully-qualified name of the LLMService's UpdateAPIKey RPC.
	LLMServiceUpdateAPIKeyProcedure = "/llm.v1.LLMService/UpdateAPIKey"
	// LLMServiceDeleteAPIKeyProcedure is the fully-qualified name of the LLMService's DeleteAPIKey RPC.
	LLMServiceDeleteAPIKeyProcedure = "/llm.v1.LLMService/DeleteAPIKey"
	// LLMServiceListTrackingLogsProcedure is the fully-qualified name of the LLMService's
	// ListTrackingLogs RPC.
	LLMServiceListTrackingLogsProcedure = "/llm.v1.LLMService/ListTrackingLogs"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	lLMServiceServiceDescriptor                     = v1.File_llm_v1_service_proto.Services().ByName("LLMService")
	lLMServiceChatCompletionsMethodDescriptor       = lLMServiceServiceDescriptor.Methods().ByName("ChatCompletions")
	lLMServiceStreamChatCompletionsMethodDescriptor = lLMServiceServiceDescriptor.Methods().ByName("StreamChatCompletions")
	lLMServiceListModelsMethodDescriptor            = lLMServiceServiceDescriptor.Methods().ByName("ListModels")
	lLMServiceListProvidersMethodDescriptor         = lLMServiceServiceDescriptor.Methods().ByName("ListProviders")
	lLMServiceCreateProviderMethodDescriptor        = lLMServiceServiceDescriptor.Methods().ByName("CreateProvider")
	lLMServiceGetProviderMethodDescriptor           = lLMServiceServiceDescriptor.Methods().ByName("GetProvider")
	lLMServiceUpsertProviderMethodDescriptor        = lLMServiceServiceDescriptor.Methods().ByName("UpsertProvider")
	lLMServiceDeleteProviderMethodDescriptor        = lLMServiceServiceDescriptor.Methods().ByName("DeleteProvider")
	lLMServiceGetProviderConfigMethodDescriptor     = lLMServiceServiceDescriptor.Methods().ByName("GetProviderConfig")
	lLMServiceListAPIKeysMethodDescriptor           = lLMServiceServiceDescriptor.Methods().ByName("ListAPIKeys")
	lLMServiceCreateAPIKeyMethodDescriptor          = lLMServiceServiceDescriptor.Methods().ByName("CreateAPIKey")
	lLMServiceGetAPIKeyMethodDescriptor             = lLMServiceServiceDescriptor.Methods().ByName("GetAPIKey")
	lLMServiceUpdateAPIKeyMethodDescriptor          = lLMServiceServiceDescriptor.Methods().ByName("UpdateAPIKey")
	lLMServiceDeleteAPIKeyMethodDescriptor          = lLMServiceServiceDescriptor.Methods().ByName("DeleteAPIKey")
	lLMServiceListTrackingLogsMethodDescriptor      = lLMServiceServiceDescriptor.Methods().ByName("ListTrackingLogs")
)

// LLMServiceClient is a client for the llm.v1.LLMService service.
type LLMServiceClient interface {
	ChatCompletions(context.Context, *connect.Request[v1.ChatCompletionRequest]) (*connect.Response[v1.ChatCompletionResponse], error)
	StreamChatCompletions(context.Context, *connect.Request[v1.ChatCompletionRequest]) (*connect.ServerStreamForClient[v1.ChatCompletionResponse], error)
	ListModels(context.Context, *connect.Request[v1.ModelRequest]) (*connect.Response[v1.ModelResponse], error)
	ListProviders(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[v1.ListProvidersResponse], error)
	CreateProvider(context.Context, *connect.Request[v1.CreateProviderRequest]) (*connect.Response[v1.CreateProviderResponse], error)
	GetProvider(context.Context, *connect.Request[v1.GetProviderRequest]) (*connect.Response[v1.GetProviderResponse], error)
	UpsertProvider(context.Context, *connect.Request[v1.UpdateProviderRequest]) (*connect.Response[v1.UpdateProviderResponse], error)
	DeleteProvider(context.Context, *connect.Request[v1.DeleteProviderRequest]) (*connect.Response[emptypb.Empty], error)
	GetProviderConfig(context.Context, *connect.Request[v1.GetProviderConfigRequest]) (*connect.Response[v1.GetProviderConfigResponse], error)
	ListAPIKeys(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[v1.ListAPIKeysResponse], error)
	CreateAPIKey(context.Context, *connect.Request[v1.CreateAPIKeyRequest]) (*connect.Response[v1.CreateAPIKeyResponse], error)
	GetAPIKey(context.Context, *connect.Request[v1.GetAPIKeyRequest]) (*connect.Response[v1.GetAPIKeyResponse], error)
	UpdateAPIKey(context.Context, *connect.Request[v1.UpdateAPIKeyRequest]) (*connect.Response[v1.UpdateAPIKeyResponse], error)
	DeleteAPIKey(context.Context, *connect.Request[v1.DeleteAPIKeyRequest]) (*connect.Response[emptypb.Empty], error)
	ListTrackingLogs(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[v1.LogResponse], error)
}

// NewLLMServiceClient constructs a client for the llm.v1.LLMService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewLLMServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) LLMServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &lLMServiceClient{
		chatCompletions: connect.NewClient[v1.ChatCompletionRequest, v1.ChatCompletionResponse](
			httpClient,
			baseURL+LLMServiceChatCompletionsProcedure,
			connect.WithSchema(lLMServiceChatCompletionsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		streamChatCompletions: connect.NewClient[v1.ChatCompletionRequest, v1.ChatCompletionResponse](
			httpClient,
			baseURL+LLMServiceStreamChatCompletionsProcedure,
			connect.WithSchema(lLMServiceStreamChatCompletionsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listModels: connect.NewClient[v1.ModelRequest, v1.ModelResponse](
			httpClient,
			baseURL+LLMServiceListModelsProcedure,
			connect.WithSchema(lLMServiceListModelsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listProviders: connect.NewClient[emptypb.Empty, v1.ListProvidersResponse](
			httpClient,
			baseURL+LLMServiceListProvidersProcedure,
			connect.WithSchema(lLMServiceListProvidersMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createProvider: connect.NewClient[v1.CreateProviderRequest, v1.CreateProviderResponse](
			httpClient,
			baseURL+LLMServiceCreateProviderProcedure,
			connect.WithSchema(lLMServiceCreateProviderMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getProvider: connect.NewClient[v1.GetProviderRequest, v1.GetProviderResponse](
			httpClient,
			baseURL+LLMServiceGetProviderProcedure,
			connect.WithSchema(lLMServiceGetProviderMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		upsertProvider: connect.NewClient[v1.UpdateProviderRequest, v1.UpdateProviderResponse](
			httpClient,
			baseURL+LLMServiceUpsertProviderProcedure,
			connect.WithSchema(lLMServiceUpsertProviderMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteProvider: connect.NewClient[v1.DeleteProviderRequest, emptypb.Empty](
			httpClient,
			baseURL+LLMServiceDeleteProviderProcedure,
			connect.WithSchema(lLMServiceDeleteProviderMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getProviderConfig: connect.NewClient[v1.GetProviderConfigRequest, v1.GetProviderConfigResponse](
			httpClient,
			baseURL+LLMServiceGetProviderConfigProcedure,
			connect.WithSchema(lLMServiceGetProviderConfigMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listAPIKeys: connect.NewClient[emptypb.Empty, v1.ListAPIKeysResponse](
			httpClient,
			baseURL+LLMServiceListAPIKeysProcedure,
			connect.WithSchema(lLMServiceListAPIKeysMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createAPIKey: connect.NewClient[v1.CreateAPIKeyRequest, v1.CreateAPIKeyResponse](
			httpClient,
			baseURL+LLMServiceCreateAPIKeyProcedure,
			connect.WithSchema(lLMServiceCreateAPIKeyMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getAPIKey: connect.NewClient[v1.GetAPIKeyRequest, v1.GetAPIKeyResponse](
			httpClient,
			baseURL+LLMServiceGetAPIKeyProcedure,
			connect.WithSchema(lLMServiceGetAPIKeyMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateAPIKey: connect.NewClient[v1.UpdateAPIKeyRequest, v1.UpdateAPIKeyResponse](
			httpClient,
			baseURL+LLMServiceUpdateAPIKeyProcedure,
			connect.WithSchema(lLMServiceUpdateAPIKeyMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteAPIKey: connect.NewClient[v1.DeleteAPIKeyRequest, emptypb.Empty](
			httpClient,
			baseURL+LLMServiceDeleteAPIKeyProcedure,
			connect.WithSchema(lLMServiceDeleteAPIKeyMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		listTrackingLogs: connect.NewClient[emptypb.Empty, v1.LogResponse](
			httpClient,
			baseURL+LLMServiceListTrackingLogsProcedure,
			connect.WithSchema(lLMServiceListTrackingLogsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// lLMServiceClient implements LLMServiceClient.
type lLMServiceClient struct {
	chatCompletions       *connect.Client[v1.ChatCompletionRequest, v1.ChatCompletionResponse]
	streamChatCompletions *connect.Client[v1.ChatCompletionRequest, v1.ChatCompletionResponse]
	listModels            *connect.Client[v1.ModelRequest, v1.ModelResponse]
	listProviders         *connect.Client[emptypb.Empty, v1.ListProvidersResponse]
	createProvider        *connect.Client[v1.CreateProviderRequest, v1.CreateProviderResponse]
	getProvider           *connect.Client[v1.GetProviderRequest, v1.GetProviderResponse]
	upsertProvider        *connect.Client[v1.UpdateProviderRequest, v1.UpdateProviderResponse]
	deleteProvider        *connect.Client[v1.DeleteProviderRequest, emptypb.Empty]
	getProviderConfig     *connect.Client[v1.GetProviderConfigRequest, v1.GetProviderConfigResponse]
	listAPIKeys           *connect.Client[emptypb.Empty, v1.ListAPIKeysResponse]
	createAPIKey          *connect.Client[v1.CreateAPIKeyRequest, v1.CreateAPIKeyResponse]
	getAPIKey             *connect.Client[v1.GetAPIKeyRequest, v1.GetAPIKeyResponse]
	updateAPIKey          *connect.Client[v1.UpdateAPIKeyRequest, v1.UpdateAPIKeyResponse]
	deleteAPIKey          *connect.Client[v1.DeleteAPIKeyRequest, emptypb.Empty]
	listTrackingLogs      *connect.Client[emptypb.Empty, v1.LogResponse]
}

// ChatCompletions calls llm.v1.LLMService.ChatCompletions.
func (c *lLMServiceClient) ChatCompletions(ctx context.Context, req *connect.Request[v1.ChatCompletionRequest]) (*connect.Response[v1.ChatCompletionResponse], error) {
	return c.chatCompletions.CallUnary(ctx, req)
}

// StreamChatCompletions calls llm.v1.LLMService.StreamChatCompletions.
func (c *lLMServiceClient) StreamChatCompletions(ctx context.Context, req *connect.Request[v1.ChatCompletionRequest]) (*connect.ServerStreamForClient[v1.ChatCompletionResponse], error) {
	return c.streamChatCompletions.CallServerStream(ctx, req)
}

// ListModels calls llm.v1.LLMService.ListModels.
func (c *lLMServiceClient) ListModels(ctx context.Context, req *connect.Request[v1.ModelRequest]) (*connect.Response[v1.ModelResponse], error) {
	return c.listModels.CallUnary(ctx, req)
}

// ListProviders calls llm.v1.LLMService.ListProviders.
func (c *lLMServiceClient) ListProviders(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[v1.ListProvidersResponse], error) {
	return c.listProviders.CallUnary(ctx, req)
}

// CreateProvider calls llm.v1.LLMService.CreateProvider.
func (c *lLMServiceClient) CreateProvider(ctx context.Context, req *connect.Request[v1.CreateProviderRequest]) (*connect.Response[v1.CreateProviderResponse], error) {
	return c.createProvider.CallUnary(ctx, req)
}

// GetProvider calls llm.v1.LLMService.GetProvider.
func (c *lLMServiceClient) GetProvider(ctx context.Context, req *connect.Request[v1.GetProviderRequest]) (*connect.Response[v1.GetProviderResponse], error) {
	return c.getProvider.CallUnary(ctx, req)
}

// UpsertProvider calls llm.v1.LLMService.UpsertProvider.
func (c *lLMServiceClient) UpsertProvider(ctx context.Context, req *connect.Request[v1.UpdateProviderRequest]) (*connect.Response[v1.UpdateProviderResponse], error) {
	return c.upsertProvider.CallUnary(ctx, req)
}

// DeleteProvider calls llm.v1.LLMService.DeleteProvider.
func (c *lLMServiceClient) DeleteProvider(ctx context.Context, req *connect.Request[v1.DeleteProviderRequest]) (*connect.Response[emptypb.Empty], error) {
	return c.deleteProvider.CallUnary(ctx, req)
}

// GetProviderConfig calls llm.v1.LLMService.GetProviderConfig.
func (c *lLMServiceClient) GetProviderConfig(ctx context.Context, req *connect.Request[v1.GetProviderConfigRequest]) (*connect.Response[v1.GetProviderConfigResponse], error) {
	return c.getProviderConfig.CallUnary(ctx, req)
}

// ListAPIKeys calls llm.v1.LLMService.ListAPIKeys.
func (c *lLMServiceClient) ListAPIKeys(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[v1.ListAPIKeysResponse], error) {
	return c.listAPIKeys.CallUnary(ctx, req)
}

// CreateAPIKey calls llm.v1.LLMService.CreateAPIKey.
func (c *lLMServiceClient) CreateAPIKey(ctx context.Context, req *connect.Request[v1.CreateAPIKeyRequest]) (*connect.Response[v1.CreateAPIKeyResponse], error) {
	return c.createAPIKey.CallUnary(ctx, req)
}

// GetAPIKey calls llm.v1.LLMService.GetAPIKey.
func (c *lLMServiceClient) GetAPIKey(ctx context.Context, req *connect.Request[v1.GetAPIKeyRequest]) (*connect.Response[v1.GetAPIKeyResponse], error) {
	return c.getAPIKey.CallUnary(ctx, req)
}

// UpdateAPIKey calls llm.v1.LLMService.UpdateAPIKey.
func (c *lLMServiceClient) UpdateAPIKey(ctx context.Context, req *connect.Request[v1.UpdateAPIKeyRequest]) (*connect.Response[v1.UpdateAPIKeyResponse], error) {
	return c.updateAPIKey.CallUnary(ctx, req)
}

// DeleteAPIKey calls llm.v1.LLMService.DeleteAPIKey.
func (c *lLMServiceClient) DeleteAPIKey(ctx context.Context, req *connect.Request[v1.DeleteAPIKeyRequest]) (*connect.Response[emptypb.Empty], error) {
	return c.deleteAPIKey.CallUnary(ctx, req)
}

// ListTrackingLogs calls llm.v1.LLMService.ListTrackingLogs.
func (c *lLMServiceClient) ListTrackingLogs(ctx context.Context, req *connect.Request[emptypb.Empty]) (*connect.Response[v1.LogResponse], error) {
	return c.listTrackingLogs.CallUnary(ctx, req)
}

// LLMServiceHandler is an implementation of the llm.v1.LLMService service.
type LLMServiceHandler interface {
	ChatCompletions(context.Context, *connect.Request[v1.ChatCompletionRequest]) (*connect.Response[v1.ChatCompletionResponse], error)
	StreamChatCompletions(context.Context, *connect.Request[v1.ChatCompletionRequest], *connect.ServerStream[v1.ChatCompletionResponse]) error
	ListModels(context.Context, *connect.Request[v1.ModelRequest]) (*connect.Response[v1.ModelResponse], error)
	ListProviders(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[v1.ListProvidersResponse], error)
	CreateProvider(context.Context, *connect.Request[v1.CreateProviderRequest]) (*connect.Response[v1.CreateProviderResponse], error)
	GetProvider(context.Context, *connect.Request[v1.GetProviderRequest]) (*connect.Response[v1.GetProviderResponse], error)
	UpsertProvider(context.Context, *connect.Request[v1.UpdateProviderRequest]) (*connect.Response[v1.UpdateProviderResponse], error)
	DeleteProvider(context.Context, *connect.Request[v1.DeleteProviderRequest]) (*connect.Response[emptypb.Empty], error)
	GetProviderConfig(context.Context, *connect.Request[v1.GetProviderConfigRequest]) (*connect.Response[v1.GetProviderConfigResponse], error)
	ListAPIKeys(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[v1.ListAPIKeysResponse], error)
	CreateAPIKey(context.Context, *connect.Request[v1.CreateAPIKeyRequest]) (*connect.Response[v1.CreateAPIKeyResponse], error)
	GetAPIKey(context.Context, *connect.Request[v1.GetAPIKeyRequest]) (*connect.Response[v1.GetAPIKeyResponse], error)
	UpdateAPIKey(context.Context, *connect.Request[v1.UpdateAPIKeyRequest]) (*connect.Response[v1.UpdateAPIKeyResponse], error)
	DeleteAPIKey(context.Context, *connect.Request[v1.DeleteAPIKeyRequest]) (*connect.Response[emptypb.Empty], error)
	ListTrackingLogs(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[v1.LogResponse], error)
}

// NewLLMServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewLLMServiceHandler(svc LLMServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	lLMServiceChatCompletionsHandler := connect.NewUnaryHandler(
		LLMServiceChatCompletionsProcedure,
		svc.ChatCompletions,
		connect.WithSchema(lLMServiceChatCompletionsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	lLMServiceStreamChatCompletionsHandler := connect.NewServerStreamHandler(
		LLMServiceStreamChatCompletionsProcedure,
		svc.StreamChatCompletions,
		connect.WithSchema(lLMServiceStreamChatCompletionsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	lLMServiceListModelsHandler := connect.NewUnaryHandler(
		LLMServiceListModelsProcedure,
		svc.ListModels,
		connect.WithSchema(lLMServiceListModelsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	lLMServiceListProvidersHandler := connect.NewUnaryHandler(
		LLMServiceListProvidersProcedure,
		svc.ListProviders,
		connect.WithSchema(lLMServiceListProvidersMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	lLMServiceCreateProviderHandler := connect.NewUnaryHandler(
		LLMServiceCreateProviderProcedure,
		svc.CreateProvider,
		connect.WithSchema(lLMServiceCreateProviderMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	lLMServiceGetProviderHandler := connect.NewUnaryHandler(
		LLMServiceGetProviderProcedure,
		svc.GetProvider,
		connect.WithSchema(lLMServiceGetProviderMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	lLMServiceUpsertProviderHandler := connect.NewUnaryHandler(
		LLMServiceUpsertProviderProcedure,
		svc.UpsertProvider,
		connect.WithSchema(lLMServiceUpsertProviderMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	lLMServiceDeleteProviderHandler := connect.NewUnaryHandler(
		LLMServiceDeleteProviderProcedure,
		svc.DeleteProvider,
		connect.WithSchema(lLMServiceDeleteProviderMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	lLMServiceGetProviderConfigHandler := connect.NewUnaryHandler(
		LLMServiceGetProviderConfigProcedure,
		svc.GetProviderConfig,
		connect.WithSchema(lLMServiceGetProviderConfigMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	lLMServiceListAPIKeysHandler := connect.NewUnaryHandler(
		LLMServiceListAPIKeysProcedure,
		svc.ListAPIKeys,
		connect.WithSchema(lLMServiceListAPIKeysMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	lLMServiceCreateAPIKeyHandler := connect.NewUnaryHandler(
		LLMServiceCreateAPIKeyProcedure,
		svc.CreateAPIKey,
		connect.WithSchema(lLMServiceCreateAPIKeyMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	lLMServiceGetAPIKeyHandler := connect.NewUnaryHandler(
		LLMServiceGetAPIKeyProcedure,
		svc.GetAPIKey,
		connect.WithSchema(lLMServiceGetAPIKeyMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	lLMServiceUpdateAPIKeyHandler := connect.NewUnaryHandler(
		LLMServiceUpdateAPIKeyProcedure,
		svc.UpdateAPIKey,
		connect.WithSchema(lLMServiceUpdateAPIKeyMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	lLMServiceDeleteAPIKeyHandler := connect.NewUnaryHandler(
		LLMServiceDeleteAPIKeyProcedure,
		svc.DeleteAPIKey,
		connect.WithSchema(lLMServiceDeleteAPIKeyMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	lLMServiceListTrackingLogsHandler := connect.NewUnaryHandler(
		LLMServiceListTrackingLogsProcedure,
		svc.ListTrackingLogs,
		connect.WithSchema(lLMServiceListTrackingLogsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/llm.v1.LLMService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case LLMServiceChatCompletionsProcedure:
			lLMServiceChatCompletionsHandler.ServeHTTP(w, r)
		case LLMServiceStreamChatCompletionsProcedure:
			lLMServiceStreamChatCompletionsHandler.ServeHTTP(w, r)
		case LLMServiceListModelsProcedure:
			lLMServiceListModelsHandler.ServeHTTP(w, r)
		case LLMServiceListProvidersProcedure:
			lLMServiceListProvidersHandler.ServeHTTP(w, r)
		case LLMServiceCreateProviderProcedure:
			lLMServiceCreateProviderHandler.ServeHTTP(w, r)
		case LLMServiceGetProviderProcedure:
			lLMServiceGetProviderHandler.ServeHTTP(w, r)
		case LLMServiceUpsertProviderProcedure:
			lLMServiceUpsertProviderHandler.ServeHTTP(w, r)
		case LLMServiceDeleteProviderProcedure:
			lLMServiceDeleteProviderHandler.ServeHTTP(w, r)
		case LLMServiceGetProviderConfigProcedure:
			lLMServiceGetProviderConfigHandler.ServeHTTP(w, r)
		case LLMServiceListAPIKeysProcedure:
			lLMServiceListAPIKeysHandler.ServeHTTP(w, r)
		case LLMServiceCreateAPIKeyProcedure:
			lLMServiceCreateAPIKeyHandler.ServeHTTP(w, r)
		case LLMServiceGetAPIKeyProcedure:
			lLMServiceGetAPIKeyHandler.ServeHTTP(w, r)
		case LLMServiceUpdateAPIKeyProcedure:
			lLMServiceUpdateAPIKeyHandler.ServeHTTP(w, r)
		case LLMServiceDeleteAPIKeyProcedure:
			lLMServiceDeleteAPIKeyHandler.ServeHTTP(w, r)
		case LLMServiceListTrackingLogsProcedure:
			lLMServiceListTrackingLogsHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedLLMServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedLLMServiceHandler struct{}

func (UnimplementedLLMServiceHandler) ChatCompletions(context.Context, *connect.Request[v1.ChatCompletionRequest]) (*connect.Response[v1.ChatCompletionResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("llm.v1.LLMService.ChatCompletions is not implemented"))
}

func (UnimplementedLLMServiceHandler) StreamChatCompletions(context.Context, *connect.Request[v1.ChatCompletionRequest], *connect.ServerStream[v1.ChatCompletionResponse]) error {
	return connect.NewError(connect.CodeUnimplemented, errors.New("llm.v1.LLMService.StreamChatCompletions is not implemented"))
}

func (UnimplementedLLMServiceHandler) ListModels(context.Context, *connect.Request[v1.ModelRequest]) (*connect.Response[v1.ModelResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("llm.v1.LLMService.ListModels is not implemented"))
}

func (UnimplementedLLMServiceHandler) ListProviders(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[v1.ListProvidersResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("llm.v1.LLMService.ListProviders is not implemented"))
}

func (UnimplementedLLMServiceHandler) CreateProvider(context.Context, *connect.Request[v1.CreateProviderRequest]) (*connect.Response[v1.CreateProviderResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("llm.v1.LLMService.CreateProvider is not implemented"))
}

func (UnimplementedLLMServiceHandler) GetProvider(context.Context, *connect.Request[v1.GetProviderRequest]) (*connect.Response[v1.GetProviderResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("llm.v1.LLMService.GetProvider is not implemented"))
}

func (UnimplementedLLMServiceHandler) UpsertProvider(context.Context, *connect.Request[v1.UpdateProviderRequest]) (*connect.Response[v1.UpdateProviderResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("llm.v1.LLMService.UpsertProvider is not implemented"))
}

func (UnimplementedLLMServiceHandler) DeleteProvider(context.Context, *connect.Request[v1.DeleteProviderRequest]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("llm.v1.LLMService.DeleteProvider is not implemented"))
}

func (UnimplementedLLMServiceHandler) GetProviderConfig(context.Context, *connect.Request[v1.GetProviderConfigRequest]) (*connect.Response[v1.GetProviderConfigResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("llm.v1.LLMService.GetProviderConfig is not implemented"))
}

func (UnimplementedLLMServiceHandler) ListAPIKeys(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[v1.ListAPIKeysResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("llm.v1.LLMService.ListAPIKeys is not implemented"))
}

func (UnimplementedLLMServiceHandler) CreateAPIKey(context.Context, *connect.Request[v1.CreateAPIKeyRequest]) (*connect.Response[v1.CreateAPIKeyResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("llm.v1.LLMService.CreateAPIKey is not implemented"))
}

func (UnimplementedLLMServiceHandler) GetAPIKey(context.Context, *connect.Request[v1.GetAPIKeyRequest]) (*connect.Response[v1.GetAPIKeyResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("llm.v1.LLMService.GetAPIKey is not implemented"))
}

func (UnimplementedLLMServiceHandler) UpdateAPIKey(context.Context, *connect.Request[v1.UpdateAPIKeyRequest]) (*connect.Response[v1.UpdateAPIKeyResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("llm.v1.LLMService.UpdateAPIKey is not implemented"))
}

func (UnimplementedLLMServiceHandler) DeleteAPIKey(context.Context, *connect.Request[v1.DeleteAPIKeyRequest]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("llm.v1.LLMService.DeleteAPIKey is not implemented"))
}

func (UnimplementedLLMServiceHandler) ListTrackingLogs(context.Context, *connect.Request[emptypb.Empty]) (*connect.Response[v1.LogResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("llm.v1.LLMService.ListTrackingLogs is not implemented"))
}