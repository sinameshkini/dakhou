package grpc

import (
	"context"
	endpoint "dakhou/todo/pkg/endpoint"
	pb "dakhou/todo/pkg/grpc/pb"
	"errors"
	grpc "github.com/go-kit/kit/transport/grpc"
	context1 "golang.org/x/net/context"
)

// makeGetHandler creates the handler logic
func makeGetHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetEndpoint, decodeGetRequest, encodeGetResponse, options...)
}

// decodeGetResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Get request.
// TODO implement the decoder
func decodeGetRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Todo' Decoder is not impelemented")
}

// encodeGetResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Todo' Encoder is not impelemented")
}
func (g *grpcServer) Get(ctx context1.Context, req *pb.GetRequest) (*pb.GetReply, error) {
	_, rep, err := g.get.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetReply), nil
}

// makeAddHandler creates the handler logic
func makeAddHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.AddEndpoint, decodeAddRequest, encodeAddResponse, options...)
}

// decodeAddResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Add request.
// TODO implement the decoder
func decodeAddRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Todo' Decoder is not impelemented")
}

// encodeAddResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeAddResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Todo' Encoder is not impelemented")
}
func (g *grpcServer) Add(ctx context1.Context, req *pb.AddRequest) (*pb.AddReply, error) {
	_, rep, err := g.add.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AddReply), nil
}

// makeSetCompleteHandler creates the handler logic
func makeSetCompleteHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.SetCompleteEndpoint, decodeSetCompleteRequest, encodeSetCompleteResponse, options...)
}

// decodeSetCompleteResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain SetComplete request.
// TODO implement the decoder
func decodeSetCompleteRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Todo' Decoder is not impelemented")
}

// encodeSetCompleteResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeSetCompleteResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Todo' Encoder is not impelemented")
}
func (g *grpcServer) SetComplete(ctx context1.Context, req *pb.SetCompleteRequest) (*pb.SetCompleteReply, error) {
	_, rep, err := g.setComplete.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SetCompleteReply), nil
}

// makeRemoveCompleteHandler creates the handler logic
func makeRemoveCompleteHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.RemoveCompleteEndpoint, decodeRemoveCompleteRequest, encodeRemoveCompleteResponse, options...)
}

// decodeRemoveCompleteResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain RemoveComplete request.
// TODO implement the decoder
func decodeRemoveCompleteRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Todo' Decoder is not impelemented")
}

// encodeRemoveCompleteResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeRemoveCompleteResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Todo' Encoder is not impelemented")
}
func (g *grpcServer) RemoveComplete(ctx context1.Context, req *pb.RemoveCompleteRequest) (*pb.RemoveCompleteReply, error) {
	_, rep, err := g.removeComplete.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.RemoveCompleteReply), nil
}

// makeDeleteHandler creates the handler logic
func makeDeleteHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.DeleteEndpoint, decodeDeleteRequest, encodeDeleteResponse, options...)
}

// decodeDeleteResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Delete request.
// TODO implement the decoder
func decodeDeleteRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Todo' Decoder is not impelemented")
}

// encodeDeleteResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeDeleteResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Todo' Encoder is not impelemented")
}
func (g *grpcServer) Delete(ctx context1.Context, req *pb.DeleteRequest) (*pb.DeleteReply, error) {
	_, rep, err := g.delete.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteReply), nil
}

// makeUpdateHandler creates the handler logic
func makeUpdateHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UpdateEndpoint, decodeUpdateRequest, encodeUpdateResponse, options...)
}

// decodeUpdateResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain Update request.
// TODO implement the decoder
func decodeUpdateRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Todo' Decoder is not impelemented")
}

// encodeUpdateResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeUpdateResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Todo' Encoder is not impelemented")
}
func (g *grpcServer) Update(ctx context1.Context, req *pb.UpdateRequest) (*pb.UpdateReply, error) {
	_, rep, err := g.update.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdateReply), nil
}

// makeSetStarHandler creates the handler logic
func makeSetStarHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.SetStarEndpoint, decodeSetStarRequest, encodeSetStarResponse, options...)
}

// decodeSetStarResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain SetStar request.
// TODO implement the decoder
func decodeSetStarRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Todo' Decoder is not impelemented")
}

// encodeSetStarResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeSetStarResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Todo' Encoder is not impelemented")
}
func (g *grpcServer) SetStar(ctx context1.Context, req *pb.SetStarRequest) (*pb.SetStarReply, error) {
	_, rep, err := g.setStar.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.SetStarReply), nil
}

// makeGetChildrenHandler creates the handler logic
func makeGetChildrenHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetChildrenEndpoint, decodeGetChildrenRequest, encodeGetChildrenResponse, options...)
}

// decodeGetChildrenResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetChildren request.
// TODO implement the decoder
func decodeGetChildrenRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Todo' Decoder is not impelemented")
}

// encodeGetChildrenResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetChildrenResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Todo' Encoder is not impelemented")
}
func (g *grpcServer) GetChildren(ctx context1.Context, req *pb.GetChildrenRequest) (*pb.GetChildrenReply, error) {
	_, rep, err := g.getChildren.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetChildrenReply), nil
}

// makeGetCategoryHandler creates the handler logic
func makeGetCategoryHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetCategoryEndpoint, decodeGetCategoryRequest, encodeGetCategoryResponse, options...)
}

// decodeGetCategoryResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetCategory request.
// TODO implement the decoder
func decodeGetCategoryRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Todo' Decoder is not impelemented")
}

// encodeGetCategoryResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetCategoryResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Todo' Encoder is not impelemented")
}
func (g *grpcServer) GetCategory(ctx context1.Context, req *pb.GetCategoryRequest) (*pb.GetCategoryReply, error) {
	_, rep, err := g.getCategory.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetCategoryReply), nil
}

// makeAddCategoryHandler creates the handler logic
func makeAddCategoryHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.AddCategoryEndpoint, decodeAddCategoryRequest, encodeAddCategoryResponse, options...)
}

// decodeAddCategoryResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain AddCategory request.
// TODO implement the decoder
func decodeAddCategoryRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Todo' Decoder is not impelemented")
}

// encodeAddCategoryResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeAddCategoryResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Todo' Encoder is not impelemented")
}
func (g *grpcServer) AddCategory(ctx context1.Context, req *pb.AddCategoryRequest) (*pb.AddCategoryReply, error) {
	_, rep, err := g.addCategory.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.AddCategoryReply), nil
}

// makeUpdateCategoryHandler creates the handler logic
func makeUpdateCategoryHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.UpdateCategoryEndpoint, decodeUpdateCategoryRequest, encodeUpdateCategoryResponse, options...)
}

// decodeUpdateCategoryResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain UpdateCategory request.
// TODO implement the decoder
func decodeUpdateCategoryRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Todo' Decoder is not impelemented")
}

// encodeUpdateCategoryResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeUpdateCategoryResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Todo' Encoder is not impelemented")
}
func (g *grpcServer) UpdateCategory(ctx context1.Context, req *pb.UpdateCategoryRequest) (*pb.UpdateCategoryReply, error) {
	_, rep, err := g.updateCategory.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.UpdateCategoryReply), nil
}

// makeDeleteCategoryHandler creates the handler logic
func makeDeleteCategoryHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.DeleteCategoryEndpoint, decodeDeleteCategoryRequest, encodeDeleteCategoryResponse, options...)
}

// decodeDeleteCategoryResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain DeleteCategory request.
// TODO implement the decoder
func decodeDeleteCategoryRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Todo' Decoder is not impelemented")
}

// encodeDeleteCategoryResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeDeleteCategoryResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Todo' Encoder is not impelemented")
}
func (g *grpcServer) DeleteCategory(ctx context1.Context, req *pb.DeleteCategoryRequest) (*pb.DeleteCategoryReply, error) {
	_, rep, err := g.deleteCategory.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.DeleteCategoryReply), nil
}

// makeGetCatChildesHandler creates the handler logic
func makeGetCatChildesHandler(endpoints endpoint.Endpoints, options []grpc.ServerOption) grpc.Handler {
	return grpc.NewServer(endpoints.GetCatChildesEndpoint, decodeGetCatChildesRequest, encodeGetCatChildesResponse, options...)
}

// decodeGetCatChildesResponse is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain GetCatChildes request.
// TODO implement the decoder
func decodeGetCatChildesRequest(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Todo' Decoder is not impelemented")
}

// encodeGetCatChildesResponse is a transport/grpc.EncodeResponseFunc that converts
// a user-domain response to a gRPC reply.
// TODO implement the encoder
func encodeGetCatChildesResponse(_ context.Context, r interface{}) (interface{}, error) {
	return nil, errors.New("'Todo' Encoder is not impelemented")
}
func (g *grpcServer) GetCatChildes(ctx context1.Context, req *pb.GetCatChildesRequest) (*pb.GetCatChildesReply, error) {
	_, rep, err := g.getCatChildes.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GetCatChildesReply), nil
}
