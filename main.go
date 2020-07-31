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
	"errors"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"net/http"
	"time"
)

//@title 博客系统，这里的注解可以用来区分项目
//@version 1.0
//@descrption blog with log、swagger、viper、error/success response、dalidator
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
	global.Logger.Infof("%s", "hahah")
	s.ListenAndServe()
}

//全局变量初始化-》init-》main
func init() {
	err := setupSetting()
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
}

func setupSetting() error {
	var s *setting.ServerSettings
	if s == nil {
		return errors.New("here")
	}
	settings, err := setting.NewSetting()
	if err != nil {
		return err
	}

	err = settings.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}

	err = settings.ReadSection("App", &global.AppSetting)
	if err != nil {
		return err
	}

	err = settings.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
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
	global.Logger = logger.NewLlogger(&lumberjack.Logger{
		Filename:  global.AppSetting.LogSavePath + "/" + global.AppSetting.LogFileName + global.AppSetting.LogFileExt,
		MaxSize:   600, //允许最大的占用空间为600MB
		MaxAge:    10,  //日志文件的最大生存周期
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}
