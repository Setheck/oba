//Package oba - One Bus Away Go Api https://onebusaway.org/
// Author: Seth T <setheck@gmail.com>
package oba

import (
	"encoding/xml"
	"errors"
	"net/url"
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
	XMLName     xml.Name `xml:"response"`
	Version     string   `xml:"version"`
	Code        int      `xml:"code"`
	CurrentTime string   `xml:"currentTime"`
	Text        string   `xml:"text"`
	Data        Data     `xml:"data"`
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
	Agencies      []Agency    `xml:"agencies>agency,omitempty"`
	Routes        []Route     `xml:"routes>route,omitempty"`
	Stops         []Stop      `xml:"stops>stop,omitempty"`
	Trips         []Trip      `xml:"trips>trip,omitempty"`
	Situations    []Situation `xml:"situations>situation,omitempty"`
	List          []string    `xml:"list>string,omitempty"`
	LimitExceeded string      `xml:"limitExceeded,omitempty"`
}

//Data container object
type Data struct {
	XMLName       xml.Name   `xml:"data"`
	Class         string     `xml:"class,attr"`
	References    References `xml:"references"`
	Entry         Entry
	Time          Time
	List          List
	StopsForRoute StopsForRoute
}

type Time struct {
	XMLName      xml.Name `xml:"time"`
	Time         string   `xml:"time"`
	ReadableTime string   `xml:"readableTime"`
}

type RegisteredAlarm struct {
	AlarmID string `xml:"alarmId,omitempty"`
}

//Entry container object
type Entry struct {
	XMLName xml.Name `xml:"entry"`
	Class   string   `xml:"class,attr"`
	ID      string   `xml:"id"`
	URL     string   `xml:"url"`
	Name    string   `xml:"name"`
	Agency
	Block
	Route
	Shape
	Stop
	StopSchedule
	Trip
	TripDetails
	RegisteredAlarm
}

type ArrivalAndDeparture struct {
	RouteID                string `xml:"routeId"`
	TripID                 string `xml:"tripId"`
	ServiceDate            string `xml:"serviceDate"`
	StopID                 string `xml:"stopId"`
	StopSequance           string `xml:"stopSequence"`
	BlockTripSequence      string `xml:"blockTripSequence"`
	RouteShortName         string `xml:"routeShortName"`
	RouteLongName          string `xml:"routeLongName"`
	TripHeadSign           string `zml:"tripHeadsign"`
	ArrivalEnabled         string `xml:"arrivalEnabled"`
	DepartureEnabled       string `xml:"departureEnabled"`
	ScheduledArrivalTime   string `xml:"scheduledArrivalTime"`
	ScheduledDepartureTime string `xml:"scheduledDepartureTime"`
	Frequency              string `xml:"frequency"`
	Predicted              string `xml:"predicted"`
	PredictedArrivalTime   string `xml:"predictedArrivalTime"`
	PredictedDepartureTime string `xml:"predictedDepartureTime"`
	DistanceFromStop       string `xml:"distanceFromStop"`
	NumberOfStopsAway      string `xml:"numberOfStopsAway"`
	TripStatus             string `xml:"tripStatus"`
}

type BlockConfiguration struct {
	ActiveServiceIDs   []string    `xml:"activeServiceIds>string"`
	InactiveServiceIDs []string    `xml:"inactiveServiceIds>string"`
	Trips              []BlockTrip `xml:"trips"`
}

type BlockTrip struct {
	TripID               string     `xml:"tripId"`
	BlockStopTimes       []StopTime `xml:"blockStopTimes"`
	AccumulatedSlackTime string     `xml:"accumulatedSlackTime"`
	DistanceAlongBlock   string     `xml:"distanceAlongBlock"`
}

type StopTime struct {
	StopID        string `xml:"stopId"`
	ArrivalTime   string `xml:"arrivalTime"`
	DepartureTime string `xml:"departureTime"`
	PickupType    string `xml:"pickupType"`
	DropOffType   string `xml:"droppOffType"`
}

type Frequency struct {
	StartTime string `xml:"startTime"`
	EndTime   string `xml:"endTime"`
	Headway   string `xml:"headway"`
}

type List struct {
	XMLName       xml.Name        `xml:"list"`
	Routes        []Route         `xml:"route"`
	Stops         []Stop          `xml:"stop"`
	Strings       []string        `xml:"string"`
	TripDetails   []TripDetails   `xml:"tripDetails"`
	VehicleStatus []VehicleStatus `xml:"vehicleStatus"`
}

type VehicleStatus struct {
	VehicleID              string `xml:"vehicleId"`
	LastUpdateTime         string `xml:"lastUpdateTime"`
	LastLocationUpdateTime string `xml:"lastLocationUpdateTime"`
	LocationLat            string `xml:"location>lat"`
	LocationLon            string `xml:"location>lon"`
	TripID                 string `xml:"tripId"`
	TripStatus             string `xml:"tripStatus"`
}

//Agency container object
type Agency struct {
	TimeZone       string `xml:"timezone"`
	Lang           string `xml:"lang,omitempty"`
	Phone          string `xml:"phone,omitempty"`
	Disclaimer     string `xml:"disclaimer,omitempty"`
	PrivateService string `xml:"privateService,omitempty"`
}

type Block struct {
	Configurations []BlockConfiguration `xml:"configurations>blockConfiguration"`
}

type AgencyWithCoverage struct {
	AgencyID string `xml:"agencyId"`
	Lat      string `xml:"lat"`
	Lon      string `xml:"lon"`
	LatSpan  string `xml:"latSpan"`
	LonSpan  string `xml:"lonSpan"`
}

//Route object
type Route struct {
	ShortName   string `xml:"shortName"`
	Description string `xml:"description"`
	Type        string `xml:"type"`
	AgencyID    string `xml:"agencyId"`
}

type Situation struct {
	ID                string           `xml:"id"`
	CreationTime      string           `xml:"creationTime"`
	EnvironmentReason string           `xml:"environmentReason"`
	Summary           []string         `xml:"summary>value"`
	Description       []string         `xml:"description>value"`
	Affects           []VehicleJourney `xml:"vehicleJourneys>vehicleJourney"`
	Consequences      []Consequence    `xml:"consequences>consequence"`
}

type Consequence struct {
	Condition                          string   `xml:"condition"`
	ConditionDetailDiversionPathPoints []string `xml:"conditionDetails>diversionPath>points"`
	ConditionDetailDiversionStopIDs    []string `xml:"conditionDetails>diversionStopIds>string"`
}

type VehicleJourney struct {
	LineID      string   `xml:"lineId"`
	Direction   string   `xml:"direction"`
	CallStopIDs []string `xml:"calls>call>stopId"`
}

type Shape struct {
	Points string `xml:"points"`
	Length string `xml:"length"`
}

type Stop struct {
	Lat          string   `xml:"lat"`
	Lon          string   `xml:"lon"`
	Direction    string   `xml:"direction"`
	Code         string   `xml:"code"`
	LocationType string   `xml:"locationType"`
	RouteIDs     []string `xml:"routeIds>string"`
}

type StopSchedule struct {
	Date               string              `xml:"date"`
	StopID             string              `xml:"stopId"`
	StopRouteSchedules []StopRouteSchedule `xml:"stopRouteSchedules>stopRouteSchedule"`
	TimeZone           string              `xml:"timeZone"`
	StopCalendarDays   []StopCalendarDay   `xml:"stopCalendarDays>stopCalendarDay"`
}

type StopCalendarDay struct {
	Date  string `xml:"date"`
	Group string `xml:"group"`
}

type StopRouteSchedule struct {
	RouteID                     string              `xml:"routeId"`
	StopRouteDirectionSchedules []StopRouteSchedule `xml:"stopRouteDirectionSchedules>stopRouteDIrectionSchedule"`
}

type StopRouteDIrectionSchedule struct {
	TripHeadsign      string             `xml:"tripHeadsign"`
	ScheduleStopTimes []ScheduleStopTime `xml:"scheduleStopTimes>scheduleStopTime"`
}

type StopsForRoute struct {
	RouteID       string         `xml:"routeId"`
	StopIDs       []string       `xml:"stopIds>string"`
	StopGroupings []StopGrouping `xml:"stopGroupings>stopGrouping"`
}

type StopGrouping struct {
	Type       string      `xml:"type"`
	Ordered    string      `xml:"ordered"`
	StopGroups []StopGroup `xml:"stopGroups>stopGroup"`
}

type StopGroup struct {
	ID        string   `xml:"id"`
	NameType  string   `xml:"name>type"`
	Names     []string `xml:"name>names>string"`
	PolyLines string   `xml:"polylines"`
}

type ScheduleStopTime struct {
	ArrivalTime   string `xml:"arrivalTime"`
	DepartureTime string `xml:"departureTime"`
	ServiceID     string `xml:"serviceId"`
	TripID        string `xml:"tripId"`
}

type Trip struct {
	RouteID       string `xml:"routeId"`
	TripShortName string `xml:"tripShortName"`
	TripHeadsign  string `xml:"tripHeadsign"`
	ServiceID     string `xml:"serviceId"`
	ShapeID       string `xml:"shapeId"`
	DirectionID   string `xml:"directionId"`
}

type TripDetails struct {
	TripID                 string     `xml:"tripId"`
	ServiceDate            string     `xml:"serviceDate"`
	Frequency              string     `xml:"frequency"`
	Status                 string     `xml:"status"`
	ScheduleTimeZone       string     `xml:"schedule>timeZone"`
	ScheduleStopTimes      []StopTime `xml:"schedule>StopTimes>tripStopTime"`
	SchedulePreviousTripID []string   `xml:"schedule>previousTripId"`
	ScheduleNextTripID     []string   `xml:"schedule>nextTripId"`
	SituationIDs           []string   `xml:"sitouationIds>string"`
}

type TripStatus struct {
	XMLName                    xml.Name `xml:"status"`
	ActiveTripID               string   `xml:"activeTripId"`
	BlockTripSequence          string   `xml:"blockTripSequence"`
	ServiceDate                string   `xml:"serviceDate"`
	ScheduledDistanceAlongTrip string   `xml:"scheduledDistanceAlongTrip"`
	TotalDistanceAlongTrip     string   `xml:"totalDistanceAlongTrip"`
	PositionLat                string   `xml:"position>lat"`
	PositionLon                string   `xml:"position>lon"`
	Orientation                string   `xml:"orientation"`
	ClosestStop                string   `xml:"closestStop"`
	ClosestStopTimeOffset      string   `xml:"closestStopTimeOffset"`
	NextStop                   string   `xml:"nextStop"`
	NextStopTimeOffset         string   `xml:"nextStopTimeOffset"`
	Phase                      string   `xml:"phase"`
	Status                     string   `xml:"status"`
	Predicted                  string   `xml:"predicted"`
	LastUpdateTime             string   `xml:"lastUpdateTime"`
	LastLocationUpdateTime     string   `xml:"lastLocationUpdateTime"`
	LastKnownLocationLat       string   `xml:"lastKnownLocation>lat"`
	LastKnownLocationLon       string   `xml:"lastKnownLocation>lon"`
	LastKnownOrientation       string   `xml:"lastKnownOrientation"`
	DistanceAlongTrip          string   `xml:"distanceAlongTrip"`
	ScheduleDeviation          string   `xml:"scheduleDeviation"`
	VehicleID                  string   `xml:"vehicleId"`
	SituationIDs               []string `xml:"situationIds>string"`
}

type StopWithArrivalsAndDepartures struct {
	StipID                string                `xml:"stopId"`
	ArrivalsAndDepartures []ArrivalAndDeparture `xml:"arrivalsAndDepartures"`
	NearByStopIDs         []string              `xml:"nearbyStopIds>string"`
}

const (
	xmlPostFix                                        = ".xml"
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
	baseURL string
	apiKey  string
}

//NewDefaultClient - instantiate a new instance of a Client
func NewDefaultClient(baseURL, apiKey string) *DefaultClient {
	return &DefaultClient{baseURL, apiKey}
}

func (c *DefaultClient) SetBaseURL(b string) {
	c.baseURL = b
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
func (c DefaultClient) AgenciesWithCoverage() (Data, error) {
	return c.getData(agencyWithCoverageEndPoint, "Agencies with Coverage", nil)
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
func (c DefaultClient) Agency(id string) (Entry, error) {
	return c.getEntry(agencyEndPoint+id, "Agency")
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
func (c DefaultClient) ArrivalAndDepartureForStop(id string, params map[string]string) (Data, error) {
	return c.getData(arrivalAndDepartureForStopEndPoint+id, "Arrival and Departure for Stop", params)
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
func (c DefaultClient) ArrivalsAndDeparturesForStop(id string, params map[string]string) (Data, error) {
	return c.getData(arrivalsAndDeparturesForStopEndPoint+id, "Arrivals and Departures for Stop", params)
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
func (c DefaultClient) Block(id string) (Entry, error) {
	return c.getEntry(blockEndPoint+id, "Block")
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
	u, err := c.buildRequestURL(cancelAlarmEndPoint+id, nil)
	if err != nil {
		return err
	}
	_, err = requestAndHandle(u.String(), "Failed to Cancel Alarm for ID: ")
	if err != nil {
		return err
	}
	return nil
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
func (c DefaultClient) CurrentTime() (Time, error) {
	u, err := c.buildRequestURL(currentTimeEndPoint, nil)
	if err != nil {
		return Time{}, err
	}
	response, err := requestAndHandle(u.String(), "Failed to get Current Time: ")
	if err != nil {
		return Time{}, err
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
func (c DefaultClient) RegisterAlarmForArrivalAndDepartureAtStop(id string, params map[string]string) (RegisteredAlarm, error) {
	u, err := c.buildRequestURL(registerAlarmForArrivalAndDepartureAtStopEndPoint+id, params)
	if err != nil {
		return RegisteredAlarm{}, err
	}
	response, err := requestAndHandle(u.String(), "Failed to Register Alarm for Arrival and Departure at Stop: ")
	if err != nil {
		return RegisteredAlarm{}, err
	}
	return response.Data.Entry.RegisteredAlarm, nil
}

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
	u, err := c.buildRequestURL(reportPoblemWithTripEndPoint, params)
	if err != nil {
		return err
	}
	_, err = requestAndHandle(u.String(), "Failed to Report Problem with Trip: ")
	if err != nil {
		return err
	}
	return nil
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
func (c DefaultClient) RouteIdsForAgency(id string) ([]string, error) {
	u, err := c.buildRequestURL(reportPoblemWithTripEndPoint+id, nil)
	if err != nil {
		return nil, err
	}
	response, err := requestAndHandle(u.String(), "Failed to get Route IDs for Agency: ")
	if err != nil {
		return nil, err
	}
	return response.Data.List.Strings, nil
}

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
func (c DefaultClient) Route(id string) (Entry, error) {
	return c.getEntry(routeEndPoint+id, "Route")
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
func (c DefaultClient) RoutesForAgency(id string) ([]Route, error) {
	u, err := c.buildRequestURL(routeForAgencyEndPoint+id, nil)
	if err != nil {
		return []Route{}, err
	}
	response, err := requestAndHandle(u.String(), "Failed to get Routes for Agency: ")
	if err != nil {
		return []Route{}, err
	}
	return response.Data.List.Routes, nil
}

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
func (c DefaultClient) RoutesForLocation(params map[string]string) (Data, error) {
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
func (c DefaultClient) ScheduleForStop(id string) (Data, error) {
	return c.getData(scheduleForStopEndPoint+id, "Schedule for Stop", nil)
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
func (c DefaultClient) Shape(id string) (Entry, error) {
	return c.getEntry(shapeEndPoint+id, "Shape")
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
func (c DefaultClient) StopIDsForAgency(id string) ([]string, error) {
	u, err := c.buildRequestURL(stopIDsForAgencyEndPoint+id, nil)
	if err != nil {
		return nil, err
	}
	response, err := requestAndHandle(u.String(), "Failed to get Stop IDs for Agency: ")
	if err != nil {
		return nil, err
	}
	return response.Data.List.Strings, nil
}

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
func (c DefaultClient) Stop(id string) (Entry, error) {
	return c.getEntry(stopEndPoint+id, "Stop")
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
func (c DefaultClient) StopsForLocation(params map[string]string) (Data, error) {
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
func (c DefaultClient) StopsForRoute(id string) (Entry, error) {
	return c.getEntry(stopsForRouteEndPoint+id, "StopsForRoute")
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
func (c DefaultClient) TripDetails(id string) (Entry, error) {
	return c.getEntry(tripDetailsEndPoint+id, "TripDetails")
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
// 						Defaults to fale.
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
//
func (c DefaultClient) TripForVehicle(id string, params map[string]string) (Data, error) {
	return c.getData(tripForVehicleEndPoint+id, "TripDetails for Vehicle", params)
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
func (c DefaultClient) Trip(id string) (Entry, error) {
	return c.getEntry(tripEndPoint+id, "Trip")
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
func (c DefaultClient) TripsForLocation(params map[string]string) (Data, error) {
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
func (c DefaultClient) TripsForRoute(id string) (Data, error) {
	return c.getData(tripsForRouteEndPoint+id, "TripDetails for Route", nil)
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
func (c DefaultClient) VehiclesForAgency(id string) (Data, error) {
	return c.getData(vehiclesForAgencyEndPoint+id, "Vehicles for Agency", nil)
}

func (c DefaultClient) getEntry(requestString, requestType string) (Entry, error) {
	u, err := c.buildRequestURL(requestString+xmlPostFix, nil)
	if err != nil {
		return Entry{}, err
	}
	response, err := requestAndHandle(u.String(), "Failed to get "+requestType+": ")
	if err != nil {
		return Entry{}, err
	}
	return response.Data.Entry, nil
}

func (c DefaultClient) getData(requestString, errMessage string, params map[string]string) (Data, error) {
	u, err := c.buildRequestURL(requestString+xmlPostFix, params)
	if err != nil {
		return Data{}, err
	}
	response, err := requestAndHandle(u.String(), "Failed to get "+errMessage+": ")
	if err != nil {
		return Data{}, err
	}
	return response.Data, nil
}

func (c DefaultClient) buildRequestURL(endpoint string, params map[string]string) (*url.URL, error) {
	requestURL := c.baseURL + endpoint
	u, err := url.Parse(requestURL)
	if err != nil {
		return nil, errors.New("Failed to parse URL: " + err.Error())
	}
	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	q.Set("key", c.apiKey)
	u.RawQuery = q.Encode()
	return u, nil
}
