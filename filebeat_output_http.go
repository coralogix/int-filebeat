package fshttp

// import (
// 	"github.com/elastic/beats/libbeat/beat"
// 	"github.com/elastic/beats/libbeat/common"
// 	"github.com/elastic/beats/libbeat/outputs"
// )

// func init() {
//         outputs.RegisterType("http", newHTTPOutput)
// }

// // func newHTTPOutput(_ outputs.IndexManager, _ beat.Info, stats outputs.Observer, cfg *common.Config) (outputs.Group, error) {
// //         clients := []outputs.NetworkClient{}
// //         return outputs.Success(batchSize, retryLimit, clients...)
// // }

// type clientConfig struct {
//         // Number of worker goroutines publishing log events
//         Workers int `config:"workers" validate:"min=1"`
//         // Max number of events in a batch to send to a single client
//         BatchSize int `config:"batch_size" validate:"min=1"`
//         // Max number of retries for single batch of events
//         RetryLimit int `config:"retry_limit"`
//         // The endpoint our client should be POSTing to
//         Endpoint string `config:"endpoint"`
// }

// func newHTTPOutput(_ outputs.IndexManager, _ beat.Info, stats outputs.Observer, cfg *common.Config) (outputs.Group, error) {
//         config := clientConfig{}
//         config.Workers = 1;
//         config.BatchSize = 1;
//         config.RetryLimit = 1;
//         config.Endpoint = "localhost:8000"

//         if err := cfg.Unpack(&config); err != nil {
//                 return outputs.Fail(err)
//         }

//         clients := make([]outputs.NetworkClient, config.Workers)
//         for i := 0; i < config.Workers; i++ {
//                 clients[i] = &httpClient{
//                         stats:    stats,
//                         endpoint: config.Endpoint,
//                 }
//         }

//         return outputs.Success(config.BatchSize, config.RetryLimit, clients...)

// }

// /////////////////////////////////////////////////
// // HTTP CLIENT
// /////////////////////////////////////////////////

// type NetworkClient interface {
// 	String() string
// 	Connect() error
// 	Close() error
// 	Publish(publisher.Batch) error
// }

// type httpClient struct {
// 	stats    outputs.Observer
// 	endpoint string
// 	client   *http.Client
// }

// func (h *httpClient) String() string {
// 	return "fshttp"
// }

// func (h *httpClient) Connect() error {
// 	h.client = &http.Client{
// 		Timeout:   2 * time.Second,
// 		Transport: &http.Transport{
// 			Dial: (&net.Dialer{Timeout: 2 * time.Second}).Dial,
// 		},
// 	}

// 	return nil
// }

// func (h *httpClient) Close() error {
// 	h.client = nil
// 	return nil
// }

// func (h *httpClient) Publish(batch publisher.Batch) error {
// 	events := batch.Events()
// 	s.stats.NewBatch(len(events))

// 	entries := make([]logEntry, 0, len(events))
// 	for _, event := range events {
// 		entries = append(entries, buildLogEntry(event))
// 	}

// 	var buf bytes.Buffer
// 	enc := json.NewEncoder(&buf)
// 	if err := enc.Encode(entries); err != nil {
// 		h.stats.Dropped(len(entries))
// 		batch.Drop()
// 		return nil
// 	}

// 	req := http.NewRequest("POST", h.endpoint, &buf)
// 	req.Header.Set("content-type", "application/json")
// 	resp, err := h.client.Do(req)

// 	if err != nil {
// 		batch.RetryEvents(events)
// 	} else {
// 		h.stats.Acked(len(events))
// 		batch.ACK()
// 	}

// 	// Request/Response cleanup
// 	defer resp.Body.Close()
// 	_, _ = ioutil.ReadAll(resp.Body)
// 	return nil
// }
