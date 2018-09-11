//Package oba - One Bus Away Go Api https://onebusaway.org/
// Author: Seth T <setheck@gmail.com>
package oba_test

import (
	"fmt"
	"github.com/setheck/oba"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
)

const (
	testdata = "testdata/"
)

func FakeServer(t *testing.T) *httptest.Server {

	file := ConvertToFilename(t.Name())
	buf, e := ioutil.ReadFile(fmt.Sprint(testdata, file))
	if e != nil {
		t.Error(e)
	}

	handler := http.HandlerFunc(func(r http.ResponseWriter, req *http.Request) {
		r.Write(buf)
	})
	return httptest.NewServer(handler)
}

func TestAgenciesWithCoverage(t *testing.T) {
	server := FakeServer(t)
	defer server.Close()

	client := oba.NewDefaultClient(server.URL, "key")
	d, e := client.AgenciesWithCoverage()
	if e != nil {
		t.Fail()
	}

	if *d.Class != "listWithReferences" {
		t.Fail()
	}
	for _, agency := range d.References.Agencies {
		if agency.TimeZone == "" {
			t.Fail()
		}
	}
	fmt.Println(d)
}

func ConvertToFilename(s string) string {

	if strings.HasPrefix(s, "Test") {
		s = strings.TrimPrefix(s, "Test")
	}
	r := regexp.MustCompile("^(?:.)|([A-Z][a-z])")
	matches := r.FindAllStringIndex(s, 10)

	slider := 0
	for _, m := range matches {
		if m[0] > 0 {
			s = fmt.Sprintf("%s-%s", s[0:m[0]+slider], s[m[0]+slider:])
			slider++
		}
	}
	return strings.ToLower(s) + ".xml"
}
