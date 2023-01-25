package main

import (
	"flag"
	"github.com/natefinch/lumberjack"
	"github.com/zhaoziliang2019/blog-service/global"
	"github.com/zhaoziliang2019/blog-service/internal/model"
	"github.com/zhaoziliang2019/blog-service/internal/routers"
	"github.com/zhaoziliang2019/blog-service/pkg/logger"
	"github.com/zhaoziliang2019/blog-service/pkg/setting"
	"github.com/zhaoziliang2019/blog-service/pkg/tracer"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	port    string
	runMode string
	config  string
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err:%v", err)
	}
	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err:%v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err:%v", err)
	}
	err = setupLogger()
	if err != nil {
		log.Fatalf("init.setupLogger err:%v", err)
	}
	err = setupTracer()
	if err != nil {
		log.Fatalf("init.setupTracer err:%v", err)
	}
	err = setupFlag()
	if err != nil {
		log.Fatalf("init.setupFlag err:%v", err)
	}
}

//@title 博客系统
//@version 1.0
//@description Go 编程之旅：一起用 GO做项目
//@termsOfService https://github.com/zhaoziliang2019/blog-service
func main() {
	router := routers.NewRouter()
	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		return
	}
}
func setupSetting() error {
	setting, err := setting.NewSetting(strings.Split(config, ",")...)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("App", &global.Appsetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("JWT", &global.JWTSetting)
	if err != nil {
		return err
	}
	err = setting.ReadSection("Email", &global.EmailSetting)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.JWTSetting.Expire *= time.Second
	if port != "" {
		global.ServerSetting.HttpPort = port
	}
	if runMode != "" {
		global.ServerSetting.RunMode = runMode
	}
	return nil
}
func setupLogger() error {
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename: global.Appsetting.LogSavePath + "/" +
			global.Appsetting.LogFileName + global.Appsetting.LogFileExt,
		MaxSize:   600,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
	return nil
}
func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}
func setupTracer() error {
	jaegerTracer, _, err := tracer.NewJaegerTracer("blog_service", "127.0.0.1:6831")
	if err != nil {
		return err
	}
	global.Tracer = jaegerTracer
	return nil
}
func setupFlag() error {
	flag.StringVar(&port, "port", "", "启动端口")
	flag.StringVar(&runMode, "mode", "", "启动模式")
	flag.StringVar(&config, "config", "configs/", "指定要使用的配置文件路径")
	flag.Parse()
	return nil
}
