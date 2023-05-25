package main

import (
	"context"
	"log"
	"net/http"

	"github.com/bufbuild/connect-go"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	myServiceV1 "github.com/suzushin54/go-api-with-connect/gen/my_service/v1"
	"github.com/suzushin54/go-api-with-connect/gen/my_service/v1/my_servicev1connect"
)

type GreetServer struct{}

func (s *GreetServer) HelloWorld(
	ctx context.Context,
	req *connect.Request[myServiceV1.HelloWorldRequest],
) (*connect.Response[myServiceV1.HelloWorldResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&myServiceV1.HelloWorldResponse{
		Message: "Hello Connect!",
	})
	res.Header().Set("Greet-Version", "v1")
	return res, nil
}

func main() {
	greeter := &GreetServer{}
	mux := http.NewServeMux()
	path, handler := my_servicev1connect.NewMyServiceHandler(greeter)
	mux.Handle(path, handler)
	http.ListenAndServe(
		"localhost:8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
