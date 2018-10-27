package oba

type List []entry
type AltList []altEntry

func (l List) toAgencies() []Agency {
	agencies := make([]Agency, 0, len(l))
	for _, entry := range l {
		agencies = append(agencies, *entry.agencyFromEntry())
	}
	return agencies
}

func (l List) toAgenciesWithCoverage(agencies []Agency) []AgencyWithCoverage {
	agmap := make(map[string]Agency)
	for _, a := range agencies {
		agmap[a.ID] = a
	}

	awcs := make([]AgencyWithCoverage, 0, len(l))
	for _, entry := range l {
		if v, ok := agmap[entry.AgencyID]; ok {
			awc := *entry.agencyWithCoverageFromEntry(v)
			awcs = append(awcs, awc)
		}
	}
	return awcs
}

func (l List) toArrivalAndDepartures(sits []Situation, st []Stop, ts []Trip) []ArrivalAndDeparture {
	aads := make([]ArrivalAndDeparture, 0, len(l))
	for _, entry := range l {
		aads = append(aads, *entry.arrivalAndDepartureFromEntry(sits, st, ts))
	}
	return aads
}

func (l List) toBlockConfigurations() []BlockConfiguration {
	bcs := make([]BlockConfiguration, 0, len(l))
	for _, c := range l {
		bcs = append(bcs, *c.blockConfigurationFromEntry(c.ActiveServiceID, c.InactiveServiceID, c.Trips.toBlockTrips()))
	}
	return bcs
}

func (l List) toBlockTrips() []BlockTrip {
	bts := make([]BlockTrip, 0, len(l))
	for _, entry := range l {
		bts = append(bts, *entry.blockTripFromEntry())
	}
	return bts
}

func (l List) toBlockStopTimes() []BlockStopTime {
	bsts := make([]BlockStopTime, 0, len(l))
	for _, entry := range l {
		bsts = append(bsts, *entry.blockStopTimeFromEntry())
	}
	return bsts
}

func (l AltList) toEncodedPolyLines() []EncodedPolyLine {
	epls := make([]EncodedPolyLine, 0, len(l))
	for _, entry := range l {
		epls = append(epls, *entry.encodedPolyLineFromEntry())
	}
	return epls
}

func (l List) toRoutes(a []Agency) []Route {
	routes := make([]Route, 0, len(l))
	for _, entry := range l {
		route := *entry.routeFromEntry(a)
		routes = append(routes, route)
	}
	return routes
}

func (l List) toSituations() []Situation {
	sits := make([]Situation, 0, len(l))
	for _, entry := range l {
		sits = append(sits, *entry.situationFromEntry())
	}
	return sits
}

func (l List) toStops(r []Route) []Stop {
	var routes map[string]Route
	stops := make([]Stop, 0, len(l))
	if r != nil {
		routes = make(map[string]Route)
		for _, route := range r {
			routes[route.ID] = route
		}
	}
	for _, entry := range l {
		stop := *entry.stopFromEntry()
		stopRoutes := make([]Route, 0, len(routes))
		for _, rid := range entry.RouteIDs {
			if route, ok := routes[rid]; ok {
				stopRoutes = append(stopRoutes, route)
			}
		}
		stop.Routes = stopRoutes
		stops = append(stops, stop)
	}
	return stops
}

func (l List) toStopRouteDirectionSchedules() []StopRouteDirectionSchedule {
	srds := make([]StopRouteDirectionSchedule, 0, len(l))
	for _, entry := range l {
		srds = append(srds, *entry.stopRouteDirectionScheduleFromEntry())
	}
	return srds
}

func (l List) toScheduleFrequencies() []ScheduleFrequency {
	sf := make([]ScheduleFrequency, 0, len(l))
	for _, entry := range l {
		sf = append(sf, *entry.scheduleFrequencyFromEntry())
	}
	return sf
}

func (l List) toStopRouteSchedules(rs []Route) []StopRouteSchedule {
	srs := make([]StopRouteSchedule, 0, len(l))
	routes := make(map[string]Route)
	for _, route := range rs {
		routes[route.ID] = route
	}
	for _, entry := range l {
		if v, ok := routes[entry.RouteID]; ok {
			srs = append(srs, *entry.stopRouteScheduleFromEntry(v))
		}
	}
	return srs
}

func (l List) toScheduleStopTimes() []ScheduleStopTime {
	ssts := make([]ScheduleStopTime, 0, len(l))
	for _, sst := range l {
		ssts = append(ssts, *sst.scheduleStopTimeFromEntry())
	}
	return ssts
}

func (l AltList) toStopGroupings(ss []Stop) []StopGrouping {
	sgs := make([]StopGrouping, 0, len(l))
	for _, entry := range l {
		sgs = append(sgs, *entry.stopGroupingFromEntry(ss))
	}
	return sgs
}

func (l AltList) toStopGroups(ss []Stop) []StopGroup {
	sgs := make([]StopGroup, 0, len(l))
	for _, entry := range l {
		sgs = append(sgs, *entry.stopGroupFromEntry(ss))
	}
	return sgs
}

func (l List) toTrips() []Trip {
	trips := make([]Trip, 0, len(l))
	for _, entry := range l {
		trips = append(trips, *entry.tripFromEntry())
	}
	return trips
}

func (l List) toTripDetails(ts []Trip, ss []Situation) []TripDetails {
	tds := make([]TripDetails, 0, len(l))
	for _, entry := range l {
		tds = append(tds, *entry.tripDetailsFromEntry(ts, ss))
	}
	return tds
}

func (l List) toVehicleStatuses(st []Stop, ts []Trip) []VehicleStatus {
	vss := make([]VehicleStatus, 0, len(l))
	for _, entry := range l {
		vhs := *entry.vehicleStatusFromEntry(st, ts)
		vss = append(vss, vhs)
	}
	return vss
}
