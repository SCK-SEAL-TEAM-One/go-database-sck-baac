package api

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func Test_GetRead_With_Stub_Input_Key_Melonbread_Should_Be_Japan(t *testing.T) {
	stubReadKeyFunc := func(string) (string, error) {
		return "Japan", nil
	}
	read := GetRead(stubReadKeyFunc)
	url := "/redis/read?key=Melonbread"
	request := httptest.NewRequest("GET", url, nil)
	responseRecorder := httptest.NewRecorder()
	expected := `Japan`

	read(responseRecorder, request)
	response := responseRecorder.Result()
	actual, _ := ioutil.ReadAll(response.Body)

	if string(actual) != expected {
		t.Errorf("Should Be %s but it got %s", expected, string(actual))
	}

}
