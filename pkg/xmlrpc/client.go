package xmlrpc

import (
	"alexejk.io/go-xmlrpc"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type PfsenseClient interface {
	BackupConfigSection(section string) (Config, error)
	RestoreConfigSection(cfg Config) error
	ConfigureUnbound() error
	ConfigureDhcpd() error
	Close() error
}

type OperationResult struct {
	Success bool
}

type client struct {
	*xmlrpc.Client
}

func NewPfsenseClient(url string, username string, password string, insecure bool) (PfsenseClient, error) {
	httpClient := &http.Client{Timeout: time.Second * 10}
	if insecure {
		httpClient.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
	headers := map[string]string{
		"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte(username+":"+password)),
	}
	c, err := xmlrpc.NewClient(url, xmlrpc.HttpClient(httpClient), xmlrpc.Headers(headers))
	if err != nil {
		return nil, fmt.Errorf("failed to create xmlrpc client; %w", err)
	}
	return &client{c}, nil
}

func (c *client) RestoreConfigSection(cfg Config) error {
	res := &OperationResult{}
	if err := c.Call("pfsense.restore_config_section", cfg, res); err != nil {
		return fmt.Errorf("failed to restore config section; %w", err)
	}
	if !res.Success {
		return errors.New("pfsense return 'false' as a result of config restoring")
	}
	return nil
}

func (c *client) BackupConfigSection(section string) (Config, error) {
	req := &struct{ Data []string }{Data: []string{section}}
	res := &Config{}
	if err := c.Call("pfsense.backup_config_section", req, res); err != nil {
		return Config{}, fmt.Errorf("failed to call %s; %w", "backup_config_section", err)
	}
	return *res, nil
}

func (c *client) ConfigureUnbound() error {
	code := `$toreturn = services_unbound_configure(false);`
	if err := c.execPhp(code); err != nil {
		return fmt.Errorf("failed to exec php to configure unbound")
	}
	return nil
}

func (c *client) ConfigureDhcpd() error {
	code := `$toreturn = services_dhcpd_configure();`
	if err := c.execPhp(code); err != nil {
		return fmt.Errorf("failed to exec php to configure dhcpd")
	}
	return nil
}

func (c *client) execPhp(code string) error {
	req := &struct{ Data string }{Data: code}
	res := &OperationResult{}
	if err := c.Call("pfsense.exec_php", req, res); err != nil {
		return fmt.Errorf("failed to exec php; %w", err)
	}
	if !res.Success {
		return errors.New("pfsense return 'false' as a result of exec php")
	}
	return nil
}
