package fshttp

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/elastic/beats/libbeat/publisher"
)

func (h *httpClient) Publish(batch publisher.Batch) error {
	events := batch.Events()
	h.stats.NewBatch(len(events))

	entries := make([]logEntry, 0, len(events))
	for _, event := range events {
		entries = append(entries, buildLogEntry(event))
	}

	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	if err := enc.Encode(entries); err != nil {
		h.stats.Dropped(len(entries))
		batch.Drop()
		return nil
	}

	req, _ := http.NewRequest("POST", h.endpoint, &buf)
	req.Header.Set("content-type", "application/json")
	resp, err := h.client.Do(req)

	if err != nil {
		batch.RetryEvents(events)
	} else {
		h.stats.Acked(len(events))
		batch.ACK()
	}

	// Request/Response cleanup
	defer resp.Body.Close()
	_, _ = ioutil.ReadAll(resp.Body)
	return nil
}
