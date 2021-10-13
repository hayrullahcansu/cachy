package utility

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//IsNullOrEmpty checks input value is empty which is string
func IsNullOrEmpty(data string) bool {
	return len(strings.TrimSpace(data)) <= 0
}

//IsNullOrEmptyB checks input value is empty which is bytes
func IsNullOrEmptyB(data []byte) bool {
	return len(strings.TrimSpace(string(data))) <= 0
}

// FormatRequest generates ascii representation of a request
func FormatRequest(r *http.Request) string {
	// Create return string
	var request []string
	// Add the request string
	url := fmt.Sprintf("%v %v %v", r.Method, r.URL, r.Proto)
	request = append(request, url)
	// Add the host
	request = append(request, fmt.Sprintf("Host: %v", r.Host))
	// Loop through headers
	for name, headers := range r.Header {
		name = strings.ToLower(name)
		for _, h := range headers {
			request = append(request, fmt.Sprintf("%v: %v", name, h))
		}
	}

	// If this is a POST, add post data
	if r.Method == "POST" {
		r.ParseForm()
		request = append(request, "\n")
		request = append(request, r.Form.Encode())
	}
	// Return the request as a string
	return strings.Join(request, "\n")
}

// ReadBody returns body of a request as string
func ReadBody(r *http.Request) (string, error) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	return string(bytes[:]), nil
}

// ReadBody returns body of a request as string
func ReadBodyAsBytes(r *http.Request) ([]byte, error) {
	bytes, err := ioutil.ReadAll(r.Body)
	return bytes, err
}

//ToJson returns string as formatted json
func ToJson(d interface{}) string {
	msg, err := json.Marshal(d)
	if err == nil {
		return string(msg[:])
	} else {
		return ""
	}
}
