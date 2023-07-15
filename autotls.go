package autotls

import (
	"crypto/tls"
	"github.com/go-kratos/kratos/v2/transport/http"
	"golang.org/x/crypto/acme/autocert"
	"log"
	"time"
)

func newTlsConfig(domains []string, opts ...Option) *tls.Config {
	m := &autocert.Manager{
		Prompt: autocert.AcceptTOS,
	}
	if len(domains) > 0 {
		m.HostPolicy = autocert.HostWhitelist(domains...)
	}
	if dir, err := getCacheDir(); err != nil {
		log.Printf("warning: autocert.NewListener not using a cache: %v", err)
	} else {
		m.Cache = dir
	}

	opt := &options{}
	for _, o := range opts {
		o.apply(opt)
	}
	if opt.email != "" {
		m.Email = opt.email
	}
	// only apply this when bigger or eq 1 min
	if opt.renewBefore >= time.Minute {
		m.RenewBefore = opt.renewBefore
	}

	return m.TLSConfig()
}

func WithAutoTlS(domains []string, opts ...Option) []http.ServerOption {
	serverOpts := make([]http.ServerOption, 0, 2)
	serverOpts = append(serverOpts,
		http.TLSConfig(newTlsConfig(domains, opts...)),
		http.Address(":https"),
	)
	return serverOpts
}
