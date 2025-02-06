package httpserver

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"url-short/internal/transport/http_server/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type test struct {
 name           string
 requestBody    UrlRequest
 msgErr         string
 mockReturn     string
 mockError      error
 expectedCode   int
 expectedAlias  string
}

func TestCreate(t *testing.T) {
 mockService := new(mocks.UrlService)

 tests := []test{
  {
   name: "valid request without alias",
   requestBody: UrlRequest{
    Url: "http://example.com",
   },
   mockReturn:    "ddsfs",
   mockError:     nil,
   expectedCode:  http.StatusOK,
   expectedAlias: "ddsfs",
   msgErr:        "",
  },
  {
   name: "valid request with alias",
   requestBody: UrlRequest{
    Url:   "http://example.com",
    Alias: "sdsdsd",
   },
   mockReturn:    "sdsdsd",
   mockError:     nil,
   expectedCode:  http.StatusOK,
   expectedAlias: "sdsdsd",
   msgErr:        "",
  },
  {
   name: "validation error require",
   requestBody: UrlRequest{
    Url:   "",
    Alias: "sdsdsd",
   },
   mockReturn:    "",
   mockError:     fmt.Errorf("validation error"),
   expectedCode:  http.StatusBadRequest,
   expectedAlias: "",
   msgErr:        "invalid request: field is required: Url",
  },
  {
   name: "validation error url",
   requestBody: UrlRequest{
    Url:   "sdfsdfsdf",
    Alias: "sdsdsd",
   },
   mockReturn:    "",
   mockError:     fmt.Errorf("validation error"),
   expectedCode:  http.StatusBadRequest,
   expectedAlias: "",
   msgErr:        "invalid request: field doesn`t type url",
  },
  {
   name: "server error",
   requestBody: UrlRequest{
    Url:   "https://www.yandex.com/",
    Alias: "sddd",
   },
   mockReturn:    "",
   mockError:     fmt.Errorf("server error"),
   expectedCode:  http.StatusInternalServerError,
   expectedAlias: "",
   msgErr:        "server error",
  },
 }

 for _, tt := range tests {
  t.Run(tt.name, func(t *testing.T) {
   
   if tt.expectedCode == http.StatusOK || tt.expectedCode == http.StatusInternalServerError {
    mockService.On("Create", mock.Anything, tt.requestBody.Url, tt.requestBody.Alias).Return(tt.mockReturn, tt.mockError)
   }

   server := HttpServer{
    UrlService: mockService,
   }

   body, _ := json.Marshal(tt.requestBody)
   
   req := httptest.NewRequest(http.MethodPost, "/create", bytes.NewBuffer(body))
   
   w := httptest.NewRecorder()
   
   server.Create(w, req)

   resp := w.Result()

   if tt.expectedCode == http.StatusOK {
    var response ResponseSuccess
    err := json.NewDecoder(resp.Body).Decode(&response)
    assert.NoError(t, err)
    assert.Equal(t, tt.expectedAlias, response.Alias)
   } else {
    var response ResponseErr
    err := json.NewDecoder(resp.Body).Decode(&response)
    assert.NoError(t, err)
    assert.Equal(t, tt.msgErr, response.Msg)
    assert.Equal(t, tt.expectedCode, resp.StatusCode) // Проверка кода статуса
   }

   mockService.AssertExpectations(t)
  })
 }
}
