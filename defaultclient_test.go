//Package oba - One Bus Away Go Api https://onebusaway.org/
// Author: Seth T <setheck@gmail.com>
package oba_test

import (
	"github.com/Setheck/oba"
	"testing"
)

const (
	DEBUG    = 0
	testdata = "testdata/"
)

func TestAgenciesWithCoverage(t *testing.T) {
	contents := RetrieveTestXmlFileContent(t)
	VerifyMarshalling(t, contents)

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

func TestAgency(t *testing.T) {
	contents := RetrieveTestXmlFileContent(t)
	VerifyMarshalling(t, contents)

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

func TestArrivalAndDepartureForStop(t *testing.T) {
	t.SkipNow()
	// TODO
	contents := RetrieveTestXmlFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestArrivalsAndDeparturesForStop(t *testing.T) {
	contents := RetrieveTestXmlFileContent(t)
	VerifyMarshalling(t, contents)

	//server := FakeServer(t, contents)
	//defer server.Close()
	//
	//client := oba.NewDefaultClient(nil, "key")
	//client.SetBaseURL(server.URL)
	//_, e := client.ArrivalsAndDeparturesForStop("1")
	//if e != nil {
	//	t.Error(e)
	//}
}

func TestBlock(t *testing.T) {
	t.SkipNow()
	// TODO
	contents := RetrieveTestXmlFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestCancelAlarm(t *testing.T) {
	contents := RetrieveTestXmlFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestCurrentTime(t *testing.T) {
	contents := RetrieveTestXmlFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestRegisterAlarmForArrivalAndDepartureAtStop(t *testing.T) {
	contents := RetrieveTestXmlFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestReportProblemWithStop(t *testing.T) {
	t.SkipNow()
	// TODO
	contents := RetrieveTestXmlFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestReportProblemWithTrip(t *testing.T) {
	contents := RetrieveTestXmlFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestRoute(t *testing.T) {
	contents := RetrieveTestXmlFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestRouteIdsForAgency(t *testing.T) {
	contents := RetrieveTestXmlFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestRoutesForLocation(t *testing.T) {
	contents := RetrieveTestXmlFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestScheduleForStop(t *testing.T) {
	contents := RetrieveTestXmlFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestShape(t *testing.T) {
	contents := RetrieveTestXmlFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestStop(t *testing.T) {
	contents := RetrieveTestXmlFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestStopIdsForAgency(t *testing.T) {
	contents := RetrieveTestXmlFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestStopsForLocation(t *testing.T) {
	contents := RetrieveTestXmlFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestStopsForRoute(t *testing.T) {
	t.SkipNow()
	// TODO
	contents := RetrieveTestXmlFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestTrip(t *testing.T) {
	contents := RetrieveTestXmlFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestTripDetails(t *testing.T) {
	t.SkipNow()
	// TODO
	contents := RetrieveTestXmlFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestTripForVehicle(t *testing.T) {
	t.SkipNow()
	// TODO
	contents := RetrieveTestXmlFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestTripsForLocation(t *testing.T) {
	contents := RetrieveTestXmlFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestTripsForRoute(t *testing.T) {
	contents := RetrieveTestXmlFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestVehiclesForAgency(t *testing.T) {
	t.SkipNow()
	// TODO
	contents := RetrieveTestXmlFileContent(t)
	VerifyMarshalling(t, contents)
}
