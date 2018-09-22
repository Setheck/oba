//Package oba - One Bus Away Go Api https://onebusaway.org/
// Author: Seth T <setheck@gmail.com>
package oba_test

import (
	"github.com/Setheck/oba"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	TestApiKey = "key"
	TestID     = "1"
)

var TestParameters = map[string]string{
	"one": "two",
}

func TestDefaultClient_AgenciesWithCoverage(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClient(nil, TestApiKey)
	client.SetBaseURL(server.URL)
	awcs, e := client.AgenciesWithCoverage()
	if e != nil {
		t.Error(e)
	}

	for _, awc := range awcs {
		assert.NotEmpty(t, awc.Agency.ID)
		assert.NotZero(t, awc.Lat)
		assert.NotZero(t, awc.LatSpan)
		assert.NotZero(t, awc.Lon)
		assert.NotZero(t, awc.LonSpan)
	}
}

func TestDefaultClient_Agency(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClient(nil, TestApiKey)
	client.SetBaseURL(server.URL)
	a, e := client.Agency("1")
	if e != nil {
		t.Error(e)
	}

	assert.NotNil(t, a)
	assert.NotEmpty(t, a.Disclaimer, "Disclaimer")
	assert.NotEmpty(t, a.Email, "Email")
	assert.NotEmpty(t, a.FareURL, "FareURL")
	assert.NotEmpty(t, a.ID, "ID")
	assert.NotEmpty(t, a.Lang, "Lang")
	assert.NotEmpty(t, a.Name, "Name")
	assert.NotEmpty(t, a.Phone, "Phone")
	assert.NotNil(t, a.PrivateService, "PrivateService")
	assert.NotEmpty(t, a.TimeZone, "TimeZone")
	assert.NotEmpty(t, a.URL, "URL")
}

func TestDefaultClient_ArrivalAndDepartureForStop(t *testing.T) {
	t.Skip("TODO")

	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClient(nil, TestApiKey)
	client.SetBaseURL(server.URL)
	//aads, e := client.ArrivalsAndDeparturesForStop(TestID, TestParameters)
	//if e != nil {
	//	t.Error(e)
	//}

	//for aad := range aads {
	//
	//}
}

func TestDefaultClient_ArrivalsAndDeparturesForStop(t *testing.T) {
	t.Skip("TODO")
}

func TestDefaultClient_Block(t *testing.T) {
	t.Skip("TODO")
}

func TestDefaultClient_CancelAlarm(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClient(nil, TestApiKey)
	client.SetBaseURL(server.URL)
	e := client.CancelAlarm(TestID)
	if e != nil {
		t.Error(e)
	}
}

func TestDefaultClient_CurrentTime(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClient(nil, TestApiKey)
	client.SetBaseURL(server.URL)
	_, e := client.CurrentTime()
	if e != nil {
		t.Error(e)
	}
}

func TestDefaultClient_RegisterAlarmForArrivalAndDepartureAtStop(t *testing.T) {
	t.Skip("TODO")
}

func TestDefaultClient_ReportProblemWithStop(t *testing.T) {
	t.Skip("TODO")
}

func TestDefaultClient_ReportProblemWithTrip(t *testing.T) {
	t.Skip("TODO")
}

func TestDefaultClient_RouteIdsForAgency(t *testing.T) {
	t.Skip("TODO")
}

func TestDefaultClient_Route(t *testing.T) {
	t.Skip("TODO")
}

func TestDefaultClient_RoutesForAgency(t *testing.T) {
	t.Skip("TODO")
}

func TestDefaultClient_RoutesForLocation(t *testing.T) {
	t.Skip("TODO")
}

func TestDefaultClient_ScheduleForStop(t *testing.T) {
	t.Skip("TODO")
}

func TestDefaultClient_Shape(t *testing.T) {
	t.Skip("TODO")
}

func TestDefaultClient_StopIdsForAgency(t *testing.T) {
	t.Skip("TODO")
}

func TestDefaultClient_Stop(t *testing.T) {
	t.Skip("TODO")
}

func TestDefaultClient_StopsForLocation(t *testing.T) {
	t.Skip("TODO")
}

func TestDefaultClient_StopsForRoute(t *testing.T) {
	t.Skip("TODO")
}

func TestDefaultClient_TripDetails(t *testing.T) {
	t.Skip("TODO")
}

func TestDefaultClient_TripForVehicle(t *testing.T) {
	t.Skip("TODO")
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

func TestDefaultClient_TripsForLocation(t *testing.T) {
	t.Skip("TODO")
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

func TestDefaultClient_VehiclesForAgency(t *testing.T) {
	t.Skip("TODO")
}
