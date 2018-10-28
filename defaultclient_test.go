// Package oba - One Bus Away Go Api https://onebusaway.org/
// Author: Seth T <setheck@gmail.com>
package oba_test

import (
	"net/url"
	"testing"

	"github.com/Setheck/oba"
	"github.com/stretchr/testify/assert"
)

const (
	TestApiKey = "key"
	TestID     = "1"
)

var TestParameters = map[string]string{
	"one": "two",
}

func TestNewDefaultClient(t *testing.T) {
	u := &url.URL{}
	client := oba.NewDefaultClient(u, TestApiKey)
	assert.NotNil(t, client, "NewDefaultClient does not return a valid client")
}

func TestDefaultClient_AgenciesWithCoverage(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClientS(server.URL, TestApiKey)
	awcs, e := client.AgenciesWithCoverage()
	if e != nil {
		t.Error(e)
	}

	for _, awc := range awcs {
		VerifyAgency(t, &awc.Agency)
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

	client := oba.NewDefaultClientS(server.URL, TestApiKey)
	a, e := client.Agency(TestID)
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
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClientS(server.URL, TestApiKey)
	aad, e := client.ArrivalAndDepartureForStop(TestID, TestParameters)
	if e != nil {
		t.Error(e)
	}

	assert.NotZero(t, aad.BlockTripSequence, "BlockTripSequence")
	assert.NotEmpty(t, aad.RouteID, "RouteID")
	assert.NotEmpty(t, aad.TripID, "TripID")
	assert.NotZero(t, aad.ServiceDate, "ServiceDate")
	assert.NotEmpty(t, aad.StopID, "StopId")
	assert.NotZero(t, aad.StopSequence, "StopSequence")
	assert.NotEmpty(t, aad.RouteShortName, "RouteShortName")
	assert.NotEmpty(t, aad.RouteLongName, "RouteLongName")
	assert.NotEmpty(t, aad.TripHeadSign, "TripHeadSign")
	assert.NotNil(t, aad.ArrivalEnabled, "ArrivalEnabled")
	assert.NotNil(t, aad.DepartureEnabled, "DepartureEnabled")
	assert.NotZero(t, aad.ScheduledArrivalTime, "ScheduledArrivalTime")
	assert.NotZero(t, aad.Frequency, "Frequency")
	assert.NotZero(t, aad.Predicted, "Predicted")
	assert.NotZero(t, aad.PredictedArrivalTime, "PredictedArrivalTime")
	assert.NotZero(t, aad.PredictedDepartureTime, "PredictedDepartureTime")
	assert.NotZero(t, aad.DistanceFromStop, "DistanceFromStop")
	assert.NotZero(t, aad.NumberOfStopsAway, "NumberOfStopsAway")
	assert.NotNil(t, aad.TripStatus, "TripStatus")
}

func TestDefaultClient_ArrivalsAndDeparturesForStop(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClientS(server.URL, TestApiKey)
	swaad, e := client.ArrivalsAndDeparturesForStop(TestID, TestParameters)
	if e != nil {
		t.Error(e)
	}

	assert.NotEmpty(t, swaad.StopID, "StopID")
	assert.NotEmpty(t, swaad.ArrivalsAndDepartures, "ArrivalsAndDepartures")
	assert.NotEmpty(t, swaad.NearByStopIDs, "NearByStopIDs")
}

func TestDefaultClient_Block(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClientS(server.URL, TestApiKey)

	b, err := client.Block(TestID)
	if err != nil {
		t.Error(err)
	}

	VerifyBlock(t, b)
	assert.NotEmpty(t, b.ID, "ID")
	assert.NotEmpty(t, b.Configurations, "Configurations")
}

func TestDefaultClient_CancelAlarm(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClientS(server.URL, TestApiKey)
	e := client.CancelAlarm(TestID)
	if e != nil {
		t.Error(e)
	}
}

func TestDefaultClient_CurrentTime(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClientS(server.URL, TestApiKey)
	time, e := client.CurrentTime()
	if e != nil {
		t.Error(e)
	}

	assert.NotEmpty(t, time.ReadableTime, "ReadableTime")
	assert.NotZero(t, time.Time, "Time")
}

func TestDefaultClient_RegisterAlarmForArrivalAndDepartureAtStop(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClientS(server.URL, TestApiKey)
	alarm, e := client.RegisterAlarmForArrivalAndDepartureAtStop(TestID, TestParameters)
	if e != nil {
		t.Error(e)
	}

	assert.NotEmpty(t, alarm.AlarmID)
}

func TestDefaultClient_ReportProblemWithStop(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClientS(server.URL, TestApiKey)

	e := client.ReportProblemWithStop(TestID, TestParameters)
	if e != nil {
		t.Error(e)
	}
}

func TestDefaultClient_ReportProblemWithTrip(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClientS(server.URL, TestApiKey)

	e := client.ReportProblemWithTrip(TestID, TestParameters)
	if e != nil {
		t.Error(e)
	}
}

func TestDefaultClient_RouteIdsForAgency(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClientS(server.URL, TestApiKey)

	routes, e := client.RouteIdsForAgency(TestID)
	if e != nil {
		t.Error(e)
	}
	assert.NotEmpty(t, routes, "RouteIdsForAgency")
}

func TestDefaultClient_Route(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClientS(server.URL, TestApiKey)

	r, e := client.Route(TestID)
	if e != nil {
		t.Error(e)
	}

	VerifyRoute(t, r)
}

func TestDefaultClient_RoutesForAgency(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClientS(server.URL, TestApiKey)

	routes, e := client.RoutesForAgency(TestID)
	if e != nil {
		t.Error(e)
	}

	assert.NotEmpty(t, routes, "Routes")
	for _, r := range routes {
		VerifyRoute(t, &r)
	}
}

func TestDefaultClient_RoutesForLocation(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClientS(server.URL, TestApiKey)

	routes, e := client.RoutesForLocation(TestParameters)
	if e != nil {
		t.Error(e)
	}

	for _, r := range routes {
		VerifyRoute(t, &r)
	}
}

func TestDefaultClient_ScheduleForStop(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClientS(server.URL, TestApiKey)

	ss, e := client.ScheduleForStop(TestID)
	if e != nil {
		t.Error(e)
	}

	VerifyStopSchedule(t, ss)

}

func TestDefaultClient_Shape(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClientS(server.URL, TestApiKey)

	shape, e := client.Shape(TestID)
	if e != nil {
		t.Error(e)
	}

	VerifyShape(t, shape)
}

func TestDefaultClient_StopIdsForAgency(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClientS(server.URL, TestApiKey)

	stops, e := client.StopIDsForAgency(TestID)
	if e != nil {
		t.Error(e)
	}

	assert.NotEmpty(t, stops)
}

func TestDefaultClient_Stop(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClientS(server.URL, TestApiKey)

	stop, e := client.Stop(TestID)
	if e != nil {
		t.Error(e)
	}

	VerifyStop(t, stop)
}

func TestDefaultClient_StopsForLocation(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClientS(server.URL, TestApiKey)

	stops, e := client.StopsForLocation(TestParameters)
	if e != nil {
		t.Error(e)
	}

	assert.NotEmpty(t, stops, "Stops")
}

func TestDefaultClient_StopsForRoute(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClientS(server.URL, TestApiKey)

	sfr, e := client.StopsForRoute(TestID)
	if e != nil {
		t.Error(e)
	}
	VerifyStopsForRoute(t, sfr)
}

func TestDefaultClient_TripDetails(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClientS(server.URL, TestApiKey)

	td, e := client.TripDetails(TestID)
	if e != nil {
		t.Error(e)
	}

	VerifyTripDetails(t, td)
}

func TestDefaultClient_TripForVehicle(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClientS(server.URL, TestApiKey)

	td, e := client.TripForVehicle(TestID, TestParameters)
	if e != nil {
		t.Error(e)
	}

	VerifyTripDetails(t, td)
}

func TestDefaultClient_Trip(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClientS(server.URL, "key")

	trip, e := client.Trip("1")
	if e != nil {
		t.Error(e)
	}

	VerifyTrip(t, trip)
}

func TestDefaultClient_TripsForLocation(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClientS(server.URL, "key")

	tds, e := client.TripsForLocation(TestParameters)
	if e != nil {
		t.Error(e)
	}

	for _, td := range tds {
		VerifyTripDetails(t, &td)
	}
}

func TestDefaultClient_TripsForRoute(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClientS(server.URL, "key")

	tds, e := client.TripsForRoute(TestID)
	if e != nil {
		t.Error(e)
	}

	for _, td := range tds {
		VerifyTripDetails(t, &td)
	}
}

func TestDefaultClient_VehiclesForAgency(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClientS(server.URL, "key")
	vss, e := client.VehiclesForAgency(TestID)
	if e != nil {
		t.Error(e)
	}

	for _, vs := range vss {
		VerifyVehicleStatus(t, &vs)
	}
}
