package oba_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Setheck/oba"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
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

func VerifyUnMarshalling(t *testing.T, data []byte) {
	t.Helper()

	var resp oba.Response
	if err := json.Unmarshal(data, &resp); err != nil {
		t.Error(err)
	}

	//m, err := json.MarshalIndent(resp, "", "  ")
	//if err != nil {
	//	t.Error(err)
	//}
	//
	//m = FixJSON(m) // TODO SetEscapeHTML(false) see https://golang.org/pkg/encoding/json/
	//
	//m = bytes.TrimSpace(m)
	//o := bytes.TrimSpace(data)
	//
	//if bytes.Compare(o, m) != 0 {
	//	t.Error("VerifyMarshalling Failed")
	//	fmt.Println(string(m))
	//}
}

func FixJSON(b []byte) []byte {
	return bytes.Replace(b, []byte("\u0026"), []byte("&"), -1)
}

func VerifyAgency(t *testing.T, a *oba.Agency) {
	t.Helper()
	assert.NotNil(t, a, "Agency")
	assert.NotEmpty(t, a.ID, "AgencyID")
	assert.NotEmpty(t, a.Name, "AgencyName")
	assert.NotEmpty(t, a.URL, "AgencyURL")
	assert.NotEmpty(t, a.TimeZone, "AgencyTimeZone")
	assert.NotNil(t, a.PrivateService, "AgencyPrivateService")
	assert.NotEmpty(t, a.Phone, "AgencyPhone")
	assert.NotEmpty(t, a.FareURL, "AgencyFareURL")
	assert.NotEmpty(t, a.Lang, "AgencyLang")
	assert.NotEmpty(t, a.Email, "AgencyEmail")
	assert.NotEmpty(t, a.Disclaimer, "AgencyDisclaimer")
}

func VerifyRoute(t *testing.T, r *oba.Route) {
	t.Helper()
	assert.NotNil(t, r, "Route")
	assert.NotEmpty(t, r.Color, "Color")
	assert.NotEmpty(t, r.Description, "Description")
	assert.NotEmpty(t, r.LongName, "LongName")
	assert.NotEmpty(t, r.ShortName, "ShortName")
	assert.NotEmpty(t, r.TextColor, "TextColor")
	assert.NotZero(t, r.Type, "Type")
	VerifyAgency(t, &r.Agency)
}

func VerifyStop(t *testing.T, s *oba.Stop) {
	t.Helper()
	assert.NotNil(t, s, "Stop")
	assert.NotNil(t, s.Code, "Code")
	assert.NotNil(t, s.Direction, "Direction")
	assert.NotNil(t, s.ID, "ID")
	assert.NotNil(t, s.Lat, "Lat")
	assert.NotNil(t, s.LocationType, "LocationType")
	assert.NotNil(t, s.Lon, "Lon")
	assert.NotNil(t, s.Name, "Name")
	assert.NotNil(t, s.WheelChairBoarding, "WheelChairBoarding")

	for _, r := range s.Routes {
		VerifyRoute(t, &r)
	}
}

func VerifyStopSchedule(t *testing.T, s *oba.StopSchedule) {
	t.Helper()
	assert.NotNil(t, s, "StopSchedule")
	assert.NotEmpty(t, s.Date, "Date")
	assert.NotEmpty(t, s.StopRouteSchedules, "StopSchedule")
	for _, srs := range s.StopRouteSchedules {
		VerifyStopRouteSchedule(t, &srs)
	}

}

func VerifyStopRouteSchedule(t *testing.T, s *oba.StopRouteSchedule) {
	t.Helper()
	assert.NotNil(t, s, "StopRouteSchedule")
	assert.NotEmpty(t, s.StopRouteDirectionSchedules)
	VerifyRoute(t, &s.Route)
	for _, srds := range s.StopRouteDirectionSchedules {
		VerifyStopRouteDirectionSchedule(t, &srds)
	}
}

func VerifyStopRouteDirectionSchedule(t *testing.T, s *oba.StopRouteDirectionSchedule) {
	t.Helper()
	assert.NotNil(t, s, "StopRouteDirectionSchedule")
	assert.NotEmpty(t, s.TripHeadsign, "TripHeadSign")
	for _, sst := range s.ScheduleStopTimes {
		VerifyScheduleStopTime(t, &sst)
	}
}

func VerifyScheduleStopTime(t *testing.T, s *oba.ScheduleStopTime) {
	t.Helper()
	assert.NotNil(t, s, "ScheduleStopTime")
	assert.NotNil(t, s.ArrivalEnabled, "ArrivalEnabled")
	assert.NotZero(t, s.ArrivalTime, "ArrivalTime")
	assert.NotNil(t, s.DepartureEnabled, "DepartureEnabled")
	assert.NotZero(t, s.DepartureTime, "DepartureTime")
	assert.NotEmpty(t, s.StopHeadsign, "StopHeadsign")
	assert.NotEmpty(t, s.ServiceID, "ServiceID")
	assert.NotEmpty(t, s.TripID, "TripID")
}
