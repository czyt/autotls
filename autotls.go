package autotls

import (
	"crypto/tls"
	"github.com/go-kratos/kratos/v2/transport/http"
	"golang.org/x/crypto/acme/autocert"
	"log"
)

func newTlsConfig(domains ...string) *tls.Config {
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
	return m.TLSConfig()
}

func WithAutoTlS(domains ...string) []http.ServerOption {
	opts := make([]http.ServerOption, 0, 2)
	opts = append(opts,
		http.TLSConfig(newTlsConfig(domains...)),
		http.Address(":https"),
	)
	return opts
}
