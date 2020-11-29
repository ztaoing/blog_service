/**
* @Author:zhoutao
* @Date:2020/7/29 下午2:24
 */

package main

import (
	"blog_service/global"
	"blog_service/internal/model"
	"blog_service/internal/routers"
	"blog_service/pkg/logger"
	"blog_service/pkg/setting"
	"blog_service/pkg/tracer"
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

var (
	port    string
	runMode string
	config  string
)

//@title 博客系统
//@version 1.0
//@descrption blog with log、swagger、viper、error/success response
func main() {
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	//优雅
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("s.ListenAndServe err:%v", err)
		}
	}()

	//等待中断信号
	quit := make(chan os.Signal)

	//接收优雅停止信号信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("Shuting down server...")

	//最大时间控制，通知该服务它有5s的时间来处理原有的请求
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	//超时停止
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("server forced to shutdown:", err)
	}
	//已停止
	log.Println("server exiting.")

}

//全局变量初始化-》init-》main
func init() {
	//通过指定的方式，获取参数及配置文件
	err := setupFlag()
	if err != nil {
		log.Fatalf("init.setupFlag err:%v", err)
	}
	//初始化配置
	err = setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err:%v", err)
	}

	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err:%v", err)
	}

}

//go run main.go -port=8001 -model=release -config=configs/
func setupFlag() error {
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&runMode, "mode", "", "启动模式")
	flag.StringVar(&config, "config", "configs/", "指定使用的配置文件路径")
	flag.Parse()
	return nil
}
func setupSetting() error {
	s, err := setting.NewSetting(strings.Split(config, ",")...)
	if err != nil {
		return err
	}

	err = s.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}

	err = s.ReadSection("Jwt", &global.JwtSetting)
	if err != nil {
		return err
	}
	global.JwtSetting.Expire *= time.Second
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.AppSetting.DefaultContextTimeout *= time.Second

	if port != "" {
		global.ServerSetting.HttpPort = port
	}
	if runMode != "" {
		global.ServerSetting.RunMode = runMode
	}
	return nil
}

//初始化db
func setupDBEngine() error {
	var err error
	//注意此处 不是:= ，否则会重新声明并创建左侧的新局部变量
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600, //允许最大的占用空间为600MB
		MaxAge:    10,  //日志文件的最大生存周期
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}

func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer("blog-service", "127.0.0.1:6831")
	if err != nil {
		return err
	}
	//放到全局变量中
	global.Tracer = jaegerTracer
	return nil
}
