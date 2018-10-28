// Package oba - One Bus Away Go Api https://onebusaway.org/
// Author: Seth T <setheck@gmail.com>
package oba

// Client - Interface for a One Bus Away Client
type Client interface {
	AgenciesWithCoverage() ([]AgencyWithCoverage, error)
	Agency(id string) (*Agency, error)
	ArrivalAndDepartureForStop(id string, params map[string]string) (*ArrivalAndDeparture, error)
	ArrivalsAndDeparturesForStop(id string, params map[string]string) (*StopWithArrivalsAndDepartures, error)
	Block(id string) (*Block, error)
	CancelAlarm(id string) error
	CurrentTime() (*CurrentTime, error)
	RegisterAlarmForArrivalAndDepartureAtStop(id string, params map[string]string) (*RegisteredAlarm, error)
	ReportProblemWithStop(id string, params map[string]string) error
	ReportProblemWithTrip(id string, params map[string]string) error
	RouteIdsForAgency(id string) ([]string, error)
	Route(id string) (*Route, error)
	RoutesForAgency(id string) ([]Route, error)
	RoutesForLocation(params map[string]string) ([]Route, error)
	ScheduleForStop(id string) (*StopSchedule, error)
	Shape(id string) (*Shape, error)
	StopIDsForAgency(id string) ([]string, error)
	Stop(id string) (*Stop, error)
	StopsForLocation(params map[string]string) ([]Stop, error)
	StopsForRoute(id string) (*StopsForRoute, error)
	TripDetails(id string) (*TripDetails, error)
	TripForVehicle(id string, params map[string]string) (*TripDetails, error)
	Trip(id string) (*Trip, error)
	TripsForLocation(params map[string]string) ([]TripDetails, error)
	TripsForRoute(id string) ([]TripDetails, error)
	VehiclesForAgency(id string) ([]VehicleStatus, error)
}
