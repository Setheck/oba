package oba

import "fmt"

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

type AltResponse struct {
	Code        int      `json:"code"`
	CurrentTime int      `json:"currentTime"`
	Data        *AltData `json:"data,omitempty"`
	Text        string   `json:"text"`
	Version     int      `json:"version"`
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
	Agencies   List `json:"agencies"`
	Routes     List `json:"routes"`
	Situations List `json:"situations"`
	Stops      List `json:"stops"`
	Trips      List `json:"trips"`
}

//Data container object
type Data struct {
	LimitExceeded *bool       `json:"limitExceeded,omitempty"`
	List          *List       `json:"list,omitempty"`
	Entry         *Entry      `json:"entry,omitempty"`
	OutOfRange    *bool       `json:"outOfRange,omitempty"`
	References    *References `json:"references"`
	Time          *Time       `json:",omitempty"`
	//StopsForRoute *StopsForRoute
}

type AltData struct {
	List []string `json:"list,omitempty"`
}

func (d Data) toTripDetails() []TripDetails {
	ss := d.References.Situations.toSituations()
	ts := d.References.Trips.toTrips()
	tds := d.List.toTripDetails(ts, ss)
	return tds
}

func (d Data) TripDetails() *TripDetails {
	ss := d.References.Situations.toSituations()
	ts := d.References.Trips.toTrips()
	td := d.Entry.TripDetailsFromEntry(ts, ss)
	return td
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
