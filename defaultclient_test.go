//Package oba - One Bus Away Go Api https://onebusaway.org/
// Author: Seth T <setheck@gmail.com>
package oba_test

import (
	"bytes"
	"encoding/xml"
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

func ReadFile(t *testing.T, file string) []byte {
	t.Helper()
	buf, e := ioutil.ReadFile(fmt.Sprint(testdata, file))
	if e != nil {
		t.Error(e)
	}
	return buf
}

func FakeServer(t *testing.T, body []byte) *httptest.Server {
	t.Helper()
	handler := http.HandlerFunc(func(r http.ResponseWriter, req *http.Request) {
		r.Write(body)
	})
	return httptest.NewServer(handler)
}

func TestAgenciesWithCoverage(t *testing.T) {
	file := ConvertToFilename(t.Name())
	body := ReadFile(t, file)
	marshalVerify(t, body)
	//server := FakeServer(t, body)
	//defer server.Close()
	//
	//client := oba.NewDefaultClient(nil, "key")
	//client.SetBaseURL(server.URL)
	//d, e := client.AgenciesWithCoverage()
	//if e != nil {
	//	t.Error(e)
	//}
	//
	//if d.Class != "listWithReferences" {
	//	t.Fail()
	//}
	//for _, agency := range d.References.Agencies {
	//	if agency.TimeZone == "" {
	//		t.Fail()
	//	}
	//}
	//fmt.Println(d)
	//
	//out, err := xml.Marshal(d)
	//if err != nil {
	//	t.Error(err)
	//}
	//fmt.Println(string(out))
	//
	//if bytes.Compare(out, body) != 0 {
	//	t.Fail()
	//}
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

func marshalVerify(t *testing.T, data []byte) {
	t.Helper()
	var resp oba.Response
	if err := xml.Unmarshal(data, &resp); err != nil {
		t.Error(err)
	}

	m, err := xml.MarshalIndent(resp, "", "\t")
	if err != nil {
		t.Error(err)
	}

	fmt.Println(string(data))
	fmt.Println(string(m))
	if bytes.Compare(data, m) != 0 {
		t.Fail()
	}
}
