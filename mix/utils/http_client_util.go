package utils

import (
	"io"
	"net/http"
	"time"
)

/**
 * @author       weimenghua
 * @time         2024-07-24 10:22
 * @description
 */

type HttpClientUtil struct {
	Client *http.Client
}

func NewHttpClientUtil(timeout time.Duration) *HttpClientUtil {
	return &HttpClientUtil{
		Client: &http.Client{
			Timeout: timeout,
		},
	}
}

func (h *HttpClientUtil) Get(url string, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := h.Client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
