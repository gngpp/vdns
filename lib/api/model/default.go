package model

// Domain -> www.innas.work, Domain=innas.work Subdomain=www
type Domain interface {
	GetDomain() string
	GetSubdomain() string
}
