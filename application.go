package zapix

// see: https://www.zabbix.com/documentation/5.0/en/manual/api/reference/application/object
type ApplicationObject struct {
	ApplicationID string   `json:"applicationid,omitempty"`
	HostID        string   `json:"hostid,omitempty"`
	Name          string   `json:"name,omitempty"`
	Flags         int      `json:"flags,omitempty"`
	TemplatesIDs  []string `json:"templateids,omitempty"`
}

// see: https://www.zabbix.com/documentation/5.0/en/manual/api/reference/application/get
type ApplicationGetParams struct {
	GetParameters

	ApplicationIDs []string `json:"applicationids,omitempty"`
	GroupIDs       []int    `json:"groupids,omitempty"`
	HostIDs        []int    `json:"hostids,omitempty"`
	TemplateIDs    []int    `json:"templateids,omitempty"`
}

// Structure to store creation result
type applicationCreateResult struct {
	ApplicationIDs []string `json:"applicationids"`
}

// Structure to store deletion result
type applicationDeleteResult struct {
	ApplicationIDs []string `json:"applicationids"`
}

// HostgroupGet gets hostgroups
func (z *ZabbixClient) ApplicationGet(params ApplicationGetParams) ([]ApplicationObject, error) {

	var result []ApplicationObject

	err := z.request("application.get", params, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// HostgroupCreate creates hostgroups
func (z *ZabbixClient) ApplicationCreate(params []ApplicationObject) ([]string, error) {

	var result applicationCreateResult

	err := z.request("application.create", params, &result)
	if err != nil {
		return nil, err
	}

	return result.ApplicationIDs, nil
}
