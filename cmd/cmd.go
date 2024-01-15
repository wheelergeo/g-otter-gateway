package gatewaycmd

import (
	"context"
	"sync"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/cors"
	"github.com/hertz-contrib/gzip"
	"github.com/hertz-contrib/logger/accesslog"
	"github.com/hertz-contrib/obs-opentelemetry/provider"
	hertztracing "github.com/hertz-contrib/obs-opentelemetry/tracing"
	"github.com/hertz-contrib/pprof"
	"github.com/spf13/cobra"
	"github.com/wheelergeo/g-otter-gateway/biz/dal"
	"github.com/wheelergeo/g-otter-gateway/biz/dal/mysql"
	"github.com/wheelergeo/g-otter-gateway/biz/dal/redis"
	"github.com/wheelergeo/g-otter-gateway/biz/router"
	"github.com/wheelergeo/g-otter-gateway/biz/rpc"
	"github.com/wheelergeo/g-otter-gateway/conf"
	"github.com/wheelergeo/g-otter-gateway/pkg/auth"
	"github.com/wheelergeo/g-otter-gateway/pkg/limiter"
	"github.com/wheelergeo/g-otter-gateway/pkg/logger"
	"github.com/wheelergeo/g-otter-gateway/pkg/token"
	"github.com/wheelergeo/g-otter-gateway/pkg/validate"
)

var once sync.Once
var c *cobra.Command

func Command() *cobra.Command {
	once.Do(func() {
		c = &cobra.Command{
			Use:   "gateway",
			Short: "start gateway",
			Run: func(cmd *cobra.Command, args []string) {
				dal.Init()
				rpc.Init()
				auth.NewServer(
					conf.GetConf().Casbin.ModelName,
					mysql.DB,
					conf.GetConf().Casbin.PolicyTable,
					conf.GetConf().Casbin.PolicyRedis,
				)
				token.Init(
					redis.RedisClient,
					conf.GetConf().Paseto.CacheKey,
					func(tv *token.TokenValue, cd token.ClaimData) {
					},
				)
				logger.InitHlogWithLogrus(
					logger.Config{
						Mode:   logger.StdOut,
						Format: logger.Text,
						Level:  conf.LogLevel(),
						FileCfg: logger.FileConfig{
							FileName:      conf.GetConf().Hertz.LogFileName,
							MaxSize:       conf.GetConf().Hertz.LogMaxSize,
							MaxBackups:    conf.GetConf().Hertz.LogMaxBackups,
							MaxAge:        conf.GetConf().Hertz.LogMaxAge,
							FlushInterval: time.Minute,
						},
					},
				)
				h := gateway()

				registerMiddleware(h)
				router.GeneratedRegister(h)

				h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
					hlog.CtxInfof(c, ctx.ClientIP())
					ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
				})

				hlog.Infof("[%s] Server start", conf.GetEnv())
				h.Spin()
			},
		}
	})

	return c
}

func gateway() (h *server.Hertz) {
	opts := []config.Option{
		server.WithHostPorts(conf.GetConf().Hertz.Address),
		server.WithValidateConfig(validate.Config()),
	}

	if conf.GetConf().Hertz.EnableOtel &&
		conf.GetConf().Otel.Endpoint != "" {
		p := provider.NewOpenTelemetryProvider(
			provider.WithServiceName(conf.GetConf().Hertz.Service),
			provider.WithExportEndpoint(conf.GetConf().Otel.Endpoint),
			provider.WithInsecure(),
		)
		tracer, cfg := hertztracing.NewServerTracer()
		opts = append(opts, tracer)

		h = server.New(opts...)
		h.OnShutdown = append(h.OnShutdown, func(ctx context.Context) {
			p.Shutdown(context.Background())
		})
		h.Use(hertztracing.ServerMiddleware(cfg))
		return
	}
	h = server.New(opts...)
	return
}

func limiterRule() (rules []limiter.LimiterRule) {
	rules = []limiter.LimiterRule{
		{
			ApiPath:    "GET:/ping",
			Type:       limiter.HotspotQPS,
			Qps:        1,
			BurstCount: 5,
			Query: map[string]string{
				"test": "abc",
			},
		},
		{
			Type:       limiter.IpFlow,
			Qps:        1,
			BurstCount: 5,
		},
	}
	return
}

func registerMiddleware(h *server.Hertz) {
	// pprof
	if conf.GetConf().Hertz.EnablePprof {
		pprof.Register(h)
	}

	// gzip
	if conf.GetConf().Hertz.EnableGzip {
		h.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	// access log
	if conf.GetConf().Hertz.EnableAccessLog {
		h.Use(accesslog.New())
	}

	// recovery
	h.Use(recovery.Recovery())

	// cores
	h.Use(cors.Default())

	// sentinel
	h.Use(limiter.GenerateMiddleware(
		"too many request, the quoto used up!",
		10222,
		limiterRule(),
	))

}
