package unbound

import (
	"fmt"
	"github.com/slamdev/pfsenser/pkg/xmlrpc"
)

func Delete(client xmlrpc.PfsenseClient, name string) error {
	cfg, err := client.BackupConfigSection(section)
	if err != nil {
		return fmt.Errorf("failed to backup unbound config; %w", err)
	}

	h, d, err := explodeHostName(name)
	if err != nil {
		return fmt.Errorf("failed to explode hostname from %s", name)
	}

	var modifiedHosts []xmlrpc.UnboundHost
	deleted := 0
	for _, host := range cfg.Struct.Unbound.Hosts {
		if host.Host == h && host.Domain == d {
			deleted++
			continue
		}
		modifiedHosts = append(modifiedHosts, host)
	}

	if deleted == 0 {
		return fmt.Errorf("failed to find %s host and %s domain to delete", h, d)
	}

	cfg.Struct.Unbound.Hosts = modifiedHosts

	if err := client.RestoreConfigSection(cfg); err != nil {
		return fmt.Errorf("failed to restore unbound config section; %w", err)
	}
	if err := client.ConfigureUnbound(); err != nil {
		return fmt.Errorf("failed to configure unbound; %w", err)
	}
	if err := client.ConfigureDhcpd(); err != nil {
		return fmt.Errorf("failed to configure dhcpd; %w", err)
	}
	return nil
}
