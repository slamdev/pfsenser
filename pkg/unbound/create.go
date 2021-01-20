package unbound

import (
	"fmt"
	"github.com/slamdev/pfsenser/pkg/xmlrpc"
)

func Create(client xmlrpc.PfsenseClient, host Host) error {
	_, found, err := Get(client, host.Name)
	if err != nil {
		return fmt.Errorf("failed to get host to create; %w", err)
	}
	if found {
		return fmt.Errorf("host %s already exists", host.Name)
	}

	cfg, err := client.BackupConfigSection(section)
	if err != nil {
		return fmt.Errorf("faile to get unbound config section; %w", err)
	}

	h, d, err := explodeHostName(host.Name)
	if err != nil {
		return fmt.Errorf("failed to explode hostname from %s", host.Name)
	}

	unboundHost := xmlrpc.UnboundHost{
		Host:    h,
		Domain:  d,
		Ip:      host.IP,
		Descr:   host.Description,
		Aliases: host.Aliases,
	}

	cfg.Struct.Unbound.Hosts = append(cfg.Struct.Unbound.Hosts, unboundHost)

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
