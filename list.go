package oba

type List []Entry

func (l List) toAgencies() []Agency {
	agencies := make([]Agency, 0, len(l))
	for _, entry := range l {
		agencies = append(agencies, *entry.AgencyFromEntry())
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
			awc := *entry.AgencyWithCoverageFromEntry(v)
			awcs = append(awcs, awc)
		}
	}
	return awcs
}

func (l List) toArrivalAndDepartures() []ArrivalAndDeparture {
	aads := make([]ArrivalAndDeparture, 0)
	for _, entry := range l {
		aads = append(aads, *entry.ArrivalAndDepartureFromEntry())
	}
	return aads
}

func (l List) toBlockConfigurations() []BlockConfiguration {
	bcs := make([]BlockConfiguration, 0, len(l))
	for _, c := range l {
		bcs = append(bcs, *c.BlockConfigurationFromEntry())
	}
	return bcs
}

func (l List) toBlockStopTimes() []BlockStopTime {
	bsts := make([]BlockStopTime, 0, len(l))
	for _, entry := range l {
		bsts = append(bsts, *entry.BlockStopTimeFromEntry())
	}
	return bsts
}

func (l List) toRoutes(a []Agency) []Route {
	routes := make([]Route, 0, len(l))
	for _, entry := range l {
		route := *entry.RouteFromEntry(a)
		routes = append(routes, route)
	}
	return routes
}

func (l List) toSituations() []Situation {
	sits := make([]Situation, 0, len(l))
	for _, entry := range l {
		sits = append(sits, *entry.SituationFromEntry())
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
		stop := *entry.StopFromEntry()
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
		srds = append(srds, *entry.StopRouteDirectionScheduleFromEntry())
	}
	return srds
}

func (l List) toScheduleFrequencies() []ScheduleFrequency {
	sf := make([]ScheduleFrequency, 0, len(l))
	for _, entry := range l {
		sf = append(sf, *entry.ScheduleFrequencyFromEntry())
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
			srs = append(srs, *entry.StopRouteScheduleFromEntry(v))
		}
	}
	return srs
}

func (l List) toScheduleStopTimes() []ScheduleStopTime {
	ssts := make([]ScheduleStopTime, 0, len(l))
	for _, sst := range l {
		ssts = append(ssts, *sst.ScheduleStopTimeFromEntry())
	}
	return ssts
}

func (l List) toTrips() []Trip {
	trips := make([]Trip, 0, len(l))
	for _, entry := range l {
		trips = append(trips, *entry.TripFromEntry())
	}
	return trips
}
