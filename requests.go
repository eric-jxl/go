package api 

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// ClientConfig 客户端配置 
type ClientConfig struct {
	BaseURL    string 
	Timeout    time.Duration 
	Retries    int               // 最大重试次数 
	RetryDelay time.Duration     // 重试延迟 
	Headers    map[string]string // 公共请求头 
}

// APIClient 封装第三方API调用的客户端 
type APIClient struct {
	client  *http.Client 
	config  ClientConfig 
}

// RequestOption 请求配置选项类型 
type RequestOption func(*http.Request)

// NewAPIClient 创建新的API客户端 
func NewAPIClient(config ClientConfig) *APIClient {
	if config.Timeout == 0 {
		config.Timeout = 30 * time.Second 
	}

	return &APIClient{
		client: &http.Client{
			Timeout: config.Timeout,
		},
		config: config,
	}
}

// APIRequest 请求参数结构体 
type APIRequest struct {
	Method   string 
	Endpoint string 
	Body     interface{}
	Headers  map[string]string 
	Context  context.Context 
}

// APIResponse 统一响应格式 
type APIResponse struct {
	StatusCode int 
	Headers    http.Header 
	Body       []byte 
	Data       interface{} // 反序列化后的数据 
}

// Call 统一调用方法 
func (c *APIClient) Call(req *APIRequest, result interface{}) (*APIResponse, error) {
	var finalErr error 

	// 请求重试逻辑 
	for attempt := 0; attempt <= c.config.Retries; attempt++ {
		response, err := c.doRequest(req, result)
		if err == nil {
			return response, nil 
		}

		finalErr = err 

		// 如果不是最后一次尝试，等待重试延迟 
		if attempt < c.config.Retries {
			time.Sleep(c.config.RetryDelay)
		}
	}

	return nil, fmt.Errorf("after %d attempts, last error: %w", c.config.Retries, finalErr)
}

// doRequest 实际执行请求的方法 
func (c *APIClient) doRequest(req *APIRequest, result interface{}) (*APIResponse, error) {
	// 请求上下文处理 
	ctx := req.Context 
	if ctx == nil {
		ctx = context.Background()
	}

	// 序列化请求体 
	var body io.Reader 
	if req.Body != nil {
		switch v := req.Body.(type) {
		case []byte:
			body = bytes.NewReader(v)
		case io.Reader:
			body = v 
		default:
			jsonBody, err := json.Marshal(req.Body)
			if err != nil {
				return nil, fmt.Errorf("marshal request body failed: %w", err)
			}
			body = bytes.NewReader(jsonBody)
		}
	}

	// 创建HTTP请求 
	fullURL := c.config.BaseURL + req.Endpoint 
	httpReq, err := http.NewRequestWithContext(ctx, req.Method, fullURL, body)
	if err != nil {
		return nil, fmt.Errorf("create request failed: %w", err)
	}

	// 设置请求头 
	for k, v := range c.config.Headers {
		httpReq.Header.Set(k, v)
	}
	for k, v := range req.Headers {
		httpReq.Header.Set(k, v)
	}

	// 发送请求 
	resp, err := c.client.Do(httpReq)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应体 
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response body failed: %w", err)
	}

	// 处理非2xx状态码 
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return &APIResponse{
			StatusCode: resp.StatusCode,
			Headers:   resp.Header,
			Body:      respBody,
		}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// 反序列化响应数据 
	if result != nil {
		if err := json.Unmarshal(respBody, result); err != nil {
			return &APIResponse{
				StatusCode: resp.StatusCode,
				Headers:    resp.Header,
				Body:       respBody,
			}, fmt.Errorf("unmarshal response failed: %w", err)
		}
	}

	return &APIResponse{
		StatusCode: resp.StatusCode,
		Headers:    resp.Header,
		Body:       respBody,
		Data:       result,
	}, nil 
}

// 示例用法 
func main() {
	config := ClientConfig{
		BaseURL:    "https://api.example.com/v1",
		Timeout:    10 * time.Second,
		Retries:    2,
		RetryDelay: 1 * time.Second,
		Headers: map[string]string{
			"Content-Type": "application/json",
			"Accept":       "application/json",
		},
	}

	client := NewAPIClient(config)

	// 创建请求 
	req := &APIRequest{
		Method:   "POST",
		Endpoint: "/users",
		Body: map[string]interface{}{
			"name":  "John Doe",
			"email": "john@example.com",
		},
		Headers: map[string]string{
			"Authorization": "Bearer token",
		},
	}

	var response struct {
		ID string `json:"id"`
	}

	resp, err := client.Call(req, &response)
	if err != nil {
		// 处理错误 
		fmt.Printf("Error: %v\n", err)
		return 
	}

	fmt.Printf("Response Status: %d\n", resp.StatusCode)
	fmt.Printf("New User ID: %s\n", response.ID)
}
