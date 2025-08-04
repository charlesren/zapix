package zapix

// see: https://www.zabbix.com/documentation/5.0/en/manual/api/reference/drule/object
type DiscoveryObject struct {
	DruleidID   string `json:"druleid,omitempty"`
	IPRange     string `json:"iprange,omitempty"`
	Name        string `json:"name,omitempty"`
	Delay       string `json:"delay,omitempty"`
	NextCheck   string `json:"nextcheck,omitempty"`
	ProxyHostid string `json:"proxy_hostid,omitempty"`
	Status      int    `json:"status,omitempty"`
}

// see: https://www.zabbix.com/documentation/5.0/en/manual/api/reference/drule/get
type DiscoveryGetParams struct {
	GetParameters
	DhostIDs    []string `json:"dhostids,omitempty"`
	DruleIDs    []string `json:"druleids,omitempty"`
	DserviceIDs []string `json:"dserviceids,omitempty"`
}

/*
// Structure to store creation result
type DiscoveryCreateResult struct {
	ApplicationIDs []string `json:"applicationids"`
}

// Structure to store deletion result
type DiscoveryDeleteResult struct {
	ApplicationIDs []string `json:"applicationids"`
}

// HostgroupGet gets hostgroups
func (z *ZabbixClient) DiscoveryGet(params ApplicationGetParams) ([]ApplicationObject, error) {

	var result []ApplicationObject

	err := z.request("application.get", params, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// HostgroupCreate creates hostgroups
func (z *ZabbixClient) DiscoveryCreate(params []ApplicationObject) ([]string, error) {

	var result applicationCreateResult

	err := z.request("application.create", params, &result)
	if err != nil {
		return nil, err
	}

	return result.ApplicationIDs, nil
}
*/
