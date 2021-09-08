package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreatePostHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	url := "/api/v1/post"
	r.POST(url, CreatePostHandler)

	body := `{
		"title": "上梯子10",
    "content": "12345，上山打老虎",
    "community_id": 3
	}
	`

	req, _ := http.NewRequest(http.MethodPost, url, bytes.NewReader([]byte(body)))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// 判断相应内容是否包含指定字符串
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "需要登录")
}
