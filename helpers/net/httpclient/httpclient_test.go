package httpclient

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	// "log"
)

var (
	reqBody = `{
		"test": "test",
	}`
	dummyHandler = func(w http.ResponseWriter, r *http.Request) {
		var bodyData struct {
			Sleep int `json:"sleep"`
		}
		switch r.Method {
		case "GET":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Hello"))
		case "POST":
			body, err := ioutil.ReadAll(r.Body)
			json.Unmarshal(body, &bodyData)
			if err != nil {
				http.Error(w, "can't read body", http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(body))
		case "PUT":
			body, err := ioutil.ReadAll(r.Body)
			json.Unmarshal(body, &bodyData)
			if err != nil {
				http.Error(w, "can't read body", http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(body))
		case "PATCH":
			body, err := ioutil.ReadAll(r.Body)
			json.Unmarshal(body, &bodyData)
			if err != nil {
				http.Error(w, "can't read body", http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(body))
		case "DELETE":
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Deleted"))
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			return	
		}
	}
)

func TestHttpDoGet(t *testing.T) {
	var httpDoer HttpDoer
	var httpParam HttpParam
	json.Unmarshal([]byte(reqBody), &httpParam)
	
	server := httptest.NewServer(http.HandlerFunc(dummyHandler))
	defer server.Close()

	httpParam.Url = server.URL
	httpParam.Method = "get"
	httpDoer = &httpParam
	response, err := httpDoer.HttpDo()
	require.NoError(t, err, "should not have failed to make a Get request")
	body, err := ioutil.ReadAll(response.Body)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "Hello", string(body))
}

func TestHttpDoPost(t *testing.T) {
	var httpDoer HttpDoer
	var httpParam HttpParam
	json.Unmarshal([]byte(reqBody), &httpParam)
	
	server := httptest.NewServer(http.HandlerFunc(dummyHandler))
	defer server.Close()

	httpParam.Url = server.URL
	httpParam.Method = "post"
	httpDoer = &httpParam
	response, err := httpDoer.HttpDo()
	require.NoError(t, err, "should not have failed to make a POST request")
	body, err := ioutil.ReadAll(response.Body)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, httpParam.Body, string(body))
}

func TestHttpDoPut(t *testing.T) {
	var httpDoer HttpDoer
	var httpParam HttpParam
	json.Unmarshal([]byte(reqBody), &httpParam)
	
	server := httptest.NewServer(http.HandlerFunc(dummyHandler))
	defer server.Close()

	httpParam.Url = server.URL
	httpParam.Method = "put"
	httpDoer = &httpParam
	response, err := httpDoer.HttpDo()
	require.NoError(t, err, "should not have failed to make a POST request")
	body, err := ioutil.ReadAll(response.Body)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, httpParam.Body, string(body))
}

func TestHttpDoPatch(t *testing.T) {
	var httpDoer HttpDoer
	var httpParam HttpParam
	json.Unmarshal([]byte(reqBody), &httpParam)
	
	server := httptest.NewServer(http.HandlerFunc(dummyHandler))
	defer server.Close()

	httpParam.Url = server.URL
	httpParam.Method = "patch"
	httpDoer = &httpParam
	response, err := httpDoer.HttpDo()
	require.NoError(t, err, "should not have failed to make a POST request")
	body, err := ioutil.ReadAll(response.Body)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, httpParam.Body, string(body))
}

func TestHttpDoDelete(t *testing.T) {
	var httpDoer HttpDoer
	var httpParam HttpParam
	json.Unmarshal([]byte(reqBody), &httpParam)
	
	server := httptest.NewServer(http.HandlerFunc(dummyHandler))
	defer server.Close()

	httpParam.Url = server.URL
	httpParam.Method = "delete"
	httpDoer = &httpParam
	response, err := httpDoer.HttpDo()
	require.NoError(t, err, "should not have failed to make a Get request")
	body, err := ioutil.ReadAll(response.Body)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, "Deleted", string(body))
}

func TestHttpDoDefault(t *testing.T) {
	var httpDoer HttpDoer
	var httpParam HttpParam
	json.Unmarshal([]byte(reqBody), &httpParam)
	
	server := httptest.NewServer(http.HandlerFunc(dummyHandler))
	defer server.Close()

	httpParam.Url = server.URL
	httpParam.Method = ""
	httpDoer = &httpParam
	response, err := httpDoer.HttpDo()
	require.NoError(t, err, "should not have failed to make a POST request")
	body, err := ioutil.ReadAll(response.Body)
	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, httpParam.Body, string(body))
}

func TestMakeHeader(t *testing.T) {
	var httpParam HttpParam

	httpParam.Header = make(map[string]string)
	httpParam.Header["Content-Type"]="application/json"
	
	header := makeHeader(httpParam.Header)
	assert.Equal(t, "application/json", header.Get("Content-Type"))
}
