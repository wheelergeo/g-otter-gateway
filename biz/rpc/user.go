package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/transport"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/wheelergeo/g-otter-gateway/conf"
	"github.com/wheelergeo/g-otter-gen/user/userservice"
)

var UserClient userservice.Client

func initUser() {
	epInfo := &rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Hertz.Service,
	}
	opts := []client.Option{
		client.WithClientBasicInfo(epInfo),
		client.WithHostPorts(conf.GetConf().Rpc[0].Address),
		client.WithTransportProtocol(transport.TTHeader),
	}

	if conf.GetConf().Hertz.EnableOtel &&
		conf.GetConf().Otel.Endpoint != "" {
		opts = append(opts, client.WithSuite(tracing.NewClientSuite()))
	}

	c, err := userservice.NewClient(
		conf.GetConf().Rpc[0].Service,
		opts...,
	)
	if err != nil {
		panic(err)
	}
	UserClient = c
}
