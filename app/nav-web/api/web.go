package main

import (
	"flag"
	"fmt"
	"net/http"

	"nav-go-zero/app/nav-web/api/internal/config"
	"nav-go-zero/app/nav-web/api/internal/handler"
	"nav-go-zero/app/nav-web/api/internal/svc"
	"nav-go-zero/app/pkg/errorx"
	"nav-go-zero/pkg/middleware/cors"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var configFile = flag.String("f", "etc/web.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 跨域配置：Register go-zero-cors-middleware handler to handle preflight request
	cors := cors.NewCORSMiddleware(&cors.Options{})

	// 跨域配置：Add run option WithNotAllowedHandler and register `.Handler()` to handle `OPTIONS` request (preflight)
	server := rest.MustNewServer(c.RestConf, rest.WithNotAllowedHandler(cors.Handler()))
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 跨域配置：添加跨域请求配置
	server.Use(cors.Handle)

	// 自定义错误
	httpx.SetErrorHandler(func(err error) (int, interface{}) {
		switch e := err.(type) {
		case *errorx.CodeError:
			return http.StatusOK, e.Data()
		default:
			return http.StatusInternalServerError, nil
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
