package routes

import (
"net/http"
"net/http/httptest"
"testing"
"github.com/stretchr/testify/assert"
"github.com/gin-gonic/gin"
 "github.com/go-resty/resty/v2"
 "github.com/jarcoal/httpmock"
)

type Player struct {
 Player_id string `json:"player_id"`
 Player_name string `json:"player_name"`
 Player_position string `json:"player_position"`
}

var respRec *httptest.ResponseRecorder

var err error

var req *http.Request

func newResponder(s int, c string, ct string) httpmock.Responder {
	resp := httpmock.NewStringResponse(s, c)
	resp.Header.Set("Content-Type", ct)

	return httpmock.ResponderFromResponse(resp)
}

func Test_GetSinglePlayer(t *testing.T) {
w := httptest.NewRecorder()

// Init Router
router := gin.Default()
Routes(router)


req, _ := http.NewRequest("GET", "/api/team/2", nil)
router.ServeHTTP(w, req)

assert.Equal(t, 200, w.Code)
}

func Test_GetSinglePlayerWithResponseBody(t *testing.T) {
    router := gin.Default()
    Routes(router)
    respRec = httptest.NewRecorder()
    rst := resty.New()
    httpmock.ActivateNonDefault(rst.GetClient())
    defer httpmock.DeactivateAndReset()

    httpmock.RegisterResponder(
      "GET",
      "/api/team/1",
      newResponder(200, `{"player_id":"7" , "player_name":"Ronaldo", "player_position":"striker"}`, "application/json"),
     )

    resp, _ := rst.R().
    SetResult(&Player{}).
    Get("/api/team/1")

    req, err = http.NewRequest("GET", "/api/team/1", nil)

    if err != nil {
        t.Fatal("Creating 'GET /api/team/1' request failed!")
    }

    router.ServeHTTP(respRec, req)

    if respRec.Code != http.StatusOK {
        t.Fatal("Server error: Returned ", respRec.Code, " instead of ", http.StatusBadRequest)
    }

  assert.Equal(t, &Player{Player_id:"7", Player_name:"Ronaldo", Player_position:"striker"}, resp.Result().(*Player))
  }