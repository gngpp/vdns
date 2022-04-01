package server

import (
	"vdns/lib/vlog"

	"github.com/kardianos/service"
)

type Vdns struct{}

func (p *Vdns) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.run()
	return nil
}
func (p *Vdns) run() {
	// Do work here
	vlog.Info("vdns starts to execute the service")
}
func (p *Vdns) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	vlog.Info("vdns stop service")
	return nil
}
