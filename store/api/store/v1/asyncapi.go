package storev1

import (
	"embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed index.html
//go:embed css/*
//go:embed js/*
var asyncAPI embed.FS

func RegisterAsyncAPI(r *gin.Engine) error {
	const specRoot = "/stores-asyncapi/"

	// mount the swagger specifications
	r.StaticFS(specRoot, http.FS(asyncAPI))

	return nil
}
