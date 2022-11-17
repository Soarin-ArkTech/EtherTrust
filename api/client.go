package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type APICallBuilder struct {
	method      string
	url         string
	token       string
	contenttype string
	key         string
	cookie      string
	payload     strings.Reader
}

type APICall struct {
	method      string
	url         string
	token       string
	contenttype string
	key         string
	cookie      string
	payload     *strings.Reader
}

func (call *APICallBuilder) SetMethod(method string) {
	call.method = method
}

func (call *APICallBuilder) SetURL(url string) {
	call.url = url
}

func (call *APICallBuilder) SetToken(token string) {
	call.token = "Bearer " + token
}

func (call *APICallBuilder) SetContentType(contenttype string) {
	call.contenttype = contenttype
}

func (call *APICallBuilder) SetKey(key string) {
	call.key = key
}

func (call *APICallBuilder) SetCookie(cookie string) {
	call.token = cookie
}

func (call *APICallBuilder) SetPayload(payload string) {
	call.payload = *strings.NewReader(payload)
}

func (call *APICallBuilder) Build() APICall {
	return APICall{
		method:      call.method,
		url:         call.url,
		token:       call.token,
		contenttype: call.contenttype,
		key:         call.key,
		cookie:      call.cookie,
		payload:     &call.payload,
	}
}

func (call APICall) Call() (*http.Response, error) {
	req, err := http.NewRequest(call.method, call.url, call.payload)
	if err != nil {
		fmt.Println("Failed to create a new HTTP request. Error: \n", err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", call.contenttype)
	req.Header.Add("Authorization", call.token)
	req.Header.Add("cookie", call.cookie)
	req.Header.Add("key", call.key)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Client failed to send HTTP request. Error:", err)
	}

	return res, err
}

func ParseResults(res *http.Response, Results interface{}) ([]byte, error) {
	response, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Unable to parse API results. Error:\n", err)
		return nil, err
	}
	err = json.Unmarshal(response, &Results)

	return response, err
}
