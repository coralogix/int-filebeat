package fshttp

import (
	"net"
	"net/http"
	"time"

	"github.com/elastic/beats/libbeat/outputs"
)

type httpClient struct {
	stats    outputs.Observer
	endpoint string
	client   *http.Client
}

func (h *httpClient) String() string {
	return "fshttp"
}

func (h *httpClient) Connect() error {
	h.client = &http.Client{
		Timeout: 2 * time.Second,
		Transport: &http.Transport{
			Dial: (&net.Dialer{Timeout: 2 * time.Second}).Dial,
		},
	}

	return nil
}

func (h *httpClient) Close() error {
	h.client = nil
	return nil
}
