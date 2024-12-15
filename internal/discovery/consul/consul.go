package consul

import (
	"context"
	"fmt"
	"time"

	"github.com/Chengxufeng1994/event-driven-arch-in-go/internal/discovery"
	consulapi "github.com/hashicorp/consul/api"
)

type ConsulDiscoverClientConfig struct {
	ConsulHost  string `mapstructure:"consul_host"`
	ConsulPort  int    `mapstructure:"consul_port"`
	ServiceName string `mapstructure:"service_name"`
	InstanceID  string `mapstructure:"instance_id"`
}

type ConsulDiscoverClient struct {
	client *consulapi.Client
}

var _ discovery.Discovery = (*ConsulDiscoverClient)(nil)

func New(cfg *ConsulDiscoverClientConfig) *ConsulDiscoverClient {
	config := consulapi.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%d", cfg.ConsulHost, cfg.ConsulPort)
	client, err := consulapi.NewClient(config)
	if err != nil {
		return nil
	}

	return &ConsulDiscoverClient{
		client: client,
	}
}

func (cli ConsulDiscoverClient) Register(ctx context.Context, serviceName string, instanceID string, host string, port int) error {
	// Check: &consulapi.AgentServiceCheck{
	// 	HTTP:                           fmt.Sprintf("http://%s:%d/health", parts[0], port),
	// 	DeregisterCriticalServiceAfter: "30s",
	// 	Interval:                       "15s",
	// 	Timeout:                        "1s",
	// },

	Check := &consulapi.AgentServiceCheck{
		DeregisterCriticalServiceAfter: "30s",
		TLSSkipVerify:                  true,
		TTL:                            "5s",
		CheckID:                        instanceID,
	}

	registration := &consulapi.AgentServiceRegistration{
		ID:      instanceID,
		Name:    serviceName,
		Address: host,
		Port:    port,
		Check:   Check,
	}

	go cli.updateHealthCheckStatus(instanceID)

	return cli.client.Agent().ServiceRegister(registration)
}

func (cli ConsulDiscoverClient) updateHealthCheckStatus(instanceID string) {
	ticker := time.NewTicker(5 * time.Second)
	for range ticker.C {
		err := cli.client.Agent().UpdateTTL(instanceID, "online", consulapi.HealthPassing)
		if err != nil {
			break
		}
	}
}

func (cli ConsulDiscoverClient) DeRegister(ctx context.Context, serviceName string, instanceID string) error {
	return cli.client.Agent().ServiceDeregister(instanceID)
}
