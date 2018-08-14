package config

import (
	"fmt"
	"net"
)

type Config struct {
	Port uint8
	Host net.IP
}

func (c *Config) GetConnectionString() string {
	return fmt.Sprintf("%s:%d", c.Host.String(), c.Port)
}
