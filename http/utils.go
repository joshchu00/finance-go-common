package http

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func Get(url string, referer string) (bytes []byte, err error) {

	client := &http.Client{}

	var request *http.Request
	request, err = http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return
	}

	request.Header.Add("Referer", referer)

	var response *http.Response
	response, err = client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	bytes, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	if 200 != response.StatusCode {
		err = errors.New("Not 200")
		return
	}

	return
}
