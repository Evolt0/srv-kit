package endpoint

import (
	"context"
	"github.com/Evolt0/def-kit/proto/book"
	"github.com/Evolt0/def-kit/proto/hello"
	"github.com/Evolt0/srv-kit/pkg/service"
	"github.com/go-kit/kit/endpoint"
)

func GetGrpcEndpointForGetBookIDs(inter service.BookInter) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		bookResp := new(book.BookResponse)
		bookResp.Name = inter.GetBookInfoByID(request.(*book.BookRequest).Id)
		return bookResp, err
	}
}

func GetGrpcEndpointForHello(data service.Hello) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		helloResp := new(hello.HelloResp)
		helloResp.Resp = data.Hello()
		return helloResp, err
	}

}
