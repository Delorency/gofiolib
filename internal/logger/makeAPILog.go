package logger

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func GetLogEntry(r *http.Request, status int) string {
	var requestBody map[string]string

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil && err != io.EOF {
		log.Println(err)
		return ""
	}

	bodyBytes, _ := json.Marshal(requestBody)
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	logEntry := APILogEntry{
		Method:      r.Method,
		URL:         r.URL.String(),
		IP:          r.RemoteAddr,
		Headers:     make(map[string]string),
		RequestBody: requestBody,
		Status:      status,
	}
	for key, values := range r.Header {
		logEntry.Headers[key] = values[0]
	}

	jsonBytes, err := json.Marshal(logEntry)
	if err != nil {
		log.Println(err)
		return ""
	}

	return string(jsonBytes)
}
