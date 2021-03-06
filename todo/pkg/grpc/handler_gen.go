// THIS FILE IS AUTO GENERATED BY GK-CLI DO NOT EDIT!!
package grpc

import (
	endpoint "dakhou/todo/pkg/endpoint"
	pb "dakhou/todo/pkg/grpc/pb"
	grpc "github.com/go-kit/kit/transport/grpc"
)

// NewGRPCServer makes a set of endpoints available as a gRPC AddServer
type grpcServer struct {
	get            grpc.Handler
	add            grpc.Handler
	setComplete    grpc.Handler
	removeComplete grpc.Handler
	delete         grpc.Handler
	update         grpc.Handler
	setStar        grpc.Handler
	getChildren    grpc.Handler
	getCategory    grpc.Handler
	addCategory    grpc.Handler
	updateCategory grpc.Handler
	deleteCategory grpc.Handler
	getCatChildes  grpc.Handler
}

func NewGRPCServer(endpoints endpoint.Endpoints, options map[string][]grpc.ServerOption) pb.TodoServer {
	return &grpcServer{
		add:            makeAddHandler(endpoints, options["Add"]),
		addCategory:    makeAddCategoryHandler(endpoints, options["AddCategory"]),
		delete:         makeDeleteHandler(endpoints, options["Delete"]),
		deleteCategory: makeDeleteCategoryHandler(endpoints, options["DeleteCategory"]),
		get:            makeGetHandler(endpoints, options["Get"]),
		getCatChildes:  makeGetCatChildesHandler(endpoints, options["GetCatChildes"]),
		getCategory:    makeGetCategoryHandler(endpoints, options["GetCategory"]),
		getChildren:    makeGetChildrenHandler(endpoints, options["GetChildren"]),
		removeComplete: makeRemoveCompleteHandler(endpoints, options["RemoveComplete"]),
		setComplete:    makeSetCompleteHandler(endpoints, options["SetComplete"]),
		setStar:        makeSetStarHandler(endpoints, options["SetStar"]),
		update:         makeUpdateHandler(endpoints, options["Update"]),
		updateCategory: makeUpdateCategoryHandler(endpoints, options["UpdateCategory"]),
	}
}
