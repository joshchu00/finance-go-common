package http

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func do(method string, url string, headers map[string]string, body []byte) (out []byte, err error) {

	client := &http.Client{}

	var bodyIOReader io.Reader
	if body != nil {
		bodyIOReader = bytes.NewBuffer(body)
	}

	var request *http.Request
	request, err = http.NewRequest(method, url, bodyIOReader)
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
		err = fmt.Errorf("Not 200: %d", response.StatusCode)
		return
	}

	return
}

type Method func(string, map[string]string, []byte) ([]byte, error)

func Get(url string, headers map[string]string, body []byte) ([]byte, error) {
	return do(http.MethodGet, url, headers, body)
}

func Post(url string, headers map[string]string, body []byte) ([]byte, error) {
	return do(http.MethodPost, url, headers, body)
}

func Put(url string, headers map[string]string, body []byte) ([]byte, error) {
	return do(http.MethodPut, url, headers, body)
}

func Delete(url string, headers map[string]string, body []byte) ([]byte, error) {
	return do(http.MethodDelete, url, headers, body)
}
