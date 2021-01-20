package unbound

import (
	"fmt"
	"strings"
)

func buildHostName(host, domain string) (string, error) {
	if strings.Count(host, ".") != 0 {
		return "", fmt.Errorf("host can have only one part, got %+v", strings.Split(host, "."))
	}
	var name string
	if host != "" {
		name = strings.Join([]string{host, domain}, ".")
	} else {
		name = domain
	}
	return name, nil
}

func explodeHostName(hostName string) (string, string, error) {
	if strings.Count(hostName, ".") == 1 {
		return "", hostName, nil
	}
	parts := strings.SplitN(hostName, ".", 2)
	if len(parts) != 2 {
		return "", "", fmt.Errorf("host name should be in form of [<sub> <domain>], got %+v", parts)
	}
	return parts[0], parts[1], nil
}
