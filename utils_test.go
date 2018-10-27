package oba_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"

	"github.com/Setheck/oba"
	"github.com/stretchr/testify/assert"
)

const (
	TestDataPath = "testdata/"
)

func FakeServer(t *testing.T, body []byte) *httptest.Server {
	t.Helper()
	handler := http.HandlerFunc(func(r http.ResponseWriter, req *http.Request) {
		r.Write(body)
	})
	return httptest.NewServer(handler)
}

func RetrieveTestJsonFileContent(t *testing.T) []byte {
	t.Helper()
	file := ConvertToFilename(t.Name())
	return ReadFile(t, file)
}

func ReadFile(t *testing.T, file string) []byte {
	t.Helper()
	buf, e := ioutil.ReadFile(fmt.Sprint(TestDataPath, file))
	if e != nil {
		t.Error(e)
	}
	return buf
}

func ConvertToFilename(s string) string {
	if strings.Contains(s, "_") {
		split := strings.Split(s, "_")
		s = split[len(split)-1]
	}
	if strings.HasPrefix(s, "Test") {
		s = strings.TrimPrefix(s, "Test")
	}
	r := regexp.MustCompile("^(?:.)|([A-Z][a-z])")
	matches := r.FindAllStringIndex(s, 10)

	slider := 0
	for _, m := range matches {
		if m[0] > 0 {
			s = fmt.Sprintf("%s-%s", s[0:m[0]+slider], s[m[0]+slider:])
			slider++
		}
	}
	return strings.ToLower(s) + ".json"
}

func VerifyAltUnMarshalling(t *testing.T, data []byte) {
	t.Helper()

	var resp oba.AltResponse
	if err := json.Unmarshal(data, &resp); err != nil {
		t.Error(err)
	}
}

func VerifyUnMarshalling(t *testing.T, data []byte) {
	t.Helper()

	var resp oba.Response
	if err := json.Unmarshal(data, &resp); err != nil {
		t.Error(err)
	}
}

func FixJSON(b []byte) []byte {
	return bytes.Replace(b, []byte("\u0026"), []byte("&"), -1)
}

func VerifyAgency(t *testing.T, a *oba.Agency) {
	t.Helper()
	assert.NotNil(t, a, "Agency")
	assert.NotEmpty(t, a.ID, "Agency - ID")
	assert.NotEmpty(t, a.Name, "Agency - Name")
	assert.NotEmpty(t, a.URL, "Agency - URL")
	assert.NotEmpty(t, a.TimeZone, "Agency - TimeZone")
	assert.NotNil(t, a.PrivateService, "Agency - PrivateService")
	assert.NotEmpty(t, a.Phone, "Agency - Phone")
	assert.NotEmpty(t, a.FareURL, "Agency - FareURL")
	assert.NotEmpty(t, a.Lang, "Agency - Lang")
	assert.NotEmpty(t, a.Email, "Agency - Email")
	assert.NotEmpty(t, a.Disclaimer, "Agency - Disclaimer")
}

func VerifyBlock(t *testing.T, b *oba.Block) {
	t.Helper()
	assert.NotNil(t, b, "Block")
	for _, c := range b.Configurations {
		VerifyBlockConfiguration(t, &c)
	}
}

func VerifyBlockConfiguration(t *testing.T, c *oba.BlockConfiguration) {
	t.Helper()
	assert.NotNil(t, c, "BlockConfiguration")
	assert.NotEmpty(t, c.ActiveServiceIDs, "BlockConfiguration - ActiveServiceIds")
	assert.NotEmpty(t, c.InactiveServiceIDs, "BlockConfiguration - InActiveServiceIds")
	for _, tr := range c.Trips {
		VerifyBlockTrip(t, &tr)
	}
}

func VerifyBlockStopTime(t *testing.T, b *oba.BlockStopTime) {
	t.Helper()
	assert.NotNil(t, b, "BlockStopTime")
	assert.NotZero(t, b.DistanceAlongBlock, "BlockStopTime - DistanceAlongBlock")
	assert.NotZero(t, b.AccumulatedSlackTime, "BlockStopTime - AccumulatedSlackTime")
	assert.NotZero(t, b.BlockSequence, "BlockStopTime - BlockSequence")
	VerifyStopTime(t, &b.StopTime)
}

func VerifyStopTime(t *testing.T, s *oba.StopTime) {
	t.Helper()
	assert.NotNil(t, s, "StopTime")
	assert.NotZero(t, s.ArrivalTime, "StopTime - ArrivalTime")
	assert.NotZero(t, s.DepartureTime, "StopTime - DepartureTime")
	assert.NotZero(t, s.DropOffType, "StopTime - DropOffType")
	assert.NotZero(t, s.PickupType, "StopTime - PickupType")
	assert.NotEmpty(t, s.StopID, "StopTime - StopID")
}

func VerifyBlockTrip(t *testing.T, b *oba.BlockTrip) {
	t.Helper()
	assert.NotNil(t, b, "BlockTrip")
	assert.NotEmpty(t, b.TripID, "BlockTrip - TripID")
	for _, bs := range b.BlockStopTimes {
		VerifyBlockStopTime(t, &bs)
	}
}

func VerifyConsequences(t *testing.T, c *oba.Consequence) {
	assert.NotNil(t, c, "Consequences")
	assert.NotEmpty(t, c.Condition, "Consequences - Condition")
	assert.NotEmpty(t, c.ConditionDetailDiversionPathPoints, "Consequences - ConditionDetailDiversionPathPoints")
	assert.NotEmpty(t, c.ConditionDetailDiversionStopIDs, "Consequences - ConditionDetailDiversionStopIDs")
}

func VerifyRoute(t *testing.T, r *oba.Route) {
	t.Helper()
	assert.NotNil(t, r, "Route")
	assert.NotEmpty(t, r.Color, "Route - Color")
	assert.NotEmpty(t, r.Description, "Route - Description")
	assert.NotEmpty(t, r.LongName, "Route - LongName")
	assert.NotEmpty(t, r.ShortName, "Route - ShortName")
	assert.NotEmpty(t, r.TextColor, "Route - TextColor")
	assert.NotZero(t, r.Type, "Route - Type")
	VerifyAgency(t, &r.Agency)
}

func VerifyStop(t *testing.T, s *oba.Stop) {
	t.Helper()
	assert.NotNil(t, s, "Stop")
	assert.NotNil(t, s.Code, "Stop - Code")
	assert.NotNil(t, s.Direction, "Stop - Direction")
	assert.NotNil(t, s.ID, "Stop - ID")
	assert.NotNil(t, s.Lat, "Stop - Lat")
	assert.NotNil(t, s.LocationType, "Stop - LocationType")
	assert.NotNil(t, s.Lon, "Stop - Lon")
	assert.NotNil(t, s.Name, "Stop - Name")
	assert.NotNil(t, s.WheelChairBoarding, "Stop - WheelChairBoarding")

	for _, r := range s.Routes {
		VerifyRoute(t, &r)
	}
}

func VerifyStopsForRoute(t *testing.T, s *oba.StopsForRoute) {
	t.Helper()
	assert.NotNil(t, s, "StopsForRoute")
	VerifyRoute(t, &s.Route)
	for _, stop := range s.Stops {
		VerifyStop(t, &stop)
	}
}

func VerifyStopSchedule(t *testing.T, s *oba.StopSchedule) {
	t.Helper()
	assert.NotNil(t, s, "StopSchedule")
	assert.NotEmpty(t, s.Date, "StopSchedule - Date")
	for _, scd := range s.StopCalendarDays {
		VerifyStopCalendarDay(t, &scd)
	}
	for _, srs := range s.StopRouteSchedules {
		VerifyStopRouteSchedule(t, &srs)
	}
}

func VerifyStopCalendarDay(t *testing.T, scd *oba.StopCalendarDay) {
	t.Helper()
	assert.NotNil(t, scd, "StopCalendarDay")
	assert.NotEmpty(t, scd.Date, "StopCalendarDay - Date")
	assert.NotEmpty(t, scd.Group, "StopCalendarDay - Group")
}

func VerifyStopRouteSchedule(t *testing.T, s *oba.StopRouteSchedule) {
	t.Helper()
	assert.NotNil(t, s, "StopRouteSchedule")
	assert.NotEmpty(t, s.StopRouteDirectionSchedules, "StopRouteSchedule - StopRouteDirectionSchedules")
	VerifyRoute(t, &s.Route)
	for _, srds := range s.StopRouteDirectionSchedules {
		VerifyStopRouteDirectionSchedule(t, &srds)
	}
}

func VerifyStopRouteDirectionSchedule(t *testing.T, s *oba.StopRouteDirectionSchedule) {
	t.Helper()
	assert.NotNil(t, s, "StopRouteDirectionSchedule")
	assert.NotEmpty(t, s.TripHeadsign, "StopRouteDirectionSchedule - TripHeadSign")
	for _, sst := range s.ScheduleStopTimes {
		VerifyScheduleStopTime(t, &sst)
	}
}

func VerifyShape(t *testing.T, s *oba.Shape) {
	t.Helper()
	assert.NotNil(t, s, "Shape")
	assert.NotEmpty(t, s.Length, "Shape - Length")
	assert.NotEmpty(t, s.Points, "Shape - Points")
}

func VerifySituation(t *testing.T, s *oba.Situation) {
	t.Helper()
	assert.NotNil(t, s, "Situation")
	assert.NotEmpty(t, s.ID, "Situation - ID")
	assert.NotEmpty(t, s.CreationTime, "Situation - CreationTime")
	assert.NotEmpty(t, s.Description, "Situation - Description")
	assert.NotEmpty(t, s.EnvironmentReason, "Situation -EnvironmentReason")
	assert.NotEmpty(t, s.Summary, "Situation - Summary")
	for _, vj := range s.Affects {
		VerifyVehicleJourney(t, &vj)
	}
	for _, c := range s.Consequences {
		VerifyConsequences(t, &c)
	}
}

func VerifyScheduleStopTime(t *testing.T, s *oba.ScheduleStopTime) {
	t.Helper()
	assert.NotNil(t, s, "ScheduleStopTime")
	assert.NotNil(t, s.ArrivalEnabled, "ScheduleStopTime - ArrivalEnabled")
	assert.NotZero(t, s.ArrivalTime, "ScheduleStopTime - ArrivalTime")
	assert.NotNil(t, s.DepartureEnabled, "ScheduleStopTime - DepartureEnabled")
	assert.NotZero(t, s.DepartureTime, "ScheduleStopTime - DepartureTime")
	assert.NotEmpty(t, s.StopHeadsign, "ScheduleStopTime - StopHeadsign")
	assert.NotEmpty(t, s.ServiceID, "ScheduleStopTime - ServiceID")
	assert.NotEmpty(t, s.TripID, "ScheduleStopTime - TripID")
}

func VerifyTripDetails(t *testing.T, td *oba.TripDetails) {
	t.Helper()
	assert.NotNil(t, td, "TripDetails")
	assert.NotEmpty(t, td.Frequency, "TripDetails - Frequency")
	assert.NotZero(t, td.ServiceDate, "TripDetails - ServiceDate")
	assert.NotEmpty(t, td.Status, "TripDetails - Status")
	VerifyTrip(t, &td.Trip)
	for _, s := range td.Situations {
		VerifySituation(t, &s)
	}
}

func VerifyTrip(t *testing.T, tr *oba.Trip) {
	t.Helper()
	assert.NotNil(t, tr, "Trip")
	assert.NotEmpty(t, tr.ID, "Trip - ID")
	assert.NotEmpty(t, tr.BlockID, "Trip - BlockID")
	assert.NotEmpty(t, tr.DirectionID, "Trip - DirectionID")
	assert.NotEmpty(t, tr.RouteID, "Trip - RouteID")
	assert.NotEmpty(t, tr.RouteShortName, "Trip - RouteShortName")
	assert.NotEmpty(t, tr.ServiceID, "Trip - ServiceID")
	assert.NotEmpty(t, tr.ShapeID, "Trip - ShapeID")
	assert.NotEmpty(t, tr.TimeZone, "Trip - TimeZone")
	assert.NotEmpty(t, tr.TripHeadsign, "Trip - TripHeadsign")
	assert.NotEmpty(t, tr.TripShortName, "Trip - TripShortName")
}

func VerifyVehicleJourney(t *testing.T, vj *oba.VehicleJourney) {
	t.Helper()
	assert.NotNil(t, vj, "VehicleJourney")
	assert.NotEmpty(t, vj.CallStopIDs, "VehicleJourney - CallStopIDs")
	assert.NotEmpty(t, vj.Direction, "VehicleJourney - Direction")
	assert.NotEmpty(t, vj.LineID, "VehicleJourney - LineID")
}

func VerifyVehicleStatus(t *testing.T, vs *oba.VehicleStatus) {
	t.Helper()
	assert.NotNil(t, vs, "VehicleStatus")
	assert.NotZero(t, vs.LastLocationUpdateTime, "VehicleStatus - LastLocationUpdateTime")
	assert.NotZero(t, vs.LastUpdateTime, "VehicleStatus - LastUpdateTime")
	assert.NotEmpty(t, vs.Phase, "VehicleStatus - Phase")
	assert.NotEmpty(t, vs.Status, "VehicleStatus - Status")
	VerifyLocation(t, &vs.Location)
	VerifyTrip(t, &vs.Trip)
	VerifyTripStatus(t, &vs.TripStatus)
}

func VerifyLocation(t *testing.T, l *oba.Location) {
	t.Helper()
	assert.NotNil(t, l, "Location")
	assert.NotZero(t, l.Lon, "Location - Lon")
	assert.NotZero(t, l.Lat, "Location - Lat")
}

func VerifyTripStatus(t *testing.T, ts *oba.TripStatus) {
	t.Helper()
	assert.NotNil(t, ts, "TripStatus")
	assert.NotEmpty(t, ts.ActiveTripID, "TripStatus - ActiveTripID")
	assert.NotZero(t, ts.BlockTripSequence, "TripStatus - BlockTripSequence")
	VerifyStop(t, &ts.ClosestStop)
	assert.NotZero(t, ts.ClosestStopTimeOffset, "TripStatus - ClosestStopTimeOffset")
	assert.NotZero(t, ts.DistanceAlongTrip, "TripStatus - DistanceAlongTrip")
	assert.NotEmpty(t, ts.Frequency, "TripStatus - Frequency")
	assert.NotZero(t, ts.LastKnownDistanceAlongTrip, "TripStatus - LastKnownDistanceAlongTrip")
	assert.NotZero(t, ts.LastKnownOrientation, "TripStatus - LastKnownOrientation")
	VerifyLocation(t, &ts.LastKnownLocation)
	assert.NotZero(t, ts.LastLocationUpdateTime, "TripStatus - LastLocationUpdateTime")
	assert.NotZero(t, ts.LastUpdateTime, "TripStatus - LastUpdateTime")
	VerifyStop(t, &ts.NextStop)
	assert.NotZero(t, ts.NextStopTimeOffset, "TripStatus - NextStopTimeOffset")
	assert.NotEmpty(t, ts.Phase, "TripStatus - Phase")
	assert.NotNil(t, ts.Predicted, "TripStatus - Predicted")
	assert.NotZero(t, ts.ScheduleDeviation, "TripStatus - ScheduleDeviation")
	assert.NotZero(t, ts.ServiceDate, "TripStatus - ServiceDate")
	assert.NotEmpty(t, ts.Status, "TripStatus  - Status")
	assert.NotEmpty(t, ts.VehicleID, "TripStatus - VehicleID")
	assert.NotEmpty(t, ts.SituationIDs, "TripStatus - SituationIDs")
}
