package fshttp

type clientConfig struct {
	// Number of worker goroutines publishing log events
	Workers int `config:"workers" validate:"min=1"`
	// Max number of events in a batch to send to a single client
	BatchSize int `config:"batch_size" validate:"min=1"`
	// Max number of retries for single batch of events
	RetryLimit int `config:"retry_limit"`
	// The endpoint our client should be POSTing to
	Endpoint string `config:"endpoint"`
}
