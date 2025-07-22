package zabbix

// For `HostObject` field: `Available`
const (
	HostAvailableUnknown     = 0
	HostAvailableAvailable   = 1
	HostAvailableUnavailable = 2
)

// For `HostObject` field: `Flags`
const (
	HostFlagsPlain      = 0
	HostFlagsDiscovered = 4
)

// For `HostObject` field: `Flags`
const (
	HostInventoryModeDisabled  = -1
	HostInventoryModeManual    = 0
	HostInventoryModeAutomatic = 1
)

// For `HostObject` field: `IpmiAuthtype`
const (
	HostIpmiAuthtypeDefault  = -1
	HostIpmiAuthtypeNone     = 0
	HostIpmiAuthtypeMD2      = 1
	HostIpmiAuthtypeMD5      = 2
	HostIpmiAuthtypeStraight = 4
	HostIpmiAuthtypeOEM      = 5
	HostIpmiAuthtypeRMCP     = 6
)

// For `HostObject` field: `IpmiAvailable`
const (
	HostIpmiAvailableUnknown     = 0
	HostIpmiAvailableAvailable   = 1
	HostIpmiAvailableUnavailable = 2
)

// For `HostObject` field: `IpmiPrivilege`
const (
	HostIpmiPrivilegeCallback = 1
	HostIpmiPrivilegeUser     = 2
	HostIpmiPrivilegeOperator = 3
	HostIpmiPrivilegeAdmin    = 4
	HostIpmiPrivilegeOEM      = 5
)

// For `HostObject` field: `JmxAvailable`
const (
	HostJmxAvailableUnknown     = 0
	HostJmxAvailableAvailable   = 1
	HostJmxAvailableUnavailable = 2
)

// For `HostObject` field: `MaintenanceStatus`
const (
	HostMaintenanceStatusDisable = 0
	HostMaintenanceStatusEnable  = 1
)

// For `HostObject` field: `MaintenanceType`
const (
	HostMaintenanceTypeDataCollectionEnabled  = 0
	HostMaintenanceTypeDataCollectionDisabled = 1
)

// For `HostObject` field: `SnmpAvailable`
const (
	HostSnmpAvailableUnknown     = 0
	HostSnmpAvailableAvailable   = 1
	HostSnmpAvailableUnavailable = 2
)

// For `HostObject` field: `Status`
const (
	HostStatusMonitored   = 0
	HostStatusUnmonitored = 1
)

// For `HostObject` field: `TLSConnect`
const (
	TLSConnectNoEncryption = 1
	TLSConnectPSK          = 2
	TLSConnectCertificate  = 4
)

// For `HostObject` field: `TLSAccept`
const (
	TLSAcceptNoEncryption = 1
	TLSAcceptPSK          = 2
	TLSAcceptCertificate  = 4
)

// For `HostGetParams` field: `Evaltype`
const (
	HostEvaltypeAndOr = 0
	HostEvaltypeOr    = 2
)

// For `HostTag` field: `Operator`
const (
	HostTagOperatorContains = 0
	HostTagOperatorEquals   = 1
)

// HostObject struct is used to store host operations results
//
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/host/object#host
type HostObject struct {
	HostID            string `json:"hostid,omitempty"`
	Host              string `json:"host,omitempty"`
	Available         int    `json:"available,string,omitempty"` // has defined consts, see above
	Description       string `json:"description,omitempty"`
	DisableUntil      int    `json:"disable_until,string,omitempty"`
	Error             string `json:"error,omitempty"`
	ErrorsFrom        int    `json:"errors_from,string,omitempty"`
	Flags             int    `json:"flags,string,omitempty"`          // has defined consts, see above
	InventoryMode     int    `json:"inventory_mode,string,omitempty"` // has defined consts, see above
	IpmiAuthtype      int    `json:"ipmi_authtype,string,omitempty"`  // has defined consts, see above
	IpmiAvailable     int    `json:"ipmi_available,string,omitempty"` // has defined consts, see above
	IpmiDisableUntil  int    `json:"ipmi_disable_until,string,omitempty"`
	IpmiError         string `json:"ipmi_error,omitempty"`
	IpmiErrorsFrom    int    `json:"ipmi_errors_from,string,omitempty"`
	IpmiPassword      string `json:"ipmi_password,omitempty"`
	IpmiPrivilege     int    `json:"ipmi_privilege,string,omitempty"` // has defined consts, see above
	IpmiUsername      string `json:"ipmi_username,omitempty"`
	JmxAvailable      int    `json:"jmx_available,string,omitempty"` // has defined consts, see above
	JmxDisableUntil   int    `json:"jmx_disable_until,string,omitempty"`
	JmxError          string `json:"jmx_error,omitempty"`
	JmxErrorsFrom     int    `json:"jmx_errors_from,string,omitempty"`
	MaintenanceFrom   int    `json:"maintenance_from,string,omitempty"`
	MaintenanceStatus int    `json:"maintenance_status,string,omitempty"` // has defined consts, see above
	MaintenanceType   int    `json:"maintenance_type,string,omitempty"`   // has defined consts, see above
	MaintenanceID     int    `json:"maintenanceid,string,omitempty"`
	Name              string `json:"name,omitempty"`
	ProxyHostID       int    `json:"proxy_hostid,string,omitempty"`
	SnmpAvailable     int    `json:"snmp_available,string,omitempty"` // has defined consts, see above
	SnmpDisableUntil  int    `json:"snmp_disable_until,string,omitempty"`
	SnmpError         string `json:"snmp_error,omitempty"`
	SnmpErrorsFrom    int    `json:"snmp_errors_from,string,omitempty"`
	Status            int    `json:"status,string,omitempty"`      // has defined consts, see above
	TLSConnect        int    `json:"tls_connect,string,omitempty"` // has defined consts, see above
	TLSAccept         int    `json:"tls_accept,string,omitempty"`  // has defined consts, see above
	TLSIssuer         string `json:"tls_issuer,omitempty"`
	TLSSubject        string `json:"tls_subject,omitempty"`
	TLSPSKIdentity    string `json:"tls_psk_identity,omitempty"`
	TLSPSK            string `json:"tls_psk,omitempty"`

	Groups          []HostgroupObject     `json:"groups,omitempty"`
	Interfaces      []HostInterfaceObject `json:"interfaces,omitempty"`
	Tags            []HostTagObject       `json:"tags,omitempty"`
	InheritedTags   []HostTagObject       `json:"inheritedTags,omitempty"`
	Macros          []UsermacroObject     `json:"macros,omitempty"`
	Templates       []TemplateObject      `json:"templates,omitempty"`       // Used for `create` operations
	ParentTemplates []TemplateObject      `json:"parentTemplates,omitempty"` // Used to store result for `get` operations
}

// HostTagObject struct is used to store host tag
//
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/host/object#host_tag
type HostTagObject struct {
	Tag   string `json:"tag"`
	Value string `json:"value,omitempty"`

	Operator int `json:"operator,omitempty"` // Used for `get` operations, has defined consts, see above
}

// HostGetParams struct is used for host get requests
//
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/host/get#parameters
type HostGetParams struct {
	GetParameters

	GroupIDs       []int    `json:"groupids,omitempty"`
	ApplicationIDs []int    `json:"applicationids,omitempty"`
	DserviceIDs    []int    `json:"dserviceids,omitempty"`
	GraphIDs       []int    `json:"graphids,omitempty"`
	HostIDs        []string `json:"hostids,omitempty"`
	HttptestIDs    []int    `json:"httptestids,omitempty"`
	InterfaceIDs   []int    `json:"interfaceids,omitempty"`
	ItemIDs        []int    `json:"itemids,omitempty"`
	MaintenanceIDs []int    `json:"maintenanceids,omitempty"`
	MonitoredHosts bool     `json:"monitored_hosts,omitempty"`
	ProxyHosts     bool     `json:"proxy_hosts,omitempty"`
	ProxyIDs       []int    `json:"proxyids,omitempty"`
	TemplatedHosts bool     `json:"templated_hosts,omitempty"`
	TemplateIDs    []int    `json:"templateids,omitempty"`
	TriggerIDs     []int    `json:"triggerids,omitempty"`

	WithItems                     bool            `json:"with_items,omitempty"`
	WithItemPrototypes            bool            `json:"with_item_prototypes,omitempty"`
	WithSimpleGraphItemPrototypes bool            `json:"with_simple_graph_item_prototypes,omitempty"`
	WithApplications              bool            `json:"with_applications,omitempty"`
	WithGraphs                    bool            `json:"with_graphs,omitempty"`
	WithGraphPrototypes           bool            `json:"with_graph_prototypes,omitempty"`
	WithHttptests                 bool            `json:"with_httptests,omitempty"`
	WithMonitoredHttptests        bool            `json:"with_monitored_httptests,omitempty"`
	WithMonitoredItems            bool            `json:"with_monitored_items,omitempty"`
	WithMonitoredTriggers         bool            `json:"with_monitored_triggers,omitempty"`
	WithSimpleGraphItems          bool            `json:"with_simple_graph_items,omitempty"`
	WithTriggers                  bool            `json:"with_triggers,omitempty"`
	WithProblemsSuppressed        bool            `json:"withProblemsSuppressed,omitempty"`
	Evaltype                      int             `json:"evaltype,omitempty"` // has defined consts, see above
	Severities                    []int           `json:"severities,omitempty"`
	Tags                          []HostTagObject `json:"tags,omitempty"`
	InheritedTags                 bool            `json:"inheritedTags,omitempty"`

	// SelectApplications    SelectQuery `json:"selectApplications,omitempty"` // not implemented yet
	// SelectDiscoveries     SelectQuery `json:"selectDiscoveries,omitempty"` // not implemented yet
	// SelectDiscoveryRule   SelectQuery `json:"selectDiscoveryRule ,omitempty"` // not implemented yet
	// SelectGraphs          SelectQuery `json:"selectGraphs,omitempty"` // not implemented yet
	SelectGroups SelectQuery `json:"selectGroups,omitempty"`
	// SelectHostDiscovery   SelectQuery `json:"selectHostDiscovery ,omitempty"` // not implemented yet
	// SelectHTTPTests       SelectQuery `json:"selectHttpTests,omitempty"` // not implemented yet
	SelectInterfaces SelectQuery `json:"selectInterfaces,omitempty"`
	// SelectInventory       SelectQuery `json:"selectInventory,omitempty"` // not implemented yet
	// SelectItems           SelectQuery `json:"selectItems,omitempty"` // not implemented yet
	SelectMacros          SelectQuery `json:"selectMacros,omitempty"`
	SelectParentTemplates SelectQuery `json:"selectParentTemplates,omitempty"`
	// SelectScreens         SelectQuery `json:"selectScreens,omitempty"` // not implemented yet
	SelectTags          SelectQuery `json:"selectTags,omitempty"`
	SelectInheritedTags SelectQuery `json:"selectInheritedTags,omitempty"`
	// SelectTriggers        SelectQuery `json:"selectTriggers,omitempty"` // not implemented yet
}

// Structure to store creation result
type hostCreateResult struct {
	HostIDs []string `json:"hostids"`
}

// Structure to store updation result
type hostUpdateResult struct {
	HostIDs []string `json:"hostids"`
}

// Structure to store deletion result
type hostDeleteResult struct {
	HostIDs []string `json:"hostids"`
}

// HostGet gets hosts
func (z *ZabbixClient) HostGet(params HostGetParams) ([]HostObject, error) {

	var result []HostObject

	err := z.request("host.get", params, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// HostCreate creates hosts
func (z *ZabbixClient) HostCreate(params []HostObject) ([]string, error) {

	var result hostCreateResult

	err := z.request("host.create", params, &result)
	if err != nil {
		return nil, err
	}

	return result.HostIDs, nil
}

// HostUpdate updates hosts
func (z *ZabbixClient) HostUpdate(params []HostObject) ([]string, error) {

	var result hostUpdateResult

	err := z.request("host.update", params, &result)
	if err != nil {
		return nil, err
	}

	return result.HostIDs, nil
}

// HostDelete deletes hosts
func (z *ZabbixClient) HostDelete(hostIDs []int) ([]string, error) {

	var result hostDeleteResult

	err := z.request("host.delete", hostIDs, &result)
	if err != nil {
		return nil, err
	}

	return result.HostIDs, nil
}

// GetHostFromHostid gets hosts
func (z *ZabbixClient) GetHostFromHostid(hostid string) ([]HostObject, error) {
	var hgp HostGetParams
	hgp.SelectParentTemplates = SelectFields{"templateid", "name"}
	hgp.HostIDs = append(hgp.HostIDs, hostid)
	result, err := z.HostGet(hgp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// AddTempsForHosts add specfic templates for hosts
func (z *ZabbixClient) AddTempsForHosts(hosts []HostObject, templateids []string) ([]string, error) {

	for _, id := range templateids {
		newtemp := TemplateObject{TemplateID: id}
		for i, host := range hosts {
			exist := false
			for _, pt := range host.ParentTemplates {
				if pt.TemplateID == id {
					exist = true
				}
			}
			if !exist {
				host.ParentTemplates = append(host.ParentTemplates, newtemp)

				hosts[i].Templates = host.ParentTemplates
			}
		}
	}

	result, err := z.HostUpdate(hosts)
	if err != nil {
		return nil, err
	}

	return result, nil

}
