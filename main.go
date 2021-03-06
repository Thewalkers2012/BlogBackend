package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Thewalkers2012/BlogBackend/controller"
	"github.com/Thewalkers2012/BlogBackend/logger"
	"github.com/Thewalkers2012/BlogBackend/repository/mysql"
	"github.com/Thewalkers2012/BlogBackend/repository/redis"
	"github.com/Thewalkers2012/BlogBackend/routes"
	"github.com/Thewalkers2012/BlogBackend/settings"
	"github.com/Thewalkers2012/BlogBackend/util/snowflake"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// @title Blog
// @version 1.0
// @description 用 Go 语言实现的博客后端代码 Api
// @termsOfService http://swagger.io/terms/

// @contact.name thewalkers
// @contact.url http://www.swagger.io/support
// @contact.email 1518832673@qq.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1
// @BasePath /swagger
func main() {
	// 1. 加载配置文件
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed, err: %v\n", err)
		return
	}

	// 2. 初始化日志
	if err := logger.Init(settings.Config.LogConfig, settings.Config.Mode); err != nil {
		fmt.Printf("init logger failed, err: %v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger init success ...")

	// 3. 初始化 mysql 连接
	if err := mysql.Init(settings.Config.MySQLConfig); err != nil {
		fmt.Printf("init mysql failed, err: %v\n", err)
		return
	}
	defer mysql.Close()

	// 4. 初始化 redis 连接
	if err := redis.Init(settings.Config.RedisConfig); err != nil {
		fmt.Printf("init redis failed, err: %v\n", err)
		return
	}

	// 5. 注册路由
	r := routes.Setup(settings.Config.Mode)

	// 6. 启动服务 （优雅关机操作）
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("port")),
		Handler: r,
	}

	// 7. 初始化雪花算法
	if err := snowflake.Init(settings.Config.StartTime, settings.Config.MachineID); err != nil {
		zap.L().Fatal("init snowflake failed", zap.Error(err))
		return
	}

	// 8. 初始化 gin 框架内置的校验器使用的翻译器
	if err := controller.InitTrans("zh"); err != nil {
		fmt.Printf("init validator trans failed, err: %v\n", err)
		return
	}

	go func() {
		// 开启一个 goroutine 启动服务
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	zap.L().Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")

}
