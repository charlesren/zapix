package zapix

// see: https://www.zabbix.com/documentation/5.0/en/manual/api/reference/itemprototype/object
type ItemPrototypeObject struct {
	ItemID                string                       `json:"itemid,omitempty"`
	Delay                 string                       `json:"delay,omitempty"`
	HostID                string                       `json:"hostid,omitempty"`
	RuleID                string                       `json:"ruleid,omitempty"`
	InterfaceID           string                       `json:"interfaceid,omitempty"`
	Key                   string                       `json:"key_,omitempty"`
	Name                  string                       `json:"name,omitempty"`
	Type                  int                          `json:"type"`
	Url                   string                       `json:"url,omitempty"`
	ValueType             int                          `json:"value_type"`
	Units                 string                       `json:"units,omitempty"`
	AllowTraps            int                          `json:"allow_traps,omitempty"`
	AuthType              int                          `json:"authtype,omitempty"`
	Description           string                       `json:"description,omitempty"`
	FollowRedirects       int                          `json:"follow_redirects,omitempty"`
	History               string                       `json:"history,omitempty"`
	Params                string                       `json:"params,omitempty"`
	Status                int                          `json:"status,omitempty"`
	Applications          []ApplicationObject          `json:"applications,omitempty"`          //used for `create` operations
	ApplicationPrototypes []ApplicationPrototypeObject `json:"applicationPrototypes,omitempty"` //used for `create` operations
}

type ApplicationPrototypeObject struct {
	Name string `json:"name,omitempty"`
}

// see: https://www.zabbix.com/documentation/5.0/en/manual/api/reference/itemprototype/get
type ItemPrototypeGetParams struct {
	GetParameters

	DiscoveryIDs []string `json:"discoveryids,omitempty"`
	GraphIDs     []string `json:"graphids,omitempty"`
	HostIDs      []string `json:"hostids,omitempty"`
	Inherited    bool     `json:"inherited,omitempty"`
	ItemIDs      []string `json:"itemids,omitempty"`
	Monitored    bool     `json:"monitored,omitempty"`
	Templated    bool     `json:"templated,omitempty"`
	TemplateIDs  []string `json:"templateids,omitempty"`
	TriggerIDs   string   `json:"triggerids,omitempty"`
}

// Structure to store creation result
type ItemPrototypeCreateResult struct {
	ItemIDs []string `json:"itemids"`
}

// Structure to store deletion result
type ItemPrototypeDeleteResult struct {
	ItemIDs []string `json:"itemids"`
}

// HostgroupGet gets hostgroups
func (z *ZabbixClient) ItemPrototypeGet(params ItemPrototypeGetParams) ([]ItemPrototypeObject, error) {

	var result []ItemPrototypeObject

	err := z.request("itemprototype.get", params, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// HostgroupCreate creates hostgroups
func (z *ZabbixClient) ItemPrototypeCreate(params []ItemPrototypeObject) ([]string, error) {

	var result ItemPrototypeCreateResult

	err := z.request("itemprototype.create", params, &result)
	if err != nil {
		return nil, err
	}

	return result.ItemIDs, nil
}
