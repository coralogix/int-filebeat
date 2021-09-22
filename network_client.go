package fshttp

import "github.com/elastic/beats/libbeat/publisher"

type NetworkClient interface {
	String() string
	Connect() error
	Close() error
	Publish(publisher.Batch) error
}
