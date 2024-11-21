package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/basket"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/cosec"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/customer"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/depot"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/config"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/logger"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/monolith"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/system"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/migrations"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/notification"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/ordering"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/payment"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/search"
	"github.com/Chengxufeng1994/event-driven-arch-in-go/store"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	_ "go.uber.org/automaxprocs"
)

func main() {
	rand.NewSource(time.Now().UTC().UnixNano())
	if err := newMonolith().Run(); err != nil {
		os.Exit(1)
	}
}

func newMonolith() *monolith.MonolithApplication {
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

	s, err := system.NewSystem("MALL BOTS", "mallbots-system", cfg, log)
	if err != nil {
		log.Errorf("failed to new system: %v", err)
		return nil
	}

	modules := []system.Module{
		basket.NewModule(),
		customer.NewModule(),
		depot.NewModule(),
		notification.NewModule(),
		ordering.NewModule(),
		payment.NewModule(),
		store.NewModule(),
		search.NewModule(),
		cosec.NewModule(),
	}

	m := monolith.NewMonolithApplication(s, modules)
	if err := m.MigrateDB(migrations.FS); err != nil {
		return nil
	}

	return m
}
