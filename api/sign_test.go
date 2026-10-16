package api

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestGenerateSignature(t *testing.T) {
	params := map[string]string{
		"X-Sign-Version": "v1",
		"X-Device-Id":    "550e8400-e29b-41d4-a716-446655440000",
		"X-Request-Id":   "6221c4b7-0ccf-465f-9651-0a7b6c4588d7",
		"X-Timestamp":    "1694853023",
	}
	secret := "2025509161199975670809202jj0nc"

	signature := GenerateSignature(params, secret)
	t.Logf("X-Signature: %s", signature)
}

type ApiSignTestSuit struct {
	suite.Suite
	*assert.Assertions
	secret string
}

func (s *ApiSignTestSuit) SetupSuite() {
	t := s.T()
	s.Assertions = assert.New(t)
	s.secret = "2025509191199999970809202jj0nd"
}

func (s *ApiSignTestSuit) TestApiRequest() {
	// 使用go-uuid包生成唯一的设备ID和请求ID
	deviceID := uuid.NewString()
	requestID := uuid.NewString()
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	// 创建请求头参数
	params := map[string]string{
		"X-Sign-Version": "v1",
		"X-Device-Id":    deviceID,
		"X-Request-Id":   requestID,
		"X-Timestamp":    timestamp,
	}

	// 生成签名
	signature := GenerateSignature(params, s.secret)
	s.T().Logf("生成的签名: %s", signature)

	// 创建请求体
	requestBody := map[string]interface{}{
		"logs": []map[string]interface{}{{
			"duration":     88,
			"end_time":     1782458511573,
			"event_source": "click",
			"event_time":   1767653420940,
			"event_type":   "aute Lorem in Excepteur",
			"start_time":   1755141977189,
		}},
	}

	// 转换为JSON
	jsonBody, err := json.Marshal(requestBody)
	s.NoError(err, "转换请求体失败")

	// 创建HTTP请求（POST方法）
	apiURL := "http://localhost:16100/api/v1/event/metrics"
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonBody))
	s.NoError(err, "创建请求失败")

	// 添加请求头
	for k, v := range params {
		req.Header.Set(k, v)
	}
	req.Header.Set("X-Signature", signature)
	req.Header.Set("Content-Type", "application/json")

	// 发送请求
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	s.NoError(err, "发送请求失败")
	defer func() {
		_ = resp.Body.Close()
	}()

	// 读取响应
	body, err := io.ReadAll(resp.Body)
	s.NoError(err, "读取响应失败")
	s.T().Logf("响应状态码: %d", resp.StatusCode)
	s.T().Logf("响应内容: %s", string(body))

	// 验证响应状态码
	s.Equal(http.StatusOK, resp.StatusCode, "期望状态码为200")
}

func TestApiSign(t *testing.T) {
	suite.Run(t, new(ApiSignTestSuit))
}
