package oba

import "fmt"

type ArrivalsAndDepartures []ArrivalAndDeparture

type ArrivalAndDeparture struct {
	ArrivalEnabled               *bool
	BlockTripSequence            int
	DepartureEnabled             *bool
	DistanceFromStop             float64
	Frequency                    string
	LastUpdateTime               int
	NumberOfStopsAway            int
	Predicted                    *bool
	PredictedArrivalInterval     int
	PredictedArrivalTime         int
	PredictedDepartureInterval   int
	PredictedDepartureTime       int
	ScheduleDeviationHistogramID string
	RouteID                      string
	RouteLongName                string
	RouteShortName               string
	ScheduledArrivalInterval     int
	ScheduledArrivalTime         int
	ScheduledDepartureInterval   int
	ScheduledDepartureTime       int
	ServiceDate                  int
	SituationIDs                 []string
	Status                       string
	StopID                       string
	StopSequence                 int
	TotalStopsInTrip             int
	TripHeadSign                 string
	TripID                       string
	TripStatus                   *TripStatus
	VehicleID                    string
}

type Histogram struct {
	Counts   []int     `json:"counts"`
	Labels   []string  `json:"labels"`
	MaxCount int       `json:"maxCount"`
	Values   []float64 `json:"values"`
}

type BlockConfiguration struct {
	ActiveServiceIDs   []string    `json:"activeServiceIds>string"`
	InactiveServiceIDs []string    `json:"inactiveServiceIds>string"`
	Trips              []BlockTrip `json:"trips"`
}

type BlockStopTime struct {
	AccumulatedSlackTime float64
	BlockSequence        int
	DistanceAlongBlock   float64
	StopTime             StopTime
}

type BlockTrip struct {
	TripID               string          `json:"tripId"`
	BlockStopTimes       []BlockStopTime `json:"blockStopTimes"`
	AccumulatedSlackTime float64         `json:"accumulatedSlackTime"`
	DistanceAlongBlock   float64         `json:"distanceAlongBlock"`
}

type StopTime struct {
	StopID        string
	ArrivalTime   int
	DepartureTime int
	PickupType    int
	DropOffType   int
}

type Frequency struct {
	StartTime int
	EndTime   int
	Headway   int
}

type VehicleStatus struct {
	VehicleID              string      `json:"vehicleId"`
	LastUpdateTime         string      `json:"lastUpdateTime"`
	LastLocationUpdateTime string      `json:"lastLocationUpdateTime"`
	LocationLat            string      `json:"location>lat"`
	LocationLon            string      `json:"location>lon"`
	TripID                 string      `json:"tripId"`
	TripStatus             *TripStatus `json:"tripStatus,omitempty"`
}

// Agency container object
type Agency struct {
	Disclaimer     string
	Email          string
	FareURL        string
	ID             string
	Lang           string
	Name           string
	Phone          string
	PrivateService *bool
	TimeZone       string
	URL            string
}

func (a Agency) String() string {
	return fmt.Sprintf("Disclaimer: %s\nEmail: %s\nFareURL: %s\nID: %s\nLang: %s\nName: %s\nPhone: %s\nPrivateService: %v\nTimeZone: %s\nURL: %s\n",
		a.Disclaimer, a.Email, a.FareURL, a.ID, a.Lang, a.Name, a.Phone, a.PrivateService, a.TimeZone, a.URL)
}

type Block struct {
	ID             string
	Configurations []BlockConfiguration
}

type AgencyWithCoverage struct {
	Agency  Agency
	Lat     float64
	LatSpan float64
	Lon     float64
	LonSpan float64
}

func (a AgencyWithCoverage) String() string {
	return fmt.Sprintf("Agency: %s\nLat: %f\nLatSpan: %f\nLon: %f\nLonSpan: %f\n",
		a.Agency.String(), a.Lat, a.LatSpan, a.Lon, a.LonSpan)
}

type Coverage struct {
	AgencyID string
	Lat      float64
	LatSpan  float64
	Lon      float64
	LonSpan  float64
}

type CurrentTime struct {
	ReadableTime string
	Time         int
}

//Route object
type Route struct {
	Agency      Agency
	Color       string
	Description string
	ID          string
	LongName    string
	ShortName   string
	TextColor   string
	Type        int
	URL         string
}

func (r Route) String() string {
	return fmt.Sprintf("AgencyID: %s\nColor: %s\nDescription: %s\nID: %s\nLongName: %s\nShortName: %s\nTextColor: %s\nType: %d\nURL: %s",
		r.Agency.String(), r.Color, r.Description, r.ID, r.LongName, r.ShortName, r.TextColor, r.Type, r.URL)
}

type RegisteredAlarm struct {
	AlarmID string
}

type Situation struct {
	ID                string
	CreationTime      string
	EnvironmentReason string
	Summary           []string
	Description       []string
	Affects           []VehicleJourney
	Consequences      []Consequence
}

type Consequence struct {
	Condition                          string
	ConditionDetailDiversionPathPoints []string
	ConditionDetailDiversionStopIDs    []string
}

type VehicleJourney struct {
	LineID      string   `json:"lineId"`
	Direction   string   `json:"direction"`
	CallStopIDs []string `json:"calls>call>stopId"`
}

type Shape struct {
	Points string
	Length int
}

type Stop struct {
	Code               string
	Direction          string
	ID                 string
	Lat                float64
	LocationType       int
	Lon                float64
	Name               string
	Routes             []Route
	WheelChairBoarding string
}

type StopSchedule struct {
	Date               int
	Stop               Stop
	StopRouteSchedules []StopRouteSchedule
	TimeZone           string
	StopCalendarDays   []StopCalendarDay
}

type StopCalendarDay struct {
	Date  string `json:"date"`
	Group string `json:"group"`
}

type StopRouteSchedule struct {
	Route                       Route
	StopRouteDirectionSchedules []StopRouteDirectionSchedule
}

type StopRouteDirectionSchedule struct {
	ScheduleFrequencies []ScheduleFrequency `json:"scheduleFrequencies"`
	ScheduleStopTimes   []ScheduleStopTime  `json:"scheduleStopTimes,omitempty"`
	TripHeadsign        string              `json:"tripHeadsign,omitempty"`
}

func NewStopRouteSchedulesFromEntry(r Route, srds []StopRouteDirectionSchedule) *StopRouteSchedule {
	return &StopRouteSchedule{
		Route: r,
		StopRouteDirectionSchedules: srds,
	}
}

type ScheduleFrequency struct {
	*Frequency
}

type StopsForRoute struct {
	Route         Route
	Stops         []Stop
	StopGroupings []StopGrouping
}

type StopGrouping struct {
	Type       int
	Ordered    *bool
	StopGroups []StopGroup `json:"stopGroups>stopGroup,omitempty"`
}

type StopGroup struct {
	ID        string
	Type      int
	Names     []Name
	Stops     []Stop
	PolyLines []EncodedPolyLine
}

type Name struct {
	Name  string
	Names []string
	Type  string
}

type EncodedPolyLine struct {
	Length int
	Levels string
	Points string
}

type ScheduleStopTime struct {
	ArrivalEnabled   *bool
	ArrivalTime      int
	DepartureEnabled *bool
	DepartureTime    int
	ServiceID        string
	StopHeadsign     string
	TripID           string
}

type Trip struct {
	BlockID        string
	DirectionID    string
	ID             string
	RouteID        string
	RouteShortName string
	ServiceID      string
	ShapeID        string
	TimeZone       string
	TripHeadsign   string
	TripShortName  string
}

type TripDetails struct {
	Trip        Trip
	ServiceDate int
	Frequency   string
	Status      string
	Situations  []Situation
}

type TripStatus Entry

type Location struct {
	Lat float64
	Lon float64
}

type StopWithArrivalsAndDepartures struct {
	StopID                string
	ArrivalsAndDepartures ArrivalsAndDepartures
	NearByStopIDs         []string
}
