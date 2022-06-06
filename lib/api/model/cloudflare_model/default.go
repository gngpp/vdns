package cloudflare_model

import "time"

// Response is a template.  There will also be a result struct.  There will be a
// unique response type for each response, which will include this type.
type Response struct {
	Success  bool          `json:"success,omitempty"`
	Errors   []interface{} `json:"errors,omitempty"`
	Messages []interface{} `json:"messages,omitempty"`
}

//goland:noinspection SpellCheckingInspection
type Result struct {
	ID         *string   `json:"id,omitempty"`
	ZoneID     *string   `json:"zone_id,omitempty"`
	ZoneName   *string   `json:"zone_name,omitempty"`
	Name       *string   `json:"name,omitempty"`
	Type       *string   `json:"type,omitempty"`
	Content    *string   `json:"content,omitempty"`
	Proxiable  bool      `json:"proxiable,omitempty"`
	Proxied    bool      `json:"proxied,omitempty"`
	TTL        *int64    `json:"ttl,omitempty"`
	Locked     bool      `json:"locked,omitempty"`
	Meta       Meta      `json:"meta,omitempty"`
	CreatedOn  time.Time `json:"created_on,omitempty"`
	ModifiedOn time.Time `json:"modified_on,omitempty"`
}

type Meta struct {
	AutoAdded           bool   `json:"auto_added,omitempty"`
	ManagedByApps       bool   `json:"managed_by_apps,omitempty"`
	ManagedByArgoTunnel bool   `json:"managed_by_argo_tunnel,omitempty"`
	Source              string `json:"source,omitempty"`
}
