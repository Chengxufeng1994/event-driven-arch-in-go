package docs

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed index.html
//go:embed api.swagger.json
var swaggerUI embed.FS

func RegisterSwagger(mux *gin.Engine) error {
	const specRoot = "/search-spec/"

	mux.Any(fmt.Sprintf("%s*any", specRoot), gin.WrapH(
		http.StripPrefix(specRoot, http.FileServer(http.FS(swaggerUI)))),
	)

	return nil
}
