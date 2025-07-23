package zapix

// see: https://www.zabbix.com/documentation/5.0/en/manual/api/reference/discoveryrule/object
type LLDObject struct {
	ItemID          string `json:"itemid,omitempty"`
	Delay           string `json:"delay,omitempty"`
	HostID          string `json:"hostid,omitempty"`
	InterfaceID     string `json:"interfaceid,omitempty"`
	Key             string `json:"key_,omitempty"`
	Name            string `json:"name,omitempty"`
	Type            int    `json:"type,omitempty"`
	Url             string `json:"url,omitempty"`
	AllowTraps      int    `json:"allow_traps,omitempty"`
	AuthType        int    `json:"authtype,omitempty"`
	Description     string `json:"description,omitempty"`
	Error           string `json:"error,omitempty"`
	FollowRedirects int    `json:"follow_redirects,omitempty"`
	LifeTime        string `json:"lifetime,omitempty"`
	Params          string `json:"params,omitempty"`
	Status          int    `json:"status,omitempty"`
}

// see: https://www.zabbix.com/documentation/5.0/en/manual/api/reference/discoveryrule/get
type LLDGetParams struct {
	GetParameters

	ItemIDs      []string `json:"itemids,omitempty"`
	GroupIDs     []string `json:"groupids,omitempty"`
	HostIDs      []string `json:"hostids,omitempty"`
	Delay        string   `json:"delay,omitempty"`
	HostID       string   `json:"hostid,omitempty"`
	InterfaceIDs string   `json:"interfaceids,omitempty"`
	Monitored    bool     `json:"monitored,omitempty"`
	Inherited    bool     `json:"inherited,omitempty"`
	Templated    bool     `json:"templated,omitempty"`
	TemplateIDs  []string `json:"templateids,omitempty"`
}

// Structure to store creation result
type LLDCreateResult struct {
	ItemIDs []string `json:"itemids"`
}

// Structure to store deletion result
type LLDDeleteResult struct {
	ItemIDs []string `json:"itemids"`
}

// HostgroupGet gets hostgroups
func (z *ZabbixClient) LLDGet(params LLDGetParams) ([]LLDObject, error) {

	var result []LLDObject

	err := z.request("discoveryrule.get", params, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// HostgroupCreate creates hostgroups
func (z *ZabbixClient) LLDCreate(params []LLDObject) ([]string, error) {

	var result LLDCreateResult

	err := z.request("discoveryrule.create", params, &result)
	if err != nil {
		return nil, err
	}

	return result.ItemIDs, nil
}
