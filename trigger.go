package zabbix

import (
	"bytes"
	"encoding/json"
)

// Trigger represents a Zabbix Trigger returned from the Zabbix API.
//
// See: https://www.zabbix.com/documentation/5.0/en/manual/api/reference/trigger/object
type TriggerObject struct {
	TriggerID string `json:"triggerid,omitempty"`

	Description string `json:"description,omitempty"`

	Expression string `json:"expression,omitempty"`

	Opdata string `json:"opdata,omitempty"`

	Priority int `json:"priority,omitempty"`

	Comments string `json:"comments,omitempty"`

	ManualClose int `json:"manual_close,omitempty"`

	State int `json:"state,omitempty"`

	Status int `json:"status,omitempty"`

	Tags []TriggerTagObject `json:"tags,omitempty"`
}

// TriggerTag is trigger tag
type TriggerTagObject struct {
	Tag   string `json:"tag,omitempty"`
	Value string `json:"value,omitempty"`
}

// Structure to store creation result
type TriggerCreateResult struct {
	TriggerIDs []string `json:"triggerids"`
}

// HostgroupCreate creates hostgroups
func (z *ZabbixClient) TriggerCreate(params []TriggerObject) ([]string, error) {

	var result TriggerCreateResult

	resp := Response{
		Result: result,
	}

	req := &Request{
		Jsonrpc: "2.0",
		Method:  "trigger.create",
		Params:  params,
		Auth:    z.Auth,
		Id:      z.Id,
	}
	// marshal req to json manually , do not using HTMLEscape, so "<", ">", "&" won't be replaced.
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(req)
	if err != nil {
		return nil, err
	}
	err = z.
		Post("").
		SetHeader("Content-Type", "application/json-rpc").
		SetHeader("User-Agent", "github.com/AlekSi/zabbix").
		SetHeader("Cache-Control", "no-cache").
		SetBodyJsonString(string(buffer.Bytes())).
		Do().
		Into(&resp)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, resp.Error
	}

	return result.TriggerIDs, nil
}
