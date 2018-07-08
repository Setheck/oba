package oba

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func makeGetRequest(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, errors.New("Error making request: " + err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New("Error reading body: " + err.Error())
	}
	return body, nil
}

func handleResponse(r Response) error {
	if r.Code != "200" {
		return fmt.Errorf("Code: %v %v", r.Code, r.Text)
	}
	return nil
}

func unmarshalResponse(data []byte) Response {
	response := Response{}
	err := xml.Unmarshal(data, &response)
	if err != nil {
		log.Fatal("Error unmarshaling ", err)
	}
	return response
}

func requestAndHandle(u, errmsg string) (Response, error) {
	body, err := makeGetRequest(u)
	log.Printf(string(body))
	if err != nil {
		return Response{}, errors.New(errmsg + err.Error())
	}
	response := unmarshalResponse(body)
	if err := handleResponse(response); err != nil {
		return Response{}, errors.New(errmsg + err.Error())
	}
	return response, nil
}
