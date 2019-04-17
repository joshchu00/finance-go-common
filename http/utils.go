package http

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

func do(method string, url string, headers map[string]string, body []byte) (out []byte, err error) {

	client := &http.Client{}

	var bodyBuffer *bytes.Buffer
	bodyBuffer = nil
	if body != nil {
		bodyBuffer = bytes.NewBuffer(body)
	}

	var request *http.Request
	request, err = http.NewRequest(method, url, bodyBuffer)
	if err != nil {
		return
	}

	if headers != nil {
		for k, v := range headers {
			request.Header.Add(k, v)
		}
	}

	var response *http.Response
	response, err = client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	out, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	if 200 != response.StatusCode {
		err = errors.New("Not 200")
		return
	}

	return
}

func Get(url string, headers map[string]string) ([]byte, error) {
	return do(http.MethodGet, url, headers, nil)
}

func Post(url string, headers map[string]string, body []byte) ([]byte, error) {
	return do(http.MethodPost, url, headers, body)
}
