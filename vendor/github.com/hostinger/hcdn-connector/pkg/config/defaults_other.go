//go:build !darwin && !windows

package config

var (
	DefaultConfigPath  = "/etc/hcdn-connector/config.yml"
	DefaultAdminSocket = "/run/hcdn-connector/admin.sock"
)
