package oba_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Setheck/oba"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
)

const (
	TestDataPath = "testdata/"
)

func FakeServer(t *testing.T, body []byte) *httptest.Server {
	t.Helper()
	handler := http.HandlerFunc(func(r http.ResponseWriter, req *http.Request) {
		r.Write(body)
	})
	return httptest.NewServer(handler)
}

func RetrieveTestJsonFileContent(t *testing.T) []byte {
	t.Helper()
	file := ConvertToFilename(t.Name())
	return ReadFile(t, file)
}

func ReadFile(t *testing.T, file string) []byte {
	t.Helper()
	buf, e := ioutil.ReadFile(fmt.Sprint(TestDataPath, file))
	if e != nil {
		t.Error(e)
	}
	return buf
}

func ConvertToFilename(s string) string {
	if strings.Contains(s, "_") {
		split := strings.Split(s, "_")
		s = split[len(split)-1]
	}
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
	return strings.ToLower(s) + ".json"
}

func VerifyMarshalling(t *testing.T, data []byte) {
	t.Helper()

	var resp oba.Response
	if err := json.Unmarshal(data, &resp); err != nil {
		t.Error(err)
	}

	m, err := json.MarshalIndent(resp, "", "  ")
	if err != nil {
		t.Error(err)
	}

	m = FixJSON(m) // TODO SetEscapeHTML(false) see https://golang.org/pkg/encoding/json/

	m = bytes.TrimSpace(m)
	o := bytes.TrimSpace(data)

	if bytes.Compare(o, m) != 0 {
		t.Error("VerifyMarshalling Failed")
		fmt.Println(string(m))
	}
}

func FixJSON(b []byte) []byte {
	return bytes.Replace(b, []byte("\u0026"), []byte("&"), -1)
}
