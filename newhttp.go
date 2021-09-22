package fshttp

// import (
// 	"github.com/elastic/beats/libbeat/beat"
// 	"github.com/elastic/beats/libbeat/common"
// 	"github.com/elastic/beats/libbeat/outputs"
// )

// func newHTTPOutput(_ outputs.IndexManager, _ beat.Info,	stats outputs.Observer, cfg *common.Config) (outputs.Group, error) {
// 	config := clientConfig{}
// 	if err := cfg.Unpack(&config); err != nil {
// 		return outputs.Fail(err)
// 	}

// 	clients := make([]outputs.NetworkClient, config.Workers)
// 	for i := 0; i < config.Workers; i++ {
// 		clients[i] = &httpClient{
// 			stats:    stats,
// 			endpoint: config.Endpoint,
// 		}
// 	}

// 	return outputs.Success(config.BatchSize, config.RetryLimit, clients...)
// }
