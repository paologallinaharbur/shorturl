package middlewares

import (
	"github.com/paologallinaharbur/shorturl/storage"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestMiddlewareUI(t *testing.T) {

	test := UIMiddleware(nil)
	req, _ := http.NewRequest("GET", "http://test:90/", nil)
	response := httptest.NewRecorder()
	test.ServeHTTP(response, req)

	//We expect the handler to redirect us
	assert.Equal(t, 302, response.Code)
	assert.Equal(t, "/swagger-ui/", response.HeaderMap["Location"][0])

}

func TestMiddlewarePrometeus(t *testing.T) {

	test := PrometheusMiddleware(nil)
	req, _ := http.NewRequest("GET", "http://test:90/metrics", nil)
	response := httptest.NewRecorder()
	test.ServeHTTP(response, req)

	body, _ := ioutil.ReadAll(response.Body)
	assert.True(t, strings.Contains(string(body), "url_redirected{version=\"1.0.0\"}"))

}

func TestMiddlewareRedirect(t *testing.T) {
	st := storage.NewStorageDB("test")
	defer os.Remove("test")

	st.Write("short-test", "https://logURLTest")

	test := RedirectMiddleware(nil, st)
	req, _ := http.NewRequest("GET", "http://test:90/short-test", nil)
	response := httptest.NewRecorder()
	test.ServeHTTP(response, req)

	//We expect the handler to redirect us since we added the short url
	assert.Equal(t, 303, response.Code)
	assert.Equal(t, "https://logURLTest", response.HeaderMap["Location"][0])

	test = RedirectMiddleware(nil, st)
	req, _ = http.NewRequest("GET", "http://test:90/short-notExisting", nil)
	response = httptest.NewRecorder()
	test.ServeHTTP(response, req)

	//We expect the handler to redirect us to http://www.notfound.com since the short url was not present
	assert.Equal(t, 303, response.Code)
	assert.Equal(t, "http://www.notfound.com", response.HeaderMap["Location"][0])

}
