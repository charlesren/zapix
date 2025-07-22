package zabbix

import (
	"errors"
)

// For `HostinterfaceObject` field: `Main`
const (
	HostinterfaceMainNotDefault = 0
	HostinterfaceMainDefault    = 1
)

// For `HostinterfaceObject` field: `Type`
const (
	HostinterfaceTypeAgent = 1
	HostinterfaceTypeSNMP  = 2
	HostinterfaceTypeIPMI  = 3
	HostinterfaceTypeJMX   = 4
)

// For `HostinterfaceObject` field: `UseIP`
const (
	HostinterfaceUseipDNS = 0
	HostinterfaceUseipIP  = 1
)

// For `HostinterfaceDetailsTagObject` field: `Bulk`
const (
	HostinterfaceDetailsTagBulkDontUse = 0
	HostinterfaceDetailsTagBulkUse     = 1
)

// For `HostinterfaceDetailsTagObject` field: `Version`
const (
	HostinterfaceDetailsTagVersionSNMPv1  = 1
	HostinterfaceDetailsTagVersionSNMPv2c = 2
	HostinterfaceDetailsTagVersionSNMPv3  = 3
)

// For `HostinterfaceDetailsTagObject` field: `SecurityLevel`
const (
	HostinterfaceDetailsTagSecurityLevelNoAuthNoPriv = 0
	HostinterfaceDetailsTagSecurityLevelAuthNoPriv   = 1
	HostinterfaceDetailsTagSecurityLevelAuthPriv     = 2
)

// For `HostinterfaceDetailsTagObject` field: `AuthProtocol`
const (
	HostinterfaceDetailsTagAuthProtocolMD5 = 0
	HostinterfaceDetailsTagAuthProtocolSHA = 1
)

// For `HostinterfaceDetailsTagObject` field: `PrivProtocol`
const (
	HostinterfaceDetailsTagPrivProtocolDES = 0
	HostinterfaceDetailsTagPrivProtocolAES = 1
)

// HostinterfaceObject struct is used to store hostinterface operations results
//
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/hostinterface/object#hostinterface
type HostInterfaceObject struct {
	InterfaceID string                          `json:"interfaceid,omitempty"`
	DNS         string                          `json:"dns"`
	HostID      string                          `json:"hostid"`
	IP          string                          `json:"ip"`
	Main        int                             `json:"main,string"` // has defined consts, see above
	Port        string                          `json:"port"`
	Type        int                             `json:"type,string"`  // has defined consts, see above
	UseIP       int                             `json:"useip,string"` // has defined consts, see above
	Details     []HostInterfaceDetailsTagObject `json:"details,omitempty"`

	// Items []ItemObject `json:"items,omitempty"` // not implemented yet
	Hosts []HostObject `json:"hosts,omitempty"`
}

// HostinterfaceDetailsTagObject struct is used to store hostinterface details
//
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/hostinterface/object#details_tag
type HostInterfaceDetailsTagObject struct {
	Version        int    `json:"version,omitempty"` // has defined consts, see above
	Bulk           int    `json:"bulk,omitempty"`    // has defined consts, see above
	Community      string `json:"community,omitempty"`
	SecurityName   string `json:"securityname,omitempty"`
	SecurityLevel  int    `json:"securitylevel,omitempty"` // has defined consts, see above
	AuthPassPhrase string `json:"authpassphrase,omitempty"`
	PrivPassPhrase string `json:"privpassphrase,omitempty"`
	AuthProtocol   int    `json:"authprotocol,omitempty"` // has defined consts, see above
	PrivProtocol   int    `json:"privprotocol,omitempty"` // has defined consts, see above
	ContextName    string `json:"contextname,omitempty"`
}

// HostinterfaceGetParams struct is used for hostinterface get requests
//
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/hostinterface/get#parameters
type HostInterfaceGetParams struct {
	GetParameters

	HostIDs      []int `json:"hostids,omitempty"`
	InterfaceIDs []int `json:"interfaceids,omitempty"`
	ItemIDs      []int `json:"itemids,omitempty"`
	TriggerIDs   []int `json:"triggerids,omitempty"`

	// SelectItems SelectQuery `json:"selectItems,omitempty"` // not implemented yet
	SelectHosts SelectQuery `json:"selectHosts,omitempty"`
}

// Structure to store creation result
type hostInterfaceCreateResult struct {
	InterfaceIDs []int `json:"interfaceids"`
}

// Structure to store deletion result
type hostinterfaceDeleteResult struct {
	InterfaceIDs []int `json:"interfaceids"`
}

// HostinterfaceGet gets hostinterfaces
func (z *ZabbixClient) HostInterfaceGet(params HostInterfaceGetParams) ([]HostInterfaceObject, error) {

	var result []HostInterfaceObject

	err := z.request("hostinterface.get", params, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// HostinterfaceCreate creates hostinterfaces
func (z *ZabbixClient) HostInterfaceCreate(params []HostInterfaceObject) ([]int, error) {

	var result hostInterfaceCreateResult

	err := z.request("hostinterface.create", params, &result)
	if err != nil {
		return nil, err
	}

	return result.InterfaceIDs, nil
}

// HostinterfaceDelete deletes hostinterfaces
func (z *ZabbixClient) HostInterfaceDelete(hostinterfaceIDs []int) ([]int, error) {

	var result hostinterfaceDeleteResult

	err := z.request("hostinterface.delete", hostinterfaceIDs, &result)
	if err != nil {
		return nil, err
	}

	return result.InterfaceIDs, nil
}

// GetHostInterfaceFromIP get  hostinterfaces
func (z *ZabbixClient) GetHostInterfaceFromIP(ip string) ([]HostInterfaceObject, error) {
	if ip == "" {
		return nil, errors.New("ip address is nil")
	}
	var hgp HostInterfaceGetParams
	hgp.Filter = make(map[string]interface{})
	hgp.Filter["type"] = 1
	hgp.Filter["ip"] = ip
	his, err := z.HostInterfaceGet(hgp)
	if err != nil {
		return nil, err
	}
	return his, nil
}
