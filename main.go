// CaddyBanIP is a Caddy 2 module that allows you to ban IPs from visiting a part or the entirety of your website.
package caddybanip

import (
	"fmt"
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"

	"go.uber.org/zap"

	"regexp"
)

type CaddyBanIP struct {
	BannedIPs string `json:"banned_ips,omitempty"`
	Message   string `json:"message,omitempty"`

	logger *zap.Logger
}

func init() {
	caddy.RegisterModule(CaddyBanIP{})
	httpcaddyfile.RegisterHandlerDirective("caddybanip", parseCaddyfile)
}

func (CaddyBanIP) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.caddybanip",
		New: func() caddy.Module { return new(CaddyBanIP) },
	}
}

func (c *CaddyBanIP) Provision(ctx caddy.Context) error {
	switch c.BannedIPs {
	case "":
		return fmt.Errorf("Missing IPs to ban")
	default:
		return nil
	}
}

func (c *CaddyBanIP) Validate() error {
	if c.BannedIPs == "" {
		return fmt.Errorf("no banned ip(s) specified")
	}
	return nil
}

func (c CaddyBanIP) ServeHTTP(w http.ResponseWriter, r *http.Request, next caddyhttp.Handler) error {
	fmt.Println("lol")
	match, _ := regexp.MatchString(c.BannedIPs, r.RemoteAddr)
	if match {
		if c.Message == "" {
			http.Error(w, "Your IP ("+r.RemoteAddr+") is banned from accessing this part or the entirety of this website. (CaddyBanIP from Drivet)", http.StatusForbidden)
			fmt.Println(r.RemoteAddr + " tried to access " + r.RequestURI + ", but they're banned and were served a 403.")
		} else {
			http.Error(w, c.Message, http.StatusForbidden)
			fmt.Println(r.RemoteAddr + " tried to access " + r.RequestURI + ", but they're banned and were served a 403.")
		}
	}
	return next.ServeHTTP(w, r)
}

func (c *CaddyBanIP) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
		if d.NextArg() {
			return d.ArgErr()
		}
		for d.NextBlock(0) {
			switch d.Val() {
			case "banned_ips":
				if !d.NextArg() {
					return d.ArgErr()
				}
				c.BannedIPs = d.Val()
			case "message":
				if !d.NextArg() {
					return d.ArgErr()
				}
				c.Message = d.Val()
			default:
				return d.Errf("theres a bit too many subdirectives here, remove: '%s'", d.Val())
			}
		}
	}
	return nil
}

// parseCaddyfile unmarshals tokens from h into a new Middleware.
func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	var c CaddyBanIP
	err := c.UnmarshalCaddyfile(h.Dispenser)
	return c, err
}

// Interface guards
var (
	_ caddy.Provisioner           = (*CaddyBanIP)(nil)
	_ caddy.Validator             = (*CaddyBanIP)(nil)
	_ caddyhttp.MiddlewareHandler = (*CaddyBanIP)(nil)
	_ caddyfile.Unmarshaler       = (*CaddyBanIP)(nil)
)
