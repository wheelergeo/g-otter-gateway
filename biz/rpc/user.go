package rpc

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/transport"
	"github.com/wheelergeo/g-otter-gateway/conf"
	"github.com/wheelergeo/g-otter-gen/user/userservice"
)

var UserClient userservice.Client

func initUser() {
	epInfo := &rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Hertz.Service,
	}

	c, err := userservice.NewClient(
		conf.GetConf().Rpc[0].Service,
		client.WithClientBasicInfo(epInfo),
		client.WithHostPorts(conf.GetConf().Rpc[0].Address),
		client.WithTransportProtocol(transport.TTHeader),
	)
	if err != nil {
		panic(err)
	}
	UserClient = c
}
