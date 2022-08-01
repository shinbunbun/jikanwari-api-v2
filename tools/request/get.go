package request

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func createNewHttpRequest(url string) (*http.Request, error) {
	return http.NewRequest(http.MethodGet, url, nil)
}

func setHttpHeader(req *http.Request, header map[string]string) {
	for k, v := range header {
		req.Header.Set(k, v)
	}
}

func doRequest(req *http.Request) (*http.Response, error) {
	client := new(http.Client)

	resp, err := client.Do(req)
	if err != nil {
		return &http.Response{}, err
	}

	return resp, nil
}

func readBody(resp *http.Response) ([]byte, error) {
	return ioutil.ReadAll(resp.Body)
}

func byteArrToString(byteArr []byte) string {
	return string(byteArr)
}

func GetRequest(url string, header map[string]string) (string, error) {

	req, err := createNewHttpRequest(url)
	if err != nil {
		return "", err
	}

	setHttpHeader(req, header)

	resp, err := doRequest(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	respBody, err := readBody(resp)
	if err != nil {
		return "", err
	}

	stringBody := byteArrToString(respBody)

	if resp.StatusCode != 200 {
		return "", errors.New(stringBody)
	}

	return stringBody, nil
}
