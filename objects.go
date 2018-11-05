package oba

import (
	"encoding/json"
)

type ArrivalsAndDepartures []ArrivalAndDeparture

type ArrivalAndDeparture struct {
	ArrivalEnabled               *bool
	BlockTripSequence            int
	DepartureEnabled             *bool
	DistanceFromStop             float64
	Frequency                    *string
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

func (a ArrivalAndDeparture) String() string {
	return jsonStringer(a)
}

type Histogram struct {
	Counts   []int
	Labels   []string
	MaxCount int
	Values   []float64
}

func (h Histogram) String() string {
	return jsonStringer(h)
}

type BlockConfiguration struct {
	ActiveServiceIDs   []string
	InactiveServiceIDs []string
	Trips              []BlockTrip
}

func (b BlockConfiguration) String() string {
	return jsonStringer(b)
}

type BlockStopTime struct {
	AccumulatedSlackTime float64
	BlockSequence        int
	DistanceAlongBlock   float64
	StopTime             StopTime
}

func (b BlockStopTime) String() string {
	return jsonStringer(b)
}

type BlockTrip struct {
	TripID         string
	BlockStopTimes []BlockStopTime
}

func (b BlockTrip) String() string {
	return jsonStringer(b)
}

type StopTime struct {
	StopID        string
	ArrivalTime   int
	DepartureTime int
	PickupType    int
	DropOffType   int
}

func (s StopTime) String() string {
	return jsonStringer(s)
}

type Frequency struct {
	StartTime int
	EndTime   int
	Headway   int
}

func (f Frequency) String() string {
	return jsonStringer(f)
}

type VehicleStatus struct {
	VehicleID              string
	LastUpdateTime         int
	LastLocationUpdateTime int
	Location               Location
	Phase                  string
	Status                 string
	Trip                   Trip
	TripStatus             TripStatus
}

func (v VehicleStatus) String() string {
	return jsonStringer(v)
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
	return jsonStringer(a)
}

type Block struct {
	ID             string
	Configurations []BlockConfiguration
}

func (b Block) String() string {
	return jsonStringer(b)
}

type AgencyWithCoverage struct {
	Agency  Agency
	Lat     float64
	LatSpan float64
	Lon     float64
	LonSpan float64
}

func (a AgencyWithCoverage) String() string {
	return jsonStringer(a)
}

type Coverage struct {
	AgencyID string
	Lat      float64
	LatSpan  float64
	Lon      float64
	LonSpan  float64
}

func (c Coverage) String() string {
	return jsonStringer(c)
}

type CurrentTime struct {
	ReadableTime string
	Time         int
}

func (c CurrentTime) String() string {
	return jsonStringer(c)
}

// Route object
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
	return jsonStringer(r)
}

type RegisteredAlarm struct {
	AlarmID string
}

func (r RegisteredAlarm) String() string {
	return jsonStringer(r)
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

func (s Situation) String() string {
	return jsonStringer(s)
}

type Consequence struct {
	Condition                          string
	ConditionDetailDiversionPathPoints []string
	ConditionDetailDiversionStopIDs    []string
}

func (c Consequence) String() string {
	return jsonStringer(c)
}

type VehicleJourney struct {
	LineID      string
	Direction   string
	CallStopIDs []string
}

func (v VehicleJourney) String() string {
	return jsonStringer(v)
}

type Shape struct {
	Points string
	Length int
}

func (s Shape) String() string {
	return jsonStringer(s)
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

func (s Stop) String() string {
	return jsonStringer(s)
}

type StopSchedule struct {
	Date               int
	Stop               Stop
	StopRouteSchedules []StopRouteSchedule
	TimeZone           string
	StopCalendarDays   []StopCalendarDay
}

func (s StopSchedule) String() string {
	return jsonStringer(s)
}

type StopCalendarDay struct {
	Date  string
	Group string
}

func (s StopCalendarDay) String() string {
	return jsonStringer(s)
}

type StopRouteSchedule struct {
	Route                       Route
	StopRouteDirectionSchedules []StopRouteDirectionSchedule
}

func (s StopRouteSchedule) String() string {
	return jsonStringer(s)
}

type StopRouteDirectionSchedule struct {
	ScheduleFrequencies []ScheduleFrequency
	ScheduleStopTimes   []ScheduleStopTime
	TripHeadsign        string
}

func (s StopRouteDirectionSchedule) String() string {
	return jsonStringer(s)
}

type ScheduleFrequency struct {
	*Frequency
}

func (s ScheduleFrequency) String() string {
	return jsonStringer(s)
}

type StopsForRoute struct {
	Route         Route
	Stops         []Stop
	StopGroupings []StopGrouping
}

func (s StopsForRoute) String() string {
	return jsonStringer(s)
}

type StopGrouping struct {
	Type       string
	Ordered    *bool
	StopGroups []StopGroup
}

func (s StopGrouping) String() string {
	return jsonStringer(s)
}

type StopGroup struct {
	ID        string
	Name      Name
	Stops     []Stop
	PolyLines []EncodedPolyLine
}

func (s StopGroup) String() string {
	return jsonStringer(s)
}

type Name struct {
	Name  string
	Names []string
	Type  string
}

func (n Name) String() string {
	return jsonStringer(n)
}

type EncodedPolyLine struct {
	Length int
	Levels string
	Points string
}

func (e EncodedPolyLine) String() string {
	return jsonStringer(e)
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

func (s ScheduleStopTime) String() string {
	return jsonStringer(s)
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

func (t Trip) String() string {
	return jsonStringer(t)
}

type TripDetails struct {
	Trip        Trip
	ServiceDate int
	Frequency   *string
	Status      string
	Situations  []Situation
}

func (t TripDetails) String() string {
	return jsonStringer(t)
}

type TripStatus struct {
	ActiveTripID               string
	BlockTripSequence          int
	ClosestStop                Stop
	ClosestStopTimeOffset      int
	DistanceAlongTrip          float64
	Frequency                  *string
	LastKnownDistanceAlongTrip float64
	LastKnownLocation          Location
	LastKnownOrientation       int
	LastLocationUpdateTime     int
	LastUpdateTime             int
	NextStop                   Stop
	NextStopTimeOffset         int
	Orientation                float64
	Phase                      string
	Position                   Location
	Predicted                  *bool
	ScheduleDeviation          int
	ScheduledDistanceAlongTrip float64
	ServiceDate                int
	SituationIDs               []string
	Status                     string
	TotalDistanceAlongTrip     float64
	VehicleID                  string
}

func (t TripStatus) String() string {
	return jsonStringer(t)
}

type Location struct {
	Lat float64
	Lon float64
}

func (l Location) String() string {
	return jsonStringer(l)
}

type StopWithArrivalsAndDepartures struct {
	StopID                string
	ArrivalsAndDepartures ArrivalsAndDepartures
	NearByStopIDs         []string
}

func (s StopWithArrivalsAndDepartures) String() string {
	return jsonStringer(s)
}

func jsonStringer(i interface{}) string {
	s, err := json.Marshal(i)
	if err != nil {
		panic(err)
	}
	return string(s)
}
