package models

import "vdns/lib/api/models/cloudflare_model"

func NewCloudflareZones() *cloudflare_model.CloudFlareZoneResult {
	return &cloudflare_model.CloudFlareZoneResult{}
}

// Domain -> www.innas.work, Domain=innas.work Subdomain=www
type Domain interface {
	GetDomain() string
	GetSubdomain() string
}
