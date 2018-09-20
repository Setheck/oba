//Package oba - One Bus Away Go Api https://onebusaway.org/
// Author: Seth T <setheck@gmail.com>
package oba_test

import (
	"github.com/Setheck/oba"
	"testing"
)

const (
	DEBUG    = 1
	testdata = "testdata/"
)

func TestDefaultClient_AgenciesWithCoverage(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)

	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClient(nil, "key")
	client.SetBaseURL(server.URL)
	_, e := client.AgenciesWithCoverage()
	if e != nil {
		t.Error(e)
	}

	// TODO: add tests to validate retrieval of objects
}

//func TestDefaultClient_Agency(t *testing.T) {
//	contents := RetrieveTestJsonFileContent(t)
//	VerifyMarshalling(t, contents)
//
//	server := FakeServer(t, contents)
//	defer server.Close()
//
//	client := oba.NewDefaultClient(nil, "key")
//	client.SetBaseURL(server.URL)
//	_, e := client.Agency("1")
//	if e != nil {
//		t.Error(e)
//	}
//
//	// TODO: add tests to validate retrieval of objects
//}
