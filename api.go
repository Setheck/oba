//Package oba - One Bus Away Go Api https://onebusaway.org/
// Author: Seth T <setheck@gmail.com>
package oba

//Client - Interface for a One Bus Away Client
type Client interface {
	AgenciesWithCoverage() (AgencyWithCoverage, error)
	Agency(id string) (Entry, error)
	ArrivalAndDepartureForStop(id string, params map[string]string) (Data, error)
	ArrivalsAndDeparturesForStop(id string) (Data, error)
	Block(id string) (Entry, error)
	CancelAlarm(id string) error
	CurrentTime() (Time, error)
	RegisterAlarmForArrivalAndDepartureAtStop(id string, params map[string]string) (RegisteredAlarm, error)
	//ReportProblemWithStop() - Appears to be broken?
	ReportProblemWithTrip(id string, params map[string]string) error
	RouteIdsForAgency(id string) ([]string, error)
	Route(id string) (Entry, error)
	RoutesForAgency(id string) ([]Route, error)
	RoutesForLocation(params map[string]string) (Data, error)
	ScheduleForStop(id string) (Data, error)
	Shape(id string) (Entry, error)
	StopIDsForAgency(id string) ([]string, error)
	Stop(id string) (Entry, error)
	StopsForLocation(params map[string]string) (Data, error)
	StopsForRoute(id string) (Entry, error)
	TripDetails(id string) (Entry, error)
	TripForVehicle(id string, params map[string]string) (Data, error)
	Trip(id string) (Entry, error)
	TripsForLocation(params map[string]string) (Data, error)
	TripsForRoute(id string) (Data, error)
	VehiclesForAgency(id string) (Data, error)
}
