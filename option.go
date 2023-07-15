package autotls

import (
	"time"
)

type options struct {
	renewBefore time.Duration
	email       string
}

type Option interface {
	apply(opt *options)
}

type renewOption time.Duration

func (r renewOption) apply(opt *options) {
	opt.renewBefore = time.Duration(r)
}

// WithRenewBefore RenewBefore optionally specifies how early certificates should be renewed before they expire.
// If zero , they're renewed 30 days before expiration.
func WithRenewBefore(dur time.Duration) Option {
	return renewOption(dur)
}

type emailOption string

func (e emailOption) apply(opt *options) {
	opt.email = string(e)
}

// WithEmail  Email optionally specifies a contact email address.
// This is used by CAs, such as Let's Encrypt, to notify about problems with issued certificates.
func WithEmail(email string) Option {
	return emailOption(email)
}
