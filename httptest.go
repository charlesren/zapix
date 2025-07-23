package zapix

// see: https://www.zabbix.com/documentation/5.0/en/manual/api/reference/httptest/object
type HttptestObject struct {
	HttptestID    string       `json:"httptestid,omitempty"`
	HostID        string       `json:"hostid,omitempty"`
	Name          string       `json:"name,omitempty"`
	Delay         string       `json:"delay,omitempty"`
	ApplicationID string       `json:"applicationid,omitempty"`
	Retries       int          `json:"retries,omitempty"`
	Steps         []StepObject `json:"steps,omitempty"`
}
type StepObject struct {
	HttpstepID  string `json:"httpstepid,omitempty"`
	No          int    `json:"no,omitempty"`
	Name        string `json:"name,omitempty"`
	Url         string `json:"url,omitempty"`
	Timeout     string `json:"timeout,omitempty"`
	StatusCodes string `json:"status_codes,omitempty"`
	Required    string `json:"required,omitempty"`
}

// Structure to store creation result
type HttptestCreateResult struct {
	HttptestIDs []string `json:"httptestids"`
}

// HostgroupCreate creates hostgroups
func (z *ZabbixClient) HttptestCreate(params []HttptestObject) ([]string, error) {

	var result HttptestCreateResult

	err := z.request("httptest.create", params, &result)
	if err != nil {
		return nil, err
	}

	return result.HttptestIDs, nil
}
