package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	m.Run()
}

func testRequest(t *testing.T, method string, url string, reqData interface{}, wantHTTPStatusCode int, wantCode int) {
	w := httptest.NewRecorder()

	body, err := json.Marshal(reqData)
	assert.NoError(t, err)

	buffer := bytes.NewBuffer(body)

	req, err := http.NewRequest(method, url, buffer)
	assert.NoError(t, err)
	//token, err := jwt.GenerateToken(conf.ROOT.System.Auth.Username)
	//assert.NoError(t, err)
	//req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", token))
	eng.ServeHTTP(w, req)
	assert.Equal(t, wantHTTPStatusCode, w.Code)
	type RespData struct {
		Code int `json:"code"`
	}
	var resp RespData
	raw := w.Body.Bytes()
	t.Logf("resp body = %v", string(raw))
	err = json.Unmarshal(raw, &resp)
	assert.NoError(t, err)
	assert.Equal(t, int(wantCode), resp.Code)
}
