package v1

import (
	"context"
	"fmt"
	"net/http"

	orderv1 "github.com/Chengxufeng1994/event-driven-arch-in-go/ordering/api/order/v1"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RegisterGateway(ctx context.Context, mux *gin.Engine, endpoint string) error {
	const apiRoot = "/api/v1/ordering"

	gwMux := runtime.NewServeMux(
		runtime.WithErrorHandler(func(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, writer http.ResponseWriter, request *http.Request, err error) {
			// creating a new HTTTPStatusError with a custom status, and passing error
			newError := runtime.HTTPStatusError{
				HTTPStatus: 400,
				Err:        err,
			}
			// using default handler to do the rest of heavy lifting of marshaling error and adding headers
			runtime.DefaultHTTPErrorHandler(ctx, mux, marshaler, writer, request, &newError)
		}),
	)

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	err := orderv1.RegisterOrderingServiceHandlerFromEndpoint(ctx, gwMux, endpoint, opts)
	if err != nil {
		return err
	}

	mux.Any(apiRoot, gin.WrapH(gwMux))
	mux.Any(fmt.Sprintf("%s/*any", apiRoot), gin.WrapH(gwMux))

	return nil
}
