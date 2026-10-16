package api

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"sort"
	"strings"
)

// GenerateSignature 根据参数和 secret 计算 HMAC-SHA256，再 Base64 编码后转大写返回签名
func GenerateSignature(params map[string]string, secret string) string {
	// 把参数按键排序
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 拼接成 plaintext 字符串
	parts := make([]string, 0, len(params))
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%s=%s", strings.ToLower(k), params[k]))
	}
	plaintext := strings.Join(parts, "&")

	// 计算 HMAC-SHA256
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(plaintext))
	signature := mac.Sum(nil)

	// Base64 编码转大写
	return strings.ToUpper(base64.StdEncoding.EncodeToString(signature))
}
