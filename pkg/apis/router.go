package apis

import (
	"context"
	"github.com/Evolt0/def-kit/proto/book"
	"github.com/Evolt0/def-kit/proto/hello"
	"github.com/Evolt0/srv-kit/pkg/endpoint"
	"github.com/Evolt0/srv-kit/pkg/service"
	"github.com/Evolt0/srv-kit/pkg/transport"
	kitgrpc "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

//调用该方法将GetBookInfoByID方法的实现注册到gRPC服务中
func InitGRPCRouter(gs *grpc.Server) {
	book.RegisterBookFunServer(gs, NewGrpcBook())
	hello.RegisterHelloFunServer(gs, NewGrpcHello())
}

//创建实现接口的函数，此处实现了endpoint的创建，并赋值给结构体中的参数
func NewGrpcBook() book.BookFunServer {
	b := &GrpcBook{}
	b.Handler = kitgrpc.NewServer(endpoint.GetGrpcEndpointForGetBookIDs(new(service.Book)),
		transport.DecodeBook, transport.EncodeBook)
	return b
}

type GrpcBook struct {
	Handler kitgrpc.Handler
}

//实现proto中的接口
func (g *GrpcBook) GetBookInfoByID(ctx context.Context, request *book.BookRequest) (*book.BookResponse, error) {
	_, res, err := g.Handler.ServeGRPC(ctx, request)
	return res.(*book.BookResponse), err
}

func NewGrpcHello() hello.HelloFunServer {
	b := &GrpcHello{}
	b.Handler = kitgrpc.NewServer(endpoint.GetGrpcEndpointForHello(new(service.HelloImpl)),
		transport.DecodeHello, transport.EncodeHello)
	return b
}

type GrpcHello struct {
	Handler kitgrpc.Handler
}

func (g *GrpcHello) HelloWorld(ctx context.Context, req *hello.HelloReq) (*hello.HelloResp, error) {
	_, resp, err := g.Handler.ServeGRPC(ctx, req)
	return resp.(*hello.HelloResp), err
}
