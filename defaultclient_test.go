//Package oba - One Bus Away Go Api https://onebusaway.org/
// Author: Seth T <setheck@gmail.com>
package oba_test

import (
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
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClient(nil, TestApiKey)
	client.SetBaseURL(server.URL)
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

	client := oba.NewDefaultClient(nil, TestApiKey)
	client.SetBaseURL(server.URL)
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

	client := oba.NewDefaultClient(nil, TestApiKey)
	client.SetBaseURL(server.URL)

	b, err := client.Block(TestID)
	if err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, b.ID, "ID")
	assert.NotEmpty(t, b.Configurations, "Configurations")
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

	client := oba.NewDefaultClient(nil, TestApiKey)
	client.SetBaseURL(server.URL)
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

	client := oba.NewDefaultClient(nil, TestApiKey)
	client.SetBaseURL(server.URL)
	e := client.ReportProblemWithStop(TestID, TestParameters)
	if e != nil {
		t.Error(e)
	}
}

func TestDefaultClient_ReportProblemWithTrip(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClient(nil, TestApiKey)
	client.SetBaseURL(server.URL)
	e := client.ReportProblemWithTrip(TestID, TestParameters)
	if e != nil {
		t.Error(e)
	}
}

func TestDefaultClient_RouteIdsForAgency(t *testing.T) {
	t.Skip("TODO...Last!")
	//contents := RetrieveTestJsonFileContent(t)
	//server := FakeServer(t, contents)
	//defer server.Close()
	//
	//client := oba.NewDefaultClient(nil, TestApiKey)
	//client.SetBaseURL(server.URL)
	//routes, e := client.RouteIdsForAgency(TestID)
	//if e != nil {
	//	t.Error(e)
	//}
	//assert.NotEmpty(t, routes, "RouteIdsForAgency")
}

func TestDefaultClient_Route(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClient(nil, TestApiKey)
	client.SetBaseURL(server.URL)
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

	client := oba.NewDefaultClient(nil, TestApiKey)
	client.SetBaseURL(server.URL)
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

	client := oba.NewDefaultClient(nil, TestApiKey)
	client.SetBaseURL(server.URL)
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

	client := oba.NewDefaultClient(nil, TestApiKey)
	client.SetBaseURL(server.URL)
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

	client := oba.NewDefaultClient(nil, TestApiKey)
	client.SetBaseURL(server.URL)
	shape, e := client.Shape(TestID)
	if e != nil {
		t.Error(e)
	}

	VerifyShape(t, shape)
}

func TestDefaultClient_StopIdsForAgency(t *testing.T) {
	t.Skip("TODO")
}

func TestDefaultClient_Stop(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClient(nil, TestApiKey)
	client.SetBaseURL(server.URL)
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

	client := oba.NewDefaultClient(nil, TestApiKey)
	client.SetBaseURL(server.URL)
	stops, e := client.StopsForLocation(TestParameters)
	if e != nil {
		t.Error(e)
	}

	assert.NotEmpty(t, stops, "Stops")
}

func TestDefaultClient_StopsForRoute(t *testing.T) {
	t.Skip("TODO")
}

func TestDefaultClient_TripDetails(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	server := FakeServer(t, contents)
	defer server.Close()

	client := oba.NewDefaultClient(nil, TestApiKey)
	client.SetBaseURL(server.URL)
	td, e := client.TripDetails(TestID)
	if e != nil {
		t.Error(e)
	}

	VerifyTripDetails(t, td)
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
	trip, e := client.Trip("1")
	if e != nil {
		t.Error(e)
	}

	VerifyTrip(t, trip)
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
