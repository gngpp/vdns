package parameter

import (
	"vdns/lib/auth"
	"vdns/lib/sign/compose"
	"vdns/lib/standard"
)

type CloudflareParameter struct {
	credential        auth.Credential
	signatureComposer compose.SignatureComposer
	version           *standard.Standard
}
