package oba

//Entry container object
type Entry struct {
	AccumulatedSlackTime         float64          `json:"accumulatedSlackTime,omitempty"`
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
	DropOffTime                  int              `json:"dropOffTime,omitempty"`
	Email                        string           `json:"email,omitempty"`
	EndTime                      int              `json:"entTime,omitempty"`
	EnvironmentReason            string           `json:"environmentReason"`
	FareURL                      string           `json:"fareUrl,omitempty"`
	Frequency                    string           `json:"frequency,omitempty"`
	Headway                      int              `json:"headway,omitempty"`
	ID                           string           `json:"id,omitempty"`
	Lang                         string           `json:"lang,omitempty"`
	LastKnownDistanceAlongTrip   int              `json:"lastKnownDistanceAlongTrip,omitempty"`
	LastKnownLocation            Location         `json:"lastKnownLocation,omitempty"`
	LastKnownOrientation         int              `json:"lastKnownOrientation,omitempty"`
	LastLocationUpdateTime       int              `json:"lastLocationUpdateTime,omitempty"`
	LastUpdateTime               int              `json:"lastUpdateTime,omitempty"`
	Lat                          float64          `json:"lat,omitempty"`
	LatSpan                      float64          `json:"latSpan,omitempty"`
	LocationType                 int              `json:"locationType,omitempty"`
	Lon                          float64          `json:"lon,omitempty"`
	LongName                     string           `json:"longName,omitempty"`
	LonSpan                      float64          `json:"lonSpan,omitempty"`
	Name                         string           `json:"name,omitempty"`
	NearbyStopIds                []string         `json:"nearbyStopIds,omitempty"`
	NextStop                     string           `json:"nextStop,omitempty"`
	NextStopTimeOffset           int              `json:"nextStopTimeOffset,omitempty"`
	NumberOfStopsAway            int              `json:"numberOfStopsAway,omitempty"`
	Orientation                  float64          `json:"orientation,omitempty"`
	Phase                        string           `json:"phase,omitempty"`
	Phone                        string           `json:"phone,omitempty"`
	PickupType                   int              `json:"pickupType,omitempty"`
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
	SituationIDs                 []string         `json:"situationIds,omitempty"`
	StartTime                    int              `json:"startTime,omitempty"`
	Status                       string           `json:"status,omitempty"`
	StopCalendarDays             List             `json:"stopCalendarDays,omitempty"`
	StopHeadsign                 string           `json:"stopHeadsign,omitempty"`
	StopID                       string           `json:"stopId,omitempty"`
	StopRouteSchedules           List             `json:"stopRouteSchedules,omitempty"`
	StopRouteDirectionSchedules  List             `json:"stopRouteDirectionSchedules,omitempty"`
	StopSequence                 int              `json:"stopSequence,omitempty"`
	StopTime                     *Entry           `json:"stopTime,omitempty"`
	Summary                      []string         `json:"summary,omitempty"`
	TextColor                    string           `json:"textColor,omitempty"`
	Time                         int              `json:"time,omitempty"`
	TimeZone                     string           `json:"timezone,omitempty"`
	TotalDistanceAlongTrip       float64          `json:"totalDistanceAlongTrip"`
	TotalStopsInTrip             int              `json:"totalStopsInTrip,omitempty"`
	TripHeadSign                 string           `json:"tripHeadsign,omitempty"`
	TripID                       string           `json:"tripId,omitempty"`
	TripShortName                string           `json:"tripShortName,omitempty"`
	TripStatus                   *Entry           `json:"tripStatus,omitempty"`
	Type                         int              `json:"type,omitempty"`
	URL                          string           `json:"url,omitempty"`
	VehicleID                    string           `json:"vehicleId,omitempty"`
	WheelChairBoarding           string           `json:"wheelchairBoarding,omitempty"`
	//Description       []string         `json:"description>value"` what the fuck
}

func (e Entry) AgencyFromEntry() *Agency {
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

func (e Entry) AgencyWithCoverageFromEntry(a Agency) *AgencyWithCoverage {
	return &AgencyWithCoverage{
		Agency:  a,
		Lat:     e.Lat,
		LatSpan: e.LatSpan,
		Lon:     e.Lon,
		LonSpan: e.LonSpan,
	}
}

func (e Entry) ArrivalAndDepartureFromEntry() *ArrivalAndDeparture {
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
		TripStatus:                   e.TripStatusFromEntry(),
		VehicleID:                    e.VehicleID,
	}
}

func (e Entry) BlockFromEntry() *Block {
	bcs := e.Configurations.toBlockConfigurations()
	return &Block{
		ID:             e.ID,
		Configurations: bcs,
	}
}

func (e Entry) BlockConfigurationFromEntry() *BlockConfiguration {
	return &BlockConfiguration{}
}

func (e Entry) BlockStopTimeFromEntry() *BlockStopTime {
	return &BlockStopTime{
		BlockSequence:        e.BlockSequence,
		DistanceAlongBlock:   e.DistanceAlongBlock,
		AccumulatedSlackTime: e.AccumulatedSlackTime,
		StopTime:             *e.StopTime.StopTimeFromEntry(),
	}
}

func (e Entry) BlockTripFromEntry() *BlockTrip {
	bsts := e.BlockStopTimes.toBlockStopTimes()
	return &BlockTrip{
		AccumulatedSlackTime: e.AccumulatedSlackTime,
		BlockStopTimes:       bsts,
		DistanceAlongBlock:   e.DistanceAlongBlock,
		TripID:               e.TripID,
	}
}

func (e Entry) CoverageFromEntry() Coverage {
	return Coverage{
		Lat:     e.Lat,
		Lon:     e.Lon,
		LatSpan: e.LatSpan,
		LonSpan: e.LonSpan,
	}
}

func (e Entry) CurrentTimeFromEntry() *CurrentTime {
	return &CurrentTime{
		ReadableTime: e.ReadableTime,
		Time:         e.Time,
	}
}

func (e Entry) FrequencyFromEntry() *Frequency {
	return &Frequency{
		StartTime: e.StartTime,
		EndTime:   e.EndTime,
		Headway:   e.Headway,
	}
}

func (e Entry) LocationFromEntry() *Location {
	return &Location{
		Lat: e.Lat,
		Lon: e.Lon,
	}
}

func (e Entry) RegisteredAlarmFromEntry() *RegisteredAlarm {
	return &RegisteredAlarm{
		AlarmID: e.AlarmID,
	}
}

func (e Entry) RouteFromEntry(agencies []Agency) *Route {
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

func (e Entry) ScheduleStopTimeFromEntry() *ScheduleStopTime {
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

func (e Entry) SituationFromEntry() *Situation {
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

func (e Entry) StopFromEntry() *Stop {
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

func (e Entry) StopWithArrivalsAndDeparturesFromEntry(a ArrivalsAndDepartures) *StopWithArrivalsAndDepartures {
	return &StopWithArrivalsAndDepartures{
		StopID:                e.StopID,
		ArrivalsAndDepartures: a,
		NearByStopIDs:         e.NearbyStopIds,
	}
}

func (e Entry) StopRouteScheduleFromEntry(r Route) *StopRouteSchedule {
	return &StopRouteSchedule{
		StopRouteDirectionSchedules: e.StopRouteDirectionSchedules.toStopRouteDirectionSchedules(),
		Route: r,
	}
}

func (e Entry) StopRouteDirectionScheduleFromEntry() *StopRouteDirectionSchedule {
	return &StopRouteDirectionSchedule{
		ScheduleFrequencies: e.ScheduleFrequencies.toScheduleFrequencies(),
		TripHeadsign:        e.TripHeadSign,
		ScheduleStopTimes:   e.ScheduleStopTimes.toScheduleStopTimes(),
	}
}

func (e Entry) ScheduleFrequencyFromEntry() *ScheduleFrequency {
	return &ScheduleFrequency{}
}

func (e Entry) StopScheduleFromEntry(ss []Stop) *StopSchedule {
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

func (e Entry) StopTimeFromEntry() *StopTime {
	return &StopTime{
		ArrivalTime:   e.ArrivalTime,
		DepartureTime: e.DepartureTime,
		DropOffType:   e.DropOffTime,
		PickupType:    e.PickupType,
		StopID:        e.StopID,
	}
}

func (e Entry) TripFromEntry() *Trip {
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

func (e Entry) TripStatusFromEntry() *TripStatus {
	return &TripStatus{
		ActiveTripID:               e.ActiveTripID,
		BlockTripSequence:          e.BlockTripSequence,
		ClosestStop:                e.ClosestStop,
		ClosestStopTimeOffset:      e.ClosestStopTimeOffset,
		DistanceAlongTrip:          e.DistanceAlongTrip,
		Frequency:                  e.Frequency,
		LastKnownDistanceAlongTrip: e.LastKnownDistanceAlongTrip,
		LastKnownLocation:          e.LastKnownLocation,
		LastKnownOrientation:       e.LastKnownOrientation,
		LastLocationUpdateTime:     e.LastLocationUpdateTime,
		LastUpdateTime:             e.LastUpdateTime,
		NextStop:                   e.NextStop,
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
