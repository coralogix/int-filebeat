package fshttp

import (
        "https://github.com/elastic/beats/tree/master/libbeat/beat"
        "https://github.com/elastic/beats/libbeat/beat"
        "https://github.com/elastic/beats/libbeat/common"
        "https://github.com/elastic/beats/libbeat/outputs"
)

func init() {
        outputs.RegisterType("http", newHTTPOutput)
}

func newHTTPOutput(_ outputs.IndexManager, _ beat.Info, stats outputs.Observer, cfg *common.Config) (outputs.Group, error) {
        clients := []outputs.NetworkClient{}
        return outputs.Success(batchSize, retryLimit, clients...)
}

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

func newHTTPOutput(_ outputs.IndexManager, _ beat.Info, stats outputs.Observer, cfg *common.Config) (outputs.Group, error) {
        config := clientConfig{}
        if err := cfg.Unpack(&config); err != nil {
                return outputs.Fail(err)
        }

        clients := make([]outputs.NetworkClient, config.Workers)
        for i := 0; i < config.Workers; i++ {
                clients[i] = &httpClient{
                        stats:    stats,
                        endpoint: config.Endpoint,
                }
        }

        return outputs.Success(confg.BatchSize, config.RetryLimit, clients...)
}
