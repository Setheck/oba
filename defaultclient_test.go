//Package oba - One Bus Away Go Api https://onebusaway.org/
// Author: Seth T <setheck@gmail.com>
package oba_test

import (
	"github.com/Setheck/oba"
	"testing"
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

func TestDefaultClient_Agency(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClient(nil, "key")
	client.SetBaseURL(server.URL)
	_, e := client.Agency("1")
	if e != nil {
		t.Error(e)
	}

	// TODO: add tests to validate retrieval of objects
}

func TestDefaultClient_CurrentTime(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClient(nil, "key")
	client.SetBaseURL(server.URL)
	_, e := client.CurrentTime()
	if e != nil {
		t.Error(e)
	}
}

func TestDefaultClient_CancelAlarm(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClient(nil, "key")
	client.SetBaseURL(server.URL)
	e := client.CancelAlarm("1")
	if e != nil {
		t.Error(e)
	}
}

func TestDefaultClient_TripsForRoute(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClient(nil, "key")
	client.SetBaseURL(server.URL)
	_, e := client.TripsForRoute("1")
	if e != nil {
		t.Error(e)
	}
}

func TestDefaultClient_Trip(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClient(nil, "key")
	client.SetBaseURL(server.URL)
	_, e := client.Trip("1")
	if e != nil {
		t.Error(e)
	}
}
