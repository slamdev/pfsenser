package xmlrpc

type Config struct {
	Struct UnboundStruct
}

type UnboundStruct struct {
	Unbound Unbound `xml:"unbound"`
}

type Unbound struct {
	Enable                    string        `xml:"enable"`
	Dnssec                    string        `xml:"dnssec"`
	ActiveInterface           string        `xml:"active_interface"`
	OutgoingInterface         string        `xml:"outgoing_interface"`
	CustomOptions             string        `xml:"custom_options"`
	Hideidentity              string        `xml:"hideidentity"`
	Hideversion               string        `xml:"hideversion"`
	Dnssecstripped            string        `xml:"dnssecstripped"`
	Hosts                     []UnboundHost `xml:"hosts"`
	Acls                      []UnboundAcl  `xml:"acls"`
	Port                      string        `xml:"port"`
	Sslport                   string        `xml:"sslport"`
	Sslcertref                string        `xml:"sslcertref"`
	SystemDomainLocalZoneType string        `xml:"system_domain_local_zone_type"`
}

type UnboundHost struct {
	Host    string `xml:"host"`
	Domain  string `xml:"domain"`
	Ip      string `xml:"ip"`
	Descr   string `xml:"descr"`
	Aliases string `xml:"aliases"`
}

type UnboundAcl struct {
	Aclid       string          `xml:"aclid"`
	Aclname     string          `xml:"aclname"`
	Aclaction   string          `xml:"aclaction"`
	Description string          `xml:"description"`
	Row         []UnboundAclRow `xml:"row"`
}

type UnboundAclRow struct {
	AclNetwork  string `xml:"acl_network"`
	Mask        string `xml:"mask"`
	Description string `xml:"description"`
}
