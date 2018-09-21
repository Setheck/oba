//Package oba - One Bus Away Go Api https://onebusaway.org/
// Author: Seth T <setheck@gmail.com>
package oba_test

import (
	"testing"
)

func TestAgenciesWithCoverage(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestAgency(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestArrivalAndDepartureForStop(t *testing.T) {
	t.SkipNow()
	// TODO
	contents := RetrieveTestJsonFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestArrivalsAndDeparturesForStop(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestBlock(t *testing.T) {
	t.SkipNow()
	// TODO
	contents := RetrieveTestJsonFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestCancelAlarm(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestCurrentTime(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestRegisterAlarmForArrivalAndDepartureAtStop(t *testing.T) {
	t.SkipNow()
	// TODO
	contents := RetrieveTestJsonFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestReportProblemWithStop(t *testing.T) {
	t.SkipNow()
	// TODO
	contents := RetrieveTestJsonFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestReportProblemWithTrip(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestRoute(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestRouteIdsForAgency(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestRoutesForLocation(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestScheduleForStop(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestShape(t *testing.T) {
	t.SkipNow()
	// TODO
	contents := RetrieveTestJsonFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestStop(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestStopIdsForAgency(t *testing.T) {
	t.SkipNow()
	// TODO: why is the format completely different?!
	contents := RetrieveTestJsonFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestStopsForLocation(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestStopsForRoute(t *testing.T) {
	t.SkipNow()
	// TODO
	contents := RetrieveTestJsonFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestTrip(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestTripDetails(t *testing.T) {
	t.SkipNow()
	// TODO
	contents := RetrieveTestJsonFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestTripForVehicle(t *testing.T) {
	t.SkipNow()
	// TODO
	contents := RetrieveTestJsonFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestTripsForLocation(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestTripsForRoute(t *testing.T) {
	contents := RetrieveTestJsonFileContent(t)
	VerifyMarshalling(t, contents)
}

func TestVehiclesForAgency(t *testing.T) {
	t.SkipNow()
	// TODO
	contents := RetrieveTestJsonFileContent(t)
	VerifyMarshalling(t, contents)
}
