package zabbix

import (
	"fmt"

	"encoding/json"

	"github.com/imroc/req/v3"
)

// ZabbixClient is the go client for Zabbix API.
type ZabbixClient struct {
	*req.Client
	isLogged bool
	Auth     string
	Id       int32
}

type Params map[string]interface{}

// NewZabbixClient create a Zabbix client.
func NewZabbixClient() *ZabbixClient {
	c := req.C().
		// All Zabbix API requests need this header.
		SetCommonHeader("Accept", "*/*").
		// Enable dump at the request level for each request, which dump content into
		// memory (not print to stdout), we can record dump content only when unexpected
		// exception occurs, it is helpful to troubleshoot problems in production.
		EnableDumpEachRequest().
		// Unmarshal all Zabbix error response into struct.
		SetCommonErrorResult(&ZabbixAPIError).
		// Handle common exceptions in response middleware.
		EnableInsecureSkipVerify().
		OnAfterResponse(func(client *req.Client, resp *req.Response) error {
			if resp.Err != nil { // There is an underlying error, e.g. network error or unmarshal error.
				return nil
			}
			if apiErr, ok := resp.ErrorResult().(*APIError); ok {
				// Server returns an error message, convert it to human-readable go error.
				resp.Err = apiErr
				return nil
			}
			// Corner case: neither an error state response nor a success state response,
			// dump content to help troubleshoot.
			if !resp.IsSuccessState() {
				return fmt.Errorf("bad response, raw dump:\n%s", resp.Dump())
			}
			return nil
		})

	return &ZabbixClient{
		Client: c,
	}
}

// APIError represents the error message that Zabbix API returns.
type APIError string

var ZabbixAPIError APIError

// Error convert APIError to a human readable error and return.
func (e *APIError) Error() string {
	msg := fmt.Sprintf("Zabbix API error: %v", e)
	return msg
}

// LoginWithToken login with Zabbix personal access token.
func (c *ZabbixClient) LoginWithToken(token string) *ZabbixClient {
	//c.SetCommonHeader("Authorization", token)
	c.Auth = token
	c.isLogged = true
	return c
}

// IsLogged return true is user is logged in, otherwise false.
func (c *ZabbixClient) IsLogged() bool {
	return c.isLogged
}

// SetDebug enable debug if set to true, disable debug if set to false.
func (c *ZabbixClient) SetDebug(enable bool) *ZabbixClient {
	if enable {
		c.EnableDebugLog()
		c.EnableDumpAll()
	} else {
		c.DisableDebugLog()
		c.DisableDumpAll()
	}
	return c
}

type Request struct {
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Auth    string      `json:"auth,omitempty"`
	Id      int32       `json:"id"`
}

type Response struct {
	Jsonrpc string          `json:"jsonrpc"`
	Error   *ZabbixApiError `json:"error"`
	Result  interface{}     `json:"result"`
	Id      int32           `json:"id"`
}

// RawResponse format of zabbix api
type RawResponse struct {
	Jsonrpc string          `json:"jsonrpc"`
	Error   *ZabbixApiError `json:"error"`
	Result  json.RawMessage `json:"result"`
	ID      int32           `json:"id"`
}

type ZabbixApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

func (e *ZabbixApiError) Error() string {
	return fmt.Sprintf("%d (%s): %s", e.Code, e.Message, e.Data)
}

// request request zabbix api,save date to result
func (c *ZabbixClient) request(method string, params interface{}, result interface{}) error {
	resp := Response{
		Result: result,
	}

	req := &Request{
		Jsonrpc: "2.0",
		Method:  method,
		Params:  params,
		Auth:    c.Auth,
		Id:      c.Id,
	}

	err := c.
		Post("").
		SetHeader("Content-Type", "application/json-rpc").
		SetHeader("User-Agent", "github.com/AlekSi/zabbix").
		SetHeader("Cache-Control", "no-cache").
		SetBody(req).
		Do().
		Into(&resp)
	if err != nil {
		return err
	}
	if resp.Error != nil {
		return resp.Error
	}
	return nil

}

func (c *ZabbixClient) Login(url, username, password string) error {
	var token string
	c.Client.SetBaseURL(url)
	method := "user.login"
	params := map[string]string{"user": username, "password": password}
	err := c.request(method, params, &token)
	if err != nil {
		return err
	}
	c.LoginWithToken(token)
	return nil
}
