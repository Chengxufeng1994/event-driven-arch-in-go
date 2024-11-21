package monolith

import (
	"context"
	"os"
	"runtime"
	"time"

	"github.com/fatih/color"
	"github.com/gin-contrib/static"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/system"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/web"
)

var progressMessage = color.GreenString("==>")

type RunFunc func(basename string) error

type MonolithApplication struct {
	*system.System
	modules []system.Module
}

var _ system.Service = (*MonolithApplication)(nil)

func NewMonolithApplication(
	system *system.System,
	modules []system.Module,
) *MonolithApplication {
	return &MonolithApplication{
		System:  system,
		modules: modules,
	}
}

// FIXME:
func (app MonolithApplication) startupModules() error {
	for _, module := range app.modules {
		if err := module.Startup(context.Background(), app); err != nil {
			app.Logger().Errorf("%v %v\n", "prepare run modules: ", err)
			return err
		}
	}
	return nil
}

func (app MonolithApplication) Run() error {
	app.Logger().Info("started mallbots application")
	defer app.Logger().Info("stopped mallbots application")
	app.printWorkingDir()

	if err := app.startupModules(); err != nil {
		app.Logger().Errorf("%v %v\n", "prepare run modules: ", err)
		return err
	}

	staticFiles := static.EmbedFolder(web.WebUI, ".")
	app.Gin().Use(static.Serve("/", staticFiles))
	// static := http.FileServer(http.FS(web.WebUI))
	// app.gin.GET("/", func(ctx *gin.Context) {
	// 	static.ServeHTTP(ctx.Writer, ctx.Request)
	// })
	// app.gin.GET("/swagger-ui/*filepath", func(ctx *gin.Context) {
	// 	static.ServeHTTP(ctx.Writer, ctx.Request)
	// })

	app.Waiter().Add(app.WaitForWeb, app.WaitForRPC, app.WaitForStream)

	go func() {
		for {
			var mem runtime.MemStats
			runtime.ReadMemStats(&mem)
			app.Logger().Debugf("Alloc = %v  TotalAlloc = %v  Sys = %v  NumGC = %v", mem.Alloc/1024, mem.TotalAlloc/1024, mem.Sys/1024, mem.NumGC)
			time.Sleep(10 * time.Second)
		}
	}()

	return app.Waiter().Wait()
}

func (app MonolithApplication) printWorkingDir() {
	wd, _ := os.Getwd()
	app.Logger().Infof("%v workingDir: %s", progressMessage, wd)
}
