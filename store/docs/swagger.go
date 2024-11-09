package docs

import (
	"embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed index.html
//go:embed api.swagger.json
var swaggerUI embed.FS

func RegisterSwagger(mux *gin.Engine) error {
	const specRoot = "/stores-spec/"
	mux.StaticFS(specRoot, http.FS(swaggerUI))
	return nil
}
