package router

import (
	"context"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
	"go.elastic.co/apm"
	"go.elastic.co/apm/module/apmechov4"

	. "letinvr.com/echo-web/conf"
	"letinvr.com/echo-web/middleware/metrics/prometheus"
	"letinvr.com/echo-web/middleware/opentracing"
	"letinvr.com/echo-web/middleware/pprof"
	"letinvr.com/echo-web/module/log"
	"letinvr.com/echo-web/router/api"
	"letinvr.com/echo-web/router/socket"
	"letinvr.com/echo-web/router/web"
)

type (
	Host struct {
		Echo *echo.Echo
	}
)

func InitRoutes() map[string]*Host {
	// Hosts
	hosts := make(map[string]*Host)

	hosts[Conf.Server.DomainWeb] = &Host{web.Routers()}
	hosts[Conf.Server.DomainApi] = &Host{api.Routers()}
	hosts[Conf.Server.DomainSocket] = &Host{socket.Routers()}

	return hosts
}

// 子域名部署
func RunSubdomains(confFilePath string) {
	// 配置初始化
	if err := InitConfig(confFilePath); err != nil {
		log.Panic(err)
	}

	// 全局日志级别
	log.SetLevel(GetLogLvl())

	// Server
	e := echo.New()

	// pprof
	e.Pre(pprof.Serve())

	e.Pre(mw.RemoveTrailingSlash())

	// Elastic APM
	// Requires APM Server 6.5.0 or newer
	apm.DefaultTracer.Service.Name = Conf.Opentracing.ServiceName
	apm.DefaultTracer.Service.Version = Conf.App.Version
	e.Use(apmechov4.Middleware(
		apmechov4.WithRequestIgnorer(func(request *http.Request) bool {
			return false
		}),
	))

	// OpenTracing
	otCtf := opentracing.Configuration{
		Disabled: Conf.Opentracing.Disable,
		Type:     opentracing.TracerType(Conf.Opentracing.Type),
	}
	if closer := otCtf.InitGlobalTracer(
		opentracing.ServiceName(Conf.Opentracing.ServiceName),
		opentracing.Address(Conf.Opentracing.Address),
	); closer != nil {
		defer closer.Close()
	}

	// 日志级别
	e.Logger.SetLevel(GetLogLvl())

	// Metrics
	if !Conf.Metrics.Disable {
		e.Use(prometheus.MetricsFunc(
			prometheus.Namespace("echo_web"),
		))

		/* Push模式
		m := metrics.NewMetrics(metrics.Prefix(""))
		e.Use(metrics.MetricsFunc(m))
		m.MemStats()

		hostname, err := os.Hostname()
		if err != nil {
			log.Warnf("os.Hostname() error:%v", err)
			hostname = "-"
		}
		// Graphite
		addr, _ := net.ResolveTCPAddr("tcp", Conf.Metrics.Address)
		m.Graphite(Conf.Metrics.FreqSec*time.Second, "echo-web.node."+hostname, addr)

		// InfluxDB
		m.InfluxDBWithTags(
			Conf.Metrics.FreqSec*time.Second,
			"http://127.0.0.1:8086",
			"metrics",
			"test",
			"test",
			map[string]string{"node": hostname})
		*/
	}

	// Secure, XSS/CSS HSTS
	e.Use(mw.SecureWithConfig(mw.DefaultSecureConfig))
	e.Use(mw.MethodOverride())

	// CORS
	e.Use(mw.CORSWithConfig(mw.CORSConfig{
		AllowOrigins: []string{"http://" + Conf.Server.DomainWeb, "http://" + Conf.Server.DomainApi},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding, echo.HeaderAuthorization},
	}))

	hosts := InitRoutes()
	e.Any("/*", func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()

		u, _err := url.Parse(c.Scheme() + "://" + req.Host)
		if _err != nil {
			e.Logger.Errorf("Request URL parse error:%v", _err)
		}

		host := hosts[u.Hostname()]
		if host == nil {
			e.Logger.Info("Host not found")
			err = echo.ErrNotFound
		} else {
			host.Echo.ServeHTTP(res, req)
		}

		return
	})

	if !Conf.Server.Graceful {
		e.Logger.Fatal(e.Start(Conf.Server.Addr))
	} else {
		// Graceful Shutdown
		// Start server
		go func() {
			if err := e.Start(Conf.Server.Addr); err != nil {
				e.Logger.Errorf("Shutting down the server with error:%v", err)
			}
		}()

		// Wait for interrupt signal to gracefully shutdown the server with
		// a timeout of 10 seconds.
		quit := make(chan os.Signal)
		signal.Notify(quit, os.Interrupt)
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			e.Logger.Fatal(err)
		}
	}
}
