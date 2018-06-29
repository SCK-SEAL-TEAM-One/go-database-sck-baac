package api

import (
	"bytes"
	"encoding/json"
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
func Test_GetBreadCreate_With_Stub_Input_Key_Applepie_Value_USA_Should_Be_Create_Successful(t *testing.T) {
	stubWriteKeyFunc := func(string, interface{}) error {
		return nil
	}
	read := GetBreadCreate(stubWriteKeyFunc)
	url := "/redis/create"
	requestBody := map[string]string{
		"key":   "Applepie",
		"value": "USA",
	}
	requestBodyString, _ := json.Marshal(requestBody)
	request := httptest.NewRequest("POST", url, bytes.NewBuffer(requestBodyString))
	responseRecorder := httptest.NewRecorder()
	expected := `Create Successful`

	read(responseRecorder, request)
	response := responseRecorder.Result()
	actual, _ := ioutil.ReadAll(response.Body)

	if string(actual) != expected {
		t.Errorf("Should Be %s but it got %s", expected, string(actual))
	}

}

func Test_GetDelete_With_Stub_Input_Key_Applepie_Should_Be_Delete_Successful(t *testing.T) {
	stubDeleteKeyFunc := func(string) (int64, error) {
		return 1, nil
	}
	read := GetDelete(stubDeleteKeyFunc)
	url := "/redis/delete"
	requestBody := map[string]string{
		"key": "Applepie",
	}
	requestBodyString, _ := json.Marshal(requestBody)
	request := httptest.NewRequest("DELETE", url, bytes.NewBuffer(requestBodyString))
	responseRecorder := httptest.NewRecorder()
	expected := `Delete Successful`

	read(responseRecorder, request)
	response := responseRecorder.Result()
	actual, _ := ioutil.ReadAll(response.Body)

	if string(actual) != expected {
		t.Errorf("Should Be %s but it got %s", expected, string(actual))
	}

}
