package main

import (
	"log"
	"os"

	"github.com/kardianos/service"
)

var logger service.Logger

type program struct{}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *program) run() {
	// Do work here
	log.Println("开始服务")
}
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	log.Println("中止服务")
	return nil
}

func main() {
	//服务的配置信息
	cfg := &service.Config{
		Name:        "server",
		DisplayName: "vdns server",
		Description: "This is an vdns Go service.",
	}
	// Interface 接口
	prg := &program{}
	// 构建服务对象
	s, err := service.New(prg, cfg)
	if err != nil {
		log.Fatal(err)
	}
	// logger 用于记录系统日志
	logger, err := s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}
	// 若是有命令则执行
	if len(os.Args) == 2 {
		err = service.Control(s, os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	} else {
		// 不然说明是方法启动了
		err = s.Run()
		if err != nil {
			logger.Error(err)
		}
	}
	if err != nil {
		logger.Error(err)
	}
}
