package unbound

import (
	"fmt"
	"github.com/slamdev/pfsenser/pkg/xmlrpc"
)

func List(client xmlrpc.PfsenseClient) ([]Host, error) {
	config, err := client.BackupConfigSection(section)
	if err != nil {
		return nil, fmt.Errorf("failed to backup unbound config; %w", err)
	}
	hosts := make([]Host, len(config.Struct.Unbound.Hosts))
	for i, h := range config.Struct.Unbound.Hosts {
		name, err := buildHostName(h.Host, h.Domain)
		if err != nil {
			return nil, fmt.Errorf("failed to build hostname from %s; %w", h, err)
		}
		hosts[i] = Host{
			Name:        name,
			IP:          h.Ip,
			Description: h.Descr,
			Aliases:     h.Aliases,
		}
	}
	return hosts, nil
}
