package server

import (
	"github.com/kardianos/service"
	"io"
	"os"
	"runtime"
	"time"
	"vdns/config"
	"vdns/lib/vlog"
)

type Vdns struct {
	interval int
	debug    bool
	ddns     *DDNS
}

func NewVdns(interval int, debug bool) Vdns {
	if interval <= 0 {
		return Vdns{
			interval: 5,
		}
	}
	return Vdns{
		interval: interval,
		debug:    debug,
		ddns:     new(DDNS),
	}
}

func (p *Vdns) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	c, err := config.ReadVdnsConfig()
	if err != nil {
		return s.Stop()
	}

	multiWriter := io.MultiWriter(os.Stdout, c.ToVlogTimeWriter())
	vlog.SetOutput(multiWriter)
	vlog.SetSyncOutput(true)
	if p.debug {
		vlog.SetLevel(vlog.Level.DEBUG)
	}
	vlog.Debugf("running args: %v", os.Args)
	vlog.Infof("running platform: %v", s.Platform())
	// Minimal use of resources
	runtime.GOMAXPROCS(1)
	go p.run()
	return nil
}

func (p *Vdns) run() {
	// Do work here
	vlog.Infof("vdns start to execute the service - [interval execution time: %v minute]", p.interval)
	timer := time.NewTicker(time.Minute * time.Duration(p.interval))
	p.resolveHandler()
	for {
		t := <-timer.C
		vlog.Debugf("now resovle time: %v, do...\n", t)
		p.resolveHandler()
	}
}
func (p *Vdns) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	vlog.Info("vdns stop service")
	return nil
}

func (p *Vdns) resolveHandler() {
	err := p.ddns.Resolve()
	if err != nil {
		vlog.Errorf("resovle error: %v", err)
	}
}
