package util

import (
	"fmt"
	"github.com/cdnetworks-api/cdnetworks-sdk-go/common/model"
	"log"
	"net/http"
	"strings"
)

func Call(requestMsg model.HttpRequestMsg) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(requestMsg.Method, requestMsg.Url, strings.NewReader(requestMsg.Body))
	if err != nil {
		log.Printf("Error: %v", err)
	}
	for key := range requestMsg.Headers {
		req.Header[key] = []string{requestMsg.Headers[key]}
	}
	resp, err := client.Do(req)

	// Check if the response status code is not in the 2xx range
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return resp, fmt.Errorf("unexpected HTTP status code: %d", resp.StatusCode)
	}

	if err != nil {
		log.Printf("Error: %v", err)
	}
	return resp, nil
}
