package zabbix


// see: https://www.zabbix.com/documentation/5.0/en/manual/api/reference/proxy/object
type ProxyObject struct {
	ProxyHostid string `json:"proxy_hostid,omitempty"`
	Host string `json:"host,omitempty"`
	Status int `json:"status,string,omitempty"`
	DisableUntil string `json:"disable_until,omitempty"`
	Error string `json:"error,omitempty"`
	Available string `json:"available,omitempty"`
	ErrorsFrom string `json:"errors_from,omitempty"`
	Lastaccess string `json:"lastaccess,omitempty"`
	IpmiAuthtype string `json:"ipmi_authtype,omitempty"`
	IpmiPrivilege string `json:"ipmi_privilege,omitempty"`
	IpmiUsername string `json:"ipmi_username,omitempty"`
	IpmiPassword string `json:"ipmi_password,omitempty"`
	IpmiDisableUntil string `json:"ipmi_disable_until,omitempty"`
	IpmiAvailable string `json:"ipmi_available,omitempty"`
	SnmpDisableUntil string `json:"snmp_disable_until,omitempty"`
	SnmpAvailable string `json:"snmp_available,omitempty"`
	Maintenanceid string `json:"maintenanceid,omitempty"`
	MaintenanceStatus string `json:"maintenance_status,omitempty"`
	MaintenanceType string `json:"maintenance_type,omitempty"`
	MaintenanceFrom string `json:"maintenance_from,omitempty"`
	IpmiErrorsFrom string `json:"ipmi_errors_from,omitempty"`
	SnmpErrorsFrom string `json:"snmp_errors_from,omitempty"`
	IpmiError string `json:"ipmi_error,omitempty"`
	SnmpError string `json:"snmp_error,omitempty"`
	JmxDisableUntil string `json:"jmx_disable_until,omitempty"`
	JmxAvailable string `json:"jmx_available,omitempty"`
	JmxErrorsFrom string `json:"jmx_errors_from,omitempty"`
	JmxError string `json:"jmx_error,omitempty"`
	Name string `json:"name,omitempty"`
	Flags string `json:"flags,omitempty"`
	Templateid string `json:"templateid,omitempty"`
	Description string `json:"description,omitempty"`
	TlsConnect int `json:"tls_connect,string,omitempty"`
	TlsAccept int `json:"tls_accept,string,omitempty"`
	TlsIssuer string `json:"tls_issuer,omitempty"`
	TlsSubject string `json:"tls_subject,omitempty"`
	TlsPskIdentity string `json:"tls_psk_identity,omitempty"`
	TlsPsk string `json:"tls_psk,omitempty"`
	ProxyAddress string `json:"proxy_address,omitempty"`
	AutoCompress int `json:"auto_compress,string,omitempty"`
	Discover string `json:"discover,omitempty"`
	Proxyid string `json:"proxyid,omitempty"`
}


//
// see: https://www.zabbix.com/documentation/5.0/en/manual/api/reference/proxy/get
type ProxyGetParams struct {
	GetParameters

	ProxyHostids       []int    `json:"proxyids,omitempty"`
	
	SelectHosts SelectQuery `json:"selectHosts,omitempty"`

	SelectInterface SelectQuery `json:"selectInterface,omitempty"`

}


func (z *ZabbixClient) ProxyGet(params ProxyGetParams) ([]ProxyObject, error) {

	var result []ProxyObject

	err := z.request("proxy.get", params, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (z *ZabbixClient) GetProxyFormHost(host string) ([]ProxyObject, error) {
	var pgp ProxyGetParams
	pgp.Filter = map[string]interface{}{"host": host}
	result, err := z.ProxyGet(pgp)
	if err != nil {
		return nil, err
	}
	return result, nil
}