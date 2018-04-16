package chartmuseum

import (
	"time"
)

type (
	// Option allows specifying various settings
	Option func(*options)

	// options specify optional settings
	options struct {
		url         string
		username    string
		password    string
		contextPath string
		timeout     time.Duration
	}
)

// URL specifies the chart repo URL
func URL(url string) Option {
	return func(opts *options) {
		opts.url = url
	}
}

// Username is HTTP basic auth username
func Username(username string) Option {
	return func(opts *options) {
		opts.username = username
	}
}

// Password is HTTP basic auth password
func Password(password string) Option {
	return func(opts *options) {
		opts.password = password
	}
}

// ContextPath is the URL prefix for ChartMuseum installation
func ContextPath(contextPath string) Option {
	return func(opts *options) {
		opts.contextPath = contextPath
	}
}

// Timeout specifies the duration (in seconds) before timing out request
func Timeout(timeout int64) Option {
	return func(opts *options) {
		opts.timeout = time.Duration(timeout) * time.Second
	}
}