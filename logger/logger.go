package logger

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"strings"
	"sync"
	"time"
)

type saasBatchHandler struct {
	url    string
	appKey string
	client *http.Client
	mu     sync.Mutex
	buffer []string
}

func NewSaaSBatchHandler(url, appKey string) *saasBatchHandler {
	return &saasBatchHandler{
		url:    url,
		appKey: appKey,
		client: &http.Client{Timeout: 5 * time.Second},
	}
}

func (h *saasBatchHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return true
}

func (h *saasBatchHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *saasBatchHandler) WithGroup(name string) slog.Handler {
	return h
}

func (h *saasBatchHandler) Handle(ctx context.Context, rec slog.Record) error {
	var rec2 slog.Record
	if cid, ok := ctx.Value("correlation_id").(string); ok {
		rec2 = rec.Clone()
		rec2.AddAttrs(slog.String(string("correlation_id"), cid))
	} else {
		rec2 = rec
	}

	buf := &bytes.Buffer{}
	jsonH := slog.NewJSONHandler(buf, &slog.HandlerOptions{AddSource: true})
	if err := jsonH.Handle(ctx, rec2); err != nil {
		return err
	}

	line := strings.TrimSpace(buf.String())

	h.mu.Lock()
	h.buffer = append(h.buffer, line)
	h.mu.Unlock()
	return nil
}

func (h *saasBatchHandler) Flush() error {
	h.mu.Lock()
	logs := h.buffer
	h.buffer = nil
	h.mu.Unlock()

	if len(logs) == 0 {
		log.Println("no logs to flush")
		return nil
	}

	payload := map[string]interface{}{
		"logs":     logs,
		"logsType": "json",
	}
	body, err := json.Marshal(payload)
	if err != nil {
		log.Println("error marshalling logs:", err)
		return err
	}

	req, err := http.NewRequest("POST", h.url, bytes.NewBuffer(body))
	if err != nil {
		log.Println("error creating request:", err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-app-key", h.appKey)

	resp, err := h.client.Do(req)
	if err != nil {
		log.Println("error sending request:", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		log.Println("error response from server:", resp.Status)
		return fmt.Errorf("batch webhook returned %s", resp.Status)
	}

	log.Println("logs flushed successfully")
	return nil
}
