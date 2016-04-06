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

func PerformRequest(t *testing.T, method, path string, header http.Header, requestBody interface{}, responseBodyStruct interface{}, respErr interface{}, expectedCode int) *httptest.ResponseRecorder {
	if curGin == nil {
		t.Fatal("Please SetGin()...")
	}

	byteJson, err := json.Marshal(requestBody)
	if err != nil {
		t.Fatal(err)
	}

	req, _ := http.NewRequest(method, path, bytes.NewBuffer(byteJson))
	if req == nil {
		t.Fatalf("invalid path. path = %s", path)
	}
	if header != nil {
		req.Header = header
	}
	w := httptest.NewRecorder()
	curGin.ServeHTTP(w, req)
	assert.Equal(t, expectedCode, w.Code)

	body := w.Body.String()
	if responseBodyStruct != nil && w.Code >= 200 && w.Code <= 210 {
		if err := json.Unmarshal([]byte(body), &responseBodyStruct); err != nil {
			assert.FailNow(t, "error body ::: ",body, w.Code)
		}
	}
	if respErr != nil && w.Code >= 400 && w.Code <= 500{
		if err := json.Unmarshal([]byte(body), &respErr); err != nil {
			assert.FailNow(t, "error body ::: ",body, w.Code)
		}
	}

	return w
}