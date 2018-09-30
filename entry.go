package oba

//entry container object
type entry struct {
	AccumulatedSlackTime         float64          `json:"accumulatedSlackTime,omitempty"`
	ActiveServiceID              []string         `json:"activeServiceIds,omitempty"`
	ActiveTripID                 string           `json:"activeTripId"`
	Affects                      []VehicleJourney `json:"vehicleJourneys>vehicleJourney"`
	AgencyID                     string           `json:"agencyId,omitempty"`
	AlarmID                      string           `json:"alarmId,omitempty"`
	ArrivalEnabled               *bool            `json:"arrivalEnabled,omitempty"`
	ArrivalsAndDepartures        List             `json:"arrivalsAndDepartures,omitempty"`
	ArrivalTime                  int              `json:"arrivalTime"`
	BlockID                      string           `json:"blockId,omitempty"`
	BlockSequence                int              `json:"blockSequence,omitempty"`
	BlockStopTimes               List             `json:"blockStopTimes,omitempty"`
	BlockTripSequence            int              `json:"blockTripSequence,omitempty"`
	Trips                        List             `json:"trips,omitempty"`
	ClosestStop                  string           `json:"closestStop"`
	ClosestStopTimeOffset        int              `json:"closestStopTimeOffset"`
	Code                         string           `json:"code,omitempty"`
	Color                        string           `json:"color,omitempty"`
	Configurations               List             `json:"configurations,omitempty"`
	Consequences                 []Consequence    `json:"consequences>consequence"`
	CreationTime                 string           `json:"creationTime"`
	Date                         int              `json:"date,omitempty"`
	DepartureEnabled             *bool            `json:"departureEnabled,omitempty"`
	DepartureTime                int              `json:"departureTime,omitempty"`
	Description                  string           `json:"description,omitempty"`
	Direction                    string           `json:"direction,omitempty"`
	DirectionID                  string           `json:"directionId,omitempty"`
	Disclaimer                   string           `json:"disclaimer,omitempty"`
	DistanceAlongBlock           float64          `json:"distanceAlongBlock,omitempty"`
	DistanceAlongTrip            float64          `json:"distanceAlongTrip,omitempty"`
	DistanceFromStop             float64          `json:"distanceFromStop,omitempty"`
	DropOffType                  int              `json:"dropOffType,omitempty"`
	Email                        string           `json:"email,omitempty"`
	EndTime                      int              `json:"entTime,omitempty"`
	EnvironmentReason            string           `json:"environmentReason"`
	FareURL                      string           `json:"fareUrl,omitempty"`
	Frequency                    *string          `json:"frequency,omitempty"`
	Headway                      int              `json:"headway,omitempty"`
	ID                           string           `json:"id,omitempty"`
	InactiveServiceID            []string         `json:"inactiveServiceIds,omitempty"`
	Lang                         string           `json:"lang,omitempty"`
	LastKnownDistanceAlongTrip   float64          `json:"lastKnownDistanceAlongTrip,omitempty"`
	LastKnownLocation            Location         `json:"lastKnownLocation,omitempty"`
	LastKnownOrientation         int              `json:"lastKnownOrientation,omitempty"`
	LastLocationUpdateTime       int              `json:"lastLocationUpdateTime,omitempty"`
	LastUpdateTime               int              `json:"lastUpdateTime,omitempty"`
	Lat                          float64          `json:"lat,omitempty"`
	LatSpan                      float64          `json:"latSpan,omitempty"`
	Length                       int              `json:"length,omitempty"`
	Levels                       string           `json:"levels,omitempty"`
	Location                     *entry           `json:"location,omitempty"`
	LocationType                 int              `json:"locationType,omitempty"`
	Lon                          float64          `json:"lon,omitempty"`
	LongName                     string           `json:"longName,omitempty"`
	LonSpan                      float64          `json:"lonSpan,omitempty"`
	Name                         string           `json:"name,omitempty"`
	Names                        []string         `json:"names,omitempty"`
	NearbyStopIds                []string         `json:"nearbyStopIds,omitempty"`
	NextStop                     string           `json:"nextStop,omitempty"`
	NextStopTimeOffset           int              `json:"nextStopTimeOffset,omitempty"`
	NumberOfStopsAway            int              `json:"numberOfStopsAway,omitempty"`
	Orientation                  float64          `json:"orientation,omitempty"`
	Ordered                      *bool            `json:"ordered,omitempty"`
	Phase                        string           `json:"phase,omitempty"`
	Phone                        string           `json:"phone,omitempty"`
	PickupType                   int              `json:"pickupType,omitempty"`
	Points                       string           `json:"points,omitempty"`
	PolyLines                    List             `json:"polylines,omitempty"`
	Position                     Location         `json:"position"`
	Predicted                    *bool            `json:"predicted,omitempty"`
	PredictedArrivalInterval     int              `json:"predictedArrivalInterval,omitempty"`
	PredictedArrivalTime         int              `json:"predictedArrivalTime,omitempty"`
	PredictedDepartureInterval   int              `json:"predictedDepartureInterval,omitempty"`
	PredictedDepartureTime       int              `json:"predictedDepartureTime,omitempty"`
	PrivateService               *bool            `json:"privateService,omitempty"`
	ReadableTime                 string           `json:"readableTime,omitempty"`
	RouteID                      string           `json:"routeId,omitempty"`
	RouteIDs                     []string         `json:"routeIds,omitempty"`
	RouteLongName                string           `json:"routeLongName,omitempty"`
	RouteShortName               string           `json:"routeShortName,omitempty"`
	ScheduledArrivalInterval     int              `json:"scheduledArrivalInterval,omitempty"`
	ScheduledArrivalTime         int              `json:"scheduledArrivalTime,omitempty"`
	ScheduledDepartureInterval   int              `json:"scheduledDepartureInterval,omitempty"`
	ScheduledDepartureTime       int              `json:"scheduledDepartureTime,omitempty"`
	ScheduledDistanceAlongTrip   float64          `json:"scheduledDistanceAlongTrip"`
	ScheduleDeviation            int              `json:"scheduleDeviation"`
	ScheduleDeviationHistogramID string           `json:"scheduleDeviationHistogramId,omitempty"`
	ScheduleFrequencies          List             `json:"scheduleFrequencies,omitempty"`
	ScheduleStopTimes            List             `json:"scheduleStopTimes,omitempty"`
	ServiceDate                  int              `json:"serviceDate,omitempty"`
	ServiceID                    string           `json:"serviceId,omitempty"`
	ShapeID                      string           `json:"shapeId,omitempty"`
	ShortName                    string           `json:"shortName,omitempty"`
	SituationID                  string           `json:"situationId,omitempty"`
	SituationIDs                 []string         `json:"situationIds,omitempty"`
	StartTime                    int              `json:"startTime,omitempty"`
	Status                       string           `json:"status,omitempty"`
	StopCalendarDays             List             `json:"stopCalendarDays,omitempty"`
	StopHeadsign                 string           `json:"stopHeadsign,omitempty"`
	StopGroupings                AltList          `json:"stopGroupings,omitempty"`
	StopGroups                   AltList          `json:"stopGroups,omitempty"`
	StopID                       string           `json:"stopId,omitempty"`
	StopIDs                      []string         `json:"stopIds,omitempty"`
	StopRouteSchedules           List             `json:"stopRouteSchedules,omitempty"`
	StopRouteDirectionSchedules  List             `json:"stopRouteDirectionSchedules,omitempty"`
	StopSequence                 int              `json:"stopSequence,omitempty"`
	StopTime                     *entry           `json:"stopTime,omitempty"`
	Summary                      []string         `json:"summary,omitempty"`
	TextColor                    string           `json:"textColor,omitempty"`
	Time                         int              `json:"time,omitempty"`
	TimeZone                     string           `json:"timezone,omitempty"`
	TotalDistanceAlongTrip       float64          `json:"totalDistanceAlongTrip"`
	TotalStopsInTrip             int              `json:"totalStopsInTrip,omitempty"`
	TripHeadSign                 string           `json:"tripHeadsign,omitempty"`
	TripID                       string           `json:"tripId,omitempty"`
	TripShortName                string           `json:"tripShortName,omitempty"`
	TripStatus                   *entry           `json:"tripStatus,omitempty"`
	Type                         int              `json:"type,omitempty"`
	URL                          string           `json:"url,omitempty"`
	VehicleID                    string           `json:"vehicleId,omitempty"`
	WheelChairBoarding           string           `json:"wheelchairBoarding,omitempty"`
	//Description       []string         `json:"description>value"` what the fuck
}

type altEntry struct {
	ID         string     `json:"id,omitempty"`
	Name       *nameEntry `json:"name,omitempty"`
	Names      []string   `json:"names,omitempty"`
	PolyLines  AltList    `json:"polylines,omitempty"`
	StopIDs    []string   `json:"stopIds,omitempty"`
	SubGroups  AltList    `json:"subGroups,omitempty"`
	Length     int        `json:"length,omitempty"`
	Levels     string     `json:"levels,omitempty"`
	Points     string     `json:"points,omitempty"`
	Type       string     `json:"type,omitempty"`
	Ordered    *bool      `json:"ordered,omitempty"`
	StopGroups AltList    `json:"stopGroups,omitempty"`
}

type nameEntry struct {
	Name  string   `json:"name,omitempty"`
	Names []string `json:"names,omitempty"`
	Type  string   `json:"type,omitempty"`
}

func (e nameEntry) nameFromEntry() *Name {
	return &Name{
		Names: e.Names,
		Type:  e.Type,
		Name:  e.Name,
	}
}

func (e altEntry) stopGroupFromEntry(stops []Stop) *StopGroup {
	ss := make([]Stop, 0, len(stops))
	for _, sid := range e.StopIDs {
		for _, s := range stops {
			if sid == s.ID {
				ss = append(ss, s)
			}
		}
	}

	var name Name
	if e.Name != nil {
		name = *e.Name.nameFromEntry()
	}
	return &StopGroup{
		ID:        e.ID,
		Stops:     ss,
		Name:      name,
		PolyLines: e.PolyLines.toEncodedPolyLines(),
	}
}

func (e entry) agencyFromEntry() *Agency {
	return &Agency{
		ID:             e.ID,
		Disclaimer:     e.Disclaimer,
		Email:          e.Email,
		Lang:           e.Lang,
		FareURL:        e.FareURL,
		Name:           e.Name,
		Phone:          e.Phone,
		PrivateService: e.PrivateService,
		TimeZone:       e.TimeZone,
		URL:            e.URL,
	}
}

func (e entry) agencyWithCoverageFromEntry(a Agency) *AgencyWithCoverage {
	return &AgencyWithCoverage{
		Agency:  a,
		Lat:     e.Lat,
		LatSpan: e.LatSpan,
		Lon:     e.Lon,
		LonSpan: e.LonSpan,
	}
}

func (e entry) arrivalAndDepartureFromEntry(sis []Situation, st []Stop, ts []Trip) *ArrivalAndDeparture {
	return &ArrivalAndDeparture{
		ArrivalEnabled:               e.ArrivalEnabled,
		BlockTripSequence:            e.BlockTripSequence,
		DepartureEnabled:             e.DepartureEnabled,
		DistanceFromStop:             e.DistanceFromStop,
		Frequency:                    e.Frequency,
		LastUpdateTime:               e.LastUpdateTime,
		NumberOfStopsAway:            e.NumberOfStopsAway,
		Predicted:                    e.Predicted,
		PredictedArrivalInterval:     e.PredictedArrivalInterval,
		PredictedArrivalTime:         e.PredictedArrivalTime,
		PredictedDepartureInterval:   e.PredictedDepartureInterval,
		PredictedDepartureTime:       e.PredictedDepartureTime,
		RouteID:                      e.RouteID,
		RouteShortName:               e.RouteShortName,
		RouteLongName:                e.RouteLongName,
		ScheduledArrivalInterval:     e.ScheduledArrivalInterval,
		ScheduledArrivalTime:         e.ScheduledArrivalTime,
		ScheduledDepartureInterval:   e.ScheduledDepartureInterval,
		ScheduledDepartureTime:       e.ScheduledDepartureTime,
		ScheduleDeviationHistogramID: e.ScheduleDeviationHistogramID,
		ServiceDate:                  e.ServiceDate,
		SituationIDs:                 e.SituationIDs,
		Status:                       e.Status,
		StopID:                       e.StopID,
		StopSequence:                 e.StopSequence,
		TripID:                       e.TripID,
		TripHeadSign:                 e.TripHeadSign,
		TripStatus:                   e.tripStatusFromEntry(sis, st, ts),
		VehicleID:                    e.VehicleID,
	}
}

func (e entry) blockFromEntry() *Block {
	bcs := e.Configurations.toBlockConfigurations()
	return &Block{
		ID:             e.ID,
		Configurations: bcs,
	}
}

func (e entry) blockConfigurationFromEntry(asds, isds []string, tps []BlockTrip) *BlockConfiguration {
	return &BlockConfiguration{
		ActiveServiceIDs:   asds,
		InactiveServiceIDs: isds,
		Trips:              tps,
	}
}

func (e entry) blockStopTimeFromEntry() *BlockStopTime {
	var st StopTime
	if e.StopTime != nil {
		st = *e.StopTime.stopTimeFromEntry()
	}
	return &BlockStopTime{
		BlockSequence:        e.BlockSequence,
		DistanceAlongBlock:   e.DistanceAlongBlock,
		AccumulatedSlackTime: e.AccumulatedSlackTime,
		StopTime:             st,
	}
}

func (e entry) blockTripFromEntry() *BlockTrip {
	bsts := e.BlockStopTimes.toBlockStopTimes()
	return &BlockTrip{
		BlockStopTimes: bsts,
		TripID:         e.TripID,
	}
}

func (e entry) coverageFromEntry() Coverage {
	return Coverage{
		Lat:     e.Lat,
		Lon:     e.Lon,
		LatSpan: e.LatSpan,
		LonSpan: e.LonSpan,
	}
}

func (e entry) currentTimeFromEntry() *CurrentTime {
	return &CurrentTime{
		ReadableTime: e.ReadableTime,
		Time:         e.Time,
	}
}

func (e altEntry) encodedPolyLineFromEntry() *EncodedPolyLine {
	return &EncodedPolyLine{
		Length: e.Length,
		Levels: e.Levels,
		Points: e.Points,
	}
}

func (e entry) frequencyFromEntry() *Frequency {
	return &Frequency{
		StartTime: e.StartTime,
		EndTime:   e.EndTime,
		Headway:   e.Headway,
	}
}

func (e entry) locationFromEntry() *Location {
	return &Location{
		Lat: e.Lat,
		Lon: e.Lon,
	}
}

func (e entry) registeredAlarmFromEntry() *RegisteredAlarm {
	return &RegisteredAlarm{
		AlarmID: e.AlarmID,
	}
}

func (e entry) routeFromEntry(agencies []Agency) *Route {
	var agency Agency
	for _, a := range agencies {
		if e.AgencyID == a.ID {
			agency = a
		}
	}

	return &Route{
		Agency:      agency,
		Color:       e.Color,
		Description: e.Description,
		ID:          e.ID,
		LongName:    e.LongName,
		ShortName:   e.ShortName,
		URL:         e.URL,
		TextColor:   e.TextColor,
		Type:        e.Type,
	}
}

func (e entry) scheduleStopTimeFromEntry() *ScheduleStopTime {
	return &ScheduleStopTime{
		ArrivalEnabled:   e.ArrivalEnabled,
		ArrivalTime:      e.ArrivalTime,
		DepartureEnabled: e.DepartureEnabled,
		DepartureTime:    e.DepartureTime,
		ServiceID:        e.ServiceID,
		StopHeadsign:     e.StopHeadsign,
		TripID:           e.TripID,
	}
}

func (e entry) shapeFromEntry() *Shape {
	return &Shape{
		Points: e.Points,
		Length: e.Length,
	}
}

func (e entry) situationFromEntry() *Situation {
	return &Situation{
		//TODO: []horseshit -> Description:       e.Description,
		Affects:           e.Affects,
		Consequences:      e.Consequences,
		CreationTime:      e.CreationTime,
		EnvironmentReason: e.EnvironmentReason,
		ID:                e.ID,
		Summary:           e.Summary,
	}
}

func (e entry) stopFromEntry() *Stop {
	return &Stop{
		Code:               e.Code,
		Direction:          e.Direction,
		ID:                 e.ID,
		Lat:                e.Lat,
		LocationType:       e.LocationType,
		Lon:                e.Lon,
		Name:               e.Name,
		WheelChairBoarding: e.WheelChairBoarding,
	}
}

func (e entry) stopWithArrivalsAndDeparturesFromEntry(a ArrivalsAndDepartures) *StopWithArrivalsAndDepartures {
	return &StopWithArrivalsAndDepartures{
		StopID:                e.StopID,
		ArrivalsAndDepartures: a,
		NearByStopIDs:         e.NearbyStopIds,
	}
}

func (e entry) stopRouteScheduleFromEntry(r Route) *StopRouteSchedule {
	return &StopRouteSchedule{
		StopRouteDirectionSchedules: e.StopRouteDirectionSchedules.toStopRouteDirectionSchedules(),
		Route: r,
	}
}

func (e entry) stopRouteDirectionScheduleFromEntry() *StopRouteDirectionSchedule {
	return &StopRouteDirectionSchedule{
		ScheduleFrequencies: e.ScheduleFrequencies.toScheduleFrequencies(),
		TripHeadsign:        e.TripHeadSign,
		ScheduleStopTimes:   e.ScheduleStopTimes.toScheduleStopTimes(),
	}
}

func (e entry) scheduleFrequencyFromEntry() *ScheduleFrequency {
	return &ScheduleFrequency{}
}

func (e entry) stopsForRouteFromEntry(rs []Route, ss []Stop) *StopsForRoute {
	var route Route
	for _, r := range rs {
		if e.RouteID == r.ID {
			route = r
		}
	}
	stops := make([]Stop, 0, len(ss))
	for _, s := range ss {
		for _, sid := range e.StopIDs {
			if sid == s.ID {
				stops = append(stops, s)
			}
		}
	}
	sgs := e.StopGroupings.toStopGroupings(ss)
	return &StopsForRoute{
		Route:         route,
		Stops:         stops,
		StopGroupings: sgs,
	}
}

func (e altEntry) stopGroupingFromEntry(ss []Stop) *StopGrouping {
	return &StopGrouping{
		Type:       e.Type,
		Ordered:    e.Ordered,
		StopGroups: e.StopGroups.toStopGroups(ss),
	}
}

func (e entry) stopScheduleFromEntry(ss []Stop) *StopSchedule {
	var stop Stop
	for _, s := range ss {
		if s.ID == e.StopID {
			stop = s
			break
		}
	}
	return &StopSchedule{
		TimeZone: e.TimeZone,
		Date:     e.Date,
		Stop:     stop,
	}
}

func (e entry) stopTimeFromEntry() *StopTime {
	return &StopTime{
		ArrivalTime:   e.ArrivalTime,
		DepartureTime: e.DepartureTime,
		DropOffType:   e.DropOffType,
		PickupType:    e.PickupType,
		StopID:        e.StopID,
	}
}

func (e entry) tripFromEntry() *Trip {
	return &Trip{
		BlockID:        e.BlockID,
		DirectionID:    e.DirectionID,
		ID:             e.ID,
		RouteID:        e.RouteID,
		RouteShortName: e.RouteShortName,
		ServiceID:      e.ServiceID,
		ShapeID:        e.ShapeID,
		TimeZone:       e.TimeZone,
		TripHeadsign:   e.TripHeadSign,
		TripShortName:  e.TripShortName,
	}
}

func (e entry) tripDetailsFromEntry(ts []Trip, ss []Situation) *TripDetails {
	var trip Trip
	for _, t := range ts {
		if t.ID == e.TripID {
			trip = t
		}
	}
	return &TripDetails{
		Frequency:   e.Frequency,
		ServiceDate: e.ServiceDate,
		Situations:  ss,
		Status:      e.Status,
		Trip:        trip,
	}
}

func (e entry) tripStatusFromEntry(sis []Situation, ss []Stop, ts []Trip) *TripStatus {
	var cstop Stop
	var nstop Stop
	for _, s := range ss {
		if e.NextStop == s.ID {
			nstop = s
		}
		if e.ClosestStop == s.ID {
			cstop = s
		}
	}
	situations := make([]Situation, 0, len(sis))
	for _, si := range sis {
		for _, sid := range e.SituationIDs {
			if sid == si.ID {
				situations = append(situations, si)
			}
		}
	}

	return &TripStatus{
		ActiveTripID:               e.ActiveTripID,
		BlockTripSequence:          e.BlockTripSequence,
		ClosestStop:                cstop,
		ClosestStopTimeOffset:      e.ClosestStopTimeOffset,
		DistanceAlongTrip:          e.DistanceAlongTrip,
		Frequency:                  e.Frequency,
		LastKnownDistanceAlongTrip: e.LastKnownDistanceAlongTrip,
		LastKnownLocation:          e.LastKnownLocation,
		LastKnownOrientation:       e.LastKnownOrientation,
		LastLocationUpdateTime:     e.LastLocationUpdateTime,
		LastUpdateTime:             e.LastUpdateTime,
		NextStop:                   nstop,
		NextStopTimeOffset:         e.NextStopTimeOffset,
		Orientation:                e.Orientation,
		Phase:                      e.Phase,
		Position:                   e.Position,
		Predicted:                  e.Predicted,
		ScheduleDeviation:          e.ScheduleDeviation,
		ScheduledDistanceAlongTrip: e.ScheduledDistanceAlongTrip,
		ServiceDate:                e.ServiceDate,
		SituationIDs:               e.SituationIDs,
		Status:                     e.Status,
		TotalDistanceAlongTrip:     e.TotalDistanceAlongTrip,
		VehicleID:                  e.VehicleID,
	}
}

func (e entry) vehicleStatusFromEntry(sis []Situation, ss []Stop, ts []Trip) (ret *VehicleStatus) {
	var trip Trip
	for _, t := range ts {
		if t.ID == e.TripID {
			trip = t
			break
		}
	}

	var tstatus TripStatus
	if e.TripStatus != nil {
		tstatus = *e.TripStatus.tripStatusFromEntry(sis, ss, ts)
	}
	var loc Location
	if e.Location != nil {
		loc = *e.Location.locationFromEntry()
	}
	return &VehicleStatus{
		Location:               loc,
		LastUpdateTime:         e.LastUpdateTime,
		LastLocationUpdateTime: e.LastLocationUpdateTime,
		Phase:      e.Phase,
		Status:     e.Status,
		TripStatus: tstatus,
		Trip:       trip,
		VehicleID:  e.VehicleID,
	}
}
