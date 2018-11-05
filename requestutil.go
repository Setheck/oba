// Package oba - One Bus Away Go Api https://onebusaway.org/
// Author: Seth T <setheck@gmail.com>
package oba

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func makeGetRequest(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New("error making request: " + err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("error reading body: " + err.Error())
	}
	return body, nil
}

func handleResponse(r Response) error {
	if r.Code != http.StatusOK {
		return fmt.Errorf("code: %d %v", r.Code, r.Text)
	}
	return nil
}

func handleAltResponse(r AltResponse) error {
	if r.Code != http.StatusOK {
		return fmt.Errorf("code: %d %v", r.Code, r.Text)
	}
	return nil
}

func unmarshalResponse(data []byte) *Response {
	response := &Response{}
	err := json.Unmarshal(data, response)
	if err != nil {
		panic(err)
	}
	return response
}

func unmarshalAltResponse(data []byte) *AltResponse {
	response := &AltResponse{}
	err := json.Unmarshal(data, response)
	if err != nil {
		panic(err)
	}
	return response
}

func requestAndHandle(u, errmsg string) (*Response, error) {
	body, err := makeGetRequest(u)
	if err != nil {
		return nil, errors.New(errmsg + err.Error())
	}
	response := unmarshalResponse(body)
	if err := handleResponse(*response); err != nil {
		return nil, errors.New(errmsg + err.Error())
	}
	return response, nil
}

func requestAndHandleAlt(u, errmsg string) (*AltResponse, error) {
	body, err := makeGetRequest(u)
	if err != nil {
		return nil, errors.New(errmsg + err.Error())
	}
	response := unmarshalAltResponse(body)
	if err := handleAltResponse(*response); err != nil {
		return nil, errors.New(errmsg + err.Error())
	}
	return response, nil
}
