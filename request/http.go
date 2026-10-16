package request

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"strings"
)

// HttpPostApi http POST请求标准
func HttpPostApi(headers map[string]string, url string, json string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(json))
	if err != nil {
		return nil, err
	}
	// 添加头部信息
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")

	// 添加自定义头部信息
	if headers != nil && len(headers) > 0 {
		for key, val := range headers {
			req.Header.Set(key, val)
		}
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// httpPost http post method
func httpPost(url string, body map[string]string, params map[string]string, headers map[string]string) (*http.Response, error) {
	// add post body
	var bodyJson []byte
	var req *http.Request
	if body != nil {
		var err error
		bodyJson, err = json.Marshal(body)
		if err != nil {
			log.Println(err)
			return nil, errors.New("http post body to json failed")
		}
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyJson))
	if err != nil {
		log.Println(err)
		return nil, errors.New("new request is fail: %v \n")
	}
	req.Header.Set("Content-type", "application/json")
	// add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	// add headers
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	// http client
	client := &http.Client{}
	log.Printf("Go POST URL : %s \n", req.URL.String())
	return client.Do(req)
}

// HttpsPostApi HTTPS POST请求标准
func HttpsPostApi(headers map[string]string, url string, json string) ([]byte, error) {
	// HTTPS 默认忽略请求安全性提醒
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(json))
	if err != nil {
		return nil, err
	}
	// 添加头部信息
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded;charset=utf-8")
	// 添加自定义头部信息
	if headers != nil && len(headers) > 0 {
		for key, val := range headers {
			req.Header.Set(key, val)
		}
	}
	resp, rerr := client.Do(req)
	if rerr != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	body, berr := io.ReadAll(resp.Body)
	if berr != nil {
		return nil, err
	}
	return body, nil // 返回读取的结果值
}

// HttpGetApi http get请求的方法
func HttpGetApi(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	data, rerr := io.ReadAll(resp.Body)
	if rerr != nil {
		return nil, rerr
	}
	return data, nil
}

// httpGet http get method
func httpGet(url string, params map[string]string, headers map[string]string) (*http.Response, error) {
	// new request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Println(err)
		return nil, errors.New("new request is fail ")
	}
	// add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	// add headers
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	// http client
	client := &http.Client{}
	log.Printf("Go GET URL : %s \n", req.URL.String())
	return client.Do(req)
}

// HttpsGetApi http get请求的方法
func HttpsGetApi(url string) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, rerr := client.Do(req)
	if rerr != nil {
		return nil, rerr
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	data, berr := io.ReadAll(resp.Body)
	if berr != nil {
		return nil, berr
	}
	return data, nil
}
