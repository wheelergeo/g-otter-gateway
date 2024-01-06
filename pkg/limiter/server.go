package limiter

import (
	"context"
	"sync"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/flow"
	"github.com/alibaba/sentinel-golang/core/hotspot"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

// Only for gin
type LimiteType int
type MWFunc func(*app.RequestContext) *base.BlockError

const (
	Flow LimiteType = iota
	HotspotQPS
	HotspotConcurrency
)

var once sync.Once

type LimiterRule struct {
	ApiPath     string // such as GET:/api/v1/user
	Type        LimiteType
	Concurrency int64             // for HotspotConcurrency
	Qps         int64             // for Flow and HotspotQPS
	Param       map[string]string // for HotspotQPS and HotspotConcurrency
	Query       map[string]string // for HotspotQPS and HotspotConcurrency
	BurstCount  int64             // for HotspotQPS
}

func GenerateMiddleware(errMsg string, errCode int,
	rules []LimiterRule) app.HandlerFunc {
	mwFuncs := initSentine(rules)
	return func(c context.Context, ctx *app.RequestContext) {
		_, err := sentinel.Entry(string(ctx.Method()) + ":" + ctx.FullPath())
		if err != nil {
			ctx.AbortWithStatusJSON(400, utils.H{
				"err":  errMsg,
				"code": errCode,
			})
			return
		}

		for _, v := range mwFuncs {
			err = v(ctx)
			if err != nil {
				ctx.AbortWithStatusJSON(400, utils.H{
					"err":  errMsg,
					"code": errCode,
				})
				return
			}
		}
		ctx.Next(c)
	}
}

func initSentine(rules []LimiterRule) []MWFunc {
	once.Do(func() {
		err := sentinel.InitDefault()
		if err != nil {
			panic(err)
		}
	})

	var flowRules []*flow.Rule
	var hotspotRules []*hotspot.Rule
	var mwFuncs []MWFunc
	for _, v := range rules {
		switch v.Type {
		case Flow:
			flowRules = append(flowRules, &flow.Rule{
				Resource:               v.ApiPath,
				Threshold:              float64(v.Qps),
				TokenCalculateStrategy: flow.Direct,
				ControlBehavior:        flow.Reject,
				StatIntervalInMs:       10000,
			})
		case HotspotQPS:
			hotspotRules = append(hotspotRules, &hotspot.Rule{
				Resource:      v.ApiPath,
				MetricType:    hotspot.QPS,
				ParamIndex:    0,
				BurstCount:    v.BurstCount,
				Threshold:     v.Qps,
				DurationInSec: 1,
			})
			for k1, v1 := range v.Query {
				mwFuncs = append(mwFuncs,
					func(ctx *app.RequestContext) (err *base.BlockError) {
						if ctx.Query(k1) == v1 {
							_, err = sentinel.Entry(
								string(ctx.Method())+":"+ctx.FullPath(),
								sentinel.WithArgs(v1),
							)
						}
						return
					})
			}
			for k1, v1 := range v.Param {
				mwFuncs = append(mwFuncs,
					func(ctx *app.RequestContext) (err *base.BlockError) {
						if ctx.Param(k1) == v1 {
							_, err = sentinel.Entry(
								string(ctx.Method())+":"+ctx.FullPath(),
								sentinel.WithArgs(v1),
							)
						}
						return
					})
			}
		case HotspotConcurrency:
			hotspotRules = append(hotspotRules, &hotspot.Rule{
				Resource:      v.ApiPath,
				MetricType:    hotspot.Concurrency,
				ParamIndex:    0,
				Threshold:     v.Concurrency,
				DurationInSec: 1,
			})
			for k1, v1 := range v.Query {
				mwFuncs = append(mwFuncs,
					func(ctx *app.RequestContext) (err *base.BlockError) {
						if ctx.Query(k1) == v1 {
							_, err = sentinel.Entry(
								string(ctx.Method())+":"+ctx.FullPath(),
								sentinel.WithArgs(v1),
							)
						}
						return
					})
			}
			for k1, v1 := range v.Param {
				mwFuncs = append(mwFuncs,
					func(ctx *app.RequestContext) (err *base.BlockError) {
						if ctx.Param(k1) == v1 {
							_, err = sentinel.Entry(
								string(ctx.Method())+":"+ctx.FullPath(),
								sentinel.WithArgs(v1),
							)
						}
						return
					})
			}
		default:
			continue
		}
	}
	if len(flowRules) != 0 {
		_, err := flow.LoadRules(flowRules)
		if err != nil {
			panic(err)
		}
	}

	if len(hotspotRules) != 0 {
		_, err := hotspot.LoadRules(hotspotRules)
		if err != nil {
			panic(err)
		}
	}
	return mwFuncs
}
