package apit

import(
	"testing"
	"net/http"
	"net/http/httptest"
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"fmt"
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
	if req == nil {
		t.Fatal("invalid path")
	}
	if header != nil {
		req.Header = header
	}
	w := httptest.NewRecorder()
	curGin.ServeHTTP(w, req)
	assert.Equal(t, expectedCode, w.Code)

	if responseBodyStruct != nil {
		body := w.Body.String()
		if err := json.Unmarshal([]byte(body), &responseBodyStruct); err != nil {
			fmt.Println("error body ::: ",body)
			t.FailNow()
		}
	}

	return w
}