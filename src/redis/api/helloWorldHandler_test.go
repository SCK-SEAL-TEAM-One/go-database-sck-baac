package api

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func Test_HelloWorldReadAll_Should_Be_Get_All_Keys(t *testing.T) {
	stubReadAllKeyFunc := func() ([]map[string]string, error) {
		return []map[string]string{
			{
				"id":          "hello world",
				"description": "hello world",
			},
		}, nil
	}
	readAll := GetReadAll(stubReadAllKeyFunc)
	url := `/helloworld/readAll`
	request := httptest.NewRequest("GET", url, nil)
	responseRecorder := httptest.NewRecorder()
	expected := `[{"description":"hello world","id":"hello world"}]`

	readAll(responseRecorder, request)
	response := responseRecorder.Result()
	actual, _ := ioutil.ReadAll(response.Body)

	if string(actual) != expected {
		t.Errorf("Should Be %s but it got %s", expected, actual)
	}
}

func Test_HelloWorldReadById_Input_id_hello_world_Should_Be_hello_world(t *testing.T) {
	stubReadByIdFunc := func(string) (map[string]string, error) {
		return map[string]string{
			"id":          "hello world",
			"description": "hello world",
		}, nil
	}
	readById := GetReadById(stubReadByIdFunc)
	url := `/helloworld/readById?id=hello%20world`
	request := httptest.NewRequest("GET", url, nil)
	responseRecorder := httptest.NewRecorder()
	expected := `{"description":"hello world","id":"hello world"}`

	readById(responseRecorder, request)
	response := responseRecorder.Result()
	actual, _ := ioutil.ReadAll(response.Body)

	if string(actual) != expected {
		t.Errorf("Should Be %s but it got %s", expected, actual)
	}
}
