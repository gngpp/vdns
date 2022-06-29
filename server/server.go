package server

import (
	"fmt"
	"time"
	"vdns/lib/vlog"

	"github.com/kardianos/service"
)

type Vdns struct {
	interval int
}

func NewVdns(interval int) Vdns {
	if interval <= 0 {
		return Vdns{
			interval: 5,
		}
	}
	return Vdns{
		interval: interval,
	}
}

func (p *Vdns) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	vlog.Debug(s)
	go p.run()
	return nil
}
func (p *Vdns) run() {
	// Do work here
	vlog.Info("vdns start to execute the service")
	vlog.Infof("interval execution time: %v minute", p.interval)
	timer := time.NewTicker(time.Minute * time.Duration(p.interval))
	for {
		select {
		case <-timer.C:
			fmt.Println("do..")
		}
	}
}
func (p *Vdns) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	vlog.Info("vdns stop service")
	return nil
}
