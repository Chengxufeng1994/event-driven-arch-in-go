package main

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-contrib/static"
	"github.com/spf13/viper"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/config"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/system"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/web"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment/migrations"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("payments exitted abnormally: %s\n", err)
		os.Exit(1)
	}
}

func run() (err error) {
	log := logger.ContextUnavailable()

	cfg, err := config.LoadConfig("")
	if err != nil {
		log.Errorf("failed to load config: %v", err)
		return nil
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(_ fsnotify.Event) {
		log.Info("reloading config")
		if err := viper.Unmarshal(&cfg); err != nil {
			log.Errorf("failed to unmarshal config: %v", err)
		}
	})
	s, err := system.NewSystem("MALL BOTS", "mallbots-payments", cfg, log)
	if err != nil {
		log.Errorf("failed to new system: %v", err)
		return nil
	}
	if err = s.MigrateDB(migrations.FS); err != nil {
		return err
	}
	staticFiles := static.EmbedFolder(web.WebUI, ".")
	s.Gin().Use(static.Serve("/", staticFiles))
	// call the module composition root
	if err = payment.Root(s.Waiter().Context(), s); err != nil {
		return err
	}

	log.Info("started payments service")
	defer log.Info("stopped payments service")

	s.Waiter().Add(
		s.WaitForWeb,
		s.WaitForRPC,
		s.WaitForStream,
	)

	// go func() {
	// 	for {
	// 		var mem runtime.MemStats
	// 		runtime.ReadMemStats(&mem)
	// 		m.logger.Debug().Msgf("Alloc = %v  TotalAlloc = %v  Sys = %v  NumGC = %v", mem.Alloc/1024, mem.TotalAlloc/1024, mem.Sys/1024, mem.NumGC)
	// 		time.Sleep(10 * time.Second)
	// 	}
	// }()

	return s.Waiter().Wait()
}
