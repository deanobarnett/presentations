package data_test

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

type hostableURL struct {
	*url.URL
}

func (h *hostableURL) prependHost(link string) string {
	parsedURL, err := url.Parse(link)
	if err != nil {
		return ""
	}

	if parsedURL.Host == "" {
		return h.URL.ResolveReference(parsedURL).String()
	}

	return link
}

// TestValidHostableURL ...
func TestValidHostableURL(t *testing.T) {
	var validURLs = []struct {
		name string
		in   string
		out  string
	}{
		{"no input", "", "http://www.example.com"},
		{"includes host", "http://www.example.com/foo", "http://www.example.com/foo"},
		{"basic path", "/foo", "http://www.example.com/foo"},
		{"nested path", "/foo/bar", "http://www.example.com/foo/bar"},
	}

	url, err := url.Parse("http://www.example.com")
	assert.NoError(t, err)
	h := hostableURL{url}

	for _, tt := range validURLs {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, h.prependHost(tt.in), tt.out)
		})
	}
}
