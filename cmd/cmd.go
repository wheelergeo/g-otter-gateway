package gatewaycmd

import (
	"context"
	"io"
	"os"
	"sync"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
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
	"github.com/wheelergeo/g-otter-gateway/pkg/token"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var once sync.Once
var c *cobra.Command

func Command() *cobra.Command {
	once.Do(func() {
		c = &cobra.Command{
			Use:   "gateway",
			Short: "start gateway",
			Run: func(cmd *cobra.Command, args []string) {
				var h *server.Hertz
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
				)

				if conf.GetConf().Hertz.EnableOtle &&
					conf.GetConf().Otle.Endpoint != "" {
					p := provider.NewOpenTelemetryProvider(
						provider.WithServiceName(conf.GetConf().Hertz.Service),
						provider.WithExportEndpoint(conf.GetConf().Otle.Endpoint),
						provider.WithInsecure(),
					)

					tracer, cfg := hertztracing.NewServerTracer()
					h = server.New(
						server.WithHostPorts(conf.GetConf().Hertz.Address),
						tracer,
					)
					h.OnShutdown = append(h.OnShutdown, func(ctx context.Context) {
						p.Shutdown(context.Background())
					})
					h.Use(hertztracing.ServerMiddleware(cfg))
				} else {
					h = server.New(
						server.WithHostPorts(conf.GetConf().Hertz.Address),
					)
				}

				registerLogger(h, nil)
				registerMiddleware(h)
				router.GeneratedRegister(h)

				h.GET("/ping", func(c context.Context, ctx *app.RequestContext) {
					ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
				})
				h.GET("/ping/pong", func(c context.Context, ctx *app.RequestContext) {
					ctx.JSON(consts.StatusOK, utils.H{"ping": "pong"})
				})
				h.GET("/hello", func(c context.Context, ctx *app.RequestContext) {
					ctx.JSON(consts.StatusOK, utils.H{"hello": "world"})
				})
				h.Spin()
			},
		}
	})

	return c
}

func limiterRule() (rules []limiter.LimiterRule) {
	rules = []limiter.LimiterRule{
		{
			ApiPath: "GET:/ping/pong",
			Type:    limiter.Flow,
			Qps:     1,
		},
		{
			ApiPath: "GET:/hello",
			Type:    limiter.Flow,
			Qps:     1,
		},
		{
			ApiPath:    "GET:/ping",
			Type:       limiter.HotspotQPS,
			Qps:        1,
			BurstCount: 5,
			Query: map[string]string{
				"test": "abc",
			},
		},
	}
	return
}

func registerLogger(h *server.Hertz, logger hlog.FullLogger) {
	hlog.SetLevel(conf.LogLevel())
	if logger != nil {
		hlog.SetLogger(logger)
	}
	asyncWriter := &zapcore.BufferedWriteSyncer{
		WS: zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.GetConf().Hertz.LogFileName,
			MaxSize:    conf.GetConf().Hertz.LogMaxSize,
			MaxBackups: conf.GetConf().Hertz.LogMaxBackups,
			MaxAge:     conf.GetConf().Hertz.LogMaxAge,
		}),
		FlushInterval: time.Minute,
	}

	hlog.SetOutput(io.MultiWriter(asyncWriter, os.Stdout))
	h.OnShutdown = append(h.OnShutdown, func(ctx context.Context) {
		asyncWriter.Sync()
	})
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
