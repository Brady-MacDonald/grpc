package clients

import "github.com/Brady-MacDonald/grpc/gateway/config"

type Clients struct{}

func New(cfg *config.Config) *Clients {
	return &Clients{}
}
