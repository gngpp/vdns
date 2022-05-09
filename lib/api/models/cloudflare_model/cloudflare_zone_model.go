package cloudflare_model

import "time"

//goland:noinspection SpellCheckingInspection
type CloudFlareZoneResult struct {
	Result []struct {
		Id                  string      `json:"id,omitempty"`
		Name                string      `json:"name,omitempty"`
		Status              string      `json:"status,omitempty"`
		Paused              bool        `json:"paused,omitempty"`
		Type                string      `json:"type,omitempty"`
		DevelopmentMode     int         `json:"development_mode,omitempty"`
		NameServers         []string    `json:"name_servers,omitempty"`
		OriginalNameServers []string    `json:"original_name_servers,omitempty"`
		OriginalRegistrar   string      `json:"original_registrar,omitempty"`
		OriginalDnshost     interface{} `json:"original_dnshost,omitempty"`
		ModifiedOn          time.Time   `json:"modified_on,omitempty"`
		CreatedOn           time.Time   `json:"created_on,omitempty"`
		ActivatedOn         interface{} `json:"activated_on,omitempty"`
		Meta                struct {
			Step                    int  `json:"step,omitempty"`
			CustomCertificateQuota  int  `json:"custom_certificate_quota,omitempty"`
			PageRuleQuota           int  `json:"page_rule_quota,omitempty"`
			PhishingDetected        bool `json:"phishing_detected,omitempty"`
			MultipleRailgunsAllowed bool `json:"multiple_railguns_allowed,omitempty"`
		} `json:"meta,omitempty"`
		Owner struct {
			Id    string `json:"id,omitempty"`
			Type  string `json:"type,omitempty"`
			Email string `json:"email,omitempty"`
		} `json:"owner,omitempty"`
		Account struct {
			Id   string `json:"id,omitempty"`
			Name string `json:"name,omitempty"`
		} `json:"account,omitempty"`
		Permissions []string `json:"permissions,omitempty"`
		Plan        struct {
			Id                string `json:"id,omitempty"`
			Name              string `json:"name,omitempty"`
			Price             int    `json:"price,omitempty"`
			Currency          string `json:"currency,omitempty"`
			Frequency         string `json:"frequency,omitempty"`
			IsSubscribed      bool   `json:"is_subscribed,omitempty"`
			CanSubscribe      bool   `json:"can_subscribe,omitempty"`
			LegacyId          string `json:"legacy_id,omitempty"`
			LegacyDiscount    bool   `json:"legacy_discount,omitempty"`
			ExternallyManaged bool   `json:"externally_managed,omitempty"`
		} `json:"plan,omitempty"`
	} `json:"result,omitempty"`
	ResultInfo struct {
		Page       int `json:"page,omitempty"`
		PerPage    int `json:"per_page,omitempty"`
		TotalPages int `json:"total_pages,omitempty"`
		Count      int `json:"count,omitempty"`
		TotalCount int `json:"total_count,omitempty"`
	} `json:"result_info,omitempty"`
	Success  bool          `json:"success,omitempty"`
	Errors   []interface{} `json:"errors,omitempty"`
	Messages []interface{} `json:"messages,omitempty"`
}
