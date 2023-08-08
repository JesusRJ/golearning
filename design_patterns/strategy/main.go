package main

import (
	"context"
	"fmt"
)

type PersonGateway struct{}

func (*PersonGateway) Name() string { return "person" }

type GatewayHandler func(context.Context, any) error

type GatewayItem interface {
	Name() string
	Handler() GatewayHandler
}

type Gateway map[string]GatewayHandler

type GatewayImpl struct {
	provider map[string]GatewayHandler
}

func (h *GatewayImpl) Handle(ctx context.Context, Type string) error {
	fn, ok := h.provider[Type]
	if !ok {
		return fmt.Errorf("provider not avaiable: %s", Type)
	}

	return fn(ctx, Type)
}

func NewGatewayImpl(provider ...GatewayItem) *GatewayImpl {
	g := &GatewayImpl{
		provider: make(map[string]GatewayHandler),
	}

	for _, p := range provider {
		g.provider[p.Name()] = p.Handler()
	}

	return g
}
