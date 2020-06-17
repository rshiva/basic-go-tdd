package main

import "testing"

type StubStore interface{
	response string
}

func (s *stubsStubStore) Fetch() string{
	return s.response
}

func TestHandler(t *testing.T){
	data := "hello, world"
	svr := Server(&StubStore{data})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	svr.ServeHTTP(response, request)

	if response.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
	}
}