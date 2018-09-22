//Package oba - One Bus Away Go Api https://onebusaway.org/
// Author: Seth T <setheck@gmail.com>
package oba

import (
	"fmt"
	"log"
	"net/url"
	"path"
)

//Response Element - All responses are wrapped in a response element.
//The response element carries the following fields:
// version - response version information
// code - a machine-readable response code with the following semantics:
// 200 - Success
// 400 - The request could not be understood due to an invalid request parameter or some other error
// 401 - The application key is either missing or invalid
// 404 - The specified resource was not found
// 500 - A service exception or error occurred while processing the request
// text - a human-readable version of the response code
// currentTime - current system time on the api server as milliseconds since the unix epoch
// data - the response payload
// references see the discussion of references below
type Response struct {
	Code        int    `json:"code"`
	CurrentTime int    `json:"currentTime"`
	Data        *Data  `json:"data,omitempty"`
	Text        string `json:"text"`
	Version     int    `json:"version"`
}

func (r Response) String() string {
	return fmt.Sprintf("Code: %d\nCurrentTime: %d\nData: %s\nText: %s\nVersion: %d\n",
		r.Code, r.CurrentTime, r.Data.String(), r.Text, r.Version)
}

//References - The <references/> element contains a dictionary of objects
// referenced by the main result payload. For elements that are
// often repeated in the result payload, the elements are instead
// included in the <references/> section and the payload will refer
// to elements by and object id that can be used to lookup the
// object in the <references/> dictionary.
// They will always appear in that order, since stops and trips reference routes
// and routes reference agencies. If you are processing the result stream in
// order, you should always be able to assume that an referenced entity would
// already have been included in the references section.
// Every API method supports an optional includeReferences=true|false parameter
// that determines if the <references/> section is included in a response. If
// you don’t need the contents of the <references/> section, perhaps because
// you’ve pre-cached all the elements, then setting includeReferences=false can
// be a good way to reduce the response size.
type References struct {
	Agencies   []Agency    `json:"agencies"`
	Routes     []Route     `json:"routes"`
	Situations []Situation `json:"situations"`
	Stops      []Stop      `json:"stops"`
	Trips      []Trip      `json:"trips"`
}

//Data container object
type Data struct {
	LimitExceeded *bool       `json:"limitExceeded,omitempty"`
	List          []*List     `json:"list,omitempty"`
	Entry         *Entry      `json:"entry,omitempty"`
	OutOfRange    *bool       `json:"outOfRange,omitempty"`
	References    *References `json:"references"`
	Time          *Time       `json:",omitempty"`
	//StopsForRoute *StopsForRoute
}

func (d Data) String() string {
	return ""
	// TODO:
	//return fmt.Sprintf("LimitExceeded: %b\nList: %s\nEntry: %s\nOutOfRange: %b\nReference: %s\nTime: %s",
	//	d.LimitExceeded, d.List.String(), d.Entry.String(), d.OutOfRange.String(), d.References.String(), d.Time.String())
}

type Time struct {
	ReadableTime *string `json:"readableTime,omitempty"`
	Time         *int    `json:"time,omitempty"`
}

func (t Time) String() string {
	return fmt.Sprintf("ReadableTime: %s\nTime: %d", *t.ReadableTime, t.Time)
}

type RegisteredAlarm struct {
	AlarmID string `json:"alarmId,omitempty"`
}

//Entry container object
type Entry struct {
	AgencyID             *string                `json:"agencyId,omitempty"`
	ArrivalAndDepartures []*ArrivalAndDeparture `json:"arrivalsAndDepartures,omitempty"`
	BlockID              *string                `json:"blockId,omitempty"`
	Code                 *string                `json:"code,omitempty"`
	Color                *string                `json:"color,omitempty"`
	Date                 *int                   `json:"date,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Direction            *string                `json:"direction,omitempty"`
	DirectionID          *int                   `json:"directionId,omitempty"`
	Disclaimer           *string                `json:"disclaimer,omitempty"`
	Email                *string                `json:"email,omitempty"`
	FareURL              *string                `json:"fareUrl,omitempty"`
	ID                   *string                `json:"id,omitempty"`
	Lang                 *string                `json:"lang,omitempty"`
	Lat                  *float64               `json:"lat,omitempty"`
	LocationType         *int                   `json:"locationType,omitempty"`
	Lon                  *float64               `json:"lon,omitempty"`
	LongName             *string                `json:"longName,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	NearbyStopIds        []*string              `json:"nearbyStopIds,omitempty"`
	Phone                *string                `json:"phone,omitempty"`
	PrivateService       *bool                  `json:"privateService,omitempty"`
	ReadableTime         *string                `json:"readableTime,omitempty"`
	RouteID              *string                `json:"routeId,omitempty"`
	RouteIDs             []*string              `json:"routeIds,omitempty"`
	RouteShortName       *string                `json:"routeShortName,omitempty"`
	ServiceID            *string                `json:"serviceId,omitempty"`
	ShapeID              *string                `json:"shapeId,omitempty"`
	ShortName            *string                `json:"shortName,omitempty"`
	SituationIDs         []*string              `json:"situationIds,omitempty"`
	StopCalendarDays     []*StopCalendarDay     `json:"stopCalendarDays,omitempty"`
	StopID               *string                `json:"stopId,omitempty,omitempty"`
	StopRouteSchedules   []*StopRouteSchedule   `json:"stopRouteSchedules,omitempty"`
	TextColor            *string                `json:"textColor,omitempty"`
	Time                 *int                   `json:"time,omitempty"`
	TimeZone             *string                `json:"timezone,omitempty"`
	TripHeadsign         *string                `json:"tripHeadsign,omitempty"`
	TripShortName        *string                `json:"tripShortName,omitempty"`
	Type                 *int                   `json:"type,omitempty"`
	URL                  *string                `json:"url,omitempty"`
	WheelChairBoarding   *string                `json:"wheelchairBoarding,omitempty"`
	//ID   string `json:"id,omitempty"`
	//Name string `json:"name,omitempty"`
	//URL  string `json:"url,omitempty"`
	//Block
	//Shape
	//StopId        *string   `json:"stopId,omitempty"`
	//*StopsForRoute
	//TripDetails
	//RegisteredAlarm
}

type ArrivalAndDeparture struct {
	ArrivalEnabled             *bool       `json:"arrivalEnabled"`
	BlockTripSequence          *int        `json:"blockTripSequence"`
	DepartureEnabled           *bool       `json:"departureEnabled"`
	DistanceFromStop           *float64    `json:"distanceFromStop"`
	Frequency                  *string     `json:"frequency"`
	LastUpdateTime             *int        `json:"lastUpdateTime"`
	NumberOfStopsAway          *int        `json:"numberOfStopsAway"`
	Predicted                  *bool       `json:"predicted"`
	PredictedArrivalInterval   *int        `json:"predictedArrivalInterval"`
	PredictedArrivalTime       *int        `json:"predictedArrivalTime"`
	PredictedDepartureInterval *int        `json:"predictedDepartureInterval"`
	PredictedDepartureTime     *int        `json:"predictedDepartureTime"`
	RouteID                    *string     `json:"routeId"`
	RouteLongName              *string     `json:"routeLongName"`
	RouteShortName             *string     `json:"routeShortName"`
	ScheduledArrivalInterval   *int        `json:"scheduledArrivalInterval"`
	ScheduledArrivalTime       *int        `json:"scheduledArrivalTime"`
	ScheduledDepartureInterval *int        `json:"scheduledDepartureInterval"`
	ScheduledDepartureTime     *int        `json:"scheduledDepartureTime"`
	ServiceDate                *int        `json:"serviceDate"`
	SituationIDs               []*int      `json:"situationIds"`
	Status                     *string     `json:"status"`
	StopID                     *string     `json:"stopId"`
	StopSequence               *int        `json:"stopSequence"`
	TotalStopsInTrip           *int        `json:"totalStopsInTrip"`
	TripHeadSign               *string     `json:"tripHeadsign"`
	TripID                     *string     `json:"tripId"`
	TripStatus                 *TripStatus `json:"tripStatus"`
	VehicleID                  *string     `json:"vehicleId"`
}

type BlockConfiguration struct {
	ActiveServiceIDs   []string    `json:"activeServiceIds>string"`
	InactiveServiceIDs []string    `json:"inactiveServiceIds>string"`
	Trips              []BlockTrip `json:"trips"`
}

type BlockTrip struct {
	TripID               string     `json:"tripId"`
	BlockStopTimes       []StopTime `json:"blockStopTimes"`
	AccumulatedSlackTime string     `json:"accumulatedSlackTime"`
	DistanceAlongBlock   string     `json:"distanceAlongBlock"`
}

type StopTime struct {
	StopID        string `json:"stopId"`
	ArrivalTime   string `json:"arrivalTime"`
	DepartureTime string `json:"departureTime"`
	PickupType    string `json:"pickupType"`
	DropOffType   string `json:"droppOffType"`
}

type Frequency struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
	Headway   string `json:"headway"`
}

type List struct {
	AgencyID           *string   `json:"agencyId,omitempty"`
	Code               *string   `json:"code,omitempty"`
	Color              *string   `json:"color,omitempty"`
	Description        *string   `json:"description,omitempty"`
	Direction          *string   `json:"direction,omitempty"`
	Disclaimer         *string   `json:"disclaimer,omitempty"`
	Email              *string   `json:"email,omitempty"`
	FareURL            *string   `json:"fareUrl,omitempty"`
	Frequency          *string   `json:"frequency,omitempty"`
	ID                 *string   `json:"id,omitempty"`
	Lang               *string   `json:"lang,omitempty"`
	Lat                *float64  `json:"lat,omitempty"`
	LatSpan            *float64  `json:"latSpan,omitempty"`
	LocationType       *int      `json:"locationType,omitempty"`
	Lon                *float64  `json:"lon,omitempty"`
	LongName           *string   `json:"longName,omitempty"`
	LonSpan            *float64  `json:"lonSpan,omitempty"`
	Name               *string   `json:"name,omitempty"`
	Phone              *string   `json:"phone,omitempty"`
	PrivateService     *bool     `json:"privateService,omitempty"`
	RouteIDs           []*string `json:"routeIds,omitempty"`
	ServiceDate        *int      `json:"serviceDate,omitempty"`
	ShortName          *string   `json:"shortName,omitempty"`
	SituationIDs       []*string `json:"situationIds,omitempty"`
	Status             *string   `json:"status,omitempty"`
	TextColor          *string   `json:"textColor,omitempty"`
	TimeZone           *string   `json:"timezone,omitempty"`
	TripID             *string   `json:"tripId,omitempty"`
	Type               *int      `json:"type,omitempty"`
	URL                *string   `json:"url,omitempty"`
	WheelChairBoarding *string   `json:"wheelchairBoarding,omitempty"`
	//*VehicleStatus
}

func (l List) AgencyWithCoverage() AgencyWithCoverage {
	return AgencyWithCoverage{
		AgencyID: *l.AgencyID,
		Lat:      *l.Lat,
		Lon:      *l.Lon,
		LatSpan:  *l.LatSpan,
		LonSpan:  *l.LonSpan,
	}
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

//Agency container object
type Agency struct {
	Disclaimer     string `json:"disclaimer"`
	Email          string `json:"email"`
	FareURL        string `json:"fareUrl"`
	ID             string `json:"id"`
	Lang           string `json:"lang"`
	Name           string `json:"name"`
	Phone          string `json:"phone"`
	PrivateService bool   `json:"privateService"`
	TimeZone       string `json:"timezone"`
	URL            string `json:"url"`
}

type Block struct {
	Configurations []BlockConfiguration `json:"blockConfiguration,omitempty"`
}

type AgencyWithCoverage struct {
	AgencyID string  `json:"agencyId"`
	Lat      float64 `json:"lat"`
	LatSpan  float64 `json:"latSpan"`
	Lon      float64 `json:"lon"`
	LonSpan  float64 `json:"lonSpan"`
}

func (a AgencyWithCoverage) String() string {
	return fmt.Sprintf("AgencyID: %s\nLat: %f\nLatSpan: %f\nLon: %f\nLonSpan: %f\n",
		a.AgencyID, a.Lat, a.LatSpan, a.Lon, a.LonSpan)
}

type Coverage struct {
	AgencyID *string  `json:"agencyId"`
	Lat      *float64 `json:"lat"`
	LatSpan  *float64 `json:"latSpan"`
	Lon      *float64 `json:"lon"`
	LonSpan  *float64 `json:"lonSpan"`
}

//Route object
type Route struct {
	AgencyID    *string `json:"agencyId"`
	Color       *string `json:"color"`
	Description *string `json:"description"`
	ID          *string `json:"id,omitempty"`
	LongName    *string `json:"longName"`
	ShortName   *string `json:"shortName"`
	TextColor   *string `json:"textColor"`
	Type        *int    `json:"type"`
	URL         *string `json:"url"`
}

func (r Route) String() string {
	return fmt.Sprintf("AgencyID: %s\nColor: %s\nDescription: %s\nID: %s\nLongName: %s\nShortName: %s\nTextColor: %s\nType: %d\nURL: %s",
		*r.AgencyID, *r.Color, *r.Description, *r.ID, *r.LongName, *r.ShortName, *r.TextColor, *r.Type, *r.URL)
}

type Situation struct {
	ID                string           `json:"id"`
	CreationTime      string           `json:"creationTime"`
	EnvironmentReason string           `json:"environmentReason"`
	Summary           []string         `json:"summary>value"`
	Description       []string         `json:"description>value"`
	Affects           []VehicleJourney `json:"vehicleJourneys>vehicleJourney"`
	Consequences      []Consequence    `json:"consequences>consequence"`
}

type Consequence struct {
	Condition                          string   `json:"condition"`
	ConditionDetailDiversionPathPoints []string `json:"conditionDetails>diversionPath>points"`
	ConditionDetailDiversionStopIDs    []string `json:"conditionDetails>diversionStopIds>string"`
}

type VehicleJourney struct {
	LineID      string   `json:"lineId"`
	Direction   string   `json:"direction"`
	CallStopIDs []string `json:"calls>call>stopId"`
}

type Shape struct {
	Points string `json:"points,omitempty"`
	Length string `json:"length,omitempty"`
}

type Stop struct {
	Code               *string   `json:"code"`
	Direction          *string   `json:"direction"`
	ID                 *string   `json:"id"`
	Lat                *float64  `json:"lat"`
	LocationType       *int      `json:"locationType"`
	Lon                *float64  `json:"lon"`
	Name               *string   `json:"name"`
	RouteIDs           []*string `json:"routeIds"`
	WheelChairBoarding *string   `json:"wheelchairBoarding"`
}

type StopSchedule struct {
	Date               *int                 `json:"date,omitempty"`
	StopID             *string              `json:"stopId,omitempty"`
	StopRouteSchedules []*StopRouteSchedule `json:"stopRouteSchedules,omitempty"`
	TimeZone           *string              `json:"timeZone,omitempty"`
	StopCalendarDays   []*StopCalendarDay   `json:"stopCalendarDays,omitempty"`
}

type StopCalendarDay struct {
	Date  string `json:"date"`
	Group string `json:"group"`
}

type StopRouteSchedule struct {
	RouteID                     string                        `json:"routeId,omitempty"`
	StopRouteDirectionSchedules []*StopRouteDirectionSchedule `json:"stopRouteDirectionSchedules,omitempty"`
}

type StopRouteDirectionSchedule struct {
	ScheduleFrequencies []*ScheduleFrequency `json:"scheduleFrequencies"`
	ScheduleStopTimes   []*ScheduleStopTime  `json:"scheduleStopTimes,omitempty"`
	TripHeadsign        *string              `json:"tripHeadsign,omitempty"`
}

type ScheduleFrequency struct {
	*Frequency
}

type StopsForRoute struct {
	StopIDs        []*string `json:"stopIds,omitempty"`
	*StopGroupings `json:"stopGroupings>stopGrouping,omitempty"`
}
type StopGroupings []*StopGrouping

type StopGrouping struct {
	Type       string      `json:"type,omitempty"`
	Ordered    string      `json:"ordered,omitempty"`
	StopGroups []StopGroup `json:"stopGroups>stopGroup,omitempty"`
}

type StopGroup struct {
	ID        string             `json:"id,omitempty"`
	NameType  string             `json:"type,omitempty"`
	Names     []string           `json:"names,omitempty"`
	StopIDs   []*string          `json:"stopIds,omitempty"`
	PolyLines []*EncodedPolyLine `json:"polylines,omitempty"`
}

type EncodedPolyLine struct {
	*Shape
}

type ScheduleStopTime struct {
	ArrivalEnabled   *bool   `json:"arrivalEnabled,omitempty"`
	ArrivalTime      *int    `json:"arrivalTime,omitempty"`
	DepartureEnabled *bool   `json:"departureEnabled,omitempty"`
	DepartureTime    *int    `json:"departureTime,omitempty"`
	ServiceID        *string `json:"serviceId,omitempty"`
	StopHeadsign     *string `json:"stopHeadsign,omitempty"`
	TripID           *string `json:"tripId,omitempty"`
}

type Trip struct {
	BlockID        *string `json:"blockId"`
	DirectionID    *string `json:"directionId"`
	ID             *string `json:"id"`
	RouteID        *string `json:"routeId"`
	RouteShortName *string `json:"routeShortName"`
	ServiceID      *string `json:"serviceId"`
	ShapeID        *string `json:"shapeId"`
	TimeZone       *string `json:"timeZone"`
	TripHeadsign   *string `json:"tripHeadsign"`
	TripShortName  *string `json:"tripShortName"`
}

type TripDetails struct {
	TripID       string   `json:"tripId,omitempty"`
	ServiceDate  string   `json:"serviceDate,omitempty"`
	Frequency    string   `json:"frequency,omitempty"`
	Status       string   `json:"status,omitempty"`
	SituationIDs []string `json:"situationIds,omitempty"`
}

type TripStatus struct {
	ActiveTripID               *string   `json:"activeTripId"`
	BlockTripSequence          *int      `json:"blockTripSequence"`
	ClosestStop                *string   `json:"closestStop"`
	ClosestStopTimeOffset      *int      `json:"closestStopTimeOffset"`
	DistanceAlongTrip          *float64  `json:"distanceAlongTrip"`
	Frequency                  *string   `json:"frequency"`
	LastKnownDistanceAlongTrip *int      `json:"lastKnownDistanceAlongTrip"`
	LastKnownLocation          *Location `json:"lastKnownLocation"`
	LastKnownOrientation       *int      `json:"lastKnownOrientation"`
	LastLocationUpdateTime     *int      `json:"lastLocationUpdateTime"`
	LastUpdateTime             *int      `json:"lastUpdateTime"`
	NextStop                   *string   `json:"nextStop"`
	NextStopTimeOffset         *int      `json:"nextStopTimeOffset"`
	Orientation                *float64  `json:"orientation"`
	Phase                      *string   `json:"phase"`
	Position                   *Location `json:"position"`
	Predicted                  *bool     `json:"predicted"`
	ScheduleDeviation          *int      `json:"scheduleDeviation"`
	ScheduledDistanceAlongTrip *float64  `json:"scheduledDistanceAlongTrip"`
	ServiceDate                *int      `json:"serviceDate"`
	SituationIDs               []*string `json:"situationIds"`
	Status                     *string   `json:"status"`
	TotalDistanceAlongTrip     *float64  `json:"totalDistanceAlongTrip"`
	VehicleID                  *string   `json:"vehicleId"`
}

type Location struct {
	Lat *float64 `json:"lat,omitempty"`
	Lon *float64 `json:"lon,omitempty"`
}

type StopWithArrivalsAndDepartures struct {
	StopID                string                `json:"stopId,omitempty"`
	ArrivalsAndDepartures []ArrivalAndDeparture `json:"arrivalsAndDepartures,omitempty"`
	NearByStopIDs         []*string             `json:"nearbyStopIds,omitempty"`
}

const (
	jsonPostFix                                       = ".json"
	agencyEndPoint                                    = "agency/"
	blockEndPoint                                     = "block/"
	routeEndPoint                                     = "route/"
	shapeEndPoint                                     = "shape/"
	stopEndPoint                                      = "stop/"
	tripEndPoint                                      = "trip/"
	agencyWithCoverageEndPoint                        = "agencies-with-coverage"
	arrivalAndDepartureForStopEndPoint                = "arrival-and-departure-for-stop/"
	arrivalsAndDeparturesForStopEndPoint              = "arrivals-and-departures-for-stop/"
	cancelAlarmEndPoint                               = "cancel_alarm/"
	currentTimeEndPoint                               = "current-time"
	registerAlarmForArrivalAndDepartureAtStopEndPoint = "register-alarm-for-arrival-and-departure-at-stop/"
	reportPoblemWithTripEndPoint                      = "report-problem-with-trip/"
	routeForAgencyEndPoint                            = "routes-for-agency/"
	routeForLocationEndPoint                          = "routes-for-location"
	scheduleForStopEndPoint                           = "schedule-for-stop/"
	stopIDsForAgencyEndPoint                          = "stop-ids-for-agency/"
	stopsForLocationEndPoint                          = "stops-for-location"
	stopsForRouteEndPoint                             = "stops-for-route/"
	tripDetailsEndPoint                               = "trip-details/"
	tripForVehicleEndPoint                            = "trip-for-vehicle/"
	tripsForLocationEndPoint                          = "trips-for-location"
	tripsForRouteEndPoint                             = "trips-for-route/"
	vehiclesForAgencyEndPoint                         = "vehicles-for-agency/"
	//routeIdsForAgencyEndPoint                         = "route-ids-for-agency/"
)

type DefaultClient struct {
	baseURL *url.URL
	apiKey  string
}

//NewDefaultClient - instantiate a new instance of a Client
func NewDefaultClient(u *url.URL, apiKey string) *DefaultClient {
	return &DefaultClient{baseURL: u, apiKey: apiKey}
}

func (c *DefaultClient) SetBaseURL(b string) {
	u, e := url.Parse(b)
	if e != nil {
		log.Fatal(e)
	}
	c.baseURL = u
}

func (c *DefaultClient) SetApiKey(a string) {
	c.apiKey = a
}

//AgenciesWithCoverage - 	list all supported agencies along with the center of
// 						 	their coverage area
// http://developer.onebusaway.org/modules/onebusaway-application-modules/current/api/where/methods/agencies-with-coverage.html
//
// Method: agency-with-coverage
// 	Returns a list of all transit agencies currently supported by OneBusAway
// 	along with the center of their coverage area.
//
// Sample Request
// http://api.pugetsound.onebusaway.org/api/where/agencies-with-coverage.xml?key=TEST
//
// Sample Response
// <response>
//   <version>2</version>
//   <code>200</code>
//   <text>OK</text>
//   <currentTime>1270614730908</currentTime>
//   <data class="listWithReferences">
//     <references>...</references>
//     <list>
//       <agencyWithCoverage>
//         <agencyId>3</agencyId>
//         <lat>47.21278384769539</lat>
//         <lon>-122.45624875362905</lon>
//         <latSpan>0.3559410000000014</latSpan>
//         <lonSpan>0.9080050000000028</lonSpan>
//       </agencyWithCoverage>
//       <agencyWithCoverage>...</agencyWithCoverage>
//    </list>
//     <limitExceeded>false</limitExceeded>
//   </data>
// </response>
//
// Response
// The response has the following fields:
// 	agencyId - 				an agency id for the agency whose coverage is included.
// 							Should match an <agency/> element referenced in the
// 							<references/> section.
// lat and lon - 			indicates the center of the agency’s coverage area
// latSpan and lonSpan - 	indicate the height (lat) and width (lon) of the
// 							coverage bounding box for the agency.
func (c DefaultClient) AgenciesWithCoverage() ([]AgencyWithCoverage, error) {
	retrieved, err := c.getData(agencyWithCoverageEndPoint, "Agencies with Coverage", nil)
	if err != nil {
		return nil, err
	}
	results := make([]AgencyWithCoverage, 0)
	for _, awc := range retrieved.List {
		results = append(results, awc.AgencyWithCoverage())
	}
	return results, nil
}

//Agency - 		get details for a specific agency
// http://developer.onebusaway.org/modules/onebusaway-application-modules/current/api/where/methods/agency.html
//
// Method: agency
//  Retrieve info for a specific transit agency identified by id
//
// Sample Request
// http://api.pugetsound.onebusaway.org/api/where/agency/1.xml?key=TEST
//
// Sample Response
// <response>
//   <version>2</version>
//   <code>200</code>
//   <text>OK</text>
//   <currentTime>1270614730908</currentTime>
//   <data class="entryWithReferences">
//     <references/>
//     <entry class="agency">
//       <id>1</id>
//       <name>Metro Transit</name>
//       <url>America/Los_Angeles</url>
//       <timezone>America/Los_Angeles</timezone>
//       <lang>en</lang>
//       <phone>206-553-3000</phone>
//       <disclaimer>Transit scheduling, geographic, and real-time data provided by permission of King County</disclaimer>
//     </entry>
//   </data>
// </response>
//
// Request Parameters
// id - 	the id of the agency, encoded directly in the URL:
// 			http://api.pugetsound.onebusaway.org/api/where/agency/[ID GOES HERE].xml
//
// Response
// For more details on the fields returned for an agency, see the documentation
// for the <agency/> element.
//

func (c DefaultClient) Agency(id string) (*Entry, error) {
	return c.getEntry(fmt.Sprint(agencyEndPoint, id), "Agency")
}

//ArrivalAndDepartureForStop - 	details about a specific arrival/departure at a
// 								stop
// http://developer.onebusaway.org/modules/onebusaway-application-modules/current/api/where/methods/arrival-and-departure-for-stop.html
//
// Method: arrival-and-departure-for-stop
//  Get info about a single arrival and departure for a stop
// Sample Request
// http://api.pugetsound.onebusaway.org/api/where/arrival-and-departure-for-stop/1_75403.xml?key=TEST&tripId=1_15551341&serviceDate=1291536000000&vehicleId=1_3521&stopSequence=42
//
// Sample Response
// <response>
//   <version>2</version>
//   <code>200</code>
//   <text>OK</text>
//   <currentTime>1270614730908</currentTime>
//   <data class="entryWithReferences">
//     <references>...</references>
//     <entry class="arrivalAndDeparture">
//       <!-- See documentation for the arrivalAndDeparture element, linked below -->
//     </entry>
//   </data>
// </response>
//
// Request Parameters
// id - 			the stop id, encoded directly in the URL:
// 					http://api.pugetsound.onebusaway.org/api/where/arrival-and-departure-for-stop/[ID GOES HERE].xml
// tripId - 		the trip id of the arriving transit vehicle
// serviceDate -	the service date of the arriving transit vehicle
// vehicleId - 		the vehicle id of the arriving transit vehicle (optional)
// stopSequence - 	the stop sequence index of the stop in the transit vehicle’s
// 					trip
// time -			by default, the method returns the status of the system
// 					right now. However, the system can also be queried at a
// 					specific time. This can be useful for testing. See
// 					timestamps for details on the format of the time parameter.
//
// The key here is uniquely identifying which arrival you are interested in.
// Typically, you would first make a call to arrivals-and-departures-for-stop to
// get a list of upcoming arrivals and departures at a particular stop. You can
// then use information from those results to specify a particular arrival. At
// minimum, you must specify the trip id and service date. Additionally, you are
// also encouraged to specify the vehicle id if available to help disambiguate
// between multiple vehicles serving the same trip instance. Finally, you are
// encouraged to specify the stop sequence. This helps in the situation when a
// vehicle visits a stop multiple times during a trip (it happens) plus there is
// performance benefit on the back-end as well.
//
// Response
// The method returns an <arrivalAndDeparture/> element as its content.
//
func (c DefaultClient) ArrivalAndDepartureForStop(id string, params map[string]string) (*Data, error) {
	return c.getData(fmt.Sprint(arrivalAndDepartureForStopEndPoint, id), "Arrival and Departure for Stop", params)
}

//ArrivalsAndDeparturesForStop - 	get current arrivals and departures for a stop
// http://developer.onebusaway.org/modules/onebusaway-application-modules/current/api/where/methods/arrivals-and-departures-for-stop.html
//
// Method: arrivals-and-departures-for-stop
//  Get current arrivals and departures for a stop identified by id
//
// Sample Request
// http://api.pugetsound.onebusaway.org/api/where/arrivals-and-departures-for-stop/1_75403.xml?key=TEST
//
// Sample Response
// <response>
//   <version>2</version>
//   <code>200</code>
//   <text>OK</text>
//   <currentTime>1270614730908</currentTime>
//   <data class="listWithReferences">
//     <references>...</references>
//     <entry class="stopWithArrivalsAndDepartures">
//       <stopId>1_75403</stopId>
//       <arrivalsAndDepartures>
//         <arrivalAndDeparture>...</arrivalAndDeparture>
//         <arrivalAndDeparture>...</arrivalAndDeparture>
//         <arrivalAndDeparture>...</arrivalAndDeparture>
//       </arrivalsAndDepartures>
//       <nearbyStopIds>
//         <string>1_75414</string>
//         <string>...</string>
//       </nearbyStopIds>
//     </entry>
//   </data>
// </response>
//
// Request Parameters
// id - 				the stop id, encoded directly in the URL:
// 						http://api.pugetsound.onebusaway.org/api/where/arrivals-and-departures-for-stop/[ID GOES HERE].xml
// minutesBefore=n - 	include vehicles having arrived or departed in the
// 						previous n minutes (default=5)
// minutesAfter=n - 	include vehicles arriving or departing in the next n
// 						minutes (default=35)
// time - 				by default, the method returns the status of the system
// 						right now. However, the system can also be queried at a
// 						specific time. This can be useful for testing. See
// 						timestamps for details on the format of the time parameter.
//
// Response
// The response is primarily composed of <arrivalAndDeparture/> elements, so see
// the element documentation for specific details.
// The nearby stop list is designed to capture stops that are very close by
// (like across the street) for quick navigation.
//
func (c DefaultClient) ArrivalsAndDeparturesForStop(id string, params map[string]string) (*Data, error) {
	return c.getData(fmt.Sprint(arrivalsAndDeparturesForStopEndPoint, id), "Arrivals and Departures for Stop", params)
}

//Block - 	get block configuration for a specific block
// http://developer.onebusaway.org/modules/onebusaway-application-modules/current/api/where/methods/block.html
//
// Method: block
//  Get details of a specific block by id\
//
// Sample Request
// <response>
//   <version>2</version>
//   <code>200</code>
//   <text>OK</text>
//   <currentTime>1391465493476</currentTime>
//   <data class="entryWithReferences">
//     <references />
//     <entry class="block">
//       <id>MTA NYCT_GH_A4-Sunday_D_GH_21000_BX12-15</id>
//       <configurations>
//         <blockConfiguration>
//           <!-- See documentation for the blockConfiguration element, linked below -->
//         </blockConfiguration>
//       </configurations>
//     </entry>
//   </data>
// </response>
//
// Request Parameters
// id - 	the id of the block, encoded directly in the url:
// 			http://api.pugetsound.onebusaway.org/api/where/block/[ID GOES HERE].xml
//
// Response
// See details about the various properties of the <blockConfiguration/> element.
//

func (c DefaultClient) Block(id string) (*Entry, error) {
	return c.getEntry(fmt.Sprint(blockEndPoint, id), "Block")
}

//CancelAlarm -	cancel a registered alarm
// http://developer.onebusaway.org/modules/onebusaway-application-modules/current/api/where/methods/cancel-alarm.html
//
// Method: cancel-alarm
//  Cancel a registered alarm.
//
// Sample Request
// http://api.pugetsound.onebusaway.org/api/where/cancel_alarm/1_00859082-9b9d-4f72-a89f-c4be0e2cf01a.xml
//
// Sample Response
// <response>
//   <version>2</version>
//   <code>200</code>
//   <text>OK</text>
//   <currentTime>1270614730908</currentTime>
//   <data>
//     <references/>
//   </data>
// </response>
//
// Request Parameters
// id - 	the alarm id is encoded directly in the URL
// 			http://api.pugetsound.onebusaway.org/api/where/cancel_alarm/[ID GOES HERE].xml
// The alarm id is returned in the call to
// register-alarm-for-arrival-and-departure-at-stop API method.
//
func (c DefaultClient) CancelAlarm(id string) error {
	u := c.buildRequestURL(fmt.Sprint(cancelAlarmEndPoint, id), nil)
	_, err := requestAndHandle(u, "Failed to Cancel Alarm for ID: ")
	return err
}

//CurrentTime -	retrieve the current system time
// http://developer.onebusaway.org/modules/onebusaway-application-modules/current/api/where/methods/current-time.html
//
// Method: current-time
//  Retrieve the current system time
//
// Sample Request
// http://api.pugetsound.onebusaway.org/api/where/current-time.xml?key=TEST
//
// Sample Response
// <response>
//   <version>2</version>
//   <code>200</code>
//   <text>OK</text>
//   <currentTime>1270614730908</currentTime>
//   <data class="time">
//     <references/>
//     <time>
//       <time>1270614730908</time>
//       <readableTime>2010-04-06T21:32:10-07:00</readableTime>
//     </time>
//   </data>
// </response>
//
// Response
// time - 			current system time as milliseconds since the Unix epoch
// readableTime - 	current system time in ISO 8601 format
//

func (c DefaultClient) CurrentTime() (*Time, error) {
	u := c.buildRequestURL(currentTimeEndPoint, nil)
	response, err := requestAndHandle(u, "Failed to get Current Time: ")
	if err != nil {
		return &Time{}, err
	}
	return response.Data.Time, nil
}

//RegisterAlarmForArrivalAndDepartureAtStop -	register an alarm for an
// 												arrival-departure event
// http://developer.onebusaway.org/modules/onebusaway-application-modules/current/api/where/methods/register-alarm-for-arrival-and-departure-at-stop.html
//
// Method: register-alarm-for-arrival-and-departure-at-stop
//  Register an alarm for a single arrival and departure at a stop, with a
//  callback URL to be requested when the alarm is fired.
//
// Sample Request
// http://api.pugetsound.onebusaway.org/api/where/register-alarm-for-arrival-and-departure-at-stop/1_75403.xml?key=TEST&tripId=1_15551341&serviceDate=1291536000000&vehicleId=1_3521&stopSequence=42&alarmTimeOffset=120&url=http://host/callback_url
//
// Sample Response:
// <response>
//   <version>2</version>
//   <code>200</code>
//   <currentTime>1318879898047</currentTime>
//   <text>OK</text>
//   <data class="entryWithReferences">
//     <references/>
//     <entry class="registeredAlarm">
//       <alarmId>1_7deee53d-9eb5-4f6b-8623-8bff398fcd5b</alarmId>
//     </entry>
//   </data>
// </response>
//
// Request Parameters
// id, tripId, serviceDate, vehicleId, stopSequence - 	see discussion in
// 														arrival-and-departure-for-stop
// 														API method for discussion
// 														of how to specify a
// 														particular arrival or
// 														departure
// url - 												callback URL that will
// 														be requested when the
// 														alarm is fired
// alarmTimeOffset - 									time, in seconds, that
// 														controls how long before
// 														the arrival/departure
// 														the alarm will be fired.
// 														Default is zero.
// onArrival - 											set to true to indicate
// 														the alarm should be
// 														fired relative to
// 														vehicle arrival, false
// 														for departure. The
// 														default is false for
// 														departure.
// We provide an arrival-departure alarm callback mechanism that allows you to
// register an alarm for an arrival or departure event and received a callback
// in the form of a GET request to a URL you specify.
// In order to specify an alarm for something like "5 minutes before a bus
// departs, we provide the alarmTimeOffset which specifies when the alarm should
// be fired relative to the actual arrival or departure event. A value of 60
// indicates that the alarm should be fired 60 seconds before, while a value of
// -30 would be fired 30 seconds after.
// A note about scheduled vs real-time arrivals and departures: You can register
// alarms for trips where we don’t have any real-time data (aka a scheduled
// arrival and departure) and we will fire the alarm at the appropriate time.
// Things get a bit trickier when you’ve registered an alarm for a scheduled
// arrival and we suddenly have real-time for the trip after you’ve registered.
// In these situations, we will automatically link your alarm to the real-time
// arrival and departure.
//
// Response
// The response is the alarm id. Note that if you include #ALARM_ID# anywhere in
// your callback URL, we will automatically replace it with the id of the alarm
// being fired. This can be useful when you register multiple alarms and need to
// be able to distinguish between them.
// Also see the cancel-alarm API method, which also accepts the alarm id as an
// argument.
//
// TODO
//func (c DefaultClient) RegisterAlarmForArrivalAndDepartureAtStop(id string, params map[string]string) (*RegisteredAlarm, error) {
//	u := c.buildRequestURL(fmt.Sprint(registerAlarmForArrivalAndDepartureAtStopEndPoint, id), params)
//	response, err := requestAndHandle(u, "Failed to Register Alarm for Arrival and Departure at Stop: ")
//	if err != nil {
//		return nil, err
//	}
//	return &response.Data.Entry.RegisteredAlarm, nil
//}

//ReportProblemWithTrip -	submit a user-generated problem for a trip
// http://developer.onebusaway.org/modules/onebusaway-application-modules/current/api/where/methods/report-problem-with-trip.html
//
// Method: report-problem-with-trip
// Submit a user-generated problem report for a particular trip. The reporting
// mechanism provides lots of fields that can be specified to give more context
// about the details of the problem (which trip, stop, vehicle, etc was
// involved), making it easier for a developer or transit agency staff to
// diagnose the problem. These reports feed into the problem reporting admin
// interface.
//
// Sample Request
// http://api.pugetsound.onebusaway.org/api/where/report-problem-with-trip/1_79430293.xml?key=TEST&serviceDate=1291536000000&vehicleId=1_3521&stopId=1_75403&code=vehicle_never_came
//
// Sample Response
// <response>
//   <version>2</version>
//   <code>200</code>
//   <currentTime>1318879898047</currentTime>
//   <text>OK</text>
//   <data/>
// </response>
//
// Request Parameters
// tripId - 		the trip id, encoded directly in the URL:
// 					http://api.pugetsound.onebusaway.org/api/where/report-problem-with-trip/[ID GOES HERE].xml
// serviceDate - 	the service date of the trip
// vehicleId - 		the vehicle actively serving the trip
// stopId - 		a stop id indicating the stop where the user is experiencing
// 					the problem
// code - 			a string code identifying the nature of the problem
// 					vehicle_never_came
// 					vehicle_came_early - 			the vehicle arrived earlier
// 													than predicted
// 					vehicle_came_late - 			the vehicle arrived later
// 													than predicted
// 					wrong_headsign - 				the headsign reported by
// 													OneBusAway differed from the
// 													vehicle’s actual headsign
// 					vehicle_does_not_stop_here - 	the trip in question does
// 													not actually service the
// 													indicated stop
// 					other - 						catch-all for everythign else
// userComment - 	additional comment text supplied by the user describing the
// 					problem
// userOnVehicle - 	true/false to indicate if the user is on the transit vehicle
// 					experiencing the problem
// userVehicleNumber - the vehicle number, as reported by the user
// userLat - 		the reporting user’s current latitude
// userLon - 		the reporting user’s current longitude
// userLocationAccuracy - the reporting user’s location accuracy, in meters
//
// In general, everything but the trip id itself is optional, but generally
// speaking, providing more fields in the report will make it easier to diagnose
// the actual underlying problem. Note that while we record specific location
// information for the user, we do not store any identifying information for the
// user in order to make it hard to link the user to their location as some
// point in the future.
//
func (c DefaultClient) ReportProblemWithTrip(id string, params map[string]string) error {
	u := c.buildRequestURL(reportPoblemWithTripEndPoint, params)
	_, err := requestAndHandle(u, "Failed to Report Problem with Trip: ")
	return err
}

//RouteIdsForAgency - 	get a list of all route ids for an agency
// http://developer.onebusaway.org/modules/onebusaway-application-modules/current/api/where/methods/route-ids-for-agency.html
//
// Method: route-ids-for-agency
//  Retrieve the list of all route ids for a particular agency.
//
// Sample Request
// http://api.pugetsound.onebusaway.org/api/where/route-ids-for-agency/40.xml?key=TEST
//
// Sample Respsone
// <response>
//   <version>2</version>
//   <code>200</code>
//   <text>OK</text>
//   <currentTime>1270614730908</currentTime>
//   <data class="listWithReferences">
//     <references/>
//     <list>
//       <string>40_510</string>
//       <string>40_511</string>
//       <string>40_513</string>
//       <string>...</string>
//     </list>
//     <limitExceeded>false</limitExceeded>
//   </data>
// </response>
//
// Request Parameters
// id - 	the id of the agency, encoded directly in the URL:
// 			http://api.pugetsound.onebusaway.org/api/where/route-ids-for-agency/[ID GOES HERE].xml?key=TEST
//
// Response
// Returns a list of all route ids for routes served by the specified agency.
// Note that <route/> elements for the referenced routes will NOT be included
// in the <references/> section, since there are potentially a large number of
// routes for an agency.
//
// TODO
//func (c DefaultClient) RouteIdsForAgency(id string) ([]*string, error) {
//	u := c.buildRequestURL(fmt.Sprint(reportPoblemWithTripEndPoint, id), nil)
//	response, err := requestAndHandle(u, "Failed to get Route IDs for Agency: ")
//	if err != nil {
//		return nil, err
//	}
//	return response.Data.List.Strings, nil
//}

//Route - 	get details for a specific route
// http://developer.onebusaway.org/modules/onebusaway-application-modules/current/api/where/methods/route.html
//
// Method: route
//  Retrieve info for a specific route by id.
//
// Sample Request
// http://api.pugetsound.onebusaway.org/api/where/route/1_100224.xml?key=TEST
//
// Sample Response
// <response>
//     <version>2</version>
//     <code>200</code>
//     <currentTime>1461441898217</currentTime>
//     <text>OK</text>
//     <data class="entryWithReferences">
//         <references>
//             <agencies>
//                 <agency>
//                     <id>1</id>
//                     <name>Metro Transit</name>
//                     <url>http://metro.kingcounty.gov</url>
//                     <timezone>America/Los_Angeles</timezone>
//                     <lang>EN</lang>
//                     <phone>206-553-3000</phone>
//                     <privateService>false</privateService>
//                 </agency>
//             </agencies>
//         </references>
//         <entry class="route">
//             <id>1_100224</id>
//             <shortName>44</shortName>
//             <description>Ballard - Montlake</description>
//             <type>3</type>
//             <url>http://metro.kingcounty.gov/schedules/044/n0.html</url>
//             <agencyId>1</agencyId>
//         </entry>
//     </data>
// </response>
//
// Request Parameters
// id - 	the id of the route, encoded directly in the URL:
// 			http://api.pugetsound.onebusaway.org/api/where/route/[ID GOES HERE].xml
//
// Response
// See details about the various properties of the <route/> element.
//

func (c DefaultClient) Route(id string) (*Entry, error) {
	return c.getEntry(fmt.Sprint(routeEndPoint, id), "Route")
}

//RoutesForAgency - 	get a list of all routes for an agency
// http://developer.onebusaway.org/modules/onebusaway-application-modules/current/api/where/methods/routes-for-agency.html
//
// Method: routes-for-agency
//  Retrieve the list of all routes for a particular agency by id
//
// Sample Request
// http://api.pugetsound.onebusaway.org/api/where/routes-for-agency/1.xml?key=TEST
//
// Sample Response
// <response>
//   <version>2</version>
//   <code>200</code>
//   <text>OK</text>
//   <currentTime>1270614730908</currentTime>
//   <data class="listWithReferences">
//     <references/>
//     <list>
//       <route>
//         <id>1_1</id>
//         <shortName>1</shortName>
//         <description>kinnear</description>
//         <type>3</type>
//         <url>http://metro.kingcounty.gov/tops/bus/schedules/s001_0_.html</url>
//         <agencyId>1</agencyId>
//       </route>
//       ...
//     </list>
//     <limitExceeded>false</limitExceeded>
//   </data>
// </response>
//
// Request Parameters
// id - 	the id of the agency, encoded directly in the URL:
//			http://api.pugetsound.onebusaway.org/api/where/routes-for-agency/[ID GOES HERE].xml
//
// Response
// Returns a list of all route ids for routes served by the specified agency.
// See the full description for the <route/> element.
//
// TODO
//func (c DefaultClient) RoutesForAgency(id string) ([]*Route, error) {
//	u := c.buildRequestURL(fmt.Sprint(routeForAgencyEndPoint, id), nil)
//	response, err := requestAndHandle(u, "Failed to get Routes for Agency: ")
//	if err != nil {
//		return []*Route{}, err
//	}
//	return response.Data.List.Routes, nil
//}

//RoutesForLocation -	search for routes near a location, optionally by route name
// http://developer.onebusaway.org/modules/onebusaway-application-modules/current/api/where/methods/routes-for-location.html
//
// Method: routes-for-location
// 	Search for routes near a specific location, optionally by name
//
// Sample Request
// http://api.pugetsound.onebusaway.org/api/where/routes-for-location.xml?key=TEST&lat=47.653435&lon=-122.305641
//
// Sample Response
// <response>
//   <version>2</version>
//   <code>200</code>
//   <text>OK</text>
//   <currentTime>1270614730908</currentTime>
//   <data class="listWithReferences">
//     <references>...</references>
//     <list>
//       <route>...</route>
//       <!-- More routes -->
//     </list>
//     <limitExceeded>true</limitExceeded>
//   </data>
// </response>
//
// Request Parameters
// lat - 				The latitude coordinate of the search center
// lon - 				The longitude coordinate of the search center
// radius - 			The search radius in meters (optional)
// latSpan/lonSpan - 	An alternative to radius to set the search bounding
// 						box (optional)
// query - 				A specific route short name to search for (optional)
// If you just specify a lat,lon search location, the routes-for-location method
// will just return nearby routes. If you specify an optional query parameter,
//  we’ll search for nearby routes with the specified route short name. This is
// the primary method from going from a user-facing route name like “44” to the
// actual underlying route id unique to a route for a particular transit agency.
//
// Response
// The routes-for-location method returns a list result, so see additional
// documentation on controlling the number of elements returned and interpreting
// the results. The list contents are <route/> elements.
//
func (c DefaultClient) RoutesForLocation(params map[string]string) (*Data, error) {
	return c.getData(routeForLocationEndPoint, "Routes for Location", params)
}

//ScheduleForStop - 	get the full schedule for a stop on a particular day
// http://developer.onebusaway.org/modules/onebusaway-application-modules/current/api/where/methods/schedule-for-stop.html
//
// Method: schedule-for-stop
//  Retrieve the full schedule for a stop on a particular day
//
// Sample Request
// http://api.pugetsound.onebusaway.org/api/where/schedule-for-stop/1_75403.xml?key=TEST
//
// Sample Response
// <response>
//   <version>2</version>
//   <code>200</code>
//   <text>OK</text>
//   <currentTime>1270614730908</currentTime>
//   <data class="entryWithReferences">
//     <references>...</references>
//     <entry class="stopSchedule">
//       <date>1270623339481</date>
//       <stopId>1_75403</stopId>
//       <stopRouteSchedules>
//         <stopRouteSchedule>
//           <routeId>1_31</routeId>
//           <stopRouteDirectionSchedules>
//             <stopRouteDirectionSchedule>
//               <tripHeadsign>Central Magnolia</tripHeadsign>
//               <scheduleStopTimes>
//                 <scheduleStopTime>
//                   <arrivalTime>1270559769000</arrivalTime>
//                   <departureTime>1270559769000</departureTime>
//                   <serviceId>1_114-WEEK</serviceId>
//                   <tripId>1_11893408</tripId>
//                 </scheduleStopTime>
//                 <!-- More schduleStopTime entries... -->
//               </scheduleStopTimes>
//             </stopRouteDirectionSchedule>
//           </stopRouteDirectionSchedules>
//           <!-- More stopRouteDirectionSchedule entries -->
//         </stopRouteSchedule>
//         <!-- More stopRouteSchedule entries -->
//       </stopRouteSchedules>
//       <timeZone>America/Los_Angeles</timeZone>
//       <stopCalendarDays>
//         <stopCalendarDay>
//           <date>1276239600000</date>
//           <group>1</group>
//           </stopCalendarDay>
//         <!-- More stopCalendarDay entries -->
//       </stopCalendarDays>
//     </entry>
//   </data>
// </response>
//
// Request Parameters
// id - 	the stop id to request the schedule for, encoded directly in the URL:
// 			http://api.pugetsound.onebusaway.org/api/where/schedule-for-stop/[ID GOES HERE].xml
// date - 	The date for which you want to request a schedule of the format YYYY-MM-DD (optional, defaults to current date)
//
// Response
// The response is pretty complex, so we’ll describe the details at a high-level
// along with references to the various elements in the response.
// The response can be considered in two parts. The first part lists specific
// arrivals and departures at a stop on a given date (<stopRouteSchedules/>
// section) while the second part lists which days the stop currently has
// service defined (the <stopCalendarDays/> section). By convention, we refer
// to the arrival and departure time details for a particular trip at a stop as
// a stop time.
//
// We break up the stop time listings in a couple of ways. First, we split the
// stop times by route (corresponds to each <stopRouteSchedule/> element). We
// next split the stop times for each route by direction of travel along the
// route (corresponds to each <stopRouteDirectionSchedule/> element). Most stops
// will serve just one direction of a particular route, but some stops will
// serve both directions, and it may be useful to present those listings
// separately. Each <stopRouteDirectionSchedule/> element has a tripHeadsign
// property that indicates the direction of travel.
//
// Finally we get down to the unit of a stop time, as represented by the
// <scheduleStopTime/> element. Each element has the following set of properties:
// arrivalTime - 	time in milliseconds since the Unix epoch that the transit
// 					vehicle will arrive
// departureTime - 	time in milliseconds since the Unix epoch that the transit
// 					vehicle will depart
// tripId - 		the id for the trip of the scheduled transit vehicle
// serviceId - 		the serviceId for the schedule trip (see the GTFS spec for
// 					more details
// In addition to all the <scheduleStopTime/> elements, the response also
// contains <stopCalendarDay/> elements which list out all the days that a
// particular stop has service. This element has the following properties:
// date - 	the date of service in milliseconds since the Unix epoch
// group - 	we provide a group id that groups <stopCalendarDay/> into
// 			collections of days with similar service. For example,
// 			Monday-Friday might all have the same schedule and the same group
// 			id as result, while Saturday and Sunday have a different weekend
// 			schedule, so they’d get their own group id.
// In addition to all the <scheduleStopTime/> elements, the main entry also has
// the following properties:
// date - 		the active date for the returned calendar
// stopId - 	the stop id for the requested stop, which can be used to access
// 				the <stop/> element in the <references/> section
// timeZone - 	the time-zone the stop is located in
//
func (c DefaultClient) ScheduleForStop(id string) (*Data, error) {
	return c.getData(fmt.Sprint(scheduleForStopEndPoint, id), "Schedule for Stop", nil)
}

//Shape -	get details for a specific shape (polyline drawn on a map)
// http://developer.onebusaway.org/modules/onebusaway-application-modules/current/api/where/methods/shape.html
//
// Method: shape
//  Retrieve a shape (the path traveled by a transit vehicle) by id
//
// Sample Request
// http://api.pugetsound.onebusaway.org/api/where/shape/1_40046045.xml?key=TEST
//
// Sample Response
// <response>
//   <version>2</version>
//   <code>200</code>
//   <text>OK</text>
//   <currentTime>1270614730908</currentTime>
//   <data class="entryWithReferences">
//     <references/>
//     <entry class="encodedPolyline">
//       <points>ky`bHvwajVtDJ??|DL??fDH|@BnALVH??n@HhA\NF??DBZHzBhA??|@d@??nAl@fDdB??rBfAf@V??pCrA??
// hEvBf@N??p@Lp@F??j@DfA?xA???|A?pB@lCDtDElB???`@MVUTe@??FSFQNm@BG??
// DKHIRIfDgAbFwA|@_@x@q@b@a@j@}@??JQr@_B|AsDbA_CJWTs@??DQRgA????Ly@J_ABs@@u@?i@?YOwD??
// g@mJG_@Kc@??IUmAeCgAwBGs@???iI??CmO???M?aO???sF???cI???q@???g@?kF???cE???oA???sG???gH???cG@_@?aF?M??
// @{N???U???O??AeF@yH???oL???eN??rC???pC@??pC???D???jC???D???bC???`C???pC???pC???pCB??pCE??bCA??TFRR??
// nBkB??rCmC??bA_A??xCsC??rAuA??bD_D??nDiD??nDiD??PO^]??zCyC???sE???kG???_B???wB???qE???G?mE???sE??@aD???
// q@???sE???eFJm@??f@]fBcB??dAiA??x@y@??j@s@\a@??_@u@KUGQ??AeD???mACe@G_@Ka@??_@{@??a@}@cAcC??
// eA{B??{@eBEKEI??kBaE??oCI??AQGe@UwBSuB??E]?wD??BwF???oF??@mF???oF??DuF??DoF??@mF??@mF??@oF???gF??
// @uE???aC??@}A??@aF???qABuA?mA??AmC??^?pCDT????I?M?S?U@aA??@aF???YCi@Eg@CeA?a@??
// PgAPs@BMH_@VaA`@y@Zs@Ra@XmABQ??Fm@@}@@I?Q?_@B}FMW??@mC???q@@a@@c@D]Fi@Ne@??HUpAwC??Y_@??
// YS????WGkKQ??}B???m@A??{BC??eHC??iHG</points>
//       <length>351</length>
//     </entry>
//   </data>
// </response>
//
// Request Parameters
// id - 	the shape id, encoded directly in the URL:
// 			http://api.pugetsound.onebusaway.org/api/where/shape/[ID GOES HERE].xml
//
// Response
// The path is returned as a <shape/> element with a points in the encoded
// polyline format defined for Google Maps.
//

func (c DefaultClient) Shape(id string) (*Entry, error) {
	return c.getEntry(fmt.Sprint(shapeEndPoint, id), "Shape")
}

//StipIDsForAgency - 	get a list of all stops for an agency
// http://developer.onebusaway.org/modules/onebusaway-application-modules/current/api/where/methods/stop-ids-for-agency.html
//
// Method: stops-ids-for-agency
//  Retrieve the list of all stops for a particular agency by id
//
// Sample Request
// http://api.pugetsound.onebusaway.org/api/where/stop-ids-for-agency/40.xml?key=TEST
//
// Sample Response
// <response>
//   <version>2</version>
//   <code>200</code>
//   <text>OK</text>
//   <currentTime>1270614730908</currentTime>
//   <data class="listWithReferences">
//     <references/>
//     <list>
//       <string>40_C_1303</string>
//       <string>40_C_1305</string>
//       <string>40_C_1366</string>
//       <string>...</string>
//     </list>
//     <limitExceeded>false</limitExceeded>
//   </data>
// </response>
//
// Request Parameters
// id -  	the id of the agency, encoded directly in the URL:
// 			http://api.pugetsound.onebusaway.org/api/where/stop-ids-for-agency/[ID GOES HERE].xml
//
// Response
// Returns a list of all stop ids for stops served by the specified agency.
// Note that <stop/> elements for the referenced stops will NOT be included in
// the <references/> section, since there are potentially a large number of
// stops for an agency.
//
// TODO
//func (c DefaultClient) StopIDsForAgency(id string) ([]*string, error) {
//	u := c.buildRequestURL(fmt.Sprint(stopIDsForAgencyEndPoint, id), nil)
//	response, err := requestAndHandle(u, "Failed to get Stop IDs for Agency: ")
//	if err != nil {
//		return nil, err
//	}
//	return response.Data.List.Strings, nil
//}

//Stop - 	get details for a specific stop
// http://developer.onebusaway.org/modules/onebusaway-application-modules/current/api/where/methods/stop.html
//
// Method: stop
//  Retrieve info for a specific stop by id
//
// Sample Request
// http://api.pugetsound.onebusaway.org/api/where/stop/1_75403.xml?key=TEST
//
// Sample Response
// <response>
//   <version>2</version>
//   <code>200</code>
//   <text>OK</text>
//   <currentTime>1270614730908</currentTime>
//   <data class="entryWithReferences">
//     <references>...</references>
//     <entry class="stop">
//       <id>1_75403</id>
//       <lat>47.6543655</lat>
//       <lon>-122.305206</lon>
//       <direction>S</direction>
//       <name>Stevens Way &amp; BENTON LANE</name>
//       <code>75403</code>
//       <locationType>0</locationType>
//       <routeIds>
//         <string>1_31</string>
//         <string>...</string>
//       </routeIds>
//     </entry>
//   </data>
// </response>
//
// Request Parameters
// id -  	the id of the requested stop, encoded directly in the URL:
// 			http://api.pugetsound.onebusaway.org/api/where/stop/[ID GOES HERE].xml
//
// Response
// See details about the various properties of the <stop/> element.
//

func (c DefaultClient) Stop(id string) (*Entry, error) {
	return c.getEntry(fmt.Sprint(stopEndPoint, id), "Stop")
}

//StopsForLocation - 	search for stops near a location, optionally by stop code
// http://developer.onebusaway.org/modules/onebusaway-application-modules/current/api/where/methods/stops-for-location.html
//
// Method: stops-for-location
//  Search for stops near a specific location, optionally by stop code
//
// Sample Request
// http://api.pugetsound.onebusaway.org/api/where/stops-for-location.xml?key=TEST&lat=47.653435&lon=-122.305641
//
// Sample Response
// <response>
//   <version>2</version>
//   <code>200</code>
//   <text>OK</text>
//   <currentTime>1270614730908</currentTime>
//   <data class="listWithReferences">
//     <references>...</references>
//     <list>
//       <stop>...</stop>
//       <!-- More stops -->
//     </list>
//     <limitExceeded>true</limitExceeded>
//     <outOfRange>false</outOfRange>
//   </data>
// </response>
//
// Request Parameters
// lat - 				The latitude coordinate of the search center
// lon - 				The longitude coordinate of the search center
// radius - 			The search radius in meters (optional)
// latSpan/lonSpan - 	An alternative to radius to set the search bounding box (optional)
// query - 				A specific stop code to search for (optional)
// If you just specify a lat,lon search location, the stops-for-location method
// will just return nearby stops. If you specify an optional query parameter,
// we’ll search for nearby stops with the specified code. This is the primary
// method from going from a user-facing stop code like “75403” to the actual
// underlying stop id unique to a stop for a particular transit agency.
//
// Response
// The stops-for-location method returns a list result, so see additional
// documentation on controlling the number of elements returned and interpreting
// the results. The list contents are <stop/> elements, so see details about the
// various properties of the <stop/> element.
//
func (c DefaultClient) StopsForLocation(params map[string]string) (*Data, error) {
	return c.getData(stopsForLocationEndPoint, "Stops for Location", params)
}

//StopsForRoute - 	get the set of stops and paths of travel for a particular route
// http://developer.onebusaway.org/modules/onebusaway-application-modules/current/api/where/methods/stops-for-route.html
//
// Method: stops-for-route
//  Retrieve the set of stops serving a particular route, including groups by
//  direction of travel. The stops-for-route method first and foremost provides
//  a method for retrieving the set of stops that serve a particular route. In
//  addition to the full set of stops, we provide various “stop groupings” that
//  are used to group the stops into useful collections. Currently, the main
//  grouping provided organizes the set of stops by direction of travel for the
//  route. Finally, this method also returns a set of polylines that can be used
//  to draw the path traveled by the route.
//
// Sample Request
// http://api.pugetsound.onebusaway.org/api/where/stops-for-route/1_100224.xml?key=TEST
//
// Sample Response
// <response>
//     <version>2</version>
//     <code>200</code>
//     <currentTime>1461443625722</currentTime>
//     <text>OK</text>
//     <data class="entryWithReferences">
//          <references></references>
//          <entry class="stopsForRoute">
//               <routeId>1_100224</routeId>
//               <stopIds>
//                    <string>1_10911</string>
//                    <string>...</string>
//               </stopIds>
//               <stopGroupings>
//                   <stopGrouping>
//                        <type>direction</type>
//                        <ordered>true</ordered>
//                        <stopGroups>
//                            <stopGroup>
//                                 <id>0</id>
//                                 <name>
//                                     <type>destination</type>
//                                     <names>
//                                         <string>BALLARD WALLINGFORD</string>
//                                     </names>
//                                 </name>
//                                 <stopIds>
//                                     <string>1_25240</string>
//                                     <string>...</string>
//                                 </stopIds>
//                                 <polylines>...</polylines>
//                            </stopGroup>
//                        </stopGroups>
//                   </stopGrouping>
//               </stopGroupings>
//               <polylines>...</polylines>
//          </entry>
//     </data>
// </response>
//
// Request Parameters
// id -  	The route id, encoded directly in the URL:
// 			http://api.pugetsound.onebusaway.org/api/where/stops-for-route/[ID GOES HERE].xml
// includePolylines=true|false = Optional parameter that controls whether
// 			polyline elements are included in the response. Defaults to true.
//
// Response
//

func (c DefaultClient) StopsForRoute(id string) (*Entry, error) {
	return c.getEntry(fmt.Sprint(stopsForRouteEndPoint, id), "StopsForRoute")
}

//TripDetails - 	get extended details for a specific trip
// http://developer.onebusaway.org/modules/onebusaway-application-modules/current/api/where/methods/trip-details.html
//
// Method: trip-details
//  Get extended details for a specific trip
//
// Sample Request
// http://api.pugetsound.onebusaway.org/api/where/trip-details/1_12540399.xml?key=TEST
//
// Sample Response
// <response>
//   <version>2</version>
//   <code>200</code>
//   <text>OK</text>
//   <currentTime>1270614730908</currentTime>
//   <data class="entryWithReferences">
//     <references>...</references>
//     <entry class="tripDetails">
//       <tripId>1_12540399</tripId>
//       <serviceDate>1271401200000</serviceDate>
//       <frequency>...</frequency>
//       <status>...</status>
//       <schedule>...</schedule>
//     </entry>
//   </data>
// </response>
//
// Request Parameters
// id -  				the id of the trip, encoded directly in the URL:
// 						http://api.pugetsound.onebusaway.org/api/where/trip-details/[ID GOES HERE].xml
// serviceDate - 		the service date for the trip as unix-time in ms (optional).
// 						Used to disambiguate different versions of the same trip.
// 						See [Glossary#ServiceDate the glossary entry for service date].
// includeTrip - 		Can be true/false to determine whether full <trip/>
// 						element is included in the <references/> section.
// 						Defaults to true.
// includeSchedule - 	Can be true/false to determine whether full <schedule/>
// 						element is included in the <tripDetails/> section.
// 						Defaults to true.
// includeStatus - 		Can be true/false to determine whether the full <status/>
// 						element is include in the <tripDetails/> section.
// 						Defaults to true.
// time - 				by default, the method returns the status of the system
// 						right now. However, the system can also be queried at a
// 						specific time. This can be useful for testing. See
// 						timestamps for details on the format of the time parameter.
//
// Response
// The response <entry/> element is a <tripDetails/> element that captures
// extended details about a trip.
//

func (c DefaultClient) TripDetails(id string) (*Entry, error) {
	return c.getEntry(fmt.Sprint(tripDetailsEndPoint, id), "TripDetails")
}

//TripForVehicle - 	get extended trip details for current trip of a specific
// 					transit vehicle
// http://developer.onebusaway.org/modules/onebusaway-application-modules/current/api/where/methods/trip-for-vehicle.html
//
// Method: trip-for-vehicle
//  Get extended trip details for a specific transit vehicle. That is, given a
//  vehicle id for a transit vehicle currently operating in the field, return
// 	extended trips details about the current trip for the vehicle.
//
// Sample Request
// http://api.pugetsound.onebusaway.org/api/where/trip-for-vehicle/1_4210.xml?key=TEST
//
// Sample Response
// <response>
//   <version>2</version>
//   <code>200</code>
//   <text>OK</text>
//   <currentTime>1270614730908</currentTime>
//   <data class="entryWithReferences">
//     <references>...</references>
//     <entry class="tripDetails">
//       <tripId>1_12540399</tripId>
//       <serviceDate>1271401200000</serviceDate>
//       <frequency>...</frequency>
//       <status>...</status>
//       <schedule>...</schedule>
//       <tripId>1_15456175</tripId>
//     </entry>
//   </data>
// </response>
//
// Request Parameters
// id - 				the id of the vehicle, encoded directly in the URL:
// 						http://api.pugetsound.onebusaway.org/api/where/trip-for-vehicle/[ID GOES HERE].xml
// includeTrip - 		Can be true/false to determine whether full <trip/>
// 						element is included in the <references/> section.
// 						Defaults to false.
// includeSchedule - 	Can be true/false to determine whether full <schedule/>
// 						element is included in the <tripDetails/> section.
// 						Defaults to false.
// includeStatus - 		Can be true/false to determine whether the full
// 						<status/> element is include in the <tripDetails/>
// 						section. Defaults to true.
// time -				by default, the method returns the status of the system
// 						right now. However, the system can also be queried at a
// 						specific time. This can be useful for testing. See
// 						timestamps for details on the format of the time parameter.
//
// Response
// The response <entry/> element is a <tripDetails/> element that captures
// extended details about a trip.

func (c DefaultClient) TripForVehicle(id string, params map[string]string) (*Data, error) {
	return c.getData(fmt.Sprint(tripForVehicleEndPoint, id), "TripDetails for Vehicle", params)
}

//Trip - 	get details for a specific trip
// http://developer.onebusaway.org/modules/onebusaway-application-modules/current/api/where/methods/trip.html
//
// Method: trip
//  Get details of a specific trip by id
//
// Sample Request
// http://api.pugetsound.onebusaway.org/api/where/trip/1_12540399.xml?key=TEST
//
// Sample Response
// <response>
//   <version>2</version>
//   <code>200</code>
//   <text>OK</text>
//   <currentTime>1270614730908</currentTime>
//   <data class="entryWithReferences">
//     <references>...</references>
//     <entry class="trip">
//       <id>1_12540399</id>
//       <routeId>1_44</routeId>
//       <tripShortName>LOCAL</tripShortName>
//       <tripHeadsign>Downtown via University District</tripHeadsign>
//       <serviceId>1_114-115-WEEK</serviceId>
//       <shapeId>1_20044006</shapeId>
//       <directionId>1</directionId>
//     </entry>
//   </data>
// </response>
//
// Request Parameters
// id -  	the id of the trip, encoded directly in the URL:
// 			http://api.pugetsound.onebusaway.org/api/where/trip/[ID GOES HERE].xml
//
// Response
// See details about the various properties of the <trip/> element.

func (c DefaultClient) Trip(id string) (*Entry, error) {
	return c.getEntry(fmt.Sprint(tripEndPoint, id), "Trip")
}

//TripsForLocation - 	get active trips near a location
// http://developer.onebusaway.org/modules/onebusaway-application-modules/current/api/where/methods/trips-for-location.html
//
// Method: trips-for-location
//  Search for active trips near a specific location.
//
// Sample Request
// http://api.pugetsound.onebusaway.org/api/where/trips-for-location.xml?key=TEST&lat=47.653&lon=-122.307&latSpan=0.008&lonSpan=0.008
//
// Sample Response
// <response>
//   <version>2</version>
//   <code>200</code>
//   <text>OK</text>
//   <currentTime>1270614730908</currentTime>
//   <data class="listWithReferences">
//     <references>...</references>
//     <list>
//       <tripDetails>...</tripDetails>
//       <tripDetails>...</tripDetails>
//       ...
//     </list>
//     <limitExceeded>false</limitExceeded>
//   </data>
// </response>
//
// Request Parameters
// lat - 				The latitude coordinate of the search center
// lon - 				The longitude coordinate of the search center
// latSpan/lonSpan - 	Set the limits of the search bounding box
// includeTrip - 		Can be true/false to determine whether full <trip/>
// 						elements are included in the <references/> section.
// 						Defaults to false.
// includeSchedule - 	Can be true/false to determine whether full <schedule/>
// 						elements are included in the <tripDetails/> section.
// 						Defaults to false.
// time - 				by default, the method returns the status of the system
// 						right now. However, the system can also be queried at a
// 						specific time. This can be useful for testing. See
// 						timestamps for details on the format of the time parameter.
//
// Response
// The response is a list of <tripDetails/> element that captures extended
// details about each active trip. Active trips are ones where the transit
// vehicle is currently located within the search radius. We use real-time
// arrival data to determine the position of transit vehicles when available,
// otherwise we determine the location of vehicles from the static schedule.
//
func (c DefaultClient) TripsForLocation(params map[string]string) (*Data, error) {
	return c.getData(tripsForLocationEndPoint, "TripDetails for Location", params)
}

//TripsForRoute - 	get active trips for a route
// http://developer.onebusaway.org/modules/onebusaway-application-modules/current/api/where/methods/trips-for-route.html
//
// Method: trips-for-route
//  Search for active trips for a specific route.
//
// Sample Request
// http://api.pugetsound.onebusaway.org/api/where/trips-for-route/1_100224.xml?key=TEST
//
// Sample Response
// <response>
//   <version>2</version>
//   <code>200</code>
//   <text>OK</text>
//   <currentTime>1270614730908</currentTime>
//   <data class="listWithReferences">
//     <references>...</references>
//     <list>
//       <tripDetails>...</tripDetails>
//       <tripDetails>...</tripDetails>
//       ...
//     </list>
//     <limitExceeded>false</limitExceeded>
//   </data>
// </response>
//
// Request Parameters
// id - 				the id of the route, encoded directly in the URL:
// 						http://api.pugetsound.onebusaway.org/api/where/trips-for-route/[ID GOES HERE].xml
// includeStatus - 		Can be true/false to determine whether full <tripStatus/>
// 						elements with full real-time information are included in
// 						the <status/> section for each <tripDetails/> element.
// 						Defaults to false.
// includeSchedule - 	Can be true/false to determine whether full <schedule/>
// 						elements are included in the <tripDetails/> element.
// 						Defaults to false.
// time - 				by default, the method returns the status of the system
// 						right now. However, the system can also be queried at a
// 						specific time. This can be useful for testing. See
// 						timestamps for details on the format of the time parameter.
//
// Response
// The response is a list of <tripDetails/> element that captures extended
// details about each active trip. The set of active trips includes any trip
// that serves that specified route that is currently active.
//
func (c DefaultClient) TripsForRoute(id string) (*Data, error) {
	return c.getData(fmt.Sprint(tripsForRouteEndPoint, id), "TripDetails for Route", nil)
}

//VehiclesForAgency - 	get active vehicles for an agency
// http://developer.onebusaway.org/modules/onebusaway-application-modules/current/api/where/methods/vehicles-for-agency.html
//
// Method: vehicles-for-agency
//  Search for active vehicles for a particular agency by id.
//
// Sample Request
// http://api.onebusaway.org/api/where/vehicles-for-agency/1.xml?key=TEST
//
// Sample Response
// <response>
//   <version>2</version>
//   <code>200</code>
//   <text>OK</text>
//   <currentTime>1270614730908</currentTime>
//   <data class="listWithRangeAndReferences">
//     <references>...</references>
//     <list>
//       <vehicleStatus>...</vehicleStatus>
//       <vehicleStatus>...</vehicleStatus>
//       <vehicleStatus>...</vehicleStatus>
//       ...
//     </list>
//     <limitExceeded>false</limitExceeded>
//     <outOfRange>false</outOfRange>
//   </data>
// </response>
//
// Request Parameters
// id -  	the id of the agency, encoded directly in the URL:
// 			http://api.onebusaway.org/api/where/vehicles-for-agency/[ID GOES HERE].xml
// time -	by default, the method returns the status of the system right now.
// 			However, the system can also be queried at a specific time. This can
// 			be useful for testing. See timestamps for details on the format of
// 			the time parameter.
//
// Response
// The response is a list of <vehicleStatus/> elements that captures extended details about each active vehicle associated with the specified agency.
//
func (c DefaultClient) VehiclesForAgency(id string) (*Data, error) {
	return c.getData(fmt.Sprint(vehiclesForAgencyEndPoint, id), "Vehicles for Agency", nil)
}

func (c DefaultClient) getEntry(requestString, requestType string) (*Entry, error) {
	u := c.buildRequestURL(requestString+jsonPostFix, nil)
	response, err := requestAndHandle(u, fmt.Sprintf("Failed to get %s: ", requestType))
	if err != nil {
		return nil, err
	}
	return response.Data.Entry, nil
}

func (c DefaultClient) getData(requestString, errMessage string, params map[string]string) (*Data, error) {
	u := c.buildRequestURL(fmt.Sprint(requestString, jsonPostFix), params)
	response, err := requestAndHandle(u, fmt.Sprintf("Failed to get %s: ", errMessage))
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

func (c DefaultClient) buildRequestURL(endpoint string, params map[string]string) string {
	u := *c.baseURL
	u.Path = path.Join(u.Path, endpoint)
	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	q.Set("key", c.apiKey)
	u.RawQuery = q.Encode()
	return u.String()
}
