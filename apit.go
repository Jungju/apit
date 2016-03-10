package apit

import(
	"testing"
	"net/http"
	"net/http/httptest"
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
	"encoding/json"
)

var(
	curGin *gin.Engine
)

func SetGin(gin *gin.Engine) {
	curGin = gin
}

func PerformRequest(t *testing.T, method, path string, header http.Header, requestBody interface{}, responseBodyStruct interface{}, expectedCode int) *httptest.ResponseRecorder {
	if curGin == nil {
		t.Fatal("Please SetGin()...")
	}

	byteJson, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	req, _ := http.NewRequest(method, path, bytes.NewBuffer(byteJson))
	if header != nil {
		req.Header = header
	}
	w := httptest.NewRecorder()
	curGin.ServeHTTP(w, req)
	assert.Equal(t, expectedCode, w.Code)

	if responseBodyStruct != nil {
		if err := json.Unmarshal([]byte(w.Body.String()), &responseBodyStruct); err != nil {
			t.Fatal(err)
		}
	}

	return w
}