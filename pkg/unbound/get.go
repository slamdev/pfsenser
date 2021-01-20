package unbound

import (
	"fmt"
	"github.com/slamdev/pfsenser/pkg/xmlrpc"
)

func Get(client xmlrpc.PfsenseClient, name string) (Host, bool, error) {
	hosts, err := List(client)
	if err != nil {
		return Host{}, false, fmt.Errorf("failed to list hosts to get one; %w", err)
	}
	for _, host := range hosts {
		if host.Name == name {
			return host, true, nil
		}
	}
	return Host{}, false, nil
}
